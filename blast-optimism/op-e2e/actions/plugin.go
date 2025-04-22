package actions

import (
	"blast/blockchain"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/stretchr/testify/require"
)

type PluginBackedMiner struct {
	blockchain.Chain
	l1Building bool
	// per account, how many txs from the pool were already included in the block, since the pool is lagging
	pendingIndices map[common.Address]uint64
	l1Signer       types.Signer
	t              Testing
	handle         *ethclient.Client
}

type ChainRunner interface {
	ActL1StartBlock(timeDelta uint64) Action
	ActL1IncludeTx(from common.Address) Action
	ActL1IncludeTxByHash(sender common.Address, txHash common.Hash) Action
	ActL1EndBlock(t Testing) *types.Block
	ActL1SetFeeRecipient(common.Address)
	EthClient() *ethclient.Client
	Close() error
	RPCEndPoint() string
	BlobStore() derive.L1BlobsFetcher
	ActEmptyBlock() *types.Block
}

func NewPluginBackedMiner(
	t Testing, log log.Logger, chain blockchain.Chain,
) ChainRunner {
	p := &PluginBackedMiner{
		Chain:          chain,
		t:              t,
		pendingIndices: make(map[common.Address]uint64),
	}

	p.handle = p.EthClient()
	return p
}

func (s *PluginBackedMiner) ActEmptyBlock() *types.Block {
	s.ActL1StartBlock(12)(s.t)
	return s.ActL1EndBlock(s.t)
}

func (s *PluginBackedMiner) ActL1IncludeTx(from common.Address) Action {
	return func(t Testing) {
		if !s.l1Building {
			t.InvalidAction("no tx inclusion when not building l1 block")
			return
		}
		getPendingIndex := func(from common.Address) uint64 {
			return s.pendingIndices[from]
		}

		contentFrom := func(who common.Address) ([]*types.Transaction, []*types.Transaction) {
			pending, queued := []*types.Transaction{}, []*types.Transaction{}

			m := map[string]map[string]*types.Transaction{}

			require.NoError(s.t, s.handle.Client().Call(&m, "txpool_contentFrom", who))
			_pending, _queued := m["pending"], m["queued"]
			for _, k := range _pending {
				pending = append(pending, k)
			}

			for _, k := range _queued {
				queued = append(queued, k)
			}

			//			fmt.Println(_pending)
			return pending, queued
		}

		tx := firstValidTx(t, from, getPendingIndex, contentFrom, s.EthClient().NonceAt)
		s.IncludeTx(t, tx)
		s.pendingIndices[from] = s.pendingIndices[from] + 1 // won't retry the tx
	}
}

func (s *PluginBackedMiner) BlobStore() derive.L1BlobsFetcher {
	return nil
}

func (s *PluginBackedMiner) ActL1SetFeeRecipient(addr common.Address) {
	require.NoError(s.t, s.Chain.SetFeeRecipient(addr.Hex()))
}

func (s *PluginBackedMiner) RPCEndPoint() string {
	ws, err := s.WSEndpoint()
	require.NoError(s.t, err)
	return ws
}

func (s *PluginBackedMiner) EthClient() *ethclient.Client {
	endpoint, err := s.WSEndpoint()
	require.NoError(s.t, err)
	minerClient, err := ethclient.Dial(endpoint)
	require.NoError(s.t, err)
	return minerClient
}

func (s *PluginBackedMiner) ActL1StartBlock(timeDelta uint64) Action {
	return func(t Testing) {
		if s.l1Building {
			t.InvalidAction("not valid if we already started building a block")
		}
		if timeDelta == 0 {
			t.Fatalf("invalid time delta: %d", timeDelta)
		}
		if err := s.StartBlock(timeDelta); err != nil {
			t.Fatalf("start block died %v", err)
		}
		s.l1Building = true
	}
}

// yes should do the plugin threading it
func (s *PluginBackedMiner) ActL1IncludeTxByHash(from common.Address, txHash common.Hash) Action {
	return func(t Testing) {
		if !s.l1Building {
			t.InvalidAction("no tx inclusion when not building l1 block")
			return
		}
		// TODO come back to this

		// tx := s.Eth.TxPool().Get(txHash)
		// require.NotNil(t, tx, "cannot find tx %s", txHash)
		// s.IncludeTx(t, tx)
		// from, err := s.l1Signer.Sender(tx)
		// require.NoError(t, err)
		s.pendingIndices[from] = s.pendingIndices[from] + 1 // won't retry the tx

		require.NoError(t, s.IncludeTxByHash(txHash.Hex()))
	}
}

func (s *PluginBackedMiner) IncludeTx(t Testing, tx *types.Transaction) {
	//
}

func (s *PluginBackedMiner) ActL1EndBlock(t Testing) *types.Block {
	if !s.l1Building {
		t.InvalidAction("cannot end L1 block when not building block")
		return nil
	}

	s.l1Building = false
	//	result, err :=
	output := s.EndBlock()
	var err error

	// if output.Err.Message != "" {
	// 	err = errors.New(output.Err.Message)
	// }

	if err != nil {
		t.Fatalf("problem making block %v", err)
	}

	result := output.SerializedBlock

	var deserialized struct {
		Hdr      *types.Header
		Txs      types.Transactions
		Receipts types.Receipts
	}

	if err := json.Unmarshal(result, &deserialized); err != nil {
		t.Fatal("could not deserialize block %v raw %v", err, string(result))
	}
	// return types.NewBlockWithHeader(deserialized.Hdr).WithBody(types.Body{
	// 	Transactions: deserialized.Txs,
	// })

	return types.NewBlock(
		deserialized.Hdr, deserialized.Txs, nil, deserialized.Receipts, trie.NewStackTrie(nil),
	)
}

type pluginBuildOpt struct {
	pluginMapName, dir, mainPath, outPath, execPath string
	buildWithRace                                   bool
}

var (
	gethL1PluginOpts = pluginBuildOpt{
		pluginMapName: "blast",
		dir:           "geth-plugin",
		mainPath:      "cmd/plugin/main.go",
		outPath:       "../op-e2e/actions/l1-geth-plugin",
	}

	blastPluginOpts = pluginBuildOpt{
		pluginMapName: "blast",
		dir:           "op-geth",
		mainPath:      "cmd/plugin/main.go",
		outPath:       "../op-e2e/actions/blast-plugin",
	}
)

func buildPlugin(lgr log.Logger, opt pluginBuildOpt) error {
	args := []string{"build"}
	if opt.buildWithRace {
		args = append(args, "-race")
	}

	args = append(args, "-o", opt.outPath, opt.mainPath)

	cmd := exec.Command("go", args...)

	current, err := os.Getwd()
	if err != nil {
		return err
	}
	root, err := config.FindMonorepoRoot(current)
	if err != nil {
		return err
	}
	cmd.Dir = filepath.Join(root, opt.dir)
	lgr.Info("building plugin using dir", "dir", cmd.Dir, "used-args", args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		lgr.Error("problem bulding blast plugin", "reason", string(out))
	}
	return err
}

func loadPlugin(lgr log.Logger, opt pluginBuildOpt) (blockchain.Chain, *plugin.Client, error) {

	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}

	pth := "./" + filepath.Base(opt.outPath)
	current, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	lgr.Info("using plugin", "path", pth, "current", current)

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{opt.pluginMapName: &blockchain.ChainPlugin{}}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(pth),
		Logger: hclog.New(&hclog.LoggerOptions{
			Output:          os.Stderr,
			Color:           hclog.ForceColor,
			IncludeLocation: false,
			DisableTime:     true,
			Level:           hclog.Trace,
			ColorHeaderOnly: true,
		}),
	})

	// Connect via RPC
	rpcClient, err := client.Client()

	if err != nil {
		return nil, nil, err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(opt.pluginMapName)
	if err != nil {
		return nil, nil, err
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	blastchain := raw.(blockchain.Chain)

	return blastchain, client, nil
}

func compileAndLoadPlugin(t Testing, lgr log.Logger, opt pluginBuildOpt) blockchain.Chain {
	require.NoError(t, buildPlugin(lgr, opt))
	c, _, err := loadPlugin(lgr, opt)
	require.NoError(t, err)
	return c
}
