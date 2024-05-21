package actions

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"math/big"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/tracers"
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

func readConfigAndHeaderFromDB(t *testing.T, path string) (*core.Genesis, *params.ChainConfig, *types.Header, time.Time) {
	nodeCfg := &node.Config{
		Name:    "geth",
		DataDir: path,
		P2P:     p2p.Config{NoDiscovery: true, NoDial: true},
	}

	n, err := node.New(nodeCfg)
	require.NoError(t, err, "loading up chain-db died")
	dbHandle, err := n.OpenDatabase("chaindata", 0, 0, "", true)
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

	const (
		ethMainnetDB   = "/Volumes/eth-chaindata/mainnet"
		blastMainnetDB = "/Volumes/eth-chaindata/blast-mainnet-chaindata"
	)

	// keep it outside the repo because gopls will die trying to analyze it
	genEthereumMainnet, chainConfigEthereumMainnet, hdrEthMain, hdrTimeEthMain := readConfigAndHeaderFromDB(gt, ethMainnetDB)
	genBlastMainnet, chainConfigBlastMainnet, hdrBlastMain, hdrTimeBlastMain := readConfigAndHeaderFromDB(gt, blastMainnetDB)

	_ = hdrEthMain
	_ = hdrBlastMain
	_ = hdrTimeEthMain

	var mainnetRollup rollup.Config
	require.Nil(t, json.Unmarshal(blastMainnetRollup, &mainnetRollup), "couldnt unmarshal blast mainnet rollup")

	nodeCfgEthGeth := &node.Config{
		Name:        "geth",
		WSHost:      "127.0.0.1",
		WSPort:      2001,
		WSModules:   []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal"},
		HTTPModules: []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal"},
		DataDir:     ethMainnetDB,
		P2P:         p2p.Config{NoDiscovery: true, NoDial: true},
	}

	nodeCfgBlastGeth := &node.Config{
		Name:        "geth",
		WSHost:      "127.0.0.1",
		WSPort:      2002,
		WSModules:   []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal"},
		HTTPModules: []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal"},
		DataDir:     blastMainnetDB,
		P2P:         p2p.Config{NoDiscovery: true, NoDial: true},
	}

	nodeEth, err := node.New(nodeCfgEthGeth)
	require.NoError(t, err, "loading up eth geth died")

	nodeBlast, err := node.New(nodeCfgBlastGeth)
	require.NoError(t, err, "loading up blast geth died")

	// lets say the next block it cranks over
	ecotoneTime := uint64(hdrTimeBlastMain.Add(time.Second * time.Duration(mainnetRollup.BlockTime)).Unix())
	ethCfgL2Geth := &ethconfig.Config{
		NetworkId:                 genBlastMainnet.Config.ChainID.Uint64(),
		Genesis:                   genBlastMainnet,
		RollupDisableTxPoolGossip: true,
		OverrideOptimismEcotone:   &ecotoneTime,
	}
	ethCfgL1Geth := &ethconfig.Config{
		NetworkId:                 genEthereumMainnet.Config.ChainID.Uint64(),
		Genesis:                   genEthereumMainnet,
		RollupDisableTxPoolGossip: true,
	}

	backendL1Geth, err := eth.New(nodeEth, ethCfgL1Geth)
	require.NoError(t, err)
	backendL2Blast, err := eth.New(nodeBlast, ethCfgL2Geth)
	require.NoError(t, err)

	backendL1Geth.Merger().FinalizePoS()
	backendL2Blast.Merger().FinalizePoS()

	nodeEth.RegisterAPIs(tracers.APIs(backendL1Geth.APIBackend))
	nodeBlast.RegisterAPIs(tracers.APIs(backendL2Blast.APIBackend))

	require.NoError(t, nodeEth.Start(), "failed to start L1 geth node")
	require.NoError(t, nodeBlast.Start(), "failed to start L2 blast node")

	lg.Info("finished loading up nodes")

	l1MinerGeth := &L1Miner{
		L1Replica: L1Replica{
			log:        lg,
			node:       nodeEth,
			eth:        backendL1Geth,
			l1Chain:    backendL1Geth.BlockChain(),
			l1Database: backendL1Geth.ChainDb(),
			l1Cfg:      genEthereumMainnet,
			l1Signer:   types.LatestSigner(chainConfigEthereumMainnet),
			failL1RPC:  nil,
		}, blobStore: e2eutils.NewBlobStore(),
	}

	l2MinerBlast := &L1Miner{
		L1Replica: L1Replica{
			log:        lg,
			node:       nodeBlast,
			eth:        backendL1Geth,
			l1Chain:    backendL1Geth.BlockChain(),
			l1Database: backendL1Geth.ChainDb(),
			l1Cfg:      genBlastMainnet,
			l1Signer:   types.LatestSigner(chainConfigBlastMainnet),
			failL1RPC:  nil,
		}, blobStore: e2eutils.NewBlobStore(),
	}

	_ = l1MinerGeth
	_ = l2MinerBlast
}
