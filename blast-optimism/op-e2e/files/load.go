package files

import (
	_ "embed"
)

var (
	//go:embed sepolia-genesis.json
	BlastGenesis []byte
	//go:embed sepolia-rollup.json
	BlastRollup []byte
)
