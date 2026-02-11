package actions

import (
	"blast/blockchain"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	emath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/stretchr/testify/require"
)

func readable(p *big.Int, decimalCount int64) string {
	b := new(big.Float).SetInt(p)
	dec := emath.Exp(big.NewInt(10), big.NewInt(decimalCount))
	return new(big.Float).Quo(b, new(big.Float).SetInt(dec)).Text('f', int(decimalCount))
}

func ptr(n uint64) *uint64 {
	return &n
}

var (
	assumeL1BlockTime = uint64(12)
	offsets           = struct {
		cancun, prague, osaka, bpo1, bpo2, bpo2Blast *uint64
	}{
		// NOTE can't pick 0 because glob can't distinguish between 0 and nil
		cancun:    ptr(0),
		prague:    ptr(2 * assumeL1BlockTime),
		osaka:     ptr(330 * assumeL1BlockTime),
		bpo1:      ptr(440 * assumeL1BlockTime),
		bpo2:      ptr(550 * assumeL1BlockTime),
		bpo2Blast: ptr(660 * assumeL1BlockTime),
	}
)

func sendTx(
	t Testing,
	cl *ethclient.Client,
	fromKey *ecdsa.PrivateKey,
) func(to common.Address, data []byte, value *big.Int, forceNonce *uint64) *types.Transaction {

	fromAddr := crypto.PubkeyToAddress(fromKey.PublicKey)
	chainID, err := cl.ChainID(t.Ctx())
	require.NoError(t, err)
	signer := types.LatestSignerForChainID(chainID)

	return func(
		to common.Address,
		data []byte,
		value *big.Int,
		forceNonce *uint64,
	) *types.Transaction {
		n, err := cl.NonceAt(t.Ctx(), fromAddr, nil)
		require.NoError(t, err)
		if forceNonce != nil {
			n = *forceNonce
		}
		current, err := cl.BlockByNumber(t.Ctx(), nil)
		require.NoError(t, err)
		suggestTip, err := cl.SuggestGasTipCap(t.Ctx())
		require.NoError(t, err)
		tx := types.MustSignNewTx(fromKey, signer, &types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     n,
			GasTipCap: suggestTip,
			GasFeeCap: new(big.Int).Add(
				current.Header().BaseFee, big.NewInt(100000*params.GWei),
			),
			Gas:   2_000_000,
			To:    &to,
			Data:  data,
			Value: value,
		})
		require.NoError(t, cl.SendTransaction(t.Ctx(), tx), "nonce %v hsh %v", n, tx.Hash())
		return tx
	}
}

func sendBlob(
	t Testing,
	cl *ethclient.Client,
	fromKey *ecdsa.PrivateKey,
	l log.Logger,
) func(to common.Address, value *big.Int, howManyBlobs int) *types.Transaction {
	fromAddr := crypto.PubkeyToAddress(fromKey.PublicKey)
	chainID, err := cl.ChainID(t.Ctx())
	require.NoError(t, err)
	signer := types.LatestSignerForChainID(chainID)

	txMgr := txmgr.NewSimpleTxManagerTxCrafter(cl, l, &metrics.NoopTxMetrics{}, txmgr.Config{
		CellProofTime:  uint64(time.Now().Add((time.Duration(*offsets.osaka) * time.Second) + time.Second).Unix()),
		NetworkTimeout: 10 * time.Second,
		From:           fromAddr,
		Signer: func(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, signer, fromKey)
		},
	})

	return func(to common.Address, value *big.Int, howManyBlobs int) *types.Transaction {
		blobs := []*eth.Blob{}

		for range howManyBlobs {
			var b eth.Blob
			blobData := eth.Data(make([]byte, eth.MaxBlobDataSize))
			rand.Read(blobData)
			b.FromData(blobData)
			blobs = append(blobs, &b)
		}

		cand := txmgr.TxCandidate{
			To:     &to,
			TxData: []byte{0x00, 0x01, 0x02, 0x03},
			Blobs:  blobs,
			Value:  value,
		}

		tx, err := txMgr.CraftTx(t.Ctx(), cand)
		require.NoError(t, err)
		require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
		return tx
	}
}

func TestBlastBlobSchedulingFix(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	require.NoError(t, buildPlugin(log, gethL1PluginOpts))
	mnemonicCfg := e2eutils.DefaultMnemonicConfig
	_, err := mnemonicCfg.Secrets()
	require.NoError(t, err)

	pluginL1, _, err := loadPlugin(log, gethL1PluginOpts)
	require.NoError(t, err)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)

	dp.DeployConfig.L1CancunTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.cancun))
	dp.DeployConfig.L1PragueTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.prague))
	dp.DeployConfig.L1OsakaTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.osaka))
	dp.DeployConfig.L2GenesisDeltaTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.prague))
	dp.DeployConfig.L2GenesisEcotoneTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.prague))

	l1Gen, deployConf, l1Deployments := e2eutils.SetupPart1(t, dp, defaultAlloc)
	l1Gen.Config.CancunTime = offsets.cancun
	l1Gen.Config.PragueTime = offsets.prague
	l1Gen.Config.OsakaTime = offsets.osaka
	l1Gen.Config.BPO1Time = offsets.bpo1

	extraAllocs := map[string]blockchain.Account{}

	for addr, payload := range l1Gen.Alloc {
		if addr == (common.Address{}) {
			continue
		}

		storage := map[string]string{}
		for key, value := range payload.Storage {
			storage[key.Hex()] = value.Hex()
		}

		extraAllocs[addr.Hex()] = blockchain.Account{
			Code:    payload.Code,
			Balance: payload.Balance,
			Storage: storage,
		}
	}

	resultL1 := pluginL1.NewChain(&blockchain.NewChainStartingArgs{
		WhenActivateCancun: offsets.cancun,
		WhenActivatePrague: offsets.prague,
		WhenActivateOsaka:  offsets.osaka,
		WhenActivateBPO1:   offsets.bpo1,
		WhenActivateBPO2:   offsets.bpo2,
		ExtraAllocs:        extraAllocs,
		MinerRecommit:      10 * time.Millisecond,
	})

	require.NoError(t, err)

	var (
		deserializedHeader *types.Header
		deserializedParams *params.ChainConfig
	)
	//	log.Debug("show dump", "result", string(resultL1.SerializedHeader), "raw-result", resultL1)

	require.NoError(t, json.Unmarshal(resultL1.SerializedHeader, &deserializedHeader))
	require.NoError(t, json.Unmarshal(resultL1.SerializedChainParam, &deserializedParams))

	l1Blk := types.NewBlock(
		deserializedHeader, nil, nil, nil, trie.NewStackTrie(nil),
	)

	sd := e2eutils.SetupPart2(
		t, dp, defaultAlloc, deployConf, l1Blk, l1Gen, l1Deployments,
	)
	sd.RollupCfg.L1ChainID = big.NewInt(1337)
	sd.RollupCfg.RegolithTime = ptr(0)
	sd.RollupCfg.CanyonTime = ptr(0)
	sd.RollupCfg.DeltaTime = ptr(0)
	sd.RollupCfg.EcotoneTime = offsets.prague

	sd.L1Cfg.Config.Blast = &params.BlastOverrides{
		OsakaBlobConfigOverride: params.DefaultPragueBlobConfig,
		BPO1BlobConfigOverride:  params.DefaultPragueBlobConfig,
		BPO2BlobConfigOverride:  params.DefaultPragueBlobConfig,
		// BPO2BlastBlobConfigOverride : params.DefaultPragueBlobConfig,
	}

	l1GenHash := sd.L1Cfg.ToBlock().Hash()
	l2GenHash := sd.L2Cfg.ToBlock().Hash()
	dp.DeployConfig.L1CancunTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.cancun))
	dp.DeployConfig.L1PragueTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.prague))
	dp.DeployConfig.L1OsakaTimeOffset = (*hexutil.Uint64)(unsafe.Pointer(offsets.osaka))

	log.Info(
		"genesis hashes",
		"l1-hsh", l1GenHash.Hex(),
		"l2-hsh", l2GenHash.Hex(),
		"doing-state-hash", sd.L1Cfg.StateHash == nil,
	)

	miner := NewPluginBackedMiner(t, log, pluginL1, assumeL1BlockTime)
	_, _, sequencer, seqEngine, verifier, _, batcher := setupReorgTestActorsWithL1Plugin(
		t, dp, sd, log, miner)
	l1Client := miner.EthClient()
	l2Client := seqEngine.EthClient()
	l2TxSender := sendTx(t, l2Client, dp.Secrets.Alice)
	l1TxSendBlobBatcher := sendBlob(t, l1Client, dp.Secrets.Batcher, log)
	l1TxSendBlobAlice := sendBlob(t, l1Client, dp.Secrets.Alice, log)

	l2AliceToBobSimpleTx := func(forceNonce *uint64) (*big.Int, string, *types.Transaction) {
		sequencer.ActL2StartBlock(t)
		junk := make([]byte, 1024)
		rand.Read(junk)
		tx0 := l2TxSender(dp.Addresses.Bob, junk, e2eutils.Ether(1), forceNonce)
		seqEngine.ActL2IncludeTx(dp.Secrets.Addresses().Alice)(t)
		sequencer.ActL2EndBlock(t)
		time.Sleep(200 * time.Millisecond)
		rcpt0, err := l2Client.TransactionReceipt(t.Ctx(), tx0.Hash())
		require.NoError(t, err)
		ez := readable(rcpt0.L1Fee, 18)
		log.Info("l2 alice to bob simple tx",
			"nonce", tx0.Nonce(),
			"l1-fee", ez,
			"landed", tx0.Hash().Hex(),
		)
		return rcpt0.L1Fee, ez, tx0
	}

	// start op-nodes
	sequencer.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	log.Info("system config addr shoudl be", "addr", sd.RollupCfg.L1SystemConfigAddress.Hex())

	fee0, fee0S, tx0 := l2AliceToBobSimpleTx(nil)
	log.Info("first alice to bob tx", "hsh", tx0.Hash().Hex())
	blobFee0, err := l1Client.BlobBaseFee(t.Ctx())
	require.NoError(t, err)

	log.Info("raw dump params", "params", string(resultL1.SerializedChainParam))

	// NOTE up until osaka we used the correct blob pricing - that is prague based
	for {
		currentBlock, err := l1Client.BlockByNumber(t.Ctx(), nil)
		require.NoError(t, err)
		if currentBlock.Time() == *offsets.osaka {
			break
		} else {
			log.Info("l1 not at osaka yet - build a block",
				"current-block", currentBlock.Number(),
				"current-timestamp", currentBlock.Time(),
			)
			miner.ActL1StartBlock(assumeL1BlockTime)(t)
			assumeTime := time.Unix(
				int64(currentBlock.Time()), 0).Add(time.Duration(assumeL1BlockTime) * time.Second)
			log.Info("time now assumed", "assume", uint64(assumeTime.Unix()))
			blobsAllowed := eip4844.MaxBlobsPerBlock(deserializedParams, uint64(assumeTime.Unix()))
			require.NotZero(t, blobsAllowed, "blobs was 0 but shuldnt be %+v", deserializedParams)
			tx := l1TxSendBlobBatcher(deployConf.BatchInboxAddress, common.Big0, blobsAllowed-3)
			require.NoError(t, pluginL1.IncludeTxByHash(tx.Hash().Hex()))
			tx = l1TxSendBlobAlice(deployConf.BatchInboxAddress, common.Big0, 3)
			require.NoError(t, pluginL1.IncludeTxByHash(tx.Hash().Hex()))
			newBlock := miner.ActL1EndBlock(t)
			_ = newBlock
			rcpt0, err := l1Client.TransactionReceipt(t.Ctx(), tx.Hash())
			require.NoError(t, err)
			require.Condition(t, func() bool {
				return rcpt0.Status == types.ReceiptStatusSuccessful
			})
		}
	}

	blobFee1, err := l1Client.BlobBaseFee(t.Ctx())
	require.NoError(t, err)
	miner.ActL1StartBlock(assumeL1BlockTime)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)

	sequencer.ActL2PipelineFull(t)
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	log.Info("check blob fee after first runup",
		"blob-fee-then", blobFee0,
		"blob-fee-now", blobFee1,
	)
	require.NoError(t, err)

	fee1, fee1S, tx1 := l2AliceToBobSimpleTx(nil)

	log.Info("compare the l2 fees alice to bob second",
		"hsh", tx1.Hash().Hex(),
		"fee0", fee0S, "fee1", fee1S,
		"fee1-greater-fee0", fee1.Cmp(fee0) == 1,
	)

	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	// build empty L1 blocks, crossing the fork boundary
	miner.ActL1SetFeeRecipient(common.Address{'A', 0})
	miner.ActEmptyBlock()
	miner.ActEmptyBlock()
	miner.ActEmptyBlock()
	currentBlock, err := l1Client.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)
	l1Head := currentBlock.Header()
	require.True(t, sd.L1Cfg.Config.IsCancun(l1Head.Number, l1Head.Time),
		"Cancun should be active at %v", *sd.L1Cfg.Config.CancunTime,
	)
	require.True(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time),
		"Prague should be active at %v", *sd.L1Cfg.Config.PragueTime,
	)

	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)
	miner.ActL1StartBlock(12)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)
	miner.ActL1EndBlock(t)

}
