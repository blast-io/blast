package immutables

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/deployer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ImmutableValues represents the values to be set in immutable code.
// The key is the name of the variable and the value is the value to set in
// immutable code.
type ImmutableValues map[string]any

// ImmutableConfig represents the immutable configuration for the L2 predeploy
// contracts.
type ImmutableConfig map[string]ImmutableValues

// Check does a sanity check that the specific values that
// Optimism uses are set inside of the ImmutableConfig.
func (i ImmutableConfig) Check() error {
	if _, ok := i["L2CrossDomainMessenger"]["otherMessenger"]; !ok {
		return errors.New("L2CrossDomainMessenger otherMessenger not set")
	}
	if _, ok := i["L2StandardBridge"]["otherBridge"]; !ok {
		return errors.New("L2StandardBridge otherBridge not set")
	}
	if _, ok := i["L2ERC721Bridge"]["messenger"]; !ok {
		return errors.New("L2ERC721Bridge messenger not set")
	}
	if _, ok := i["L2ERC721Bridge"]["otherBridge"]; !ok {
		return errors.New("L2ERC721Bridge otherBridge not set")
	}
	if _, ok := i["OptimismMintableERC721Factory"]["bridge"]; !ok {
		return errors.New("OptimismMintableERC20Factory bridge not set")
	}
	if _, ok := i["OptimismMintableERC721Factory"]["remoteChainId"]; !ok {
		return errors.New("OptimismMintableERC20Factory remoteChainId not set")
	}
	if _, ok := i["SequencerFeeVault"]["recipient"]; !ok {
		return errors.New("SequencerFeeVault recipient not set")
	}
	if _, ok := i["L1FeeVault"]["recipient"]; !ok {
		return errors.New("L1FeeVault recipient not set")
	}
	if _, ok := i["BaseFeeVault"]["recipient"]; !ok {
		return errors.New("BaseFeeVault recipient not set")
	}
	if _, ok := i["Gas"]["baseClaimRate"]; !ok {
		return errors.New("Gas baseClaimRate not set")
	}
	if _, ok := i["L2BlastBridge"]["otherBridge"]; !ok {
		return errors.New("L2BlastBridge otherBridge not set")
	}

	if _, ok := i["Shares"]["reporter"]; !ok {
		return errors.New("Shares reporter not set")
	}
	if _, ok := i["Gas"]["admin"]; !ok {
		return errors.New("Gas Admin not set")
	}
	if _, ok := i["Gas"]["zeroClaimRate"]; !ok {
		return errors.New("Gas zeroClaimRate not set")
	}
	if _, ok := i["Gas"]["baseGasSeconds"]; !ok {
		return errors.New("Gas baseGasSeconds not set")
	}
	if _, ok := i["Gas"]["ceilGasSeconds"]; !ok {
		return errors.New("Gas ceilGasSeconds not set")
	}
	if _, ok := i["Gas"]["ceilClaimRate"]; !ok {
		return errors.New("Gas ceilClaimRate not set")
	}
	if _, ok := i["Blast"]["yieldContract"]; !ok {
		return errors.New("Blast yieldContract not set")
	}
	if _, ok := i["USDB"]["usdYieldManager"]; !ok {
		return errors.New("USDB usdYieldManager not set")
	}
	if _, ok := i["USDB"]["remoteToken"]; !ok {
		return errors.New("USDB remoteToken not set")
	}

	return nil
}

// DeploymentResults represents the output of deploying each of the
// contracts so that the immutables can be set properly in the bytecode.
type DeploymentResults map[string]hexutil.Bytes

// BuildOptimism will deploy the L2 predeploys so that their immutables are set
// correctly.
func BuildOptimism(immutable ImmutableConfig) (DeploymentResults, error) {
	if err := immutable.Check(); err != nil {
		return DeploymentResults{}, err
	}

	deployments := []deployer.Constructor{
		{
			Name: "GasPriceOracle",
		},
		{
			Name: "L1Block",
		},
		{
			Name: "L2CrossDomainMessenger",
			Args: []interface{}{
				immutable["L2CrossDomainMessenger"]["otherMessenger"],
			},
		},
		{
			Name: "L2StandardBridge",
			Args: []interface{}{
				immutable["L2StandardBridge"]["otherBridge"],
			},
		},
		{
			Name: "L2ToL1MessagePasser",
		},
		{
			Name: "SequencerFeeVault",
			Args: []interface{}{
				immutable["SequencerFeeVault"]["recipient"],
				immutable["SequencerFeeVault"]["minimumWithdrawalAmount"],
				immutable["SequencerFeeVault"]["withdrawalNetwork"],
			},
		},
		{
			Name: "BaseFeeVault",
			Args: []interface{}{
				immutable["BaseFeeVault"]["recipient"],
				immutable["BaseFeeVault"]["minimumWithdrawalAmount"],
				immutable["BaseFeeVault"]["withdrawalNetwork"],
			},
		},
		{
			Name: "L1FeeVault",
			Args: []interface{}{
				immutable["L1FeeVault"]["recipient"],
				immutable["L1FeeVault"]["minimumWithdrawalAmount"],
				immutable["L1FeeVault"]["withdrawalNetwork"],
			},
		},
		{
			Name: "OptimismMintableERC20Factory",
		},
		{
			Name: "DeployerWhitelist",
		},
		{
			Name: "LegacyMessagePasser",
		},
		{
			Name: "L1BlockNumber",
		},
		{
			Name: "L2ERC721Bridge",
			Args: []interface{}{
				immutable["L2ERC721Bridge"]["otherBridge"],
			},
		},
		{
			Name: "OptimismMintableERC721Factory",
			Args: []interface{}{
				predeploys.L2ERC721BridgeAddr,
				immutable["OptimismMintableERC721Factory"]["remoteChainId"],
			},
		},
		{
			Name: "LegacyERC20ETH",
		},
		{
			Name: "EAS",
		},
		{
			Name: "SchemaRegistry",
		},
		{
			Name: "Shares",
			Args: []interface{}{
				immutable["Shares"]["reporter"],
			},
		},
		{
			Name: "Gas",
			Args: []interface{}{
				immutable["Gas"]["admin"],
				common.HexToAddress(predeploys.Blast),
				common.HexToAddress(predeploys.BaseFeeVault),
				immutable["Gas"]["zeroClaimRate"],
				immutable["Gas"]["baseGasSeconds"],
				immutable["Gas"]["baseClaimRate"],
				immutable["Gas"]["ceilGasSeconds"],
				immutable["Gas"]["ceilClaimRate"],
			},
		},
		{
			Name: "Blast",
			Args: []interface{}{
				common.HexToAddress(predeploys.Gas),
				immutable["Blast"]["yieldContract"],
			},
		},
		{
			Name: "WETHRebasing",
		},
		{
			Name: "L2BlastBridge",
			Args: []interface{}{
				immutable["L2BlastBridge"]["otherBridge"],
			},
		},
		{
			Name: "USDB",
			Args: []interface{}{
				immutable["USDB"]["usdYieldManager"],
				common.HexToAddress(predeploys.L2BlastBridge),
				immutable["USDB"]["remoteToken"],
			},
		},
	}
	return BuildL2(deployments)
}

// BuildL2 will deploy contracts to a simulated backend so that their immutables
// can be properly set. The bytecode returned in the results is suitable to be
// inserted into the state via state surgery.
func BuildL2(constructors []deployer.Constructor) (DeploymentResults, error) {
	log.Info("Creating L2 state")
	deployments, err := deployer.Deploy(deployer.NewL2Backend(), constructors, l2Deployer)
	if err != nil {
		return nil, err
	}
	results := make(DeploymentResults)
	for _, dep := range deployments {
		results[dep.Name] = dep.Bytecode
	}
	return results, nil
}

func l2Deployer(backend *backends.SimulatedBackend, opts *bind.TransactOpts, deployment deployer.Constructor) (*types.Transaction, error) {
	var tx *types.Transaction
	var recipient common.Address
	var minimumWithdrawalAmount *big.Int
	var withdrawalNetwork uint8
	var err error
	switch deployment.Name {
	case "GasPriceOracle":
		_, tx, _, err = bindings.DeployGasPriceOracle(opts, backend)
	case "L1Block":
		// No arguments required for the L1Block contract
		_, tx, _, err = bindings.DeployL1Block(opts, backend)
	case "L2CrossDomainMessenger":
		otherMessenger, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherMessenger")
		}
		_, tx, _, err = bindings.DeployL2CrossDomainMessenger(opts, backend, otherMessenger)
	case "L2StandardBridge":
		otherBridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherBridge")
		}
		_, tx, _, err = bindings.DeployL2StandardBridge(opts, backend, otherBridge)
	case "L2ToL1MessagePasser":
		// No arguments required for L2ToL1MessagePasser
		_, tx, _, err = bindings.DeployL2ToL1MessagePasser(opts, backend)
	case "SequencerFeeVault":
		recipient, minimumWithdrawalAmount, withdrawalNetwork, err = prepareFeeVaultArguments(deployment)
		if err != nil {
			return nil, err
		}
		_, tx, _, err = bindings.DeploySequencerFeeVault(opts, backend, recipient, minimumWithdrawalAmount, withdrawalNetwork)
	case "BaseFeeVault":
		recipient, minimumWithdrawalAmount, withdrawalNetwork, err = prepareFeeVaultArguments(deployment)
		if err != nil {
			return nil, err
		}
		_, tx, _, err = bindings.DeployBaseFeeVault(opts, backend, recipient, minimumWithdrawalAmount, withdrawalNetwork)
	case "L1FeeVault":
		recipient, minimumWithdrawalAmount, withdrawalNetwork, err = prepareFeeVaultArguments(deployment)
		if err != nil {
			return nil, err
		}
		_, tx, _, err = bindings.DeployL1FeeVault(opts, backend, recipient, minimumWithdrawalAmount, withdrawalNetwork)
	case "OptimismMintableERC20Factory":
		_, tx, _, err = bindings.DeployOptimismMintableERC20Factory(opts, backend)
	case "DeployerWhitelist":
		_, tx, _, err = bindings.DeployDeployerWhitelist(opts, backend)
	case "LegacyMessagePasser":
		_, tx, _, err = bindings.DeployLegacyMessagePasser(opts, backend)
	case "L1BlockNumber":
		_, tx, _, err = bindings.DeployL1BlockNumber(opts, backend)
	case "L2ERC721Bridge":
		otherBridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherBridge")
		}
		_, tx, _, err = bindings.DeployL2ERC721Bridge(opts, backend, otherBridge)
	case "OptimismMintableERC721Factory":
		bridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for bridge")
		}
		remoteChainId, ok := deployment.Args[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid type for remoteChainId")
		}
		_, tx, _, err = bindings.DeployOptimismMintableERC721Factory(opts, backend, bridge, remoteChainId)
	case "LegacyERC20ETH":
		_, tx, _, err = bindings.DeployLegacyERC20ETH(opts, backend)
	case "EAS":
		_, tx, _, err = bindings.DeployEAS(opts, backend)
	case "SchemaRegistry":
		_, tx, _, err = bindings.DeploySchemaRegistry(opts, backend)
	case "Shares":
		reporter, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for reporter")
		}
		_, tx, _, err = bindings.DeployShares(opts, backend, reporter)
	case "USDB":
		usdYieldManager, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for usdYieldManager")
		}
		l2Bridge, ok := deployment.Args[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for l2Bridge")
		}
		remoteToken, ok := deployment.Args[2].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for remoteToken")
		}
		_, tx, _, err = bindings.DeployUSDB(opts, backend, usdYieldManager, l2Bridge, remoteToken)
	case "WETHRebasing":
		_, tx, _, err = bindings.DeployWETHRebasing(opts, backend)
	case "L2BlastBridge":
		otherBridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherBridge")
		}
		_, tx, _, err = bindings.DeployL2BlastBridge(opts, backend, otherBridge)
	case "Gas":
		gasAdmin, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for gasAdmin")
		}
		blastAddr, ok := deployment.Args[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for blast addr")
		}
		baseFeeVaultAddr, ok := deployment.Args[2].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for base fee vault addr")
		}
		zeroClaimRate, ok := deployment.Args[3].(*hexutil.Big)
		if !ok {
			return nil, fmt.Errorf("invalid type for zero claim rate")
		}
		baseGasSeconds, ok := deployment.Args[4].(*hexutil.Big)
		if !ok {
			return nil, fmt.Errorf("invalid type for base gas seconds")
		}
		baseClaimRate, ok := deployment.Args[5].(*hexutil.Big)
		if !ok {
			return nil, fmt.Errorf("invalid type for base claim rate")
		}
		ceilGasSeconds, ok := deployment.Args[6].(*hexutil.Big)
		if !ok {
			return nil, fmt.Errorf("invalid type for ceil gas seconds")
		}
		ceilClaimRate, ok := deployment.Args[7].(*hexutil.Big)
		if !ok {
			return nil, fmt.Errorf("invalid type for ceil claim rate")
		}
		_, tx, _, err = bindings.DeployGas(opts, backend, gasAdmin, blastAddr, baseFeeVaultAddr, zeroClaimRate.ToInt(), baseGasSeconds.ToInt(), baseClaimRate.ToInt(), ceilGasSeconds.ToInt(), ceilClaimRate.ToInt())
	case "Blast":
		gasContractAddr, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for gasContract")
		}
		yieldContractAddr, ok := deployment.Args[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for yield contract")
		}
		_, tx, _, err = bindings.DeployBlast(opts, backend, gasContractAddr, yieldContractAddr)
	default:
		return tx, fmt.Errorf("unknown contract: %s", deployment.Name)
	}

	return tx, err
}

func prepareFeeVaultArguments(deployment deployer.Constructor) (common.Address, *big.Int, uint8, error) {
	recipient, ok := deployment.Args[0].(common.Address)
	if !ok {
		return common.Address{}, nil, 0, fmt.Errorf("invalid type for recipient")
	}
	minimumWithdrawalAmountHex, ok := deployment.Args[1].(*hexutil.Big)
	if !ok {
		return common.Address{}, nil, 0, fmt.Errorf("invalid type for minimumWithdrawalAmount")
	}
	withdrawalNetwork, ok := deployment.Args[2].(uint8)
	if !ok {
		return common.Address{}, nil, 0, fmt.Errorf("invalid type for withdrawalNetwork")
	}
	return recipient, minimumWithdrawalAmountHex.ToInt(), withdrawalNetwork, nil
}
