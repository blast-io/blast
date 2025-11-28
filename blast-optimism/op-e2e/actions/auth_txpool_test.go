package actions

import (
	"blast/blockchain"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-e2e/files"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-wheel/engine"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-plugin"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func statusDataForInMemory(ctx context.Context, cl *ethclient.Client) (*engine.StatusData, error) {
	latestBlock, err := cl.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	head := latestBlock.Header()
	blkRef := eth.L1BlockRef{
		Hash:       head.Hash(),
		Number:     head.Number.Uint64(),
		Time:       head.Time,
		ParentHash: head.ParentHash,
	}
	return &engine.StatusData{
		Head:      blkRef,
		Safe:      blkRef,
		Finalized: blkRef,
	}, nil
}

func aSendToB(
	t Testing, cl *ethclient.Client,
	fromKey *ecdsa.PrivateKey, to common.Address,
	chainID *big.Int,
	signer types.Signer,
) *types.Transaction {
	fromAddr := crypto.PubkeyToAddress(fromKey.PublicKey)
	n, err := cl.NonceAt(t.Ctx(), fromAddr, nil)
	require.NoError(t, err)
	current, err := cl.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)

	tx := types.MustSignNewTx(fromKey, signer, &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     n,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: new(big.Int).Add(current.Header().BaseFee, big.NewInt(2*params.GWei)),
		Gas:       params.TxGas,
		To:        &to,
		Value:     e2eutils.Ether(2),
	})

	return tx
}

func contractAction(
	t Testing, cl *ethclient.Client,
	fromKey *ecdsa.PrivateKey,
	toAddr common.Address,
	payload []byte,
	amt *big.Int,
	chainID *big.Int,
	signer types.Signer,
) *types.Transaction {
	fromAddr := crypto.PubkeyToAddress(fromKey.PublicKey)
	n, err := cl.NonceAt(t.Ctx(), fromAddr, nil)
	require.NoError(t, err)
	current, err := cl.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)

	txData := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     n,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: new(big.Int).Add(current.Header().BaseFee, big.NewInt(2*params.GWei)),
		Gas:       3_000_000,
		Value:     common.Big0,
		Data:      payload,
	}

	if amt != nil {
		txData.Value = amt
	}

	if toAddr != (common.Address{}) {
		txData.To = &toAddr
	}

	return types.MustSignNewTx(fromKey, signer, txData)
}

func runOffInMemoryAuthPool(
	t Testing,
	log log.Logger,
	startingArgs *blockchain.NewChainStartingArgs,
	jwt []byte,
	dp *e2eutils.DeployParams,
	sd *e2eutils.SetupData,
) (blockchain.Chain, *plugin.Client) {
	blastchain, plgin, err := loadPlugin(log, blastPluginOpts)
	require.NoError(t, err)
	newChainOrErr := blastchain.NewChain(startingArgs)

	require.NotNil(t, newChainOrErr.SerializedHeader, "new-chain still died")
	var head *types.Header
	require.NoError(t, json.Unmarshal(newChainOrErr.SerializedHeader, &head))

	endpoint, err := blastchain.WSEndpoint()
	require.NoError(t, err)
	authRoute, err := blastchain.AuthEndpoint()
	require.NoError(t, err)
	log.Info("loaded up jwt", "jwt", string(jwt), "authed-route", authRoute, "ws-endpoint", endpoint)

	minerClient, err := ethclient.Dial(endpoint)
	require.NoError(t, err)

	chainID, err := minerClient.ChainID(t.Ctx())
	require.NoError(t, err)

	latestSigner := types.LatestSignerForChainID(chainID)

	auth := rpc.WithHTTPAuth(node.NewJWTAuth([32]byte(jwt)))
	opts := []client.RPCOption{
		client.WithGethRPCOptions(auth),
		client.WithDialBackoff(10),
	}

	l2NodeEngine, err := client.NewRPC(t.Ctx(), log, authRoute, opts...)
	require.NoError(t, err)

	replayChainEngine := sources.NewEngineAPIClientWithTimeout(
		l2NodeEngine, log, sd.RollupCfg, 10*time.Second,
	)

	latestBlock, err := minerClient.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)

	buildNextBlock := func(txs []eth.Data, descr string, sampleTxPoolTime time.Duration) {
		statusData, err := statusDataForInMemory(t.Ctx(), minerClient)
		require.NoError(t, err)
		blkSetting := &engine.BlockBuildingSettings{
			GasLimit:     latestBlock.GasLimit(),
			BlockTime:    uint64(2),
			Transactions: txs,
			// NOTE critical to make non-zero if you want txpool to be used at all
			BuildTime: sampleTxPoolTime,
		}

		payloadEnv, err := engine.BuildBlock(t.Ctx(), replayChainEngine, statusData, blkSetting)
		require.NoError(t, err)
		log.Info("successfully built block",
			"description", descr,
			"block-number", payloadEnv.ExecutionPayload.BlockNumber,
			"block.time", time.Unix(int64(payloadEnv.ExecutionPayload.Timestamp), 0),
			"forced-tx-count", len(payloadEnv.ExecutionPayload.Transactions),
		)
	}

	balanceOf := func(who common.Address) []byte {
		packed, err := WETH.ABI.Pack("balanceOf", who)
		require.NoError(t, err)
		return packed
	}

	incomingReceipts, incomingHead := make(chan *types.Receipt), make(chan *types.Header, 8)
	subFastReceipt, err := minerClient.SubscribeFastReceipt(t.Ctx(), incomingReceipts)
	require.NoError(t, err)
	defer subFastReceipt.Unsubscribe()

	subNewHead, err := minerClient.SubscribeNewHead(t.Ctx(), incomingHead)
	require.NoError(t, err)
	defer subNewHead.Unsubscribe()

	type fastReceive struct {
		atWhatBlockHeadShouldBeReceived, atWhatBlockHeadShouldBeContained, logsCountExpected uint64
		tag                                                                                  string
	}
	// we should get the receipt even before the block offically comes out
	fastReceiptExpectations := map[common.Hash]*fastReceive{}
	plainTxsPool := map[common.Hash]struct{}{}
	_ = plainTxsPool

	go func() {
		verifyClient, err := ethclient.Dial(endpoint)
		require.NoError(t, err)
		currentHeadNumber := uint64(0)

		for {
			select {
			case err := <-subFastReceipt.Err():
				if err != nil {
					log.Error("problem from fast receipt subscription", "err", err)
				}
				return
			case r := <-incomingReceipts:
				log.Trace("received fast receipt", "hsh", r.TxHash.Hex(), "logs-count", len(r.Logs))
				spec, ok := fastReceiptExpectations[r.TxHash]
				require.True(t, ok, "expected receipt not captured")
				require.Equal(t, spec.atWhatBlockHeadShouldBeReceived, currentHeadNumber, spec.tag)
				require.Equal(t, spec.logsCountExpected, uint64(len(r.Logs)),
					fmt.Sprintf("tagged: %s expected logs count %d but got %d",
						spec.tag, spec.logsCountExpected, len(r.Logs),
					),
				)
			case head := <-incomingHead:
				currentHeadNumber = head.Number.Uint64()
				allReceipts, err := verifyClient.BlockReceipts(
					t.Ctx(), rpc.BlockNumberOrHashWithHash(head.Hash(), true),
				)
				require.NoError(t, err)

				shouldHaveThisBlock, countHave := 0, 0
				for _, v := range fastReceiptExpectations {
					if v.atWhatBlockHeadShouldBeContained == currentHeadNumber {
						shouldHaveThisBlock++
					}
				}

				// so we know we should have them now
				for _, r := range allReceipts {
					_, doHave := fastReceiptExpectations[r.TxHash]
					if doHave {
						countHave++
					}
				}

				require.Equal(t, shouldHaveThisBlock, countHave)
			}
		}
	}()

	txFastPool := aSendToB(t, minerClient, dp.Secrets.Mallory, dp.Addresses.Alice, chainID, latestSigner)
	fastReceiptExpectations[txFastPool.Hash()] = &fastReceive{
		atWhatBlockHeadShouldBeReceived:  0,
		atWhatBlockHeadShouldBeContained: 1,
		tag:                              "plain value transfer",
	}

	createWethEncoded, err := WETH.ABI.Constructor.Inputs.Pack()
	require.NoError(t, err)

	txPlainPoolWethCreation := contractAction(
		t, minerClient, dp.Secrets.Alice, common.Address{},
		append(WETH.Bytecode, createWethEncoded...), nil,
		chainID, latestSigner,
	)
	wethAddr := crypto.CreateAddress(dp.Addresses.Alice, 0)

	require.NoError(t, minerClient.SendTransaction(t.Ctx(), txPlainPoolWethCreation), "plain tx pool send died")
	require.NoError(t, minerClient.BlastSendTransaction(t.Ctx(), txFastPool), "blast send tx died")
	require.ErrorContains(t, minerClient.BlastSendTransaction(t.Ctx(), txFastPool), txpool.ErrAlreadyKnown.Error())

	buildNextBlock(nil, "first block", 800*time.Millisecond)
	receiptMap, err := minerClient.BlastTransactionReceipt(t.Ctx(), txFastPool.Hash())
	require.NoError(t, err)
	log.Info("result received", "m", receiptMap)

	wethCreationReceipt, err := minerClient.TransactionReceipt(t.Ctx(), txPlainPoolWethCreation.Hash())
	require.NoError(t, err)
	require.Condition(t, func() bool { return wethCreationReceipt.Status == types.ReceiptStatusSuccessful }, "weth creation failed")

	txDepositWethFast := contractAction(
		t, minerClient, dp.Secrets.Alice, wethAddr, nil, e2eutils.Ether(2),
		chainID, latestSigner,
	)
	fastReceiptExpectations[txDepositWethFast.Hash()] = &fastReceive{
		atWhatBlockHeadShouldBeReceived:  1,
		atWhatBlockHeadShouldBeContained: 2,
		logsCountExpected:                1,
		tag:                              "deposit weth",
	}
	require.NoError(t, minerClient.BlastSendTransaction(t.Ctx(), txDepositWethFast), "blast send tx died")
	buildNextBlock(nil, "second block", 300*time.Millisecond)
	sendToBobAmt := new(big.Int).Div(e2eutils.Ether(1), common.Big2)

	buildNextBlock(nil, "third block", 300*time.Millisecond)
	wethTransferPacked, err := WETH.ABI.Pack("transfer", dp.Addresses.Bob, sendToBobAmt)
	require.NoError(t, err)

	txTransferWeth := contractAction(
		t, minerClient, dp.Secrets.Alice, wethAddr, wethTransferPacked, nil, chainID, latestSigner,
	)
	fastReceiptExpectations[txTransferWeth.Hash()] = &fastReceive{
		atWhatBlockHeadShouldBeReceived:  3,
		atWhatBlockHeadShouldBeContained: 4,
		logsCountExpected:                1,
		tag:                              "tranfer weth ",
	}

	log.Info("sending weth transfer eval", "hsh", txTransferWeth.Hash().Hex())

	require.NoError(t, minerClient.BlastSendTransaction(t.Ctx(), txTransferWeth), "blast send tx died")
	buildNextBlock(nil, "fourth block", 800*time.Millisecond)
	result, err := minerClient.CallContract(t.Ctx(), ethereum.CallMsg{
		To:   &wethAddr,
		Data: balanceOf(dp.Addresses.Bob),
	}, nil)
	log.Info("show balance")
	require.NoError(t, err)
	balBob, err := WETH.ABI.Unpack("balanceOf", result)
	require.NoError(t, err)
	require.Equal(t, balBob[0].(*big.Int).Cmp(sendToBobAmt) == 0, true)
	return blastchain, plgin
}

var (
	rollupcfg rollup.Config
)

func init() {
	if err := json.Unmarshal(files.BlastMainnetRollup, &rollupcfg); err != nil {
		panic("could not load rollup config blast mainnet " + err.Error())
	}
}

func runOffExistingDB(
	t Testing,
	log log.Logger,
	startingArgs *blockchain.NewChainStartingArgs,
	jwt []byte,
	dp *e2eutils.DeployParams,
	sd *e2eutils.SetupData,
) (blockchain.Chain, *plugin.Client) {
	startingArgs.UseDatadir = "/Volumes/eth-chaindata-8tb/latest-blast-mainnet-21_331_481/blast-mainnet/geth"
	mainnetKey := os.Getenv("BLAST_MAINNET_KEY")
	require.NotEmpty(t, mainnetKey, "blast mainnet key is empty, need BLAST_MAINNET_KEY env variable")
	key, err := crypto.HexToECDSA(mainnetKey)
	require.NoError(t, err, "could not parse hex of env BLAST_MAINNET_KEY")
	addr := crypto.PubkeyToAddress(key.PublicKey)
	log.Info("have mainnet key to run off blast mainnet db", "addr", addr.Hex())
	blastchain, plgin, err := loadPlugin(log, blastPluginOpts)
	require.NoError(t, err)

	newChainOrErr := blastchain.NewChain(startingArgs)
	var head *types.Header
	require.NoError(t, json.Unmarshal(newChainOrErr.SerializedHeader, &head))
	endpoint, err := blastchain.WSEndpoint()
	require.NoError(t, err)
	authRoute, err := blastchain.AuthEndpoint()
	require.NoError(t, err)
	minerClient, err := ethclient.Dial(endpoint)
	require.NoError(t, err)
	gClient := gethclient.New(minerClient.Client())
	goTo := big.NewInt(head.Number.Int64())
	require.Condition(t, func() bool { return goTo.Cmp(common.Big0) != 0 }, "head number was zero")
	defer gClient.SetHead(t.Ctx(), goTo)
	l2NodeEthRPC, err := client.NewRPC(t.Ctx(), log, endpoint)
	require.NoError(t, err)
	l2NodeEngine, err := client.NewRPC(t.Ctx(), log, authRoute)
	require.NoError(t, err)

	replayChainEngine := sources.NewEngineAPIClientWithTimeout(l2NodeEngine, log, &rollupcfg, time.Second*10)
	require.NoError(t, err)
	latestBlock, err := minerClient.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)

	buildNextBlock := func(
		txs []eth.Data, descr string, sampleTxPoolTime time.Duration,
	) *eth.ExecutionPayloadEnvelope {
		statusData, err := engine.Status(t.Ctx(), l2NodeEthRPC)
		require.NoError(t, err)
		blkSetting := &engine.BlockBuildingSettings{
			GasLimit:     latestBlock.GasLimit(),
			BlockTime:    uint64(2),
			Transactions: txs,
			// NOTE critical to make non-zero if you want txpool to be used at all
			BuildTime: sampleTxPoolTime,
		}

		payloadEnv, err := engine.BuildBlock(t.Ctx(), replayChainEngine, statusData, blkSetting)
		require.NoError(t, err)
		log.Info("successfully built block",
			"description", descr,
			"block-number", payloadEnv.ExecutionPayload.BlockNumber,
			"block.time", time.Unix(int64(payloadEnv.ExecutionPayload.Timestamp), 0),
			"forced-tx-count", len(txs),
			"tx-count-in-block", len(payloadEnv.ExecutionPayload.Transactions),
		)
		return payloadEnv
	}

	chainID, err := minerClient.ChainID(t.Ctx())
	require.NoError(t, err)
	latestSigner := types.LatestSignerForChainID(chainID)
	encodedERC20, err := ERC20.ABI.Constructor.Inputs.Pack("memecoin", uint8(18), "MEME")
	require.Nil(t, err, "failed to pack memecoin creation")
	incomingReceipts := make(chan *types.Receipt)
	subFastReceipt, err := minerClient.SubscribeFastReceipt(t.Ctx(), incomingReceipts)
	require.NoError(t, err)
	defer subFastReceipt.Unsubscribe()

	var memeAddr common.Address
	txFastMemeCreation := contractAction(
		t, minerClient, key, common.Address{}, append(ERC20.Bytecode, encodedERC20...), nil,
		chainID, latestSigner,
	)

	go func() {
		for {
			select {
			case err := <-subFastReceipt.Err():
				if err != nil {
					log.Error("problem from fast receipt subscription", "err", err)
				}
				return
			case r := <-incomingReceipts:
				log.Trace("received fast receipt", "hsh", r.TxHash.Hex(), "logs-count", len(r.Logs))
				if r.TxHash == txFastMemeCreation.Hash() && r.ContractAddress != (common.Address{}) {
					memeAddr = r.ContractAddress
				}
			}
		}
	}()

	require.NoError(t, minerClient.BlastSendTransaction(t.Ctx(), txFastMemeCreation), "blast send tx died")
	// NOTE will block for the sample tx pool time
	execPayload := buildNextBlock(nil, "block off mainnet db with erc20 meme-coin creation ", 2*time.Second)

	const (
		plainPool = iota
		fastPool
		randomMix
	)

	mintWork := func(poolChoice int) (result []*types.Transaction) {
		n, err := minerClient.NonceAt(t.Ctx(), addr, nil)
		require.NoError(t, err)
		encodedMint, err := ERC20.ABI.Pack("mint", e2eutils.Ether(2))
		require.Nil(t, err, "failed to pack memecoin creation")
		baseFee := (*uint256.Int)(&execPayload.ExecutionPayload.BaseFeePerGas).ToBig()
		time.Sleep(100 * time.Millisecond)

		for i := range 3 {
			txData := &types.DynamicFeeTx{
				ChainID:   chainID,
				Nonce:     n + uint64(i),
				To:        &memeAddr,
				GasTipCap: big.NewInt(2 * params.GWei),
				GasFeeCap: new(big.Int).Add(baseFee, big.NewInt(2*params.GWei)),
				Gas:       120_000,
				Value:     common.Big0,
				Data:      encodedMint,
			}

			readyTx := types.MustSignNewTx(key, latestSigner, txData)
			result = append(result, readyTx)

			switch poolChoice {
			case plainPool:
				require.NoError(t, minerClient.SendTransaction(t.Ctx(), readyTx), "using plain tx pool failed")
			case fastPool:
				require.NoError(t, minerClient.BlastSendTransaction(t.Ctx(), readyTx), "using fast pool tx failed")
			case randomMix:
				switch v := rand.Intn(4); v {
				case 0:
					if err := minerClient.SendTransaction(t.Ctx(), readyTx); err != nil {
						//          "plain send tx died case 0")
					}
				case 1:
					if err := minerClient.BlastSendTransaction(t.Ctx(), readyTx); err != nil {
						//
					}
				case 2:
					if err := minerClient.SendTransaction(t.Ctx(), readyTx); err != nil {
						//
					}
					if err := minerClient.BlastSendTransaction(t.Ctx(), readyTx); err != nil {
						//
					}
				case 3:
					if err := minerClient.BlastSendTransaction(t.Ctx(), readyTx); err != nil {
						//
					}
					if err := minerClient.SendTransaction(t.Ctx(), readyTx); err != nil {
						//
					}
				}
			}
		}

		return result
	}

	// NOTE need this time so tx pool can clean up pool address reservations after getting
	// notification about new head and txs that made it through
	time.Sleep(3 * time.Second)
	txs := mintWork(plainPool)
	execPayload = buildNextBlock(nil, "building second block off mainnet db", 2*time.Second)
	for _, tx := range txs {
		rcpt, err := minerClient.BlastTransactionReceipt(t.Ctx(), tx.Hash())
		require.NoError(t, err)
		require.Condition(t, func() bool { return rcpt.Status == types.ReceiptStatusSuccessful })
	}

	// NOTE so after we used the fast pool, can we still land a tx in plain pool ?
	mintWork(fastPool)
	execPayload = buildNextBlock(nil, "building third block off mainnet db", 2*time.Second)
	log.Info("able to get plain tx after fast tx too")
	return blastchain, plgin
}

type authTestCase func(
	Testing, log.Logger, *blockchain.NewChainStartingArgs, []byte,
	*e2eutils.DeployParams, *e2eutils.SetupData,
) (blockchain.Chain, *plugin.Client)

func TestAuthTxPool(gt *testing.T) {
	t := NewDefaultTesting(gt)
	lg := testlog.Logger(t, log.LvlTrace)
	jwtPath := e2eutils.WriteDefaultJWT(t)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	faucet := dp.Addresses.Deployer.Hex()
	bigWheel := e2eutils.Ether(20)
	require.NoError(t, buildPlugin(lg, blastPluginOpts))
	jwt, err := os.ReadFile(jwtPath)
	require.NoError(t, err)

	inMemoryStartingArgs := &blockchain.NewChainStartingArgs{
		JWTFilePath:         jwtPath,
		IncludeCatalystAPI:  true,
		CatalystAuthEnabled: false,
		// SerializedGenesis:  serialized,
		AuthPort: 2312,
		WSPort:   2222,
		Faucet:   faucet,
		ExtraAllocs: map[string]*big.Int{
			dp.Addresses.Mallory.Hex(): bigWheel,
			dp.Addresses.Alice.Hex():   bigWheel,
			dp.Addresses.Bob.Hex():     bigWheel,
		},
		MinerRecommit:          100 * time.Millisecond,
		MinerNewPayloadTimeout: 300 * time.Millisecond,
	}

	mainnetStartingArgs := &blockchain.NewChainStartingArgs{
		JWTFilePath:         jwtPath,
		IncludeCatalystAPI:  true,
		CatalystAuthEnabled: false,
		SerializedGenesis:   files.BlastMainnetGenesis,
		AuthPort:            2312,
		WSPort:              2222,
		// NOTE keep like mainnet
		MinerRecommit: 200 * time.Millisecond,
		// NOTE must have this be high enough so that payload gets pulled properly from txpools, 2
		// seconds is default but keeping it explicit
		MinerNewPayloadTimeout: 2000 * time.Millisecond,
	}

	for _, tstCase := range []struct {
		name              string
		kickOff           authTestCase
		startingChainArgs *blockchain.NewChainStartingArgs
		skip              bool
	}{
		{
			name:              "in memory only",
			kickOff:           runOffInMemoryAuthPool,
			startingChainArgs: inMemoryStartingArgs,
			skip:              true,
		},
		{
			name:              "mainnet db",
			kickOff:           runOffExistingDB,
			startingChainArgs: mainnetStartingArgs,
			skip:              false,
		},
	} {
		gt.Run(tstCase.name, func(gt *testing.T) {
			if tstCase.skip {
				gt.Skip("skipping test")
			}
			t := NewDefaultTesting(gt)
			log := testlog.Logger(t, log.LvlTrace)
			blkchain, _ := tstCase.kickOff(t, log, tstCase.startingChainArgs, jwt, dp, sd)
			if blkchain != nil {
				blkchain.Close()
			}
			log.Info("finished testcase", "kind", tstCase.name)
		})
	}

}
