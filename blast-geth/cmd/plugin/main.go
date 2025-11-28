package main

import (
	"blast/blockchain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"path"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/catalyst"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/hashicorp/go-plugin"
)

type workState struct {
	l1BuildingHeader       *types.Header
	l1BuildingState        *state.StateDB
	l1Receipts             []*types.Receipt
	L1Transactions         []*types.Transaction
	pendingIndices         map[common.Address]uint64
	l1TxFailed             []*types.Transaction // log of failed transactions which could not be included
	l1BuildingBlobSidecars []*types.BlobTxSidecar
	L1GasPool              *core.GasPool
}

type pluginBlast struct {
	log          log.Logger
	cfg          blockchain.NewChainStartingArgs
	node         *node.Node
	Eth          *eth.Ethereum
	prefCoinbase common.Address

	// L1 evm / chain
	l1Chain    *core.BlockChain
	l1Database ethdb.Database
	l1Cfg      *core.Genesis
	l1Signer   types.Signer
	// current work state - only non-nil when making a block
	s *workState
}

// force compliation fail
var (
	_ blockchain.Chain = (*pluginBlast)(nil)
)

// NOTE NOTE NOTE all errors returned MUST be wrapped with plugin.NewBasicError!

func (p *pluginBlast) InitExtraConfigs(cfg []byte) error {
	return nil
}

func (p *pluginBlast) IncludeTxByHash(hexHash string) error {
	hsh := common.HexToHash(hexHash)
	if hsh == (common.Hash{}) {
		return plugin.NewBasicError(ErrDidNotHaveTx)
	}
	tx := p.Eth.TxPool().Get(hsh)
	if tx == nil {
		return plugin.NewBasicError(fmt.Errorf("cannot find tx %s", hsh))
	}
	p.log.Info("had tx requested - now will process tx")
	if err := p.includeTx(tx); err != nil {
		return plugin.NewBasicError(err)
	}
	from, err := p.l1Signer.Sender(tx)
	if err != nil {
		return plugin.NewBasicError(err)
	}
	p.s.pendingIndices[from] = p.s.pendingIndices[from] + 1 // won't retry the tx
	return nil
}

func (p *pluginBlast) includeTx(tx *types.Transaction) error {
	from, err := p.l1Signer.Sender(tx)
	if err != nil {
		return plugin.NewBasicError(err)
	}
	p.log.Info("including tx", "nonce", tx.Nonce(), "from", from, "to", tx.To())
	if tx.Gas() > p.s.l1BuildingHeader.GasLimit {
		return plugin.NewBasicError(fmt.Errorf("tx consumes %d gas, more than available in L1 block %d", tx.Gas(), p.s.l1BuildingHeader.GasLimit))
	}

	if tx.Gas() > uint64(*p.s.L1GasPool) {
		return plugin.NewBasicError(fmt.Errorf("action takes too much gas: %d, only have %d", tx.Gas(), uint64(*p.s.L1GasPool)))
	}

	p.s.l1BuildingState.SetTxContext(tx.Hash(), len(p.s.L1Transactions))
	p.log.Info("about to apply tx", "hsh", tx.Hash())
	st := time.Now()
	receipt, err := core.ApplyTransaction(
		p.l1Cfg.Config, p.l1Chain, &p.s.l1BuildingHeader.Coinbase,
		p.s.L1GasPool, p.s.l1BuildingState, p.s.l1BuildingHeader, tx.WithoutBlobTxSidecar(),
		&p.s.l1BuildingHeader.GasUsed, *p.l1Chain.GetVMConfig(),
	)

	p.log.Info("applied tx", "hsh", tx.Hash().Hex(), "took", time.Since(st))

	if err != nil {
		p.s.l1TxFailed = append(p.s.l1TxFailed, tx)
		return plugin.NewBasicError(fmt.Errorf("failed to apply transaction to L1 block (tx %d): %v", len(p.s.L1Transactions), err))
	}

	p.s.l1Receipts = append(p.s.l1Receipts, receipt)
	p.s.L1Transactions = append(p.s.L1Transactions, tx.WithoutBlobTxSidecar())

	if tx.Type() == types.BlobTxType {
		if !p.l1Cfg.Config.IsCancun(p.s.l1BuildingHeader.Number, p.s.l1BuildingHeader.Time) {
			return plugin.NewBasicError(ErrNotCancunCantDoBlob)
		}
		sidecar := tx.BlobTxSidecar()
		if sidecar != nil {
			p.s.l1BuildingBlobSidecars = append(p.s.l1BuildingBlobSidecars, sidecar)
		}
		*p.s.l1BuildingHeader.BlobGasUsed += receipt.BlobGasUsed
	}

	return nil
}

func (p *pluginBlast) Close() error {
	p.log.Info("blast plugin about to call close on node")
	//	finished := make(chan error)

	// silly workaround because gob encoding, just leave as is whatever
	if err := plugin.NewBasicError(p.node.Close()); err != nil {
		//		finished <- err
		return err
	}

	if err := p.Eth.ChainDb().Close(); err != nil {
		//		finished <- plugin.NewBasicError(err)
		//		return err
		return plugin.NewBasicError(err)
	}

	//	finished <- nil

	p.log.Info("closed node and chaindb")
	//	err := <-finished

	//	return err
	return nil
}

func (p *pluginBlast) EndBlock() blockchain.NewBlockOrError {
	p.s.l1BuildingHeader.GasUsed = p.s.l1BuildingHeader.GasLimit - uint64(*p.s.L1GasPool)
	p.s.l1BuildingHeader.Root = p.s.l1BuildingState.IntermediateRoot(p.l1Cfg.Config.IsEIP158(p.s.l1BuildingHeader.Number))

	var withdrawals []*types.Withdrawal
	if p.l1Cfg.Config.IsShanghai(p.s.l1BuildingHeader.Number, p.s.l1BuildingHeader.Time) {
		withdrawals = make([]*types.Withdrawal, 0)
	}

	block := types.NewBlockWithWithdrawals(
		p.s.l1BuildingHeader, p.s.L1Transactions, nil, p.s.l1Receipts, withdrawals, trie.NewStackTrie(nil),
	)

	// Write state changes to db
	root, err := p.s.l1BuildingState.Commit(p.s.l1BuildingHeader.Number.Uint64(), p.l1Cfg.Config.IsEIP158(p.s.l1BuildingHeader.Number))
	if err != nil {
		return blockchain.NewBlockOrError{Err: plugin.NewBasicError(err)}
	}
	if err := p.s.l1BuildingState.Database().TrieDB().Commit(root, false); err != nil {
		return blockchain.NewBlockOrError{Err: plugin.NewBasicError(err)}
	}

	// TODO what about these?
	// now that the blob txs are in a canonical block, flush them to the blob store
	// for _, sidecar := range s.l1BuildingBlobSidecars {
	// 	for i, h := range sidecar.BlobHashes() {
	// 		blob := (*eth.Blob)(&sidecar.Blobs[i])
	// 		indexedHash := eth.IndexedBlobHash{Index: uint64(i), Hash: h}
	// 		s.blobStore.StoreBlob(block.Time(), indexedHash, blob)
	// 	}
	// }

	_, err = p.l1Chain.InsertChain(types.Blocks{block})
	if err != nil {
		return blockchain.NewBlockOrError{Err: plugin.NewBasicError(err)}
	}

	serialized, err := json.Marshal(struct {
		Hdr      *types.Header
		Txs      types.Transactions
		Receipts types.Receipts
	}{Hdr: block.Header(), Txs: block.Transactions()})

	return blockchain.NewBlockOrError{SerializedBlock: serialized}
}

// TODO error if its not started yet?
func (p *pluginBlast) WSEndpoint() (string, error) {
	route := p.node.WSEndpoint()
	p.log.Info("sneding back endpoint for ws", "endpoint", route)
	return route, nil
}

func (p *pluginBlast) AuthEndpoint() (string, error) {
	if p.cfg.CatalystAuthEnabled {
		return p.node.HTTPAuthEndpoint(), nil
	}
	return p.node.HTTPEndpoint(), nil
}

func (p *pluginBlast) SetFeeRecipient(addr string) error {
	coinbase := common.HexToAddress(addr)
	if coinbase == (common.Address{}) {
		return plugin.NewBasicError(fmt.Errorf("fee recipient set to empty addr"))
	}
	p.prefCoinbase = coinbase
	return nil
}

var (
	ErrEmptyAddr           = errors.New("empty addr")
	ErrDidNotHaveTx        = errors.New("did not have tx by hash")
	ErrNotCancunCantDoBlob = errors.New("not cancun yet so cant do blob")
)

func (p *pluginBlast) NewChain(startingArgs *blockchain.NewChainStartingArgs) blockchain.NewChainOrError {
	p.cfg = *startingArgs
	var gen *core.Genesis

	p.log.Info(
		"making new chain in blast-geth as plugin",
		"is-using-custom-genesis", len(startingArgs.SerializedGenesis) > 0,
	)

	if startingArgs.AssumeMainnet {
		// Then assume eth mainnet
		gen = core.DefaultGenesisBlock()
	} else if len(startingArgs.SerializedGenesis) > 0 {
		if err := json.Unmarshal(startingArgs.SerializedGenesis, &gen); err != nil {
			return blockchain.NewChainOrError{
				Err: plugin.NewBasicError(fmt.Errorf("problem deserializing genesis %w", err)),
			}
		}
		blockZeroHash := gen.ToBlock().Hash()
		p.log.Info(
			"custom genesis provided blob schedule",
			"assuming-block-zero-hash", blockZeroHash,
			"schedule", gen.Config.BlobScheduleConfig,
		)

	} else if startingArgs.UseDatadir != "" {
		//
	} else if startingArgs.Faucet != "" {
		p.log.Info("using developer genesis block", "faucet", startingArgs.Faucet)
		gen = core.DeveloperGenesisBlock(30_000_000, common.HexToAddress(startingArgs.Faucet))

		gen.Config.CancunTime = startingArgs.WhenActivateCancun
		gen.Config.PragueTime = startingArgs.WhenActivatePrague
		gen.Config.TerminalTotalDifficultyPassed = true

		for _addr, amt := range startingArgs.ExtraAllocs {
			addr := common.HexToAddress(_addr)
			if addr == (common.Address{}) {
				return blockchain.NewChainOrError{Err: plugin.NewBasicError(ErrEmptyAddr)}
			}
			gen.Alloc[addr] = core.GenesisAccount{
				Balance: amt,
			}
		}

	}

	minerCfg := miner.DefaultConfig

	if startingArgs.MinerRecommit > 0 {
		minerCfg.Recommit = startingArgs.MinerRecommit
	}

	if startingArgs.MinerNewPayloadTimeout > 0 {
		minerCfg.NewPayloadTimeout = startingArgs.MinerNewPayloadTimeout
	}

	ethCfg := &ethconfig.Config{
		NetworkId:                 gen.Config.ChainID.Uint64(),
		Genesis:                   gen,
		RollupDisableTxPoolGossip: true,
		StateScheme:               rawdb.HashScheme,
		TrieTimeout:               5 * time.Minute,
		DatabaseCache:             3096,
		TrieCleanCache:            3096,
		TrieDirtyCache:            3096,
		Miner:                     minerCfg,
	}

	namespaces := []string{
		"debug", "admin", "eth", "txpool", "net",
		"rpc", "web3", "personal", "engine", "blast",
	}

	nodeCfg := &node.Config{
		Name:     "plugin-chain",
		WSHost:   "127.0.0.1",
		WSPort:   startingArgs.WSPort,
		HTTPHost: "127.0.0.1",
		HTTPPort: 0,
		// AuthAddr:    "127.0.0.1",
		AuthPort:            startingArgs.AuthPort,
		WSModules:           namespaces,
		HTTPModules:         namespaces,
		P2P:                 p2p.Config{NoDiscovery: true, NoDial: true},
		PebbleFormatVersion: pebble.FormatNewest,
	}
	if j := startingArgs.JWTFilePath; j != "" {
		nodeCfg.JWTSecret = j
	}

	if d := startingArgs.UseDatadir; d != "" {
		nodeCfg.Name = path.Base(d)
		nodeCfg.DataDir = d
	}

	n, err := node.New(nodeCfg)
	if err != nil {
		return blockchain.NewChainOrError{Err: plugin.NewBasicError(err)}
	}
	backend, err := eth.New(n, ethCfg)
	if err != nil {
		return blockchain.NewChainOrError{Err: plugin.NewBasicError(err)}
	}
	n.RegisterAPIs(tracers.APIs(backend.APIBackend))
	filterSystem := filters.NewFilterSystem(backend.APIBackend, filters.Config{
		LogCacheSize: ethCfg.FilterLogCacheSize,
	})

	apis := []rpc.API{
		{Namespace: "eth", Service: filters.NewFilterAPI(filterSystem, false)},
		// IF YOU DO THIS THEN IT BREAKS!
		// {Namespace: "eth", Service: backend.APIBackend},
		// {Namespace: "blast", Service: backend.NewBlastAPI()},
	}

	// By doing this, we dont have to use the auth port, the regular one works fine
	if startingArgs.IncludeCatalystAPI {
		catalyst.AuthEnabled = startingArgs.CatalystAuthEnabled
		catalyst.Register(n, backend)
	}

	n.RegisterAPIs(apis)

	p.node = n
	p.Eth = backend
	p.l1Chain = backend.BlockChain()
	p.l1Database = backend.ChainDb()
	p.l1Cfg = gen
	p.l1Signer = types.LatestSigner(gen.Config)

	if err := n.Start(); err != nil {
		return blockchain.NewChainOrError{Err: plugin.NewBasicError(err)}
	}

	p.log.Info("starting chain",
		"at-block-num", backend.BlockChain().CurrentHeader().Number,
		"block-hash", backend.BlockChain().CurrentHeader().Hash().Hex(),
	)

	payload, err := json.Marshal(gen.ToBlock().Header())
	if err != nil {
		return blockchain.NewChainOrError{Err: plugin.NewBasicError(err)}
	}

	var headHash, safeHash, finalizedHash string

	if startingArgs.UseDatadir == "" {
		hsh := gen.ToBlock().Hash().Hex()
		headHash = hsh
		safeHash = hsh
		finalizedHash = hsh
	} else {
		headHash = p.l1Chain.CurrentHeader().Hash().Hex()
		safeHash = p.l1Chain.CurrentSafeBlock().Hash().Hex()
		finalizedHash = p.l1Chain.CurrentFinalBlock().Hash().Hex()
		payload, _ = json.Marshal(p.l1Chain.CurrentHeader())
	}

	return blockchain.NewChainOrError{
		SerializedHeader: payload,
		HeadHash:         headHash,
		SafeHash:         safeHash,
		FinalizedHash:    finalizedHash,
	}
}

func (p *pluginBlast) StartBlock(timeDelta uint64) error {
	parent := p.l1Chain.CurrentHeader()
	parentHash := parent.Hash()
	statedb, err := state.New(parent.Root, state.NewDatabaseWithNodeDB(p.l1Database, p.l1Chain.TrieDB()), nil)
	if err != nil {
		return plugin.NewBasicError(fmt.Errorf(
			"failed to init state db around block %s (state %s): %w", parentHash, parent.Root, err),
		)
	}
	header := &types.Header{
		ParentHash: parentHash,
		Coinbase:   p.prefCoinbase,
		Difficulty: common.Big0,
		Number:     new(big.Int).Add(parent.Number, common.Big1),
		GasLimit:   parent.GasLimit,
		Time:       parent.Time + timeDelta,
		Extra:      []byte("L1 was here"),
		MixDigest:  common.Hash{}, // TODO: maybe randomize this (prev-randao value)
	}

	if p.l1Cfg.Config.IsLondon(header.Number) {
		header.BaseFee = eip1559.CalcBaseFee(p.l1Cfg.Config, parent, header.Time)
		// At the transition, double the gas limit so the gas target is equal to the old gas limit.
		if !p.l1Cfg.Config.IsLondon(parent.Number) {
			header.GasLimit = parent.GasLimit * p.l1Cfg.Config.ElasticityMultiplier()
		}
	}

	if p.l1Cfg.Config.IsShanghai(header.Number, header.Time) {
		header.WithdrawalsHash = &types.EmptyWithdrawalsHash
	}

	if p.l1Cfg.Config.IsCancun(header.Number, header.Time) {
		header.BlobGasUsed = new(uint64)
		header.ExcessBlobGas = new(uint64)
		root := crypto.Keccak256Hash([]byte("fake-beacon-block-root"), header.Number.Bytes())
		header.ParentBeaconRoot = &root

		// Copied from op-program/client/l2/engineapi/block_processor.go
		// TODO(client-pod#826)
		// Unfortunately this is not part of any Geth environment setup,
		// we just have to apply it, like how the Geth block-builder worker does.
		context := core.NewEVMBlockContext(header, p.l1Chain, nil, p.l1Chain.Config(), statedb)
		// NOTE: Unlikely to be needed for the beacon block root, but we setup any precompile overrides anyways for forwards-compatibility
		var precompileOverrides vm.PrecompileOverrides
		if vmConfig := p.l1Chain.GetVMConfig(); vmConfig != nil && vmConfig.OptimismPrecompileOverrides != nil {
			precompileOverrides = vmConfig.OptimismPrecompileOverrides
		}
		vmenv := vm.NewEVM(context, vm.TxContext{}, statedb, p.l1Chain.Config(), vm.Config{OptimismPrecompileOverrides: precompileOverrides})
		core.ProcessBeaconBlockRoot(*header.ParentBeaconRoot, vmenv, statedb)
	}
	p.log.Info("started new block")

	p.s = &workState{
		l1BuildingHeader: header,
		l1BuildingState:  statedb,
		l1Receipts:       make([]*types.Receipt, 0),
		L1Transactions:   make([]*types.Transaction, 0),
		pendingIndices:   make(map[common.Address]uint64),
		//		l1BuildingBlobSidecars: make([]*types.BlobTxSidecar, 0),
		L1GasPool: new(core.GasPool).AddGas(header.GasLimit),
	}
	return nil
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	glogger := log.NewGlogHandler(
		log.StreamHandler(io.MultiWriter(os.Stdout, os.Stderr), log.TerminalFormat(true)),
	)
	glogger.Verbosity(log.LvlTrace)
	log.Root().SetHandler(glogger)
	log.PrintOrigins(true)

	chain := &pluginBlast{
		log: log.New("env", "blast-geth"),
	}

	pluginMap := map[string]plugin.Plugin{
		"blast": &blockchain.ChainPlugin{Impl: chain},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
