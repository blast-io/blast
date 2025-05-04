package actions

import (
	"crypto/ecdsa"
	_ "embed"
	"encoding/json"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	ethPebble "github.com/ethereum/go-ethereum/ethdb/pebble"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

//go:embed combined-contracts.json
var allEmbedded []byte

type solcThrowaway struct {
	Contracts struct {
		IERC20 struct {
			ABI json.RawMessage `json:"abi"`
		} `json:"contracts/interfaces.sol:IERC20"`
		IWETH struct {
			ABI json.RawMessage `json:"abi"`
		} `json:"contracts/interfaces.sol:IWETH"`
		ERC20 struct {
			ABI      json.RawMessage `json:"abi"`
			Bytecode string          `json:"bin"`
		} `json:"contracts/basic-erc20.sol:ERC20"`
		WETH struct {
			ABI      json.RawMessage `json:"abi"`
			Bytecode string          `json:"bin"`
		} `json:"contracts/weth.sol:WETH9"`
		UniswapV2Factory struct {
			ABI      json.RawMessage `json:"abi"`
			Bytecode string          `json:"bin"`
		} `json:"contracts/full-uniswap-v2.sol:UniswapV2Factory"`
		UniswapV2Router struct {
			ABI      json.RawMessage `json:"abi"`
			Bytecode string          `json:"bin"`
		} `json:"contracts/full-uniswap-v2.sol:UniswapV2Router02"`
		UniswapV2Pair struct {
			ABI      json.RawMessage `json:"abi"`
			Bytecode string          `json:"bin"`
		} `json:"contracts/full-uniswap-v2.sol:UniswapV2Pair"`
	} `json:"contracts"`
}

var (
	loadIt solcThrowaway
)

func init() {
	if err := json.Unmarshal(allEmbedded, &loadIt); err != nil {
		panic("not possible missing combined-abi-bin-liq.json " + err.Error())
	}
	UniswapV2Factory.Bytecode = common.Hex2Bytes(loadIt.Contracts.UniswapV2Factory.Bytecode)
	if len(UniswapV2Factory.Bytecode) == 0 {
		panic("not possible uniswapv2 factory bytecode len 0")
	}

	v2FactABI, err := abi.JSON(
		strings.NewReader(string(loadIt.Contracts.UniswapV2Factory.ABI)),
	)
	if err != nil {
		panic("no possible " + err.Error())
	}

	UniswapV2Factory.ABI = v2FactABI

	UniswapV2Router.Bytecode = common.Hex2Bytes(loadIt.Contracts.UniswapV2Router.Bytecode)
	if len(UniswapV2Router.Bytecode) == 0 {
		panic("not possible uniswap v2 router byte len 0")
	}

	v2RouterABI, err := abi.JSON(
		strings.NewReader(string(loadIt.Contracts.UniswapV2Router.ABI)),
	)
	if err != nil {
		panic("not possible " + err.Error())
	}

	UniswapV2Router.ABI = v2RouterABI

	UniswapV2Pair.Bytecode = common.Hex2Bytes(loadIt.Contracts.UniswapV2Pair.Bytecode)
	if len(UniswapV2Pair.Bytecode) == 0 {
		panic("not possible uniswap v2 pair byte len 0")
	}

	v2PairABI, err := abi.JSON(
		strings.NewReader(string(loadIt.Contracts.UniswapV2Pair.ABI)),
	)
	if err != nil {
		panic("not possible " + err.Error())
	}
	UniswapV2Pair.ABI = v2PairABI
	erc20ABI, err := abi.JSON(
		strings.NewReader(string(loadIt.Contracts.ERC20.ABI)),
	)
	if err != nil {
		panic("cant load up erc20 abi json " + err.Error())
	}
	erc20Code := common.Hex2Bytes(loadIt.Contracts.ERC20.Bytecode)
	ERC20 = contract{ABI: erc20ABI, Bytecode: erc20Code}
	wethABI, err := abi.JSON(
		strings.NewReader(string(loadIt.Contracts.WETH.ABI)),
	)
	if err != nil {
		panic("cant load up weth abi json " + err.Error())
	}
	wethCode := common.Hex2Bytes(loadIt.Contracts.WETH.Bytecode)
	if len(wethCode) == 0 {
		panic("cant be 0 length byte code")
	}
	WETH = contract{ABI: wethABI, Bytecode: wethCode}
}

func dbSize(dbDir string) (string, error) {
	output, err := exec.Command("du", "-sh", dbDir).Output()
	if err != nil {
		return "", err
	}

	all := strings.Split(strings.TrimSpace(string(output)), "\t")
	return all[0], nil
}

type contract struct {
	ABI      abi.ABI
	Bytecode []byte
}

var (
	WETH             contract
	ERC20            contract
	UniswapV2Factory contract
	UniswapV2Router  contract
	UniswapV2Pair    contract
)

type encodedPayloads struct {
	WethCreation             []byte
	MemecoinCreation         []byte
	MemecoinMint             []byte
	UniswapV2FactoryCreation []byte
	UniswapV2RouterCreation  []byte
	UniswapV2CreatePair      []byte
	GetInitHash              []byte
}

func prepareEncodedPayloads(
	t Testing, memecoinMintAmount *big.Int,
	uniswapV2Creator, uniswapV2FactoryAddr, wethAddr, memecoinAddr common.Address,
) *encodedPayloads {
	encodedERC20, err := ERC20.ABI.Constructor.Inputs.Pack("memecoin", uint8(18), "MEME")
	require.Nil(t, err, "failed to pack memecoin creation")
	encodedMint, err := ERC20.ABI.Pack("mint", memecoinMintAmount)
	require.Nil(t, err, "failed to pack memecoin mint five million")
	encodedUniswapV2, err := UniswapV2Factory.ABI.Constructor.Inputs.Pack(uniswapV2Creator)
	require.Nil(t, err, "failed to pack uniswap v2 factory creation")
	createRouterEncoded, err := UniswapV2Router.ABI.Constructor.Inputs.Pack(uniswapV2FactoryAddr, wethAddr)
	require.Nil(t, err, "failed to pack uniswap v2 router creation")
	encodedCreatePair, err := UniswapV2Factory.ABI.Pack("createPair", wethAddr, memecoinAddr)
	require.Nil(t, err, "failed to pack create pair")
	encodedGetInitHash, err := UniswapV2Factory.ABI.Pack("getInitHash")
	require.Nil(t, err, "failed to pack getInitHash")

	createWethEncoded, err := WETH.ABI.Constructor.Inputs.Pack()
	require.Nil(t, err, "Failed to pack weth")

	return &encodedPayloads{
		WethCreation:             append(WETH.Bytecode, createWethEncoded...),
		MemecoinCreation:         append(ERC20.Bytecode, encodedERC20...),
		MemecoinMint:             encodedMint,
		UniswapV2FactoryCreation: append(UniswapV2Factory.Bytecode, encodedUniswapV2...),
		UniswapV2RouterCreation:  append(UniswapV2Router.Bytecode, createRouterEncoded...),
		UniswapV2CreatePair:      encodedCreatePair,
		GetInitHash:              encodedGetInitHash,
	}
}

type encodedTransfers struct {
	transferUniswapMemeCoin, transferUniswapWethAmt []byte
	uniswapV2PairMint                               []byte
}

func prepareEncodedTransfers(
	t Testing,
	uniswapV2Pair, uniswapMinter common.Address,
	wethAmtToUniswap, memeAmtToUniswap *big.Int,
) *encodedTransfers {
	transferUniswapMemeCoin, err := ERC20.ABI.Pack("transfer", uniswapV2Pair, memeAmtToUniswap)
	require.Nil(t, err, "failed packing transfer to uniswap V2 Pair memecoin")
	transferUniswapWethAmt, err := ERC20.ABI.Pack("transfer", uniswapV2Pair, wethAmtToUniswap)
	require.Nil(t, err, "failed packing transfer to uniswap V2 Pair weth transfer")
	mintEncodedUniswap, err := UniswapV2Pair.ABI.Pack("mint", uniswapMinter)
	require.Nil(t, err, "failed packing mint on uniswap pair")
	return &encodedTransfers{transferUniswapMemeCoin, transferUniswapWethAmt, mintEncodedUniswap}
}

func someSwappers(t Testing) {
	var swappers []*ecdsa.PrivateKey
	for i := 50; i < 60; i++ {
		k := big.NewInt(int64(i))
		padded := common.LeftPadBytes(k.Bytes(), 32)
		hex := common.Bytes2Hex(padded)
		key, err := crypto.HexToECDSA(hex)
		require.NoError(t, err)
		swappers = append(swappers, key)
	}
}

func TestCompactionRates(gt *testing.T) {
	deltaTimeOffset := hexutil.Uint64(0)
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   20, // larger than L1 block time we simulate in this test (12)
		SequencerWindowSize: 3,
		ChannelTimeout:      20,
		L1BlockTime:         12,
	}
	dp := e2eutils.MakeDeployParams(t, p)
	applyDeltaTimeOffset(dp, &deltaTimeOffset)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlInfo)
	miner, seqEngine, sequencer, dbDir := setupSequencerTestWithPebble(t, sd, log)
	gt.Cleanup(func() {
		os.RemoveAll(dbDir)
		log.Info("removed temporary dir", "location", dbDir)
	})
	cl := seqEngine.EthClient()
	miner.ActL1SetFeeRecipient(common.Address{'E'})
	l2Signer := types.LatestSigner(sd.L2Cfg.Config)
	l2BlockPerL1 := p.L1BlockTime / sd.RollupCfg.BlockTime
	// lets do a full sequencing window size worth of work and see how compaction
	// holds up in the sequencer and then in the verifier
	uniswapV2FactoryAddr := crypto.CreateAddress(dp.Addresses.Alice, 0)
	memecoinAddr := crypto.CreateAddress(dp.Addresses.Alice, 1)
	// hack - pulled from logs but whatever
	v2PairAddr := common.HexToAddress("0x15dba5efed90f066250565015b5c476bdd4adc4c")
	wethAddr := crypto.CreateAddress(dp.Addresses.Alice, 5)

	fiveMillion := big.NewInt(1e18)
	fiveMillion.Mul(fiveMillion, big.NewInt(5_000_000))
	fiftyWeth := new(big.Int).Mul(big.NewInt(50), big.NewInt(params.Ether))
	readyPayloads := prepareEncodedPayloads(
		t, fiveMillion,
		dp.Addresses.Alice, uniswapV2FactoryAddr, wethAddr, memecoinAddr,
	)
	readyTransfersPayloads := prepareEncodedTransfers(
		t, v2PairAddr, dp.Addresses.Alice, new(big.Int).Div(fiftyWeth, common.Big2), fiveMillion,
	)
	twentyFiveE18 := new(big.Int).Div(fiftyWeth, common.Big2)

	var (
		v2FactoryTx, memecoinTx, memecoinMintTx, v2PairCreation       common.Hash
		wethCreationTx, wethWrapTxAlice, pairLPMintTx, simpleTransfer common.Hash
	)
	encodedBalanceOfAlice, err := ERC20.ABI.Pack("balanceOf", dp.Addresses.Alice)
	require.NoError(t, err)
	encodedTotalSupply, err := ERC20.ABI.Pack("totalSupply")
	require.NoError(t, err)

	for i := uint64(0); i < sd.RollupCfg.SeqWindowSize; i++ {
		sequencer.ActL1HeadSignal(t)
		sequencer.ActL2PipelineFull(t)
		miner.ActL1StartBlock(p.L1BlockTime)(t)

		for j := uint64(0); j < l2BlockPerL1; j++ {
			sequencer.ActL2StartBlock(t)
			// deploy factory, router, weth, memecoin, mint memecoin, create v2 pair, wrap 25 weth
			if i == 1 && j == 1 {
				defaultTx := &types.DynamicFeeTx{
					ChainID:   sd.L2Cfg.Config.ChainID,
					GasTipCap: big.NewInt(2 * params.GWei),
					GasFeeCap: new(big.Int).Add(
						miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei),
					),
					Gas:  5_000_000,
					Data: readyPayloads.UniswapV2FactoryCreation,
				}
				// v2 factory
				tx := types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				v2FactoryTx = tx.Hash()
				defaultTx.Nonce = 1
				defaultTx.Data = readyPayloads.MemecoinCreation
				// erc20 memecoin creation
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				memecoinTx = tx.Hash()
				defaultTx.Nonce = 2
				defaultTx.To = &memecoinAddr
				defaultTx.Data = readyPayloads.MemecoinMint
				// erc20 memecoin mint to alice
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				memecoinMintTx = tx.Hash()
				defaultTx.Nonce = 3
				defaultTx.To = &dp.Addresses.Bob
				defaultTx.Value = common.Big1
				defaultTx.Data = nil
				// simple value transfer
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				simpleTransfer = tx.Hash()
				defaultTx.Nonce = 4
				defaultTx.Value = nil
				defaultTx.To = &uniswapV2FactoryAddr
				defaultTx.Data = readyPayloads.UniswapV2CreatePair
				// v2 pair creation
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				v2PairCreation = tx.Hash()
				defaultTx.Nonce = 5
				defaultTx.To = nil
				defaultTx.Data = readyPayloads.WethCreation
				// weth creation
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				wethCreationTx = tx.Hash()
				defaultTx.Nonce = 6
				defaultTx.Value = fiftyWeth
				defaultTx.To = &wethAddr
				defaultTx.Data = nil
				// wrap 50 weth
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				wethWrapTxAlice = tx.Hash()
				defaultTx.Nonce = 7
				defaultTx.Value = nil
				defaultTx.To = &wethAddr
				defaultTx.Data = readyTransfersPayloads.transferUniswapWethAmt
				// send weth to the pair
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				defaultTx.Nonce = 8
				defaultTx.To = &memecoinAddr
				defaultTx.Data = readyTransfersPayloads.transferUniswapMemeCoin
				// send memecoin to the pair
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				defaultTx.Nonce = 9
				defaultTx.To = &v2PairAddr
				defaultTx.Data = readyTransfersPayloads.uniswapV2PairMint
				// mint the position on the pair
				tx = types.MustSignNewTx(dp.Secrets.Alice, l2Signer, defaultTx)
				require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
				seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
				pairLPMintTx = tx.Hash()
			}

			if i == 1 && j == 2 {
				//
			}

			sequencer.ActL2EndBlock(t)
			// lets make sure we got the right factory addresses created as expected
			if i == 1 && j == 3 {
				var gasUsedSoFar uint64
				for k, h := range []common.Hash{
					v2FactoryTx, memecoinTx, memecoinMintTx, simpleTransfer, v2PairCreation,
					wethCreationTx, wethWrapTxAlice, pairLPMintTx,
				} {
					require.NotEqual(t, common.Hash{}, h, "not legit hash at index %d", k)
					rcpt, err := cl.TransactionReceipt(t.Ctx(), h)
					require.NoError(t, err, "didnt have receipt as expected")
					require.NotNil(t, rcpt, "receipt shouldnt been nil")
					// require.Equal(t, types.ReceiptStatusSuccessful, rcpt.Status,
					// 	"%dth tx reverted but shouldnt %v gas used %v total so far %d", k, h, rcpt.GasUsed, gasUsedSoFar,
					// )
					wait.ForReceiptOK(t.Ctx(), cl, h)
					gasUsedSoFar += rcpt.GasUsed
					cAddr := rcpt.ContractAddress
					switch h {
					case v2FactoryTx:
						require.Equal(t, uniswapV2FactoryAddr, cAddr, "contract factory address wasnt as expected")
					case memecoinTx:
						require.Equal(t, memecoinAddr, cAddr, "contract memecoin address wasnt as expected")
					case memecoinMintTx:
						totalSupply := ethereum.CallMsg{To: &memecoinAddr, Data: encodedTotalSupply}
						result, err := cl.CallContract(t.Ctx(), totalSupply, nil)
						require.NoError(t, err)
						unpacked, err := ERC20.ABI.Unpack("totalSupply", result)
						require.NoError(t, err)
						putIn := abi.ConvertType(unpacked[0], new(big.Int)).(*big.Int)
						require.Condition(
							t, func() bool { return fiveMillion.Cmp(putIn) == 0 },
							"expected five million meme coin total supply but got %v raw %v", putIn, unpacked,
						)
						t.Logf("total supply minted of memecoin %v", putIn)
					case simpleTransfer:
						// simple value transfer because v2 router failed
						// and didnt feel like adjusting everything
					case v2PairCreation:
						// pair creation didn't make a new addr, it was via factory call
					case wethCreationTx:
						require.Equal(t, cAddr, wethAddr)
					case wethWrapTxAlice:
						checkBalance := ethereum.CallMsg{
							From: dp.Addresses.Alice,
							To:   &wethAddr,
							Data: encodedBalanceOfAlice,
						}
						result, err := cl.CallContract(t.Ctx(), checkBalance, nil)
						require.NoError(t, err)
						unpacked, err := ERC20.ABI.Unpack("balanceOf", result)
						require.NoError(t, err)
						putIn := abi.ConvertType(unpacked[0], new(big.Int)).(*big.Int)
						require.Condition(
							t, func() bool { return twentyFiveE18.Cmp(putIn) == 0 },
							"balanceOf Weth for alice not as expected %v", putIn,
						)
					case pairLPMintTx:
						checkBalance := ethereum.CallMsg{
							From: dp.Addresses.Alice,
							To:   &v2PairAddr,
							Data: encodedBalanceOfAlice,
						}
						result, err := cl.CallContract(t.Ctx(), checkBalance, nil)
						require.NoError(t, err)
						unpacked, err := WETH.ABI.Unpack("balanceOf", result)
						require.NoError(t, err)
						putIn := abi.ConvertType(unpacked[0], new(big.Int)).(*big.Int)
						expected, _ := new(big.Int).SetString("11180339887498948481045", 10)
						require.Condition(
							t, func() bool { return expected.Cmp(putIn) == 0 },
							"balanceOf v2 pair for alice not as expected %v", putIn,
						)

					default:
						t.Fatalf("unhandled tx confirmation logic %v", h)
					}
				}
			}

		}

		miner.ActL1EndBlock(t)
	}

	size, err := dbSize(dbDir)
	require.NoError(t, err)
	avgCompDur, totalComp := ethPebble.AverageCompactionTime()

	log.Info(
		"db persisted to disk",
		"size", size,
		"place", dbDir,
		"average-compaction", avgCompDur,
		"total-compactions", totalComp,
	)

}
