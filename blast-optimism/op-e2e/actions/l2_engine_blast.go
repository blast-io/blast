package actions

import (
	"os"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	geth "github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
)

func NewL2EnginePebble(t Testing, log log.Logger, genesis *core.Genesis, rollupGenesisL1 eth.BlockID, jwtPath string, options ...EngineOption) (*L2Engine, string) {
	n, ethBackend, apiBackend, dbDir := newBackendPebble(t, genesis, jwtPath, options)
	engineApi := engineapi.NewL2EngineAPI(log, apiBackend, ethBackend.Downloader())
	chain := ethBackend.BlockChain()
	genesisBlock := chain.Genesis()
	eng := &L2Engine{
		log:  log,
		node: n,
		eth:  ethBackend,
		rollupGenesis: &rollup.Genesis{
			L1:     rollupGenesisL1,
			L2:     eth.BlockID{Hash: genesisBlock.Hash(), Number: genesisBlock.NumberU64()},
			L2Time: genesis.Timestamp,
		},
		l2Chain:   chain,
		l2Signer:  types.LatestSigner(genesis.Config),
		engineApi: engineApi,
	}
	// register the custom engine API, so we can serve engine requests while having more control
	// over sequencing of individual txs.
	n.RegisterAPIs([]rpc.API{
		{
			Namespace:     "engine",
			Service:       eng.engineApi,
			Authenticated: true,
		},
	})
	require.NoError(t, n.Start(), "failed to start L2 op-geth node")

	return eng, dbDir
}

func newBackendPebble(t e2eutils.TestingBase, genesis *core.Genesis, jwtPath string, options []EngineOption) (*node.Node, *geth.Ethereum, *engineApiBackend, string) {
	tmpPlace, err := os.MkdirTemp("", "pebble-geth-testing")
	require.NoError(t, err)
	ethCfg := &ethconfig.Config{
		NetworkId: genesis.Config.ChainID.Uint64(),
		Genesis:   genesis,
		NoPruning: true,
		Preimages: true,
	}
	nodeCfg := &node.Config{
		Name:                "l2-geth",
		WSHost:              "127.0.0.1",
		WSPort:              0,
		AuthAddr:            "127.0.0.1",
		AuthPort:            0,
		WSModules:           []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal"},
		HTTPModules:         []string{"debug", "admin", "eth", "txpool", "net", "rpc", "web3", "personal"},
		JWTSecret:           jwtPath,
		DBEngine:            "pebble",
		DataDir:             tmpPlace,
		PebbleFormatVersion: pebble.FormatVirtualSSTables,
	}
	for i, opt := range options {
		require.NoError(t, opt(ethCfg, nodeCfg), "engine option %d failed", i)
	}
	n, err := node.New(nodeCfg)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = n.Close()
	})
	backend, err := geth.New(n, ethCfg)
	require.NoError(t, err)
	n.RegisterAPIs(tracers.APIs(backend.APIBackend))

	chain := backend.BlockChain()
	db := backend.ChainDb()
	apiBackend := &engineApiBackend{
		BlockChain: chain,
		db:         db,
		genesis:    genesis,
	}
	return n, backend, apiBackend, tmpPlace
}

func setupSequencerTestWithPebble(t Testing, sd *e2eutils.SetupData, log log.Logger) (*L1Miner, *L2Engine, *L2Sequencer, string) {
	jwtPath := e2eutils.WriteDefaultJWT(t)
	miner := NewL1Miner(t, log, sd.L1Cfg)
	l1F, err := sources.NewL1Client(miner.RPCClient(), log, nil, sources.L1ClientDefaultConfig(sd.RollupCfg, false, sources.RPCKindStandard))
	require.NoError(t, err)
	engine, dbDir := NewL2EnginePebble(t, log, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath)
	l2Cl, err := sources.NewEngineClient(engine.RPCClient(), log, nil, sources.EngineClientDefaultConfig(sd.RollupCfg))
	require.NoError(t, err)
	sequencer := NewL2Sequencer(t, log, l1F, miner.BlobStore(), l2Cl, sd.RollupCfg, sd.L1Cfg.Config, 0)
	return miner, engine, sequencer, dbDir
}
