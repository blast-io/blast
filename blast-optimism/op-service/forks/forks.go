package forks

import "github.com/ethereum/go-ethereum/params"

type Blob string

const (
	Prague    Blob = "prague"
	Osaka     Blob = "osaka"
	BPO1      Blob = "bpo1"
	BPO2      Blob = "bpo2"
	BPO2Blast Blob = "bpo2blast"
	BPO3      Blob = "bpo3"
	None      Blob = ""
)

// NOTE all the forks that affect blob pricing - in order
var All = []Blob{
	Prague,
	Osaka,
	BPO1,
	BPO2,
	BPO2Blast,
	BPO3,
}

func Schedule(cfg *params.ChainConfig, name Blob) *params.BlobConfig {
	switch name {
	case Prague:
		return cfg.Blast.PragueBlobConfigOverride
	case Osaka:
		return cfg.Blast.OsakaBlobConfigOverride
	case BPO1:
		return cfg.Blast.BPO1BlobConfigOverride
	case BPO2:
		return cfg.Blast.BPO2BlobConfigOverride
	case BPO2Blast:
		return cfg.Blast.BPO2BlastBlobConfigOverride
	case BPO3:
		return cfg.Blast.BPO3BlobConfigOverride
	}

	return nil
}

var LatestBlob = All[len(All)-1]

var next = func() map[Blob]Blob {
	m := make(map[Blob]Blob, len(All))
	for i, f := range All {
		if i == len(All)-1 {
			m[f] = None
			break
		}
		m[f] = All[i+1]
	}
	return m
}()

// IsValid returns true if the provided fork is a known fork.
func IsValid(f Blob) bool {
	_, ok := next[f]
	return ok
}

// Next returns the fork that follows the provided fork, or None if it is the last.
func Next(f Blob) Blob { return next[f] }
