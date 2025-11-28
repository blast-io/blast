package files

import (
	_ "embed"
)

var (
	//go:embed sepolia-genesis.json
	BlastGenesis []byte
	//go:embed sepolia-rollup.json
	BlastRollup []byte
	//go:embed mainnet-genesis.json
	BlastMainnetGenesis []byte
	//go:embed mainnet-rollup.json
	BlastMainnetRollup []byte
)
