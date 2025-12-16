package derive

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
)

var (
	MockDepositContractAddr               = common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeef00000000")
	_                       eth.BlockInfo = (*testutils.MockBlockInfo)(nil)
)

type infoTest struct {
	name    string
	mkInfo  func(rng *rand.Rand) *testutils.MockBlockInfo
	mkL1Cfg func(rng *rand.Rand, l1Info eth.BlockInfo) eth.SystemConfig
	seqNr   func(rng *rand.Rand) uint64
}

func randomL1Cfg(rng *rand.Rand, l1Info eth.BlockInfo) eth.SystemConfig {
	return eth.SystemConfig{
		BatcherAddr: testutils.RandomAddress(rng),
		Overhead:    [32]byte{},
		Scalar:      [32]byte{},
		GasLimit:    1234567,
	}
}

func TestPicksPragueBlobPricing(t *testing.T) {
	var (
		rollupCfg         rollup.Config
		excess            uint64 = 200000000
		blkTime           uint64 = uint64(time.Now().Unix())
		assumeL2BlockTime uint64 = blkTime
		ecotoneSepolia    uint64 = 1716843599
	)

	rollupCfg.EcotoneTime = &ecotoneSepolia

	rng := rand.New(rand.NewSource(int64(1234)))
	assumeHeader := &types.Header{
		ExcessBlobGas: &excess,
		Time:          blkTime,
		BaseFee:       big.NewInt(100),
		Number:        big.NewInt(5_000_000),
	}
	someL1Block := eth.HeaderBlockInfo(assumeHeader)

	defaultBlobBaseFee := someL1Block.BlobBaseFee(params.SepoliaChainConfig)
	t.Logf("initial blob base fee %v", defaultBlobBaseFee)
	bumpFusaka := &blkTime

	require.True(t, params.SepoliaChainConfig.IsPrague(someL1Block.Header().Number, someL1Block.Time()), "should be true already on prague")
	require.True(t, params.SepoliaChainConfig.IsOsaka(someL1Block.Header().Number, someL1Block.Time()), "should be true already on osaka")
	require.True(t, params.SepoliaChainConfig.IsBPO1(someL1Block.Header().Number, someL1Block.Time()), "should be true already on bpo1")
	require.True(t, params.SepoliaChainConfig.IsBPO2(someL1Block.Header().Number, someL1Block.Time()), "should be true already on bpo2")

	rollupCfg.FusakaBlobScheduleTime = bumpFusaka
	someL1Confg := randomL1Cfg(rng, someL1Block)
	deposit, err := L1InfoDeposit(&rollupCfg, params.SepoliaChainConfig, someL1Confg, 0, someL1Block, assumeL2BlockTime)
	require.NoError(t, err, "l1 deposit failed")
	l1InfoBack, err := L1BlockInfoFromBytes(&rollupCfg, assumeL2BlockTime, deposit.Data)
	require.NoError(t, err, "l1 info back failed")
	// NOTE So this is how it ought to be when its correct - that is if all pricing was correct
	require.True(t, l1InfoBack.BlobBaseFee.Cmp(defaultBlobBaseFee) == 0)
	// NOTE so now lets see - what if the time now was prague time - that is - just before Osaka

	betweenTime := time.Unix(int64(*params.SepoliaChainConfig.OsakaTime), 0).Add(-1 * time.Minute)
	assumeHeader.Time = uint64(betweenTime.Unix())
	someL1BlockPricePrague := eth.HeaderBlockInfo(assumeHeader)
	shouldBePrague := someL1BlockPricePrague.BlobBaseFee(params.SepoliaChainConfig)
	t.Logf("prague pricing should be %v", shouldBePrague)
	require.False(t, defaultBlobBaseFee.Cmp(shouldBePrague) == 0)

	{
		bpo2Time := time.Unix(int64(*params.SepoliaChainConfig.BPO2Time), 0)
		bpo2TimeTs := uint64(bpo2Time.Unix())
		letsSayBeforeBpo2 := uint64(bpo2Time.Add(-2 * time.Minute).Unix())
		assumeHeader.Time = letsSayBeforeBpo2
		someL1Block := eth.HeaderBlockInfo(assumeHeader)

		rollupCfg.FusakaBlobScheduleTime = &bpo2TimeTs
		rollupCfg.Bpo1BlobScheduleTime = &bpo2TimeTs
		rollupCfg.Bpo2BlobScheduleTime = &bpo2TimeTs

		deposit, err := L1InfoDeposit(&rollupCfg, params.SepoliaChainConfig, someL1Confg, 0, someL1Block, assumeL2BlockTime)
		require.NoError(t, err, "l1 deposit failed")
		l1InfoBack, err := L1BlockInfoFromBytes(&rollupCfg, assumeL2BlockTime, deposit.Data)
		require.NoError(t, err, "l1 info back failed")
		// NOTE so now this one should be prague based pricing
		require.True(t, l1InfoBack.BlobBaseFee.Cmp(shouldBePrague) == 0, "picked incorrect block base fee got %v but should be prague %v", l1InfoBack.BlobBaseFee, shouldBePrague)
	}

}

func TestParseL1InfoDepositTxData(t *testing.T) {
	randomSeqNr := func(rng *rand.Rand) uint64 {
		return rng.Uint64()
	}
	// Go 1.18 will have native fuzzing for us to use, until then, we cover just the below cases
	cases := []infoTest{
		{"random", testutils.MakeBlockInfo(nil), randomL1Cfg, randomSeqNr},
		{"zero basefee", testutils.MakeBlockInfo(func(l *testutils.MockBlockInfo) {
			l.InfoBaseFee = new(big.Int)
		}), randomL1Cfg, randomSeqNr},
		{"zero time", testutils.MakeBlockInfo(func(l *testutils.MockBlockInfo) {
			l.InfoTime = 0
		}), randomL1Cfg, randomSeqNr},
		{"zero num", testutils.MakeBlockInfo(func(l *testutils.MockBlockInfo) {
			l.InfoNum = 0
		}), randomL1Cfg, randomSeqNr},
		{"zero seq", testutils.MakeBlockInfo(nil), randomL1Cfg, func(rng *rand.Rand) uint64 {
			return 0
		}},
		{"all zero", func(rng *rand.Rand) *testutils.MockBlockInfo {
			return &testutils.MockBlockInfo{InfoBaseFee: new(big.Int)}
		}, randomL1Cfg, func(rng *rand.Rand) uint64 {
			return 0
		}},
	}
	var rollupCfg rollup.Config
	for i, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			rng := rand.New(rand.NewSource(int64(1234 + i)))
			info := testCase.mkInfo(rng)
			l1Cfg := testCase.mkL1Cfg(rng, info)
			seqNr := testCase.seqNr(rng)
			depTx, err := L1InfoDeposit(&rollupCfg, params.MergedTestChainConfig, l1Cfg, seqNr, info, 0)
			require.NoError(t, err)
			res, err := L1BlockInfoFromBytes(&rollupCfg, info.Time(), depTx.Data)
			require.NoError(t, err, "expected valid deposit info")
			assert.Equal(t, res.Number, info.NumberU64())
			assert.Equal(t, res.Time, info.Time())
			assert.True(t, res.BaseFee.Sign() >= 0)
			assert.Equal(t, res.BaseFee.Bytes(), info.BaseFee().Bytes())
			assert.Equal(t, res.BlockHash, info.Hash())
			assert.Equal(t, res.SequenceNumber, seqNr)
			assert.Equal(t, res.BatcherAddr, l1Cfg.BatcherAddr)
			assert.Equal(t, res.L1FeeOverhead, l1Cfg.Overhead)
			assert.Equal(t, res.L1FeeScalar, l1Cfg.Scalar)
		})
	}
	t.Run("no data", func(t *testing.T) {
		_, err := L1BlockInfoFromBytes(&rollupCfg, 0, nil)
		assert.Error(t, err)
	})
	t.Run("not enough data", func(t *testing.T) {
		_, err := L1BlockInfoFromBytes(&rollupCfg, 0, []byte{1, 2, 3, 4})
		assert.Error(t, err)
	})
	t.Run("too much data", func(t *testing.T) {
		_, err := L1BlockInfoFromBytes(&rollupCfg, 0, make([]byte, 4+32+32+32+32+32+1))
		assert.Error(t, err)
	})
	t.Run("invalid selector", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		info := testutils.MakeBlockInfo(nil)(rng)
		depTx, err := L1InfoDeposit(&rollupCfg, params.MergedTestChainConfig, randomL1Cfg(rng, info), randomSeqNr(rng), info, 0)
		require.NoError(t, err)
		_, err = crand.Read(depTx.Data[0:4])
		require.NoError(t, err)
		_, err = L1BlockInfoFromBytes(&rollupCfg, info.Time(), depTx.Data)
		require.ErrorContains(t, err, "function signature")
	})
	t.Run("regolith", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		info := testutils.MakeBlockInfo(nil)(rng)
		zero := uint64(0)
		rollupCfg := rollup.Config{
			RegolithTime: &zero,
		}
		depTx, err := L1InfoDeposit(&rollupCfg, params.MergedTestChainConfig, randomL1Cfg(rng, info), randomSeqNr(rng), info, 0)
		require.NoError(t, err)
		require.False(t, depTx.IsSystemTransaction)
		require.Equal(t, depTx.Gas, uint64(RegolithSystemTxGas))
	})
	t.Run("ecotone", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		info := testutils.MakeBlockInfo(nil)(rng)
		zero := uint64(0)
		rollupCfg := rollup.Config{
			RegolithTime: &zero,
			EcotoneTime:  &zero,
		}
		depTx, err := L1InfoDeposit(&rollupCfg, params.MergedTestChainConfig, randomL1Cfg(rng, info), randomSeqNr(rng), info, 1)
		require.NoError(t, err)
		require.False(t, depTx.IsSystemTransaction)
		require.Equal(t, depTx.Gas, uint64(RegolithSystemTxGas))
		require.Equal(t, L1InfoEcotoneLen, len(depTx.Data))
	})
	t.Run("first-block ecotone", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		info := testutils.MakeBlockInfo(nil)(rng)
		zero := uint64(2)
		rollupCfg := rollup.Config{
			RegolithTime: &zero,
			EcotoneTime:  &zero,
			BlockTime:    2,
		}
		depTx, err := L1InfoDeposit(&rollupCfg, params.MergedTestChainConfig, randomL1Cfg(rng, info), randomSeqNr(rng), info, 2)
		require.NoError(t, err)
		require.False(t, depTx.IsSystemTransaction)
		require.Equal(t, depTx.Gas, uint64(RegolithSystemTxGas))
		require.Equal(t, L1InfoBedrockLen, len(depTx.Data))
	})
	t.Run("genesis-block ecotone", func(t *testing.T) {
		rng := rand.New(rand.NewSource(1234))
		info := testutils.MakeBlockInfo(nil)(rng)
		zero := uint64(0)
		rollupCfg := rollup.Config{
			RegolithTime: &zero,
			EcotoneTime:  &zero,
			BlockTime:    2,
		}
		depTx, err := L1InfoDeposit(&rollupCfg, params.MergedTestChainConfig, randomL1Cfg(rng, info), randomSeqNr(rng), info, 0)
		require.NoError(t, err)
		require.False(t, depTx.IsSystemTransaction)
		require.Equal(t, depTx.Gas, uint64(RegolithSystemTxGas))
		require.Equal(t, L1InfoEcotoneLen, len(depTx.Data))
	})
}
