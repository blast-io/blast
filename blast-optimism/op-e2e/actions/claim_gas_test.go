package actions

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func loadContract(t *testing.T) ([]byte, abi.ABI, abi.ABI) {
	cwd, _ := os.Getwd()
	t.Logf("Current working directory %v", cwd)
	cmd := exec.Command(
		"solc", "--combined-json", "bin,abi", "--no-cbor-metadata",
		"gas_tester.sol",
	)
	output := bytes.Buffer{}
	oops := bytes.Buffer{}
	cmd.Stdout = &output
	cmd.Stderr = &oops
	if err := cmd.Run(); err != nil {
		t.Fatalf("solc died %v because %v", err, oops.String())
	}
	type loadIt struct {
		Contracts struct {
			Wrapper struct {
				ABI      json.RawMessage `json:"abi"`
				Bytecode string          `json:"bin"`
			} `json:"gas_tester.sol:Worker"`
			IBlast struct {
				ABI json.RawMessage `json:"abi"`
			} `json:"gas_tester.sol:IBlast"`
		} `json:"contracts"`
	}
	var l loadIt
	// t.Logf("output of solc %v", string(output.Bytes()))
	if err := json.Unmarshal(output.Bytes(), &l); err != nil {
		t.Fatalf("couldn't unmarshal %v", err)
	}
	//	t.Logf("load it %v", l)
	if len(l.Contracts.Wrapper.Bytecode) == 0 {
		t.Fatalf("contract bytecode non existent")
	}
	abiHandle, err := abi.JSON(bytes.NewReader(l.Contracts.Wrapper.ABI))
	require.Nil(t, err, "couldn't load up abi %v", err)
	iblastABI, err := abi.JSON(bytes.NewReader(l.Contracts.IBlast.ABI))

	t.Logf("worker abi looks like %v", abiHandle)
	t.Logf("iblast abi looks like %v", iblastABI)

	return common.Hex2Bytes(l.Contracts.Wrapper.Bytecode), abiHandle, iblastABI
}

func TestGasTrackerClaim(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	sd, dp, miner, sequencer, seqEngine, _, _ := setupEIP4844Test(t, log)
	contractByteCode, contractABI, iBlastABI := loadContract(gt)

	cnstr, err := contractABI.Constructor.Inputs.Pack()
	require.Nil(t, err, "constructor cant die")
	payload := append(contractByteCode, cnstr...)
	l2Client := seqEngine.EthClient()
	sayBlock := func(ctx string) {
		blkNum, err := l2Client.BlockNumber(t.Ctx())
		require.NoError(t, err, "block num cant die")
		t.Logf("%v block number is %v", ctx, blkNum)
	}

	sequencer.ActL2PipelineFull(t)
	miner.ActEmptyBlock(t)

	amt, err := l2Client.BalanceAt(t.Ctx(), dp.Addresses.Alice, nil)
	require.Nil(t, err, "l2client cant die")
	t.Logf("amt alice %v", amt)

	signer := types.LatestSigner(sd.L2Cfg.Config)
	n, err := l2Client.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
	require.NoError(t, err, "couldnt get pending nonce")
	newContractAddr := crypto.CreateAddress(dp.Addresses.Alice, n)

	tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
		ChainID:   sd.L2Cfg.Config.ChainID,
		Nonce:     n,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
		Gas:       2_000_000,
		To:        nil,
		Value:     common.Big0,
		Data:      payload,
	})

	sayBlock("before contract creation ")
	require.NoError(gt, l2Client.SendTransaction(t.Ctx(), tx))
	sequencer.ActL2StartBlock(t)
	seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
	sequencer.ActL2EndBlock(t)
	sayBlock("cranked the l2 chain ")

	runtimeCode, err := l2Client.CodeAt(t.Ctx(), newContractAddr, nil)
	require.NoError(t, err, "missing runtime contract code after deployment")
	require.Greater(t, len(runtimeCode), 0, "Zero length contract code")

	readGas := func() (*big.Int, *big.Int, *big.Int, uint8) {
		readGasEncoded, err := iBlastABI.Pack("readGasParams", newContractAddr)
		require.NoError(t, err, "not possible couldnt pack read gas params")
		readGasResult, err := l2Client.CallContract(t.Ctx(), ethereum.CallMsg{
			To:   &predeploys.BlastAddr,
			Data: readGasEncoded,
		}, nil)
		require.NoError(t, err, "call contract died")
		unpacked, err := iBlastABI.Unpack("readGasParams", readGasResult)
		require.NoError(t, err, "unpack died")
		etherSeconds, etherBalance, lastUpdated, mode :=
			unpacked[0].(*big.Int), unpacked[1].(*big.Int), unpacked[2].(*big.Int), unpacked[3].(uint8)
		return etherSeconds, etherBalance, lastUpdated, mode
	}

	wasteGas := func() {
		wasteEncoded, err := contractABI.Pack("burn_gas")
		require.NoError(t, err, "couldnt pack burn_gas")
		n, err := l2Client.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
		require.NoError(t, err, "couldnt get pending nonce")
		tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
			ChainID:   sd.L2Cfg.Config.ChainID,
			Nonce:     n,
			GasTipCap: big.NewInt(2 * params.GWei),
			GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
			Gas:       50_000,
			To:        &newContractAddr,
			Value:     common.Big0,
			Data:      wasteEncoded,
		})
		require.NoError(gt, l2Client.SendTransaction(t.Ctx(), tx))
		sequencer.ActL2StartBlock(t)
		seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
		sequencer.ActL2EndBlock(t)
		// rcpt, err := wait.ForReceiptOK(t.Ctx(), l2Client, tx.Hash())
		// rcpt, err := l2Client.TransactionReceipt(t.Ctx(), tx.Hash())
		// require.NoError(t, err, "couldnt get receipt")
		// return rcpt
	}

	sayBlock("About to waste Gas call")
	wasteGas()
	sayBlock("did waste gas")
	etherSeconds, etherBalance, lastUpdated, mode := readGas()
	// t.Logf("What values %v %v %v %v", etherSeconds, etherBalance, lastUpdated, mode)
	etherSecondsShouldBe, etherBalanceShouldBe, modeExpected :=
		big.NewInt(85306000042653), big.NewInt(143064000071532), uint8(1)

	currentBlock, err := l2Client.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)

	lastUpdatedExpected := big.NewInt(int64(currentBlock.Time()))

	if etherSeconds.Cmp(etherSecondsShouldBe) != 0 {
		t.Fatalf("incorrect etherseconds %v %v", etherSeconds, etherSecondsShouldBe)
	}

	if etherBalance.Cmp(etherBalanceShouldBe) != 0 {
		t.Fatalf("incorrect etherBalance %v %v", etherBalance, etherBalanceShouldBe)
	}

	if lastUpdated.Cmp(lastUpdatedExpected) != 0 {
		t.Fatalf("incorrect last timestamp updated %v %v", lastUpdated, lastUpdatedExpected)
	}

	if mode != modeExpected {
		t.Fatalf("incorrect mode %v %v", mode, modeExpected)
	}
}

func readConfigAndHeaderFromDB(t *testing.T, pathDB string) (*core.Genesis, *params.ChainConfig, *types.Header, time.Time) {
	nodeCfg := &node.Config{Name: "geth", DataDir: pathDB, P2P: p2p.Config{NoDiscovery: true, NoDial: true}}
	n, err := node.New(nodeCfg)
	require.NoError(t, err, "loading up chain-db died")
	dbHandle, err := n.OpenDatabase("chaindata", 0, 0, "", true, 0)
	require.NoError(t, err, "loading up database handle to chaindata died")
	gen, err := core.ReadGenesis(dbHandle)
	require.NoError(t, err, "reading genesis in chaindata died")
	chainCfg, err := core.LoadChainConfig(dbHandle, gen)
	require.NoError(t, err, "reading chain config in chaindata died")
	currentHdr := rawdb.ReadHeadHeader(dbHandle)
	require.NotNil(t, currentHdr, "header off rawdb is nil - can't proceed, is db corrupted")
	require.Nil(t, n.Close(), "couldnt close node handles")
	return gen, chainCfg, currentHdr, time.Unix(int64(currentHdr.Time), 0)
}

//go:embed blast-mainnet-rollup.json
var blastMainnetRollup []byte

func TestBlastE2EMainnet(gt *testing.T) {
	t := NewDefaultTesting(gt)
	lg := testlog.Logger(t, log.LvlDebug)

	// keep them outside the repo because gopls will die trying to analyze it
	const (
		// ethMainnetDB          = "/Volumes/eth-chaindata/mainnet"
		// ethMainnetFreezerDB   = "/Volumes/eth-chaindata/mainnet-ancient"
		blastMainnetDB        = "/Volumes/eth-chaindata/blast-mainnet-chaindata"
		blastMainnetFreezerDB = "/Volumes/eth-chaindata/blast-mainnet-chaindata/geth/chaindata/ancient"
	)

	// Note you cannot trust the state root from ethereum mainnet node loaded up because the
	// StateAccount changed in structure and hence its state root will always be computed differently
	// genEthereumMainnet, chainConfigEthereumMainnet, hdrEthMain, hdrTimeEthMain := readConfigAndHeaderFromDB(gt, ethMainnetDB)
	genBlastMainnet, chainConfigBlastMainnet, hdrBlastMain, hdrTimeBlastMain := readConfigAndHeaderFromDB(gt, blastMainnetDB)

	lg.Info(
		"loaded up existing mainnet dbs",
		// "eth-mainnet-head", hdrEthMain.Number,
		// "eth-mainnet-ts", hdrTimeEthMain,
		// "eth-chain-config in db", chainConfigEthereumMainnet,
		"blast-mainnet-head", hdrBlastMain.Number,
		"blast-mainnet-ts", hdrTimeBlastMain,
		"blast-chain-config in db", chainConfigBlastMainnet,
	)

	var mainnetRollup rollup.Config
	require.Nil(t, json.Unmarshal(blastMainnetRollup, &mainnetRollup), "couldnt unmarshal blast mainnet rollup")

	// lets say the next block it cranks over
	ecotoneTS := hdrTimeBlastMain.Add(time.Second * time.Duration(mainnetRollup.BlockTime))
	ecotoneTime := uint64(ecotoneTS.Unix())

	mainnetRollup.DeltaTime = &ecotoneTime
	mainnetRollup.EcotoneTime = &ecotoneTime

	cwd, _ := os.Getwd()
	root, _ := config.FindMonorepoRoot(cwd)
	deployConfigPath := filepath.Join(root, "packages", "contracts-bedrock", "deploy-config", "blast-mainnet.json")
	deployConfig, err := genesis.NewDeployConfig(deployConfigPath)
	require.NoError(t, err, "deploy config died")
	lg.Info("deploy config loaded", "config", deployConfig)

	sd := &e2eutils.SetupData{
		// L1Cfg:     genEthereumMainnet,
		L2Cfg:     genBlastMainnet,
		RollupCfg: &mainnetRollup,
	}

	jwtPath := e2eutils.WriteDefaultJWT(t)
	seqEngine := NewL2Engine(
		t, lg, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath,
		func(ethCfg *ethconfig.Config, nodeCfg *node.Config) error {
			nodeCfg.Name = "geth"
			nodeCfg.WSPort = 2002
			nodeCfg.DataDir = blastMainnetDB
			nodeCfg.P2P = p2p.Config{NoDiscovery: true, NoDial: true}

			ethCfg.NetworkId = sd.L2Cfg.Config.ChainID.Uint64()
			ethCfg.Genesis = genBlastMainnet
			ethCfg.RollupDisableTxPoolGossip = true
			ethCfg.OverrideCancun = &ecotoneTime
			ethCfg.OverrideOptimismEcotone = &ecotoneTime
			return nil
		},
	)
	t.Cleanup(func() {
		// Must do this otherwise it will write the trie to disk and the DB will be useless
		seqEngine.l2Chain.SetHead(hdrBlastMain.Number.Uint64())
	})

	// seqEngine.l2Chain.SetHead(3349547)

	l2Cl, err := sources.NewEngineClient(seqEngine.RPCClient(), lg, nil, sources.EngineClientDefaultConfig(sd.RollupCfg))
	require.NoError(t, err, "new engine client died")
	sequencer := NewL2Sequencer(t, lg, &mockL1Fetcher{
		lg: lg,
	}, nil, l2Cl, sd.RollupCfg, sd.L1Cfg.Config, 0)
	lg.Info(
		"finished loading up nodes, ecotone time picked for next block",
		"blast mainnet header", hdrTimeBlastMain,
		"ecotone time will be", ecotoneTS,
	)

	// hack - lying with the mock origin - TODO get eth mainnet header at this moment?
	// Or load up the block info contract call and then find it with cast, yea that's the right way
	// yes that will take time, first lets get this way maybe working
	sequencer.mockL1OriginSelector.originOverride = eth.InfoToL1BlockRef(eth.HeaderBlockInfo(hdrBlastMain))
	// yes a huge hack but let's see what we can do
	sequencer.engine.SetUnsafeHead(eth.L2BlockRef{
		Hash:           hdrBlastMain.Hash(),
		Number:         hdrBlastMain.Number.Uint64(),
		ParentHash:     hdrBlastMain.ParentHash,
		Time:           hdrBlastMain.Time,
		L1Origin:       eth.BlockID{Hash: hdrBlastMain.Hash(), Number: hdrBlastMain.Number.Uint64()},
		SequenceNumber: 0,
	})

	sequencer.ActL2StartBlock(t)
	sequencer.ActL2EndBlock(t)
	blk := seqEngine.l2Chain.GetBlockByNumber(hdrBlastMain.Number.Uint64() + 1)
	require.Equal(t, len(blk.Transactions()), 7, "should be 7 ecotone upgrade txs")
}

type mockL1Fetcher struct {
	derive.L1Fetcher
	lg                log.Logger
	realMainnetHeader *types.Header
}

type mockBlock struct {
	eth.BlockInfo
	realMainnetHeader *types.Header
}

// better to do it with this real deserialized ethereum mainnet header
// and hold onto the params to give them for these overriden methods
// but that can be done for extending with a syncing test
func newMockBlock(mainnetHeader *types.Header) *mockBlock {
	return &mockBlock{}
}

func (m *mockBlock) Time() uint64 {
	return 0
}

func (m *mockBlock) NumberU64() uint64 {
	return 0
}

func (m *mockBlock) BaseFee() *big.Int {
	return big.NewInt(params.InitialBaseFee)
}

func (m *mockBlock) Hash() common.Hash {
	return common.Hash{}
}

func (m *mockBlock) ParentBeaconRoot() *common.Hash {
	return &common.Hash{}
}

func (m *mockBlock) MixDigest() common.Hash {
	return common.Hash{}
}

func (m *mockL1Fetcher) InfoByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, error) {
	m.lg.Info("mock l1 fetcher queried", "hash", hash)
	return newMockBlock(m.realMainnetHeader), nil
}
