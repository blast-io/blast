package actions

import (
	"blast/blockchain"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-e2e/files"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func TestBuildBlastPlugin(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	require.NoError(t, buildPlugin(log, blastPluginOpts))
}

func TestLoadBlastPlugin(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	require.NoError(t, buildPlugin(log, blastPluginOpts))

	blastchain, _, err := loadPlugin(log, blastPluginOpts)
	require.NoError(t, err)
	mnemonicCfg := e2eutils.DefaultMnemonicConfig
	secrets, err := mnemonicCfg.Secrets()

	require.NoError(t, err)
	addresses := secrets.Addresses()
	initBal := e2eutils.Ether(50)
	extraAllocs := map[string]*big.Int{}
	for _, addr := range addresses.All() {
		extraAllocs[addr.Hex()] = initBal
	}

	result := blastchain.NewChain(&blockchain.NewChainStartingArgs{
		//		SerializedGenesis: files.BlastMainnetGenesis,
		ExtraAllocs: extraAllocs,
	})

	require.NoError(t, result.Err)

	endpoint, err := blastchain.WSEndpoint()
	require.NoError(t, err)
	log.Info("now have handle on blast subproc", "ws", endpoint)
	minerClient, err := ethclient.Dial(endpoint)
	require.NoError(t, err)
	miner := NewPluginBackedMiner(t, log, blastchain, 2)

	for _, addr := range addresses.All() {
		checkBal, err := minerClient.BalanceAt(t.Ctx(), addr, nil)
		require.NoError(t, err)
		require.Condition(t,
			func() bool { return checkBal.Cmp(initBal) == 0 },
			"starting balance not as expected",
		)
	}

	chainID, err := minerClient.ChainID(t.Ctx())
	require.NoError(t, err)
	signer := types.LatestSignerForChainID(chainID)
	miner.ActL1StartBlock(12)(t)
	currentBlock, err := minerClient.BlockByNumber(t.Ctx(), nil)
	require.NoError(t, err)
	baseGas := new(big.Int).Add(currentBlock.BaseFee(), big.NewInt(2*params.GWei))
	sendTx := types.MustSignNewTx(secrets.Alice, signer, &types.LegacyTx{
		To:       &addresses.Bob,
		GasPrice: baseGas,
		Gas:      150_000,
		Value:    e2eutils.Ether(1),
	})

	require.NoError(t, minerClient.SendTransaction(t.Ctx(), sendTx))
	pool := map[string]any{}
	require.NoError(t, minerClient.Client().CallContext(t.Ctx(), &pool, "txpool_content"))
	log.Info("checking tx pool content", "pool-dump", pool)
	miner.ActL1IncludeTxByHash(addresses.Alice, sendTx.Hash())(t)
	log.Info("included the tx, now ending the block")
	blk := miner.ActL1EndBlock(t)
	hdr, txs := blk.Header(), blk.Transactions()
	log.Info("finished lifecycle", "block-num", hdr.Number, "tx-count", len(txs))
}

func TestSepoliaPlugin(gt *testing.T) {
	t := NewDefaultTesting(gt)
	log := testlog.Logger(t, log.LvlDebug)
	require.NoError(t, buildPlugin(log, blastPluginOpts))
	blastchain, pluginClient, err := loadPlugin(log, blastPluginOpts)
	t.Cleanup(func() {
		// Must do this
		blastchain.Close()
		pluginClient.Kill()
	})

	require.NoError(t, err)
	jwtPath := e2eutils.WriteDefaultJWT(t)

	result := blastchain.NewChain(&blockchain.NewChainStartingArgs{
		SerializedGenesis:  files.BlastGenesis,
		IncludeCatalystAPI: true,
		JWTFilePath:        jwtPath,
		AuthPort:           8551,
		UseDatadir:         "/Users/europablast/blast-work/local-sepolia-rpc-node/plugin-chain",
	})
	log.Info("what is it", "t", result)

	require.NoError(t, result.Err)
	replayer, err := NewReplayer(t.Ctx(), log, ReplayConfig{
		// assume ssh forwarded local chain for now
		ExistingChainRPC: "http://localhost:8545",
	}, blastchain, pluginClient, jwtPath)
	require.NoError(t, err)

	require.NoError(t, replayer.Start(t.Ctx()))
	replayer.KeepImporting(t.Ctx(), 50_000)
	time.Sleep(time.Minute * 10)

	// for i := 0; i < 10; i++ {
	// 	require.NoError(t, replayer.AddNextBlock(t.Ctx()))
	// }

}
