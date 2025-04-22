package actions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"

	batcherFlags "github.com/ethereum-optimism/optimism/op-batcher/flags"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-e2e/files"
	"github.com/ethereum-optimism/optimism/op-node/cmd/batch_decoder/fetch"
	"github.com/ethereum-optimism/optimism/op-node/cmd/batch_decoder/reassemble"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
)

// TestL2BatcherBatchType run each batcher-related test case in singular batch mode and span batch mode.
func TestL2BatcherBatchType(t *testing.T) {
	tests := []struct {
		name string
		f    func(gt *testing.T, deltaTimeOffset *hexutil.Uint64, taigaTimeOffset *hexutil.Uint64)
	}{
		{"NormalBatcher", NormalBatcher},
		{"L2Finalization", L2Finalization},
		{"L2FinalizationWithSparseL1", L2FinalizationWithSparseL1},
		{"GarbageBatch", GarbageBatch},
		{"ExtendedTimeWithoutL1Batches", ExtendedTimeWithoutL1Batches},
		{"BigL2Txs", BigL2Txs},
		{"ProgressEvenWithReceiptRPCFailure", ProgressEvenWithReceiptRPCFailure},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name+"_SingularBatch", func(t *testing.T) {
			test.f(t, nil, nil)
		})
	}

	deltaTimeOffset, taigaTimeOffset := hexutil.Uint64(0), hexutil.Uint64(0)
	for _, test := range tests {
		test := test
		t.Run(test.name+"_SpanBatch", func(t *testing.T) {
			test.f(t, &deltaTimeOffset, &taigaTimeOffset)
		})
	}
}

// applyDeltaTimeOffset adjusts fork configuration to not conflict with the delta overrides
func applyDeltaTimeOffset(dp *e2eutils.DeployParams, deltaTimeOffset *hexutil.Uint64) {
	dp.DeployConfig.L2GenesisDeltaTimeOffset = deltaTimeOffset
	// configure Ecotone to not be before Delta accidentally
	if dp.DeployConfig.L2GenesisEcotoneTimeOffset != nil {
		if deltaTimeOffset == nil {
			dp.DeployConfig.L2GenesisEcotoneTimeOffset = nil
		} else if *dp.DeployConfig.L2GenesisEcotoneTimeOffset < *deltaTimeOffset {
			dp.DeployConfig.L2GenesisEcotoneTimeOffset = deltaTimeOffset
		}
	}
	// configure Fjord to not be before Delta accidentally
	if dp.DeployConfig.L2GenesisFjordTimeOffset != nil {
		if deltaTimeOffset == nil {
			dp.DeployConfig.L2GenesisFjordTimeOffset = nil
		} else if *dp.DeployConfig.L2GenesisFjordTimeOffset < *deltaTimeOffset {
			dp.DeployConfig.L2GenesisFjordTimeOffset = deltaTimeOffset
		}
	}
}

func applyTaigaTimeOffset(dp *e2eutils.DeployParams, taigaTimeOffset *hexutil.Uint64) {
	dp.DeployConfig.L2GenesisTaigaTimeOffset = taigaTimeOffset
	dp.DeployConfig.L2GenesisEcotoneTimeOffset = taigaTimeOffset
}

func NormalBatcher(gt *testing.T, deltaTimeOffset, taigaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   20, // larger than L1 block time we simulate in this test (12)
		SequencerWindowSize: 24,
		ChannelTimeout:      20,
		L1BlockTime:         12,
	}
	dp := e2eutils.MakeDeployParams(t, p)
	applyDeltaTimeOffset(dp, deltaTimeOffset)
	applyTaigaTimeOffset(dp, taigaTimeOffset)

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, seqEngine, sequencer := setupSequencerTest(t, sd, log)
	verifEngine, verifier := setupVerifier(t, sd, log, miner.L1Client(t, sd.RollupCfg), miner.BlobStore(), &sync.Config{})

	rollupSeqCl := sequencer.RollupClient()
	batcher := NewL2Batcher(log, sd.RollupCfg, DefaultBatcherCfg(dp),
		rollupSeqCl, miner.EthClient(), seqEngine.EthClient(), seqEngine.EngineClient(t, sd.RollupCfg))

	// Alice makes a L2 tx
	cl := seqEngine.EthClient()
	n, err := cl.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
	require.NoError(t, err)
	signer := types.LatestSigner(sd.L2Cfg.Config)
	tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
		ChainID:   sd.L2Cfg.Config.ChainID,
		Nonce:     n,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
		Gas:       params.TxGas,
		To:        &dp.Addresses.Bob,
		Value:     e2eutils.Ether(2),
	})
	require.NoError(t, cl.SendTransaction(t.Ctx(), tx))

	sequencer.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	// Make L2 block
	sequencer.ActL2StartBlock(t)
	seqEngine.ActL2IncludeTx(dp.Addresses.Alice)(t)
	sequencer.ActL2EndBlock(t)

	// batch submit to L1
	batcher.ActL2BatchBuffer(t)
	batcher.ActL2ChannelClose(t)
	batcher.ActL2BatchSubmit(t)

	// confirm batch on L1
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)
	bl := miner.l1Chain.CurrentBlock()
	log.Info("bl", "txs", len(miner.l1Chain.GetBlockByHash(bl.Hash()).Transactions()))

	// Now make enough L1 blocks that the verifier will have to derive a L2 block
	// It will also eagerly derive the block from the batcher
	for i := uint64(0); i < sd.RollupCfg.SeqWindowSize; i++ {
		miner.ActL1StartBlock(12)(t)
		miner.ActL1EndBlock(t)
	}

	// sync verifier from L1 batch in otherwise empty sequence window
	verifier.ActL1HeadSignal(t)
	verifier.ActL2PipelineFull(t)
	require.Equal(t, uint64(1), verifier.SyncStatus().SafeL2.L1Origin.Number)

	// check that the tx from alice made it into the L2 chain
	verifCl := verifEngine.EthClient()
	vTx, isPending, err := verifCl.TransactionByHash(t.Ctx(), tx.Hash())
	require.NoError(t, err)
	require.False(t, isPending)
	require.NotNil(t, vTx)
}

func L2Finalization(gt *testing.T, deltaTimeOffset, taigaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	applyDeltaTimeOffset(dp, deltaTimeOffset)
	applyTaigaTimeOffset(dp, taigaTimeOffset)

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, engine, sequencer := setupSequencerTest(t, sd, log)

	sequencer.ActL2PipelineFull(t)

	// build an empty L1 block (#1), mark it as justified
	miner.ActEmptyBlock(t)
	miner.ActL1SafeNext(t) // #0 -> #1

	// sequencer builds L2 chain, up to and including a block that has the new L1 block as origin
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	sequencer.ActL2PipelineFull(t)
	sequencer.ActL1SafeSignal(t)
	require.Equal(t, uint64(1), sequencer.SyncStatus().SafeL1.Number)

	// build another L1 block (#2), mark it as justified. And mark previous justified as finalized.
	miner.ActEmptyBlock(t)
	miner.ActL1SafeNext(t)     // #1 -> #2
	miner.ActL1FinalizeNext(t) // #0 -> #1
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	// continue to build L2 chain referencing the new L1 blocks
	sequencer.ActL2PipelineFull(t)
	sequencer.ActL1FinalizedSignal(t)
	sequencer.ActL1SafeSignal(t)
	require.Equal(t, uint64(2), sequencer.SyncStatus().SafeL1.Number)
	require.Equal(t, uint64(1), sequencer.SyncStatus().FinalizedL1.Number)
	require.Equal(t, uint64(0), sequencer.SyncStatus().FinalizedL2.Number, "L2 block has to be included on L1 before it can be finalized")

	batcher := NewL2Batcher(log, sd.RollupCfg, DefaultBatcherCfg(dp),
		sequencer.RollupClient(), miner.EthClient(), engine.EthClient(), engine.EngineClient(t, sd.RollupCfg))

	heightToSubmit := sequencer.SyncStatus().UnsafeL2.Number

	batcher.ActSubmitAll(t)
	// confirm batch on L1, block #3
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)

	// read the batch
	sequencer.ActL2PipelineFull(t)
	require.Equal(t, uint64(0), sequencer.SyncStatus().FinalizedL2.Number, "Batch must be included in finalized part of L1 chain for L2 block to finalize")

	// build some more L2 blocks, so there is an unsafe part again that hasn't been submitted yet
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	// submit those blocks too, block #4
	batcher.ActSubmitAll(t)
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)

	// add some more L1 blocks #5, #6
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t)

	// and more unsafe L2 blocks
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	// move safe/finalize markers: finalize the L1 chain block with the first batch, but not the second
	miner.ActL1SafeNext(t)     // #2 -> #3
	miner.ActL1SafeNext(t)     // #3 -> #4
	miner.ActL1FinalizeNext(t) // #1 -> #2
	miner.ActL1FinalizeNext(t) // #2 -> #3

	sequencer.ActL2PipelineFull(t)
	sequencer.ActL1FinalizedSignal(t)
	sequencer.ActL1SafeSignal(t)
	sequencer.ActL1HeadSignal(t)
	require.Equal(t, uint64(6), sequencer.SyncStatus().HeadL1.Number)
	require.Equal(t, uint64(4), sequencer.SyncStatus().SafeL1.Number)
	require.Equal(t, uint64(3), sequencer.SyncStatus().FinalizedL1.Number)
	require.Equal(t, heightToSubmit, sequencer.SyncStatus().FinalizedL2.Number, "finalized L2 blocks in first batch")

	// need to act with the engine on the signals still
	sequencer.ActL2PipelineFull(t)

	engCl := engine.EngineClient(t, sd.RollupCfg)
	engBlock, err := engCl.L2BlockRefByLabel(t.Ctx(), eth.Finalized)
	require.NoError(t, err)
	require.Equal(t, heightToSubmit, engBlock.Number, "engine finalizes what rollup node finalizes")

	// Now try to finalize block 4, but with a bad/malicious alternative hash.
	// If we get this false signal, we shouldn't finalize the L2 chain.
	altBlock4 := sequencer.SyncStatus().SafeL1
	altBlock4.Hash = common.HexToHash("0xdead")
	sequencer.derivation.Finalize(altBlock4)
	sequencer.ActL2PipelineFull(t)
	require.Equal(t, uint64(3), sequencer.SyncStatus().FinalizedL1.Number)
	require.Equal(t, heightToSubmit, sequencer.SyncStatus().FinalizedL2.Number, "unknown/bad finalized L1 blocks are ignored")
}

// L2FinalizationWithSparseL1 tests that safe L2 blocks can be finalized even if we do not regularly get a L1 finalization signal
func L2FinalizationWithSparseL1(gt *testing.T, deltaTimeOffset, taigaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	applyDeltaTimeOffset(dp, deltaTimeOffset)
	applyTaigaTimeOffset(dp, taigaTimeOffset)

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, engine, sequencer := setupSequencerTest(t, sd, log)

	sequencer.ActL2PipelineFull(t)

	miner.ActEmptyBlock(t)
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	startStatus := sequencer.SyncStatus()
	require.Less(t, startStatus.SafeL2.Number, startStatus.UnsafeL2.Number, "sequencer has unsafe L2 block")

	batcher := NewL2Batcher(log, sd.RollupCfg, DefaultBatcherCfg(dp),
		sequencer.RollupClient(), miner.EthClient(), engine.EthClient(), engine.EngineClient(t, sd.RollupCfg))
	batcher.ActSubmitAll(t)

	// include in L1
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)

	// Make 2 L1 blocks without batches
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t)

	// See the L1 head, and traverse the pipeline to it
	sequencer.ActL1HeadSignal(t)
	sequencer.ActL2PipelineFull(t)

	updatedStatus := sequencer.SyncStatus()
	require.Equal(t, updatedStatus.SafeL2.Number, updatedStatus.UnsafeL2.Number, "unsafe L2 block is now safe")
	require.Less(t, updatedStatus.FinalizedL2.Number, updatedStatus.UnsafeL2.Number, "submitted block is not yet finalized")

	// Now skip straight to the head with L1 signals (sequencer has traversed the L1 blocks, but they did not have L2 contents)
	headL1Num := miner.UnsafeNum()
	miner.ActL1Safe(t, headL1Num)
	miner.ActL1Finalize(t, headL1Num)
	sequencer.ActL1SafeSignal(t)
	sequencer.ActL1FinalizedSignal(t)

	// Now see if the signals can be processed
	sequencer.ActL2PipelineFull(t)

	finalStatus := sequencer.SyncStatus()
	// Verify the signal was processed, even though we signalled a later L1 block than the one with the batch.
	require.Equal(t, finalStatus.FinalizedL2.Number, finalStatus.UnsafeL2.Number, "sequencer submitted its L2 block and it finalized")
}

// GarbageBatch tests the behavior of an invalid/malformed output channel frame containing
// valid batches being submitted to the batch inbox. These batches should always be rejected
// and the safe L2 head should remain unaltered.
func GarbageBatch(gt *testing.T, deltaTimeOffset, taigaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)
	p := defaultRollupTestParams
	dp := e2eutils.MakeDeployParams(t, p)
	applyDeltaTimeOffset(dp, deltaTimeOffset)
	applyTaigaTimeOffset(dp, taigaTimeOffset)

	for _, garbageKind := range GarbageKinds {
		sd := e2eutils.Setup(t, dp, defaultAlloc)
		log := testlog.Logger(t, log.LvlError)
		miner, engine, sequencer := setupSequencerTest(t, sd, log)

		_, verifier := setupVerifier(t, sd, log, miner.L1Client(t, sd.RollupCfg), miner.BlobStore(), &sync.Config{})

		batcherCfg := DefaultBatcherCfg(dp)

		if garbageKind == MALFORM_RLP || garbageKind == INVALID_COMPRESSION {
			// If the garbage kind is `INVALID_COMPRESSION` or `MALFORM_RLP`, use the `actions` packages
			// modified `ChannelOut`.
			batcherCfg.GarbageCfg = &GarbageChannelCfg{
				useInvalidCompression: garbageKind == INVALID_COMPRESSION,
				malformRLP:            garbageKind == MALFORM_RLP,
			}
		}

		batcher := NewL2Batcher(log, sd.RollupCfg, batcherCfg, sequencer.RollupClient(), miner.EthClient(), engine.EthClient(), engine.EngineClient(t, sd.RollupCfg))

		sequencer.ActL2PipelineFull(t)
		verifier.ActL2PipelineFull(t)

		syncAndBuildL2 := func() {
			// Send a head signal to the sequencer and verifier
			sequencer.ActL1HeadSignal(t)
			verifier.ActL1HeadSignal(t)

			// Run the derivation pipeline on the sequencer and verifier
			sequencer.ActL2PipelineFull(t)
			verifier.ActL2PipelineFull(t)

			// Build the L2 chain to the L1 head
			sequencer.ActBuildToL1Head(t)
		}

		// Build an empty block on L1 and run the derivation pipeline + build L2
		// to the L1 head (block #1)
		miner.ActEmptyBlock(t)
		syncAndBuildL2()

		// Ensure that the L2 safe head has an L1 Origin at genesis before any
		// batches are submitted.
		require.Equal(t, uint64(0), sequencer.L2Safe().L1Origin.Number)
		require.Equal(t, uint64(1), sequencer.L2Unsafe().L1Origin.Number)

		// Submit a batch containing all blocks built on L2 while catching up
		// to the L1 head above. The output channel frame submitted to the batch
		// inbox will be invalid- it will be malformed depending on the passed
		// `garbageKind`.
		batcher.ActBufferAll(t)
		batcher.ActL2ChannelClose(t)
		batcher.ActL2BatchSubmitGarbage(t, garbageKind)

		// Include the batch on L1 in block #2
		miner.ActL1StartBlock(12)(t)
		miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
		miner.ActL1EndBlock(t)

		// Send a head signal + run the derivation pipeline on the sequencer
		// and verifier.
		syncAndBuildL2()

		// Verify that the L2 blocks that were batch submitted were *not* marked
		// as safe due to the malformed output channel frame. The safe head should
		// still have an L1 Origin at genesis.
		require.Equal(t, uint64(0), sequencer.L2Safe().L1Origin.Number)
		require.Equal(t, uint64(2), sequencer.L2Unsafe().L1Origin.Number)
	}
}

func ExtendedTimeWithoutL1Batches(gt *testing.T, deltaTimeOffset, taigaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   20, // larger than L1 block time we simulate in this test (12)
		SequencerWindowSize: 24,
		ChannelTimeout:      20,
		L1BlockTime:         12,
	}
	dp := e2eutils.MakeDeployParams(t, p)
	applyDeltaTimeOffset(dp, deltaTimeOffset)
	applyTaigaTimeOffset(dp, taigaTimeOffset)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlError)
	miner, engine, sequencer := setupSequencerTest(t, sd, log)

	_, verifier := setupVerifier(t, sd, log, miner.L1Client(t, sd.RollupCfg), miner.BlobStore(), &sync.Config{})

	batcher := NewL2Batcher(log, sd.RollupCfg, DefaultBatcherCfg(dp),
		sequencer.RollupClient(), miner.EthClient(), engine.EthClient(), engine.EngineClient(t, sd.RollupCfg))

	sequencer.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	// make a long L1 chain, up to just one block left for L2 blocks to be included.
	for i := uint64(0); i < p.SequencerWindowSize-1; i++ {
		miner.ActEmptyBlock(t)
	}

	// Now build a L2 chain that references all of these L1 blocks
	sequencer.ActL1HeadSignal(t)
	sequencer.ActBuildToL1Head(t)

	// Now submit all the L2 blocks in the very last L1 block within sequencer window range
	batcher.ActSubmitAll(t)
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)

	// Now sync the verifier, and see if the L2 chain of the sequencer is safe
	verifier.ActL2PipelineFull(t)
	require.Equal(t, sequencer.L2Unsafe(), verifier.L2Safe(), "all L2 blocks should have been included just in time")
	sequencer.ActL2PipelineFull(t)
	require.Equal(t, sequencer.L2Unsafe(), sequencer.L2Safe(), "same for sequencer")
}

// BigL2Txs tests a high-throughput case with constrained batcher:
//   - Fill 40 L2 blocks to near max-capacity, with txs of 120 KB each
//   - Buffer the L2 blocks into channels together as much as possible, submit data-txs only when necessary
//     (just before crossing the max RLP channel size)
//   - Limit the data-tx size to 40 KB, to force data to be split across multiple datat-txs
//   - Defer all data-tx inclusion till the end
//   - Fill L1 blocks with data-txs until we have processed them all
//   - Run the verifier, and check if it derives the same L2 chain as was created by the sequencer.
//
// The goal of this test is to quickly run through an otherwise very slow process of submitting and including lots of data.
// This does not test the batcher code, but is really focused at testing the batcher utils
// and channel-decoding verifier code in the derive package.
func BigL2Txs(gt *testing.T, deltaTimeOffset, taigaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   100,
		SequencerWindowSize: 1000,
		ChannelTimeout:      200, // give enough space to buffer large amounts of data before submitting it
		L1BlockTime:         12,
	}
	dp := e2eutils.MakeDeployParams(t, p)
	applyDeltaTimeOffset(dp, deltaTimeOffset)
	applyTaigaTimeOffset(dp, taigaTimeOffset)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlInfo)
	miner, engine, sequencer := setupSequencerTest(t, sd, log)

	_, verifier := setupVerifier(t, sd, log, miner.L1Client(t, sd.RollupCfg), miner.BlobStore(), &sync.Config{})

	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize:          0,
		MaxL1TxSize:          40_000, // try a small batch size, to force the data to be split between more frames
		BatcherKey:           dp.Secrets.Batcher,
		DataAvailabilityType: batcherFlags.CalldataType,
	}, sequencer.RollupClient(), miner.EthClient(), engine.EthClient(), engine.EngineClient(t, sd.RollupCfg))

	sequencer.ActL2PipelineFull(t)

	verifier.ActL2PipelineFull(t)
	cl := engine.EthClient()

	batcherNonce := uint64(0) // manually track batcher nonce. the "pending nonce" value in tx-pool is incorrect after we fill the pending-block gas limit and keep adding txs to the pool.
	batcherTxOpts := func(tx *types.DynamicFeeTx) {
		tx.Nonce = batcherNonce
		batcherNonce++
		tx.GasFeeCap = e2eutils.Ether(1) // be very generous with basefee, since we're spamming L1
	}

	rng := rand.New(rand.NewSource(555))

	// build many L2 blocks filled to the brim with large txs of random data
	for i := 0; i < 40; i++ {
		aliceNonce, err := cl.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
		status := sequencer.SyncStatus()
		// build empty L1 blocks as necessary, so the L2 sequencer can continue to include txs while not drifting too far out
		if status.UnsafeL2.Time >= status.HeadL1.Time+12 {
			miner.ActEmptyBlock(t)
		}
		sequencer.ActL1HeadSignal(t)
		sequencer.ActL2StartBlock(t)
		baseFee := engine.l2Chain.CurrentBlock().BaseFee // this will go quite high, since so many consecutive blocks are filled at capacity.
		// fill the block with large L2 txs from alice
		for n := aliceNonce; ; n++ {
			require.NoError(t, err)
			signer := types.LatestSigner(sd.L2Cfg.Config)
			data := make([]byte, 120_000) // very large L2 txs, as large as the tx-pool will accept
			_, err := rng.Read(data[:])   // fill with random bytes, to make compression ineffective
			require.NoError(t, err)
			gas, err := core.IntrinsicGas(data, nil, false, true, true, false)
			require.NoError(t, err)
			if gas > engine.engineApi.RemainingBlockGas() {
				break
			}
			tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
				ChainID:   sd.L2Cfg.Config.ChainID,
				Nonce:     n,
				GasTipCap: big.NewInt(2 * params.GWei),
				GasFeeCap: new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), big.NewInt(2*params.GWei)),
				Gas:       gas,
				To:        &dp.Addresses.Bob,
				Value:     big.NewInt(0),
				Data:      data,
			})
			require.NoError(t, cl.SendTransaction(t.Ctx(), tx))
			engine.ActL2IncludeTx(dp.Addresses.Alice)(t)
		}
		sequencer.ActL2EndBlock(t)
		for batcher.l2BufferedBlock.Number < sequencer.SyncStatus().UnsafeL2.Number {
			// if we run out of space, close the channel and submit all the txs
			if err := batcher.Buffer(t); errors.Is(err, derive.ErrTooManyRLPBytes) || errors.Is(err, derive.ErrCompressorFull) {
				log.Info("flushing filled channel to batch txs", "id", batcher.l2ChannelOut.ID())
				batcher.ActL2ChannelClose(t)
				for batcher.l2ChannelOut != nil {
					batcher.ActL2BatchSubmit(t, batcherTxOpts)
				}
			}
		}
	}

	// if anything is left in the channel, submit it
	if batcher.l2ChannelOut != nil {
		log.Info("flushing trailing channel to batch txs", "id", batcher.l2ChannelOut.ID())
		batcher.ActL2ChannelClose(t)
		for batcher.l2ChannelOut != nil {
			batcher.ActL2BatchSubmit(t, batcherTxOpts)
		}
	}

	// build L1 blocks until we're out of txs
	txs, _ := miner.eth.TxPool().ContentFrom(dp.Addresses.Batcher)
	for {
		if len(txs) == 0 {
			break
		}
		miner.ActL1StartBlock(12)(t)
		for range txs {
			if len(txs) == 0 {
				break
			}
			tx := txs[0]
			if miner.l1GasPool.Gas() < tx.Gas() { // fill the L1 block with batcher txs until we run out of gas
				break
			}
			log.Info("including batcher tx", "nonce", tx.Nonce())
			miner.IncludeTx(t, tx)
			txs = txs[1:]
		}
		miner.ActL1EndBlock(t)
	}
	verifier.ActL1HeadSignal(t)
	verifier.ActL2PipelineFull(t)
	require.Equal(t, sequencer.SyncStatus().UnsafeL2, verifier.SyncStatus().SafeL2, "verifier synced sequencer data even though of huge tx in block")
}

func ProgressEvenWithReceiptRPCFailure(gt *testing.T, deltaTimeOffset, taigaTimeOffset *hexutil.Uint64) {
	if deltaTimeOffset != nil {
		gt.Skipf("skipping delta for now because span batches need adjustment to compressor logic")
	}

	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxSequencerDrift:   20, // larger than L1 block time we simulate in this test (12)
		SequencerWindowSize: 24,
		ChannelTimeout:      20,
		L1BlockTime:         12,
	}
	dp := e2eutils.MakeDeployParams(t, p)
	applyDeltaTimeOffset(dp, deltaTimeOffset)
	applyTaigaTimeOffset(dp, taigaTimeOffset)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, seqEngine, sequencer := setupSequencerTest(t, sd, log)
	miner.ActL1SetFeeRecipient(common.Address{'A'})
	verifEngine, verifier := setupVerifier(t, sd, log, miner.L1Client(t, sd.RollupCfg), miner.BlobStore(), &sync.Config{})
	rollupSeqCl := sequencer.RollupClient()
	batcherCfg := DefaultBatcherCfg(dp)

	batcher := NewL2Batcher(log, sd.RollupCfg, batcherCfg,
		rollupSeqCl, miner.EthClient(), seqEngine.EthClient(),
		seqEngine.EngineClient(t, sd.RollupCfg),
	)
	verifier.ActL2PipelineFull(t)
	sequencer.ActL2PipelineFull(t)
	verifierRPC := verifEngine.EthClient()
	txs := map[common.Hash]struct{}{}

	signer := types.LatestSigner(sd.L2Cfg.Config)
	cl := seqEngine.EthClient()
	aliceTx := func() common.Hash {
		n, err := cl.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
		require.NoError(t, err)
		tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
			ChainID:   sd.L2Cfg.Config.ChainID,
			Nonce:     n,
			GasTipCap: big.NewInt(2 * params.GWei),
			GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
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

	// L1 makes a block
	miner.ActL1StartBlock(p.L1BlockTime)(t)
	miner.ActL1EndBlock(t)
	sequencer.ActL1HeadSignal(t)
	origin := miner.l1Chain.CurrentBlock()

	// L2 makes blocks to catch up
	for sequencer.SyncStatus().UnsafeL2.Time+sd.RollupCfg.BlockTime < origin.Time {
		makeL2BlockWithAliceTx()
		batcher.ActL2BatchBuffer(t)
		require.Equal(t, uint64(0), sequencer.SyncStatus().UnsafeL2.L1Origin.Number, "no L1 origin change before time matches")
	}

	batcher.ActL2ChannelClose(t)
	batcher.ActL2BatchSubmit(t)

	miner.ActL1StartBlock(p.L1BlockTime)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)

	//      require.Equal(t, uint64(1), sequencer.SyncStatus().UnsafeL2.L1Origin.Number, "L1 origin changes as soon as L2 time equals or exceeds L1 time")

	sequencer.ActL1HeadSignal(t)
	sequencer.ActL2PipelineFull(t)
	// lets say the receipts fetcher is failing
	sequencer.mockL1ReceiptsFetcher.failReceipts = true

	// Make blocks up till the sequencer drift is about to surpass
	// for sequencer.SyncStatus().UnsafeL2.Time+sd.RollupCfg.BlockTime < origin.Time+sd.RollupCfg.MaxSequencerDrift-p.L1BlockTime {
	for sequencer.SyncStatus().UnsafeL2.Time+sd.RollupCfg.BlockTime < origin.Time+sd.ChainSpec.MaxSequencerDrift(sequencer.SyncStatus().UnsafeL2.Time)-p.L1BlockTime {
		txs[makeL2BlockWithAliceTx()] = struct{}{}
		batcher.ActL2BatchBuffer(t)
		require.Equal(t, uint64(0), sequencer.SyncStatus().UnsafeL2.L1Origin.Number, "expected to keep old L1 origin")
	}

	batcher.ActL2ChannelClose(t)
	batcher.ActL2BatchSubmit(t)
	miner.ActL1StartBlock(p.L1BlockTime)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)
	sequencer.ActL1HeadSignal(t)
	sequencer.ActL2PipelineFull(t)

	sequencer.mockL1ReceiptsFetcher.failReceipts = false

	// Now one more l2 block to kick it over the drift
	makeL2BlockWithAliceTx()
	batcher.ActL2BatchBuffer(t)
	batcher.ActL2ChannelClose(t)
	batcher.ActL2BatchSubmit(t)
	miner.ActL1StartBlock(p.L1BlockTime)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)

	// Now make enough L1 blocks that the verifier will have to derive a L2 block
	// It will also eagerly derive the block from the batcher
	for i := uint64(0); i < sd.RollupCfg.SeqWindowSize; i++ {
		miner.ActL1StartBlock(p.L1BlockTime)(t)
		miner.ActL1EndBlock(t)
	}

	// sync verifier from L1 batch in otherwise empty sequence window
	verifierSyncStatus := verifier.SyncStatus()
	safeHead := verifierSyncStatus.SafeL2.Number
	sequencer.ActL1HeadSignal(t)
	sequencer.ActL2PipelineFull(t)
	seqSyncStatus := sequencer.SyncStatus()

	verifier.l1.failReceipts = true
	time.AfterFunc(time.Millisecond*200, func() { verifier.l1.failReceipts = false })

	verifier.ActL1HeadSignal(t)
	verifier.ActL2PipelineFull(t)
	verifierSyncStatus = verifier.SyncStatus()
	t.Logf(
		"verifier after deriviation kicked off unsafehead %v safehead %v before was safehead %v - sequencer sync status %v",
		verifierSyncStatus.UnsafeL2.Number, verifierSyncStatus.SafeL2.Number, safeHead, seqSyncStatus.SafeL2.Number,
	)
	require.Equal(t, seqSyncStatus.SafeL2.Number, verifierSyncStatus.SafeL2.Number, "verifier safe head did not move as necessary")

	for hsh := range txs {
		rcpt, err := verifierRPC.TransactionReceipt(t.Ctx(), hsh)
		require.NoError(t, err, "missing rcpt in verifier for tx %v", hsh)
		require.NotNil(t, rcpt, "receipt is nil for verifier but should be there")
	}
}

func getBlock(ctx context.Context, client client.RPC, method string, tag string) (*types.Block, error) {
	var raw json.RawMessage
	err := client.CallContext(ctx, &raw, method, tag, true)
	if err != nil {
		return nil, err
	}

	var head *types.Header
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, err
	}

	type rpcBlock struct {
		Hash         common.Hash          `json:"hash"`
		Transactions []*types.Transaction `json:"transactions"`
	}

	var body *rpcBlock
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}

	return types.NewBlockWithHeader(head).WithBody(body.Transactions, nil), nil
}

func TestBatchWrongInput(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)

	var rollupConfig rollup.Config
	require.NoError(t, json.Unmarshal(files.BlastRollup, &rollupConfig))

	l1Sepolia, err := ethclient.Dial(l1RPC)
	require.NoError(t, err)

	rpc, err := client.NewRPC(t.Ctx(), log, l2RPC)
	require.NoError(t, err)

	// blk, err := getBlock(t.Ctx(), rpc, methodEthGetBlockByNumber, hexutil.EncodeUint64(20154827))
	blk, err := getBlock(t.Ctx(), rpc, methodEthGetBlockByNumber, "safe")

	require.NoError(t, err)
	log.Info("pulled block from rpc", "ts", blk.Time())
	//getBlock(ctx, client, methodEthGetBlockByNumber, "latest")
	batch, l1Info, err := derive.BlockToSingularBatch(&rollupConfig, blk)
	require.NoError(t, err)
	log.Info("show batch, l1info", "batch", batch, "l1-info", l1Info)

	beaconAddr := l1RPC
	beaconClient := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(beaconAddr, nil))
	beaconCfg := sources.L1BeaconClientConfig{FetchAllSidecars: true}
	beacon := sources.NewL1BeaconClient(beaconClient, beaconCfg)
	batcher := rollupConfig.Genesis.SystemConfig.BatcherAddr
	wrk, _ := os.Getwd()

	locFrames := filepath.Join(wrk, "frames-dump")
	locBatches := filepath.Join(wrk, "batches-dump")

	totalValid, totalInvalid := fetch.Batches(l1Sepolia, beacon, fetch.Config{
		Start:              batch.Epoch().Number - 1000,
		End:                batch.Epoch().Number + 1,
		ChainID:            rollupConfig.L1ChainID,
		BatchInbox:         rollupConfig.BatchInboxAddress,
		BatchSenders:       map[common.Address]struct{}{batcher: {}},
		OutDirectory:       locFrames,
		ConcurrentRequests: 1,
	})

	log.Info("finished grabbing batches", "total-valid", totalValid, "total-invalid", totalInvalid)

	config := reassemble.Config{
		BatchInbox:    rollupConfig.BatchInboxAddress,
		InDirectory:   locFrames,
		OutDirectory:  locBatches,
		L2ChainID:     rollupConfig.L2ChainID,
		L2GenesisTime: rollupConfig.Genesis.L2Time,
		L2BlockTime:   rollupConfig.BlockTime,
	}

	onChain := reassemble.ChannelsLookingFor(config, &rollupConfig, batch)
	if onChain == nil {
		log.Error("couldnt find ours, need adjust range search")
	} else {
		log.Info("found a posted one for our safe head", "landed", onChain, "but-ours-computed", batch)
	}
}

func TestShowBadBatch(gt *testing.T) {
	decoded := common.FromHex(badBatch)
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)

	var rollupConfig rollup.Config
	require.NoError(t, json.Unmarshal(files.BlastRollup, &rollupConfig))

	var badBtch derive.SingularBatch
	require.NoError(t, rlp.DecodeBytes(decoded, &badBtch))
	// plainL1, err := ethclient.Dial(l1RPC)
	// require.NoError(t, err)
	// plainL1.BlockByHash(t.Ctx(), btch.ParentHash)
	log.Info("show bad batch", "bad", badBtch)
	rpc, err := client.NewRPC(t.Ctx(), log, l2RPC)
	require.NoError(t, err)

	// blk, err := getBlock(t.Ctx(), rpc, methodEthGetBlockByNumber, hexutil.EncodeUint64(20154827))
	blk, err := getBlock(t.Ctx(), rpc, methodEthGetBlockByNumber, "safe")
	require.NoError(t, err)
	batch, l1Info, err := derive.BlockToSingularBatch(&rollupConfig, blk)
	require.NoError(t, err)

	badPr, _ := json.MarshalIndent([]any{"bad-batch", badBtch, "computed", batch, "safe-block-hash", blk.Hash()}, " ", " ")
	_ = l1Info
	fmt.Println(string(badPr))
	if blk.Hash() == badBtch.ParentHash {
		log.Info("got it to match - what happened")
	} else {
		log.Error("not matching")
	}

	endpointProvider, err := dial.NewStaticL2EndpointProvider(t.Ctx(), log, l2RPC, localOPNode)
	require.NoError(t, err)
	l2Client, err := endpointProvider.EthClient(t.Ctx())
	require.NoError(t, err)
	block, err := l2Client.BlockByNumber(t.Ctx(), new(big.Int).SetUint64(blk.Number().Uint64()))
	require.NoError(t, err)
	log.Info("how batcher saw hsh", "hsh", block.Hash().Hex())
}

const (
	l1RPC = "https://serene-yolo-sanctuary.ethereum-sepolia.quiknode.pro/369c5d02538c5ab258019669585a02bf15351ede/"
	// l2RPC                     = "ws://localhost:8546"
	l2RPC                     = "http://localhost:8545"
	localOPNode               = "http://localhost:1545"
	methodEthGetBlockByNumber = "eth_getBlockByNumber"
	badBatch                  = "f90898a09aac1fc2dca37397e19380d1962bb76faad01a9572c461fed072143bdfdafdf7837c27baa04067a36974ab97bf6c7c25d558df8d0102f5573951b9d521a6670f96d66e2c34846800e21af9084ab9015af9015782046f841dcd65008303d0909436b2415644d47b8f646697b6c4c5a9d55400f2dd880de0b6b3a7640000b8e456591d5961726274000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f00bd29577a944eee55e5198c64e957cc4d7ba790000000000000000000000000000000000000000000000000de08e51f0c04e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a7640000841418e41ea072cd92070a2803ff026179d4985322fb6da4b4007a768370169fdb208cf02e97a0387d74ae812c8c3ed64c9e7eb9e29de295bc7689df454b4636d3aa2301ca5d0fb9016002f9015c840a0c71fd820958830f433c830f433c830266799436b2415644d47b8f646697b6c4c5a9d55400f2dd880de0b6b3a7640000b8e456591d596f707374000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d91fca033ea59bb087cc0456a7d57de9031241a20000000000000000000000000000000000000000000000000de0b31054cfc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a7640000c001a0930b47215af9970ca592747b315fc3bc94aa1a4c28f3a3d0d5620ec6153ec995a0326409c2dc91f07014f8900b7f23f90d06304604c60bed9b08e1d7a13bcba36eb9016002f9015c840a0c71fd820bc1830f433c830f433c83024b359436b2415644d47b8f646697b6c4c5a9d55400f2dd880de0b6b3a7640000b8e456591d596f70737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000034c8d658873248df5ef211b6f29dceafcb721c0a0000000000000000000000000000000000000000000000000de0b31054cfc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a7640000c001a0874ba035c90d7f0778680b4ae4c9087467bec56ee2b8f5c2016ef3b310140994a0578e043f155bdc327e1179661637214ff063cdc4d70eee6d736296c0d22bf83ab9016002f9015c840a0c71fd820948830f433c830f433c830283e09436b2415644d47b8f646697b6c4c5a9d55400f2dd880de0b6b3a7640000b8e456591d596f707374000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000964edb512fc657b07992a0ca478cdb5ca0649b2a0000000000000000000000000000000000000000000000000de0b31054cfc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a7640000c001a0ee0e576d2f1115cc1175e56ecde1c1ba33ae42a1373b78806e08793508c95e02a0509876b91bf852fb6da41b5292bdb8dfe2bf574769c5f26f401737840e1a576fb9016002f9015c840a0c71fd82166e830f433c830f433c830291a69436b2415644d47b8f646697b6c4c5a9d55400f2dd880de0b6b3a7640000b8e456591d596f70737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000088ba6591b95feb7963805bb030cf36e3e9ef57d30000000000000000000000000000000000000000000000000de0b31054cfc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a7640000c001a096c2e8e618162ab54b5cb49fc4cb6b708f896cb32e03358c1fa43a82762e931da0239955ff99ab8e92944272fee51f86ef4f37f70ecdc24e8e9d2945811bc714a6b9015e02f9015a840a0c71fd82085c82045582074983019c939436b2415644d47b8f646697b6c4c5a9d55400f2dd880de0b32a5daf0680b8e42dc4edfd0d780422e90450f2283ff40d657870b0f0adb0d21c25ec95e3b0f9afd784c5170000000000000000000000009d44f016823085f3f9a57dea7b142b37672248bf0000000000000000000000000000000000000000000000000de0b32a5daf06800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006800e1dd0000000000000000000000009d44f016823085f3f9a57dea7b142b37672248bf756e697400000000000000000000000000000000000000000000000000000000c001a0bfa9daea61490c9449d77f1c17c346c3ad3402bd3952100ebd8e0bf2d6f8db4fa05e46ea6bea97bc6fa4b44effda8829568a41f431a7ca74f58cc56f11b17cf257"
)
