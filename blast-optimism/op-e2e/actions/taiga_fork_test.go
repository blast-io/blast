package actions

import (
	"blast/blockchain"
	"encoding/json"
	"math/big"
	"testing"

	batcherFlags "github.com/ethereum-optimism/optimism/op-batcher/flags"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/stretchr/testify/require"
)

func TestTaigaSequencer(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	sd, dp, miner, sequencer, seqEngine, verifier, _ := setupEIP4844Test(t, log)

	batcher := setupBatcher(t, log, sd, dp, miner, sequencer, seqEngine, batcherFlags.BlobsType)
	_ = verifier
	_ = batcher
}

func TestTaigaActivatesSameTimeAsPrague(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	var howManyL1BlocksTillPrague uint64 = 4

	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	atGen := hexutil.Uint64(0)
	whenCancun := hexutil.Uint64(dp.DeployConfig.L1BlockTime * 1)
	whenEcotone := hexutil.Uint64(dp.DeployConfig.L1BlockTime * 2)
	whenPrague := hexutil.Uint64(dp.DeployConfig.L1BlockTime * howManyL1BlocksTillPrague)

	dp.DeployConfig.L1CancunTimeOffset = &whenCancun
	dp.DeployConfig.L1PragueTimeOffset = &whenPrague
	dp.DeployConfig.L2GenesisRegolithTimeOffset = &atGen
	dp.DeployConfig.L2GenesisDeltaTimeOffset = &atGen
	dp.DeployConfig.L2GenesisEcotoneTimeOffset = &whenEcotone
	dp.DeployConfig.L2GenesisTaigaTimeOffset = &whenPrague

	l1Gen, deployConf, l1Deployments := e2eutils.SetupPart1(t, dp, defaultAlloc)

	require.NoError(t, buildPlugin(log, gethL1PluginOpts))
	plugin, _, err := loadPlugin(log, gethL1PluginOpts)
	require.NoError(t, err)
	serialized, err := json.Marshal(l1Gen)
	require.NoError(t, err)

	newChainOrErr := plugin.NewChain(&blockchain.NewChainStartingArgs{
		SerializedGenesis: serialized,
	})

	log.Info("show result newChain", "result", newChainOrErr)
	//	require.NoError(t, newChainOrErr.Err, "new-chain still died")
	var deserialized *types.Header
	require.NoError(t, json.Unmarshal(newChainOrErr.SerializedHeader, &deserialized))
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
	_, _, sequencer, seqEngine, verifier, _, batcher := setupReorgTestActorsWithL1Plugin(t, dp, sd, log, miner)
	l1Client, cl := miner.EthClient(), seqEngine.EthClient()
	signer := types.LatestSigner(sd.L2Cfg.Config)

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

	miner.ActEmptyBlock()

	currentBlock, err = l1Client.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)
	l1Head = currentBlock.Header()

	require.True(t, sd.L1Cfg.Config.IsPrague(l1Head.Number, l1Head.Time), "Prague active")

	aliceTx := func() common.Hash {
		n, err := cl.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
		require.NoError(t, err)
		current, err := l1Client.BlockByNumber(t.Ctx(), nil)
		require.NoError(t, err)

		tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
			ChainID:   sd.L2Cfg.Config.ChainID,
			Nonce:     n,
			GasTipCap: big.NewInt(2 * params.GWei),
			GasFeeCap: new(big.Int).Add(current.Header().BaseFee, big.NewInt(2*params.GWei)),
			Gas:       params.TxGas,
			To:        &dp.Addresses.Bob,
			Value:     e2eutils.Ether(2),
		})
		require.NoError(gt, cl.SendTransaction(t.Ctx(), tx))
		return tx.Hash()
	}

	makeL2BlockWithAliceTx := func() common.Hash {
		hsh := aliceTx()
		sequencer.ActL2StartBlock(t)
		seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t) // include a test tx from alice
		sequencer.ActL2EndBlock(t)
		return hsh
	}

	origin, err := l1Client.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)

	for sequencer.SyncStatus().UnsafeL2.Time+sd.RollupCfg.BlockTime < origin.Header().Time {
		makeL2BlockWithAliceTx()
		batcher.ActL2BatchBuffer(t)
	}

	makeL2BlockWithAliceTx()
	// batcher.ActL2BatchBuffer(t)
	// batcher.ActL2ChannelClose(t)
	// batcher.ActL2BatchSubmit(t)
	// miner.ActL1StartBlock(defaultRollupTestParams.L1BlockTime)(t)
	// miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	// miner.ActL1EndBlock(t)

}
