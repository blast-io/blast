package actions

import (
	"blast/blockchain"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/files"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-wheel/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-plugin"
)

type ReplayConfig struct {
	ExistingChainRPC    string `toml:"existing-chain-rpc"`
	Datadir             string `toml:"data-dir"`
	BuildPluginWithRace bool   `toml:"build-plugin-with-race"`
}

type Replayer struct {
	cfg                              ReplayConfig
	lgr                              log.Logger
	plugin                           blockchain.Chain
	existingChain                    *ethclient.Client
	replayChain                      *ethclient.Client
	replayChainEngine                *sources.EngineAPIClient
	rollupConfig                     *rollup.Config
	pluginClient                     *plugin.Client
	existingChainRPC, replayChainRPC client.RPC
	keepImporting                    atomic.Bool
	LastKeepImportingUtil            atomic.Uint64
	L2EngineJWTSecret                [32]byte
}

func (r *Replayer) Close(ctx context.Context) error {
	r.lgr.Info("closing replayer rpc handles")
	r.existingChain.Close()
	r.replayChain.Close()
	//	r.replayChainEngine.RPC.Close()
	r.replayChainRPC.Close()
	r.existingChainRPC.Close()
	if err := r.plugin.Close(); err != nil {
		r.lgr.Error("problem clsing plugin", "err", err)
	}
	r.pluginClient.Kill()
	return nil
}

func (r *Replayer) AddNextBlock(ctx context.Context) error {
	currentBlock, err := r.replayChain.BlockByNumber(ctx, nil)
	if err != nil {
		return err
	}
	r.lgr.Info("replay chain head", "num", currentBlock.NumberU64())

	next := new(big.Int).Add(currentBlock.Header().Number, common.Big1)
	if err := engine.CopyPayload(ctx, next.Uint64(), r.existingChainRPC, r.replayChainEngine); err != nil {
		return err
	}
	blk, err := r.existingChain.BlockByNumber(ctx, next)
	if err != nil {
		return err
	}
	return engine.SetForkchoiceByHash(
		ctx, r.replayChainEngine, blk.ParentHash(), blk.ParentHash(), blk.Hash(),
	)
}

func (r *Replayer) StopImporting(ctx context.Context) error {
	r.keepImporting.Store(false)
	return nil
}

func (r *Replayer) KeepImporting(ctx context.Context, upTo uint64) error {
	r.LastKeepImportingUtil.Store(upTo)
	currentBlock, err := r.replayChain.BlockByNumber(ctx, nil)
	if err != nil {
		return err
	}

	r.lgr.Info("replay chain head", "num", currentBlock.NumberU64())

	if h := currentBlock.Number().Uint64(); upTo <= h {
		return fmt.Errorf("up to is already behind prior chain %v", h)
	}

	if r.keepImporting.Load() {
		return fmt.Errorf("already importing")
	}

	r.keepImporting.Store(true)

	go func() {
		started := time.Now()
		var added uint64

	loop:
		for {
			if r.keepImporting.Load() == false {
				r.lgr.Info("stop importing asked",
					"added-so-far", added,
					"time-spent", time.Since(started),
				)
				return
			}

			if err := r.AddNextBlock(context.Background()); err != nil {
				r.lgr.Error("problem adding next block will sleep and try again",
					"time-spent", time.Since(started),
					"err", err,
				)
				time.Sleep(time.Second * 2)
				goto loop
			}

			currentBlock, err := r.replayChain.BlockByNumber(context.Background(), nil)
			if err != nil {
				r.lgr.Error("problem asking replaychain latest block",
					"time-spent", time.Since(started),
					"err", err,
				)
				r.keepImporting.Store(false)
				return
			}

			if currentBlock.Header().Number.Uint64() == upTo {
				r.lgr.Info("finished importing block to replay chain", "upto", upTo, "took", time.Since(started))
				r.keepImporting.Store(false)
				return
			}

			added++
		}
	}()

	return nil
}

func (r *Replayer) Start(ctx context.Context) error {
	route, err := r.plugin.WSEndpoint()
	if err != nil {
		return err
	}

	r.lgr.Info("using wsendpoing on plugin", "route", route)

	replayChainHandle, err := ethclient.DialContext(ctx, route)
	if err != nil {
		return err
	}
	r.replayChain = replayChainHandle
	blk, err := replayChainHandle.BlockByNumber(ctx, nil)
	if err != nil {
		return err
	}

	authRoute, err := r.plugin.AuthEndpoint()
	if err != nil {
		return err
	}

	r.lgr.Info("replay chain started",
		"head", blk.Number(),
		"auth-endpoint", authRoute,
		"eth-endpoint", route,
	)

	auth := rpc.WithHTTPAuth(node.NewJWTAuth(r.L2EngineJWTSecret))
	opts := []client.RPCOption{
		client.WithGethRPCOptions(auth),
		client.WithDialBackoff(10),
	}

	l2Node, err := client.NewRPC(context.Background(), r.lgr, authRoute, opts...)
	if err != nil {
		return err
	}

	existing, err := client.NewRPC(context.Background(), r.lgr, r.cfg.ExistingChainRPC)
	if err != nil {
		return err
	}
	r.replayChainRPC = l2Node
	r.existingChainRPC = existing
	r.replayChainEngine = sources.NewEngineAPIClientWithTimeout(
		r.replayChainRPC, r.lgr, r.rollupConfig, time.Second*10,
	)

	return nil
}

func NewReplayer(
	ctx context.Context,
	lgr log.Logger,
	cfg ReplayConfig,
	replayChain blockchain.Chain,
	pluginClient *plugin.Client,
	jwtFile string,
) (*Replayer, error) {
	var rollupcfg rollup.Config
	if err := json.Unmarshal(files.BlastRollup, &rollupcfg); err != nil {
		return nil, err
	}

	existingChain, err := ethclient.DialContext(ctx, cfg.ExistingChainRPC)
	if err != nil {
		return nil, err
	}

	var secret [32]byte
	jwtFile = strings.TrimSpace(jwtFile)
	if jwtFile == "" {
		return nil, fmt.Errorf("file-name of jwt secret is empty")
	}
	if data, err := os.ReadFile(jwtFile); err == nil {
		jwtSecret := common.FromHex(strings.TrimSpace(string(data)))
		if len(jwtSecret) != 32 {
			return nil, fmt.Errorf("invalid jwt secret in path %s, not 32 hex-formatted bytes", jwtFile)
		}
		copy(secret[:], jwtSecret)
	}

	return &Replayer{
		cfg:               cfg,
		lgr:               lgr,
		existingChain:     existingChain,
		plugin:            replayChain,
		rollupConfig:      &rollupcfg,
		pluginClient:      pluginClient,
		L2EngineJWTSecret: secret,
	}, nil
}
