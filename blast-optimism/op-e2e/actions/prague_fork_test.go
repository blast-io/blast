package actions

import (
	"blast/blockchain"
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/transactions"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/stretchr/testify/require"
)

func setupSequencerTestWithL1Plugin(
	t Testing, sd *e2eutils.SetupData, log log.Logger, miner ChainRunner,
) (*L2Engine, *L2Sequencer) {
	jwtPath := e2eutils.WriteDefaultJWT(t)
	baseClient := client.NewBaseRPCClient(miner.EthClient().Client())
	l1Client, err := sources.NewL1Client(
		baseClient, log, nil,
		sources.L1ClientDefaultConfig(sd.RollupCfg, false, sources.RPCKindStandard),
	)
	require.NoError(t, err)
	engine := NewL2Engine(t, log, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath)
	l2Cl, err := sources.NewEngineClient(engine.RPCClient(), log, nil, sources.EngineClientDefaultConfig(sd.RollupCfg))
	require.NoError(t, err)
	sequencer := NewL2Sequencer(t, log, l1Client, miner.BlobStore(), l2Cl, sd.RollupCfg, sd.L1Cfg.Config, 0)
	return engine, sequencer
}

func setupReorgTestActorsWithL1Plugin(
	t Testing, dp *e2eutils.DeployParams,
	sd *e2eutils.SetupData, log log.Logger, miner ChainRunner,
) (*e2eutils.SetupData, *e2eutils.DeployParams, *L2Sequencer, *L2Engine, *L2Verifier, *L2Engine, *L2Batcher) {
	seqEngine, sequencer := setupSequencerTestWithL1Plugin(t, sd, log, miner)
	miner.ActL1SetFeeRecipient(common.Address{'A'})
	sequencer.ActL2PipelineFull(t)

	baseClient := client.NewBaseRPCClient(miner.EthClient().Client())
	l1Client, err := sources.NewL1Client(
		baseClient, log, nil,
		sources.L1ClientDefaultConfig(sd.RollupCfg, false, sources.RPCKindStandard),
	)
	require.NoError(t, err)
	verifEngine, verifier := setupVerifier(t, sd, log, l1Client, miner.BlobStore(), &sync.Config{})
	rollupSeqCl := sequencer.RollupClient()
	batcher := NewL2Batcher(
		log, sd.RollupCfg, DefaultBatcherCfg(dp),
		rollupSeqCl, miner.EthClient(), seqEngine.EthClient(),
		seqEngine.EngineClient(t, sd.RollupCfg),
	)

	return sd, dp, sequencer, seqEngine, verifier, verifEngine, batcher
}

func TestPragueAgainstRealL1Geth(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	require.NoError(t, buildPlugin(log, gethL1PluginOpts))
	mnemonicCfg := e2eutils.DefaultMnemonicConfig
	secrets, err := mnemonicCfg.Secrets()
	require.NoError(t, err)
	addresses := secrets.Addresses()
	plugin, _, err := loadPlugin(log, gethL1PluginOpts)
	require.NoError(t, err)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	offset := hexutil.Uint64(24)
	offsetMore := hexutil.Uint64(48)
	dp.DeployConfig.L1CancunTimeOffset = &offset
	dp.DeployConfig.L1PragueTimeOffset = &offsetMore

	l1Gen, deployConf, l1Deployments := e2eutils.SetupPart1(t, dp, defaultAlloc)
	_ = addresses
	serialized, err := json.Marshal(l1Gen)
	require.NoError(t, err)

	result := plugin.NewChain(&blockchain.NewChainStartingArgs{
		SerializedGenesis: serialized,
		//		ExtraAllocs:       extraAllocs,
	})

	require.NoError(t, err)

	var deserialized *types.Header

	require.NoError(t, json.Unmarshal(result.SerializedHeader, &deserialized))
	l1Blk := types.NewBlock(
		deserialized, nil, nil, nil, trie.NewStackTrie(nil),
	)

	sd := e2eutils.SetupPart2(t, dp, defaultAlloc, deployConf, l1Blk, l1Gen, l1Deployments)
	l1GenHash := sd.L1Cfg.ToBlock().Hash()
	l2GenHash := sd.L2Cfg.ToBlock().Hash()

	log.Info(
		"genesis hashes",
		"l1-hsh", l1GenHash.Hex(),
		"l2-hsh", l2GenHash.Hex(),
		"doing-state-hash", sd.L1Cfg.StateHash == nil,
	)

	miner := NewPluginBackedMiner(t, log, plugin, defaultRollupTestParams.L1BlockTime)
	_, _, sequencer, _, verifier, _, batcher := setupReorgTestActorsWithL1Plugin(t, dp, sd, log, miner)
	l1Client := miner.EthClient()
	currentBlock, err := l1Client.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)
	l1Head := currentBlock.Header()

	require.False(t, sd.L1Cfg.Config.IsCancun(l1Head.Number, l1Head.Time), "Cancun not active yet")
	require.False(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time), "Prague not active yet")

	// start op-nodes
	sequencer.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	// build empty L1 blocks, crossing the fork boundary
	miner.ActL1SetFeeRecipient(common.Address{'A', 0})
	miner.ActEmptyBlock()
	miner.ActEmptyBlock() // Cancun activates here
	miner.ActEmptyBlock()
	currentBlock, err = l1Client.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)
	l1Head = currentBlock.Header()
	require.True(t, sd.L1Cfg.Config.IsCancun(l1Head.Number, l1Head.Time), "Cancun active")
	require.False(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time), "Prague not active yet")

	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)
	miner.ActL1StartBlock(12)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)
	miner.ActL1EndBlock(t)

	// blk2 := miner.ActEmptyBlock()
	// blk3 := miner.ActEmptyBlock()

	_ = batcher

	// log.Info("made few empty blocks", "blk1", blk1.Number(), "blk2", blk2.Number(), "blk3", blk3.Number())
	time.Sleep(time.Second * 1)
	// TODO control when turn on cancun, when turn on prague
	// require.Nil(t, l1Head.ExcessBlobGas, "Cancun blob gas not in header")
	// require.Nil(t, l1Head.RequestsHash, "Prague RequestHash not in header")

}

func TestPragueL1ForkAfterGenesis(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	offset := hexutil.Uint64(24)
	offsetMore := hexutil.Uint64(48)
	dp.DeployConfig.L1CancunTimeOffset = &offset
	dp.DeployConfig.L1PragueTimeOffset = &offsetMore
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	_, _, miner, sequencer, _, verifier, _, batcher := setupReorgTestActors(t, dp, sd, log)

	l1Head := miner.l1Chain.CurrentBlock()
	require.False(t, sd.L1Cfg.Config.IsCancun(l1Head.Number, l1Head.Time), "Cancun not active yet")
	require.False(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time), "Prague not active yet")

	require.Nil(t, l1Head.ExcessBlobGas, "Cancun blob gas not in header")
	require.Nil(t, l1Head.RequestsHash, "Prague RequestHash not in header")

	// start op-nodes
	sequencer.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	// build empty L1 blocks, crossing the fork boundary
	miner.ActL1SetFeeRecipient(common.Address{'A', 0})
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t) // Cancun activates here
	miner.ActEmptyBlock(t)
	// verify Cancun is active
	l1Head = miner.l1Chain.CurrentBlock()
	require.True(t, sd.L1Cfg.Config.IsCancun(l1Head.Number, l1Head.Time), "Cancun active")
	require.NotNil(t, l1Head.ExcessBlobGas, "Cancun blob gas in header")

	// build L2 chain up to and including L2 blocks referencing Cancun L1 blocks
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)
	miner.ActL1StartBlock(12)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync verifier
	verifier.ActL1HeadSignal(t)
	verifier.ActL2PipelineFull(t)
	// verify verifier accepted Cancun L1 inputs
	require.Equal(t, l1Head.Hash(), verifier.SyncStatus().SafeL2.L1Origin.Hash, "verifier synced L1 chain that includes Cancun headers")
	require.Equal(t, sequencer.SyncStatus().UnsafeL2, verifier.SyncStatus().UnsafeL2, "verifier and sequencer agree")

	// Now continue to Prague , two more blocks
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t) // Prague activates here
	miner.ActEmptyBlock(t)

	//verify Prague is active
	l1Head = miner.l1Chain.CurrentBlock()
	require.True(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time), "Prague active")
	require.NotNil(t, l1Head.RequestsHash, "Prague request hash in header")

	// continue to build the l2 block referencing the Prague L1 Blocks
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)
	miner.ActL1StartBlock(12)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync verifier
	verifier.ActL1HeadSignal(t)
	verifier.ActL2PipelineFull(t)
	// verify verifier accepted Prague L1 inputs
	require.Equal(t, l1Head.Hash(), verifier.SyncStatus().SafeL2.L1Origin.Hash, "verifier synced L1 chain that includes Prague headers")
	require.Equal(t, sequencer.SyncStatus().UnsafeL2, verifier.SyncStatus().UnsafeL2, "verifier and sequencer agree")
}

func aliceSimpleSetCodeTx(t Testing, dp *e2eutils.DeployParams, chainID uint64, to common.Address, nonceSetCode, nonceAuth uint64) *types.Transaction {
	txData := transactions.CreateSetCodeTx(dp.Secrets.Bob, chainID, to)
	txData.Nonce = nonceSetCode
	txData.AuthList[0].Nonce = nonceAuth
	signer := types.NewPragueSigner(txData.ChainID.ToBig())
	tx, err := types.SignNewTx(dp.Secrets.Alice, signer, txData)
	require.NoError(t, err, "must sign tx")
	return tx
}

var junkTo = common.Address{0x02, 0x08, 0x04}

// TestDencunBlobTxRPC tries to send a SetCode tx to the L2 engine via RPC, it should not be accepted.
func TestPragueSetCodeTxRPC(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	offset := hexutil.Uint64(0)
	dp.DeployConfig.L2GenesisRegolithTimeOffset = &offset
	dp.DeployConfig.L2GenesisCanyonTimeOffset = &offset
	dp.DeployConfig.L2GenesisDeltaTimeOffset = &offset
	dp.DeployConfig.L2GenesisEcotoneTimeOffset = &offset

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	engine := newEngine(t, sd, log)
	cl := engine.EthClient()
	tx := aliceSimpleSetCodeTx(t, dp, dp.DeployConfig.L2ChainID, junkTo, 0, 0)
	err := cl.SendTransaction(context.Background(), tx)
	require.ErrorContains(t, err, "transaction type not supported")
}

// TestPragueSetCodeTxInTxPool tries to insert a setcode tx directly into the tx pool, it should not be accepted.
func TestPragueSetCodeTxInTxPool(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	offset := hexutil.Uint64(0)
	dp.DeployConfig.L2GenesisRegolithTimeOffset = &offset
	dp.DeployConfig.L2GenesisCanyonTimeOffset = &offset
	dp.DeployConfig.L2GenesisDeltaTimeOffset = &offset
	dp.DeployConfig.L2GenesisEcotoneTimeOffset = &offset

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	engine := newEngine(t, sd, log)
	tx := aliceSimpleSetCodeTx(t, dp, dp.DeployConfig.L2ChainID, junkTo, 0, 0)
	errs := engine.eth.TxPool().Add([]*types.Transaction{tx}, true, true)
	require.ErrorContains(t, errs[0], "transaction type not supported")
}

// TestPragueSetCodeTxInclusion tries to send a SetCode tx to the L2 engine, it should not be accepted.
func TestPragueSetCodeTxInclusion(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	offset := hexutil.Uint64(0)
	dp.DeployConfig.L2GenesisRegolithTimeOffset = &offset
	dp.DeployConfig.L2GenesisCanyonTimeOffset = &offset
	dp.DeployConfig.L2GenesisDeltaTimeOffset = &offset
	dp.DeployConfig.L2GenesisEcotoneTimeOffset = &offset

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)

	_, engine, sequencer := setupSequencerTest(t, sd, log)
	sequencer.ActL2PipelineFull(t)

	tx := aliceSimpleSetCodeTx(t, dp, dp.DeployConfig.L2ChainID, junkTo, 0, 0)

	sequencer.ActL2StartBlock(t)
	err := engine.engineApi.IncludeTx(tx, dp.Addresses.Alice)
	require.ErrorContains(t, err, "invalid L2 block (tx 1): failed to apply transaction to L2 block (tx 1): transaction type not supported")
}

// go test -v -count=1 -run PragueL2HandlesSetCodeOnL1ForkAfterGenesis -timeout 20m ./actions/...

func TestPragueL2HandlesSetCodeOnL1ForkAfterGenesis(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	offset := hexutil.Uint64(24)
	offsetMore := hexutil.Uint64(48)
	dp.DeployConfig.L1CancunTimeOffset = &offset
	dp.DeployConfig.L1PragueTimeOffset = &offsetMore
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlTrace)
	_, _, miner, sequencer, _, verifier, _, batcher := setupReorgTestActors(t, dp, sd, log)

	l1Head := miner.l1Chain.CurrentBlock()
	require.False(t, sd.L1Cfg.Config.IsCancun(l1Head.Number, l1Head.Time), "Cancun not active yet")
	require.False(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time), "Prague not active yet")

	require.Nil(t, l1Head.ExcessBlobGas, "Cancun blob gas not in header")
	require.Nil(t, l1Head.RequestsHash, "Prague RequestHash not in header")

	// start op-nodes
	sequencer.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	// build empty L1 blocks, crossing the fork boundary
	miner.ActL1SetFeeRecipient(common.Address{'A', 0})
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t) // Cancun activates here
	miner.ActEmptyBlock(t)
	// verify Cancun is active
	l1Head = miner.l1Chain.CurrentBlock()
	require.True(t, sd.L1Cfg.Config.IsCancun(l1Head.Number, l1Head.Time), "Cancun active")
	require.NotNil(t, l1Head.ExcessBlobGas, "Cancun blob gas in header")

	// build L2 chain up to and including L2 blocks referencing Cancun L1 blocks
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)
	miner.ActL1StartBlock(12)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync verifier
	verifier.ActL1HeadSignal(t)
	verifier.ActL2PipelineFull(t)
	// verify verifier accepted Cancun L1 inputs
	require.Equal(t, l1Head.Hash(), verifier.SyncStatus().SafeL2.L1Origin.Hash, "verifier synced L1 chain that includes Cancun headers")
	require.Equal(t, sequencer.SyncStatus().UnsafeL2, verifier.SyncStatus().UnsafeL2, "verifier and sequencer agree")

	// Now continue to Prague , two more blocks
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t) // Prague activates here
	miner.ActEmptyBlock(t)

	//verify Prague is active
	l1Head = miner.l1Chain.CurrentBlock()
	require.True(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time), "Prague active")
	require.NotNil(t, l1Head.RequestsHash, "Prague request hash in header")

	// now someone does a setcode tx on l1 - , will opnode be able to process without dying
	tx := aliceSimpleSetCodeTx(t, dp, dp.DeployConfig.L1ChainID, junkTo, 0, 0)
	batchInboxAddrTx := aliceSimpleSetCodeTx(t, dp, dp.DeployConfig.L1ChainID, dp.DeployConfig.BatchInboxAddress, 1, 1)

	miner.ActL1StartBlock(12)(t)
	miner.IncludeTx(t, tx)
	miner.IncludeTx(t, batchInboxAddrTx)
	log.Info("included two txs we ought to handle", "l1-set-code", tx.Hash().Hex(), "set-code-to-batcher", batchInboxAddrTx.Hash().Hex())
	miner.ActL1EndBlock(t)

	rcptSetCode, err := miner.EthClient().TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err, "simple-set-code-didnt land")
	require.NotNil(t, rcptSetCode, "dont have simple set code tx on l1")

	rcptSetCode, err = miner.EthClient().TransactionReceipt(context.Background(), batchInboxAddrTx.Hash())
	require.NoError(t, err, "batch inbox addr land")
	require.NotNil(t, rcptSetCode, "dont have bad setcode set to batcher inbox")

	// continue to build the l2 block referencing the Prague L1 Blocks
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)
	l1Head = miner.l1Chain.CurrentBlock()
	miner.ActL1StartBlock(12)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync verifier
	verifier.ActL1HeadSignal(t)
	verifier.ActL2PipelineFull(t)
	// verify verifier accepted Prague L1 inputs
	require.Equal(t, l1Head.Hash(), verifier.SyncStatus().SafeL2.L1Origin.Hash, "verifier synced L1 chain that includes Prague headers")
	require.Equal(t, sequencer.SyncStatus().UnsafeL2, verifier.SyncStatus().UnsafeL2, "verifier and sequencer agree")

}
