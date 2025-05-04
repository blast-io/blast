package rollup

import (
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/log"
)

// Taiga changes the max sequencer drift to a protocol constant. It was previously configurable via
// the rollup config.
// From Taiga, the max sequencer drift for a given block timestamp should be learned via the
// ChainSpec instead of reading the rollup configuration field directly.
const maxSequencerDriftTaiga = 1800

type ForkName string

const (
	Bedrock  ForkName = "bedrock"
	Regolith ForkName = "regolith"
	Canyon   ForkName = "canyon"
	Delta    ForkName = "delta"
	Ecotone  ForkName = "ecotone"
	Taiga    ForkName = "taiga"
	Fjord    ForkName = "fjord"
	// ADD NEW FORKS TO AllForks BELOW!
	None ForkName = "none"
)

var AllForks = []ForkName{
	Bedrock,
	Regolith,
	Canyon,
	Delta,
	Ecotone,
	Taiga,
	Fjord,
	// ADD NEW FORKS HERE!
}

var nextFork = func() map[ForkName]ForkName {
	m := make(map[ForkName]ForkName, len(AllForks))
	for i, f := range AllForks {
		if i == len(AllForks)-1 {
			m[f] = None
			break
		}
		m[f] = AllForks[i+1]
	}
	return m
}()

type ChainSpec struct {
	config      *Config
	currentFork ForkName
}

func NewChainSpec(config *Config) *ChainSpec {
	return &ChainSpec{config: config}
}

func (s *ChainSpec) IsFeatMaxSequencerDriftConstant(t uint64) bool {
	return s.config.IsTaiga(t)
}

// MaxSequencerDrift returns the maximum sequencer drift for the given block timestamp. Until Taiga,
// this was a rollup configuration parameter. Since Taiga, it is a constant, so its effective value
// should always be queried via the ChainSpec.
func (s *ChainSpec) MaxSequencerDrift(t uint64) uint64 {
	if s.IsFeatMaxSequencerDriftConstant(t) {
		return maxSequencerDriftTaiga
	}
	return s.config.MaxSequencerDrift
}

func (s *ChainSpec) CheckForkActivation(log log.Logger, block eth.L2BlockRef) {

	if s.currentFork == "" {
		// Initialize currentFork if it is not set yet
		s.currentFork = Bedrock
		if s.config.IsRegolith(block.Time) {
			s.currentFork = Regolith
		}
		if s.config.IsCanyon(block.Time) {
			s.currentFork = Canyon
		}
		if s.config.IsDelta(block.Time) {
			s.currentFork = Delta
		}
		if s.config.IsEcotone(block.Time) {
			s.currentFork = Ecotone
		}
		if s.config.IsTaiga(block.Time) {
			s.currentFork = Taiga
		}
		if s.config.IsFjord(block.Time) {
			s.currentFork = Fjord
		}
		log.Info("Current hardfork version detected", "forkName", s.currentFork)
		return
	}

	foundActivationBlock := false

	switch nextFork[s.currentFork] {
	case Regolith:
		foundActivationBlock = s.config.IsRegolithActivationBlock(block.Time)
	case Canyon:
		foundActivationBlock = s.config.IsCanyonActivationBlock(block.Time)
	case Delta:
		foundActivationBlock = s.config.IsDeltaActivationBlock(block.Time)
	case Ecotone:
		foundActivationBlock = s.config.IsEcotoneActivationBlock(block.Time)
	case Taiga:
		foundActivationBlock = s.config.IsTaigaActivationBlock(block.Time)
	case Fjord:
		foundActivationBlock = s.config.IsFjordActivationBlock(block.Time)
	}

	if foundActivationBlock {
		s.currentFork = nextFork[s.currentFork]
		log.Info("Detected hardfork activation block", "forkName", s.currentFork, "timestamp", block.Time, "blockNum", block.Number, "hash", block.Hash)
	}
}
