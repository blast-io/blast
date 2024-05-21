package genesis

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/immutables"
	"github.com/ethereum-optimism/optimism/op-chain-ops/squash"
	"github.com/ethereum-optimism/optimism/op-chain-ops/state"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// BuildL2DeveloperGenesis will build the L2 genesis block.
func BuildL2Genesis(config *DeployConfig, l1StartBlock *types.Block) (*core.Genesis, error) {
	genspec, err := NewL2Genesis(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	db := state.NewMemoryStateDB(genspec)
	if config.FundDevAccounts {
		log.Info("Funding developer accounts in L2 genesis")
		FundDevAccounts(db)
	}

	SetPrecompileBalances(db)

	// Set the Blast precompile's balance
	SetPrecompileBalance(db, common.BytesToAddress([]byte{1, 0}))

	storage, err := NewL2StorageConfig(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	immutable, err := NewL2ImmutableConfig(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	// Set up base optimism proxies
	err = setProxies(db, predeploys.ProxyAdminAddr, BigL2PredeployNamespace, 2048)
	if err != nil {
		return nil, err
	}
	// set up blast proxies
	err = setProxies(db, predeploys.ProxyAdminAddr, BlastBigL2PredeployNamespace, 2048)
	if err != nil {
		return nil, err
	}

	// Set up the implementations
	deployResults, err := immutables.BuildOptimism(immutable)
	if err != nil {
		return nil, err
	}
	for name, predeploy := range predeploys.Predeploys {
		addr := *predeploy
		if addr == predeploys.GovernanceTokenAddr && !config.EnableGovernance {
			// there is no governance token configured, so skip the governance token predeploy
			log.Warn("Governance is not enabled, skipping governance token predeploy.")
			continue
		}
		codeAddr := addr
		if predeploys.IsProxied(addr) {
			codeAddr, err = AddressToCodeNamespace(addr)
			if err != nil {
				return nil, fmt.Errorf("error converting to code namespace: %w", err)
			}
			db.CreateAccount(codeAddr)
			db.SetState(addr, ImplementationSlot, eth.AddressAsLeftPaddedHash(codeAddr))
			log.Info("Set proxy", "name", name, "address", addr, "implementation", codeAddr)
		} else if db.Exist(addr) {
			db.DeleteState(addr, AdminSlot)
		}

		if err := setupPredeploy(db, deployResults, storage, name, addr, codeAddr); err != nil {
			return nil, err
		}

		db.SetFlags(addr, 1)

		code := db.GetCode(codeAddr)
		if len(code) == 0 {
			return nil, fmt.Errorf("code not set for %s", name)
		}
	}

	// custom blast behaviors
	// set flag for weth rebasing contract to automatic on genesis
	db.SetFlags(predeploys.WETHRebasingAddr, 0)

	// deploy create 2 contract
	err = deployCreateTwo(db)
	if err != nil {
		return nil, err
	}

	if err := PerformUpgradeTxs(db); err != nil {
		return nil, fmt.Errorf("failed to perform upgrade txs: %w", err)
	}

	return db.Genesis(), nil
}

func PerformUpgradeTxs(db *state.MemoryStateDB) error {
	// Only the Ecotone upgrade is performed with upgrade-txs.
	if !db.Genesis().Config.IsEcotone(db.Genesis().Timestamp) {
		return nil
	}
	sim := squash.NewSimulator(db)
	ecotone, err := derive.EcotoneNetworkUpgradeTransactions()
	if err != nil {
		return fmt.Errorf("failed to build ecotone upgrade txs: %w", err)
	}
	if err := sim.AddUpgradeTxs(ecotone); err != nil {
		return fmt.Errorf("failed to apply ecotone upgrade txs: %w", err)
	}
	return nil
}
