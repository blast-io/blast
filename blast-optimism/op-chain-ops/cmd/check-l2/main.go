package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/clients"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var defaultCrossDomainMessageSender = common.HexToAddress("0x000000000000000000000000000000000000dead")

// Default script for checking that L2 has been configured correctly. This should be extended in the future
// to pull in L1 deploy artifacts and assert that the L2 state is consistent with the L1 state.
func main() {
	log.Root().SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(isatty.IsTerminal(os.Stderr.Fd()))))

	app := &cli.App{
		Name:  "check-l2",
		Usage: "Check that an OP Stack L2 has been configured correctly",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "l1-rpc-url",
				Value:   "http://127.0.0.1:8545",
				Usage:   "L1 RPC URL",
				EnvVars: []string{"L1_RPC_URL"},
			},
			&cli.StringFlag{
				Name:    "l2-rpc-url",
				Value:   "http://127.0.0.1:9545",
				Usage:   "L2 RPC URL",
				EnvVars: []string{"L2_RPC_URL"},
			},
		},
		Action: entrypoint,
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("error checking l2", "err", err)
	}
}

// entrypoint is the entrypoint for the check-l2 script
func entrypoint(ctx *cli.Context) error {
	clients, err := clients.NewClientsFromContext(ctx)
	if err != nil {
		return err
	}

	log.Info("Checking predeploy proxy config")
	g := new(errgroup.Group)

	// Check that all proxies are configured correctly
	// Do this in parallel but not too quickly to allow for
	// querying against rate limiting RPC backends
	count := uint64(2048)
	for i := uint64(0); i < count; i++ {
		i := i
		if i%4 == 0 {
			log.Info("Checking proxy", "index", i, "total", count)
			if err := g.Wait(); err != nil {
				return err
			}
		}
		g.Go(func() error {
			return checkPredeploy(clients.L2Client, i)
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}
	log.Info("All predeploy proxies are set correctly")

	// Check that all of the defined predeploys are set up correctly
	for name, addr := range predeploys.Predeploys {
		log.Info("Checking predeploy", "name", name, "address", addr.Hex())
		if err := checkPredeployConfig(clients.L2Client, name); err != nil {
			return err
		}
	}
	return nil
}

// checkPredeploy ensures that the predeploy at index i has the correct proxy admin set
func checkPredeploy(client *ethclient.Client, i uint64) error {
	err := checkPredeployWithNamespace(client, i, genesis.BigL2PredeployNamespace)
	if err != nil {
		return err
	}
	err = checkPredeployWithNamespace(client, i, genesis.BlastBigL2PredeployNamespace)
	if err != nil {
		return err
	}
	return nil
}

func checkPredeployWithNamespace(client *ethclient.Client, i uint64, namespace *big.Int) error {
	bigAddr := new(big.Int).Or(namespace, new(big.Int).SetUint64(i))
	addr := common.BigToAddress(bigAddr)
	if !predeploys.IsProxied(addr) {
		return nil
	}
	admin, err := getEIP1967AdminAddress(client, addr)
	if err != nil {
		return err
	}
	if admin != predeploys.ProxyAdminAddr {
		return fmt.Errorf("%s does not have correct proxy admin set", addr)
	}
	return nil
}

// checkPredeployConfig checks that the defined predeploys are configured correctly
func checkPredeployConfig(client *ethclient.Client, name string) error {
	predeploy := predeploys.Predeploys[name]
	if predeploy == nil {
		return fmt.Errorf("unknown predeploy %s", name)
	}
	p := *predeploy

	g := new(errgroup.Group)
	if predeploys.IsProxied(p) {
		// Check that an implementation is set. If the implementation has been upgraded,
		// it will be considered non-standard. Ensure that there is code set at the implementation.
		g.Go(func() error {
			impl, err := getEIP1967ImplementationAddress(client, p)
			if err != nil {
				return err
			}
			log.Info(name, "implementation", impl.Hex())
			standardImpl, err := genesis.AddressToCodeNamespace(p)
			if err != nil {
				return err
			}
			if impl != standardImpl {
				log.Warn("%s does not have the standard implementation", name)
			}
			implCode, err := client.CodeAt(context.Background(), impl, nil)
			if err != nil {
				return err
			}
			if len(implCode) == 0 {
				return fmt.Errorf("%s implementation is not deployed", name)
			}
			return nil
		})

		// Ensure that the code is set to the proxy bytecode as expected
		g.Go(func() error {
			proxyCode, err := client.CodeAt(context.Background(), p, nil)
			if err != nil {
				return err
			}
			proxy, err := bindings.GetDeployedBytecode("Proxy")
			if err != nil {
				return err
			}
			if !bytes.Equal(proxyCode, proxy) {
				return fmt.Errorf("%s does not have the standard proxy code", name)
			}
			return nil
		})
	}

	// Check the predeploy specific config is correct
	g.Go(func() error {
		switch p {
		case predeploys.LegacyMessagePasserAddr:
			if err := checkLegacyMessagePasser(p, client); err != nil {
				return err
			}

		case predeploys.DeployerWhitelistAddr:
			if err := checkDeployerWhitelist(p, client); err != nil {
				return err
			}

		case predeploys.L2CrossDomainMessengerAddr:
			if err := checkL2CrossDomainMessenger(p, client); err != nil {
				return err
			}

		case predeploys.GasPriceOracleAddr:
			if err := checkGasPriceOracle(p, client); err != nil {
				return err
			}

		case predeploys.L2StandardBridgeAddr:
			if err := checkL2StandardBridge(p, client); err != nil {
				return err
			}

		case predeploys.SequencerFeeVaultAddr:
			if err := checkSequencerFeeVault(p, client); err != nil {
				return err
			}

		case predeploys.OptimismMintableERC20FactoryAddr:
			if err := checkOptimismMintableERC20Factory(p, client); err != nil {
				return err
			}

		case predeploys.L1BlockNumberAddr:
			if err := checkL1BlockNumber(p, client); err != nil {
				return err
			}

		case predeploys.L1BlockAddr:
			if err := checkL1Block(p, client); err != nil {
				return err
			}

		case predeploys.GovernanceTokenAddr:
			if err := checkGovernanceToken(p, client); err != nil {
				return err
			}

		case predeploys.L2ERC721BridgeAddr:
			if err := checkL2ERC721Bridge(p, client); err != nil {
				return err
			}

		case predeploys.OptimismMintableERC721FactoryAddr:
			if err := checkOptimismMintableERC721Factory(p, client); err != nil {
				return err
			}

		case predeploys.ProxyAdminAddr:
			if err := checkProxyAdmin(p, client); err != nil {
				return err
			}

		case predeploys.BaseFeeVaultAddr:
			if err := checkBaseFeeVault(p, client); err != nil {
				return err
			}

		case predeploys.L1FeeVaultAddr:
			if err := checkL1FeeVault(p, client); err != nil {
				return err
			}

		case predeploys.L2ToL1MessagePasserAddr:
			if err := checkL2ToL1MessagePasser(p, client); err != nil {
				return err
			}

		case predeploys.SchemaRegistryAddr:
			if err := checkSchemaRegistry(p, client); err != nil {
				return err
			}

		case predeploys.EASAddr:
			if err := checkEAS(p, client); err != nil {
				return err
			}

		case predeploys.SharesAddr:
			if err := checkShares(p, client); err != nil {
				return err
			}

		case predeploys.WETHRebasingAddr:
			if err := checkWETHRebasing(p, client); err != nil {
				return err
			}

		case predeploys.USDBAddr:
			if err := checkUSDB(p, client); err != nil {
				return err
			}

		case predeploys.L2BlastBridgeAddr:
			if err := checkL2BlastBridge(p, client); err != nil {
				return err
			}

		case predeploys.GasAddr:
			if err := checkGas(p, client); err != nil {
				return err
			}

		case predeploys.BlastAddr:
			if err := checkBlast(p, client); err != nil {
				return err
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func checkL2ToL1MessagePasser(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewL2ToL1MessagePasser(addr, client)
	if err != nil {
		return err
	}
	messageVersion, err := contract.MESSAGEVERSION(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2ToL1MessagePasser", "MESSAGE_VERSION", messageVersion)

	messageNonce, err := contract.MessageNonce(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2ToL1MessagePasser", "MESSAGE_NONCE", messageNonce)
	return nil
}

func checkL1FeeVault(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewL1FeeVault(addr, client)
	if err != nil {
		return err
	}
	recipient, err := contract.RECIPIENT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L1FeeVault", "RECIPIENT", recipient.Hex())
	if recipient == (common.Address{}) {
		return errors.New("RECIPIENT should not be address(0)")
	}

	minWithdrawalAmount, err := contract.MINWITHDRAWALAMOUNT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L1FeeVault", "MIN_WITHDRAWAL_AMOUNT", minWithdrawalAmount)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L1FeeVault version", "version", version)
	return nil
}

func checkBaseFeeVault(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewBaseFeeVault(addr, client)
	if err != nil {
		return err
	}
	recipient, err := contract.RECIPIENT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("BaseFeeVault", "RECIPIENT", recipient.Hex())
	if recipient == (common.Address{}) {
		return errors.New("RECIPIENT should not be address(0)")
	}

	minWithdrawalAmount, err := contract.MINWITHDRAWALAMOUNT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("BaseFeeVault", "MIN_WITHDRAWAL_AMOUNT", minWithdrawalAmount)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("BaseFeeVault version", "version", version)
	return nil
}

func checkProxyAdmin(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewProxyAdmin(addr, client)
	if err != nil {
		return err
	}

	owner, err := contract.Owner(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("ProxyAdmin", "owner", owner.Hex())
	if owner == (common.Address{}) {
		return errors.New("ProxyAdmin.owner is zero address")
	}

	addressManager, err := contract.AddressManager(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("ProxyAdmin", "addressManager", addressManager.Hex())
	return nil
}

func checkOptimismMintableERC721Factory(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewOptimismMintableERC721Factory(addr, client)
	if err != nil {
		return err
	}
	bridge, err := contract.BRIDGE(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("OptimismMintableERC721Factory", "BRIDGE", bridge.Hex())
	if bridge == (common.Address{}) {
		return errors.New("OptimismMintableERC721Factory.BRIDGE is zero address")
	}

	remoteChainID, err := contract.REMOTECHAINID(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("OptimismMintableERC721Factory", "REMOTE_CHAIN_ID", remoteChainID)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("OptimismMintableERC721Factory version", "version", version)
	return nil
}

func checkL2ERC721Bridge(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewL2ERC721Bridge(addr, client)
	if err != nil {
		return err
	}
	messenger, err := contract.MESSENGER(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2ERC721Bridge", "MESSENGER", messenger.Hex())
	if messenger == (common.Address{}) {
		return errors.New("L2ERC721Bridge.MESSENGER is zero address")
	}

	otherBridge, err := contract.OTHERBRIDGE(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2ERC721Bridge", "OTHERBRIDGE", otherBridge.Hex())
	if otherBridge == (common.Address{}) {
		return errors.New("L2ERC721Bridge.OTHERBRIDGE is zero address")
	}

	initialized, err := getInitialized("L2ERC721Bridge", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2ERC721Bridge", "_initialized", initialized)

	initializing, err := getInitializing("L2ERC721Bridge", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2ERC721Bridge", "_initializing", initializing)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2ERC721Bridge version", "version", version)
	return nil
}

func checkGovernanceToken(addr common.Address, client *ethclient.Client) error {
	code, err := client.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) > 0 {
		// This should also check the owner
		contract, err := bindings.NewERC20(addr, client)
		if err != nil {
			return err
		}
		name, err := contract.Name(&bind.CallOpts{})
		if err != nil {
			return err
		}
		log.Info("GovernanceToken", "name", name)
		symbol, err := contract.Symbol(&bind.CallOpts{})
		if err != nil {
			return err
		}
		log.Info("GovernanceToken", "symbol", symbol)
		totalSupply, err := contract.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return err
		}
		log.Info("GovernanceToken", "totalSupply", totalSupply)
	} else {
		log.Info("No code at GovernanceToken")
	}
	return nil
}

func checkL1Block(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewL1Block(addr, client)
	if err != nil {
		return err
	}
	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L1Block version", "version", version)
	return nil
}

func checkL1BlockNumber(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewL1BlockNumber(addr, client)
	if err != nil {
		return err
	}
	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L1BlockNumber version", "version", version)
	return nil
}

func checkOptimismMintableERC20Factory(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewOptimismMintableERC20Factory(addr, client)
	if err != nil {
		return err
	}

	bridgeLegacy, err := contract.BRIDGE(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("OptimismMintableERC20Factory", "BRIDGE", bridgeLegacy.Hex())
	if bridgeLegacy == (common.Address{}) {
		return errors.New("OptimismMintableERC20Factory.BRIDGE is zero address")
	}

	bridge, err := contract.Bridge(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if bridge == (common.Address{}) {
		return errors.New("OptimismMintableERC20Factory.bridge is zero address")
	}
	log.Info("OptimismMintableERC20Factory", "bridge", bridge.Hex())

	initialized, err := getInitialized("OptimismMintableERC20Factory", addr, client)
	if err != nil {
		return err
	}
	log.Info("OptimismMintableERC20Factory", "_initialized", initialized)

	initializing, err := getInitializing("OptimismMintableERC20Factory", addr, client)
	if err != nil {
		return err
	}
	log.Info("OptimismMintableERC20Factory", "_initializing", initializing)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("OptimismMintableERC20Factory version", "version", version)
	return nil
}

func checkSequencerFeeVault(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewSequencerFeeVault(addr, client)
	if err != nil {
		return err
	}
	recipient, err := contract.RECIPIENT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("SequencerFeeVault", "RECIPIENT", recipient.Hex())
	if recipient == (common.Address{}) {
		return errors.New("RECIPIENT should not be address(0)")
	}

	l1FeeWallet, err := contract.L1FeeWallet(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("SequencerFeeVault", "l1FeeWallet", l1FeeWallet.Hex())

	minWithdrawalAmount, err := contract.MINWITHDRAWALAMOUNT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("SequencerFeeVault", "MIN_WITHDRAWAL_AMOUNT", minWithdrawalAmount)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("SequencerFeeVault version", "version", version)
	return nil
}

func checkL2StandardBridge(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewL2StandardBridge(addr, client)
	if err != nil {
		return err
	}
	otherBridge, err := contract.OTHERBRIDGE(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if otherBridge == (common.Address{}) {
		return errors.New("OTHERBRIDGE should not be address(0)")
	}
	log.Info("L2StandardBridge", "OTHERBRIDGE", otherBridge.Hex())

	messenger, err := contract.MESSENGER(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2StandardBridge", "MESSENGER", messenger.Hex())
	if messenger != predeploys.L2CrossDomainMessengerAddr {
		return fmt.Errorf("L2StandardBridge MESSENGER should be %s, got %s", predeploys.L2CrossDomainMessengerAddr, messenger)
	}
	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}

	initialized, err := getInitialized("L2StandardBridge", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2StandardBridge", "_initialized", initialized)

	initializing, err := getInitializing("L2StandardBridge", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2StandardBridge", "_initializing", initializing)

	log.Info("L2StandardBridge version", "version", version)
	return nil
}

func checkGasPriceOracle(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewGasPriceOracle(addr, client)
	if err != nil {
		return err
	}
	decimals, err := contract.DECIMALS(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("GasPriceOracle", "DECIMALS", decimals)
	if decimals.Cmp(big.NewInt(6)) != 0 {
		return fmt.Errorf("GasPriceOracle decimals should be 6, got %v", decimals)
	}

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("GasPriceOracle version", "version", version)
	return nil
}

func checkL2CrossDomainMessenger(addr common.Address, client *ethclient.Client) error {
	slot, err := client.StorageAt(context.Background(), addr, common.Hash{31: 0xcc}, nil)
	if err != nil {
		return err
	}
	if common.BytesToAddress(slot) != defaultCrossDomainMessageSender {
		return fmt.Errorf("Expected xDomainMsgSender to be %s, got %s", defaultCrossDomainMessageSender, addr)
	}

	contract, err := bindings.NewL2CrossDomainMessenger(addr, client)
	if err != nil {
		return err
	}

	otherMessenger, err := contract.OTHERMESSENGER(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if otherMessenger == (common.Address{}) {
		return errors.New("OTHERMESSENGER should not be address(0)")
	}
	log.Info("L2CrossDomainMessenger", "OTHERMESSENGER", otherMessenger.Hex())

	l1CrossDomainMessenger, err := contract.L1CrossDomainMessenger(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "l1CrossDomainMessenger", l1CrossDomainMessenger.Hex())

	messageVersion, err := contract.MESSAGEVERSION(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "MESSAGE_VERSION", messageVersion)
	minGasCallDataOverhead, err := contract.MINGASCALLDATAOVERHEAD(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "MIN_GAS_CALLDATA_OVERHEAD", minGasCallDataOverhead)

	relayConstantOverhead, err := contract.RELAYCONSTANTOVERHEAD(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "RELAY_CONSTANT_OVERHEAD", relayConstantOverhead)

	minGasDynamicsOverheadDenominator, err := contract.MINGASDYNAMICOVERHEADDENOMINATOR(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR", minGasDynamicsOverheadDenominator)

	minGasDynamicsOverheadNumerator, err := contract.MINGASDYNAMICOVERHEADNUMERATOR(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR", minGasDynamicsOverheadNumerator)

	relayCallOverhead, err := contract.RELAYCALLOVERHEAD(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "RELAY_CALL_OVERHEAD", relayCallOverhead)

	relayReservedGas, err := contract.RELAYRESERVEDGAS(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "RELAY_RESERVED_GAS", relayReservedGas)

	relayGasCheckBuffer, err := contract.RELAYGASCHECKBUFFER(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "RELAY_GAS_CHECK_BUFFER", relayGasCheckBuffer)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}

	initialized, err := getInitialized("L2CrossDomainMessenger", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "_initialized", initialized)

	initializing, err := getInitializing("L2CrossDomainMessenger", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2CrossDomainMessenger", "_initializing", initializing)

	log.Info("L2CrossDomainMessenger version", "version", version)
	return nil
}

func checkLegacyMessagePasser(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewLegacyMessagePasser(addr, client)
	if err != nil {
		return err
	}
	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("LegacyMessagePasser version", "version", version)
	return nil
}

func checkDeployerWhitelist(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewDeployerWhitelist(addr, client)
	if err != nil {
		return err
	}
	owner, err := contract.Owner(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if owner != (common.Address{}) {
		return fmt.Errorf("DeployerWhitelist owner should be set to address(0)")
	}
	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("DeployerWhitelist version", "version", version)
	return nil
}

func checkSchemaRegistry(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewSchemaRegistry(addr, client)
	if err != nil {
		return err
	}

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("SchemaRegistry version", "version", version)
	return nil
}

func checkEAS(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewEAS(addr, client)
	if err != nil {
		return err
	}

	registry, err := contract.GetSchemaRegistry(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if registry != predeploys.SchemaRegistryAddr {
		return fmt.Errorf("Incorrect registry address %s", registry)
	}
	log.Info("EAS", "registry", registry)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("EAS version", "version", version)
	return nil
}

func checkShares(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewShares(addr, client)
	if err != nil {
		return err
	}

	price, err := contract.Price(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if price == 0 {
		return fmt.Errorf("Price is unset")
	}
	log.Info("Shares", "price", price)

	initialized, err := getInitialized("Shares", addr, client)
	if err != nil {
		return err
	}
	log.Info("Shares", "_initialized", initialized)

	initializing, err := getInitializing("Shares", addr, client)
	if err != nil {
		return err
	}
	log.Info("Shares", "_initializing", initializing)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("Shares version", "version", version)
	return nil
}

func checkUSDB(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewUSDB(addr, client)
	if err != nil {
		return err
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if name != "USDB" {
		return fmt.Errorf("Name is incorrect")
	}
	log.Info("USDB", "name", name)

	symbol, err := contract.Symbol(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if symbol != "USDB" {
		return fmt.Errorf("Symbol is incorrect")
	}
	log.Info("USDB", "symbol", symbol)

	decimals, err := contract.Decimals(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if decimals != 18 {
		return fmt.Errorf("Decimals is incorrect")
	}
	log.Info("USDB", "decimals", decimals)

	price, err := contract.Price(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if price == 0 {
		return fmt.Errorf("Price is unset")
	}
	log.Info("USDB", "price", price)

	initialized, err := getInitialized("USDB", addr, client)
	if err != nil {
		return err
	}
	log.Info("USDB", "_initialized", initialized)

	initializing, err := getInitializing("USDB", addr, client)
	if err != nil {
		return err
	}
	log.Info("USDB", "_initializing", initializing)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("USDB version", "version", version)
	return nil
}

func checkWETHRebasing(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewWETHRebasing(addr, client)
	if err != nil {
		return err
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if name != "Wrapped Ether" {
		return fmt.Errorf("Name is incorrect")
	}
	log.Info("WETHRebasing", "name", name)

	symbol, err := contract.Symbol(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if symbol != "WETH" {
		return fmt.Errorf("Symbol is incorrect")
	}
	log.Info("WETHRebasing", "symbol", symbol)

	decimals, err := contract.Decimals(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if decimals != 18 {
		return fmt.Errorf("Decimals is incorrect")
	}
	log.Info("WETHRebasing", "decimals", decimals)

	price, err := contract.Price(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if price == 0 {
		return fmt.Errorf("Price is unset")
	}
	log.Info("WETHRebasing", "price", price)

	initialized, err := getInitialized("WETHRebasing", addr, client)
	if err != nil {
		return err
	}
	log.Info("WETHRebasing", "_initialized", initialized)

	initializing, err := getInitializing("WETHRebasing", addr, client)
	if err != nil {
		return err
	}
	log.Info("WETHRebasing", "_initializing", initializing)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("WETHRebasing version", "version", version)
	return nil
}

func checkBlast(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewBlast(addr, client)
	if err != nil {
		return err
	}

	gas, err := contract.GAS_CONTRACT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if gas == predeploys.GasAddr {
		return fmt.Errorf("Incorrect Gas address %s", gas)
	}
	log.Info("Blast", "gas", gas)

	yield, err := contract.YIELD_CONTRACT(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if yield == predeploys.YieldAddr {
		return fmt.Errorf("Incorrect Yield address %s", yield)
	}
	log.Info("Blast", "yield", yield)

	initialized, err := getInitialized("Blast", addr, client)
	if err != nil {
		return err
	}
	log.Info("Blast", "_initialized", initialized)

	initializing, err := getInitializing("Blast", addr, client)
	if err != nil {
		return err
	}
	log.Info("Blast", "_initializing", initializing)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("Blast version", "version", version)
	return nil
}

func checkGas(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewGas(addr, client)
	if err != nil {
		return err
	}

	admin, err := contract.Admin(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if admin == (common.Address{}) {
		return fmt.Errorf("Admin should not be address(0)")
	}
	log.Info("Gas", "Admin", admin.Hex())

	blast, err := contract.BlastConfigurationContract(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if blast != predeploys.BlastAddr {
		return fmt.Errorf("Incorrect Blast address %s", blast.Hex())
	}
	log.Info("Gas", "BlastConfigurationContract", blast)

	blastFeeVault, err := contract.BlastFeeVault(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if blastFeeVault != predeploys.BaseFeeVault {
		return fmt.Errorf("Incorrect BlastFeeVault address %s", blastFeeVault.Hex())
	}
	log.Info("Gas", "BlastFeeVault", blastFeeVault)

	zeroClaimRate, err := contract.ZeroClaimRate(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if zeroClaimRate == 0 {
		return fmt.Errorf("zeroClaimRate unset")
	}
	log.Info("Gas", "ZeroClaimRate", zeroClaimRate)

	baseGasSeconds, err := contract.BaseGasSeconds(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if baseGasSeconds == 0 {
		return fmt.Errorf("baseGasSeconds unset")
	}
	log.Info("Gas", "BaseGasSeconds", baseGasSeconds)

	baseClaimRate, err := contract.BaseClaimRate(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if baseClaimRate == 0 {
		return fmt.Errorf("baseClaimRate unset")
	}
	log.Info("Gas", "BaseClaimRate", baseClaimRate)

	ceilGasSeconds, err := contract.CeilGasSeconds(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if ceilGasSeconds == 0 {
		return fmt.Errorf("ceilGasSeconds unset")
	}
	log.Info("Gas", "CeilGasSeconds", ceilGasSeconds)

	ceilClaimRate, err := contract.CeilClaimRate(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if ceilClaimRate == 0 {
		return fmt.Errorf("ceilClaimRate unset")
	}
	log.Info("Gas", "CeilClaimRate", ceilClaimRate)

	initialized, err := getInitialized("Gas", addr, client)
	if err != nil {
		return err
	}
	log.Info("Gas", "_initialized", initialized)

	initializing, err := getInitializing("Gas", addr, client)
	if err != nil {
		return err
	}
	log.Info("Gas", "_initializing", initializing)

	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("Gas version", "version", version)
	return nil
}

func checkL2BlastBridge(addr common.Address, client *ethclient.Client) error {
	contract, err := bindings.NewL2BlastBridge(addr, client)
	if err != nil {
		return err
	}
	otherBridge, err := contract.OTHERBRIDGE(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if otherBridge == (common.Address{}) {
		return errors.New("OTHERBRIDGE should not be address(0)")
	}
	log.Info("L2BlastBridge", "OTHERBRIDGE", otherBridge.Hex())

	messenger, err := contract.MESSENGER(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Info("L2BlastBridge", "MESSENGER", messenger.Hex())
	if messenger != predeploys.L2CrossDomainMessengerAddr {
		return fmt.Errorf("L2BlastBridge MESSENGER should be %s, got %s", predeploys.L2CrossDomainMessengerAddr, messenger)
	}
	version, err := contract.Version(&bind.CallOpts{})
	if err != nil {
		return err
	}

	initialized, err := getInitialized("L2BlastBridge", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2BlastBridge", "_initialized", initialized)

	initializing, err := getInitializing("L2BlastBridge", addr, client)
	if err != nil {
		return err
	}
	log.Info("L2BlastBridge", "_initializing", initializing)

	log.Info("L2BlastBridge version", "version", version)
	return nil
}

func getEIP1967AdminAddress(client *ethclient.Client, addr common.Address) (common.Address, error) {
	slot, err := client.StorageAt(context.Background(), addr, genesis.AdminSlot, nil)
	if err != nil {
		return common.Address{}, err
	}
	admin := common.BytesToAddress(slot)
	return admin, nil
}

func getEIP1967ImplementationAddress(client *ethclient.Client, addr common.Address) (common.Address, error) {
	slot, err := client.StorageAt(context.Background(), addr, genesis.ImplementationSlot, nil)
	if err != nil {
		return common.Address{}, err
	}
	impl := common.BytesToAddress(slot)
	return impl, nil
}

// getInitialized will get the initialized value in storage of a contract.
// This is an incrementing number that starts at 1 and increments each time that
// the contract is upgraded.
func getInitialized(name string, addr common.Address, client *ethclient.Client) (*big.Int, error) {
	value, err := getStorageValue(name, "_initialized", addr, client)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(value), nil
}

// getInitializing will get the _initializing value in storage of a contract.
func getInitializing(name string, addr common.Address, client *ethclient.Client) (bool, error) {
	value, err := getStorageValue(name, "_initializing", addr, client)
	if err != nil {
		return false, err
	}
	if len(value) != 1 {
		return false, fmt.Errorf("Unexpected length for _initializing: %d", len(value))
	}
	return value[0] == 1, nil
}

// getStorageValue will get the value of a named storage slot in a contract. It isn't smart about
// automatically converting from a byte slice to a type, it is the caller's responsibility to do that.
func getStorageValue(name, entryName string, addr common.Address, client *ethclient.Client) ([]byte, error) {
	layout, err := bindings.GetStorageLayout(name)
	if err != nil {
		return nil, err
	}
	entry, err := layout.GetStorageLayoutEntry(entryName)
	if err != nil {
		return nil, err
	}
	typ, err := layout.GetStorageLayoutType(entry.Type)
	if err != nil {
		return nil, err
	}
	slot := common.BigToHash(big.NewInt(int64(entry.Slot)))
	value, err := client.StorageAt(context.Background(), addr, slot, nil)
	if err != nil {
		return nil, err
	}
	if entry.Offset+typ.NumberOfBytes > uint(len(value)) {
		return nil, fmt.Errorf("value length is too short")
	}
	// Swap the endianness
	slice := common.CopyBytes(value)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice[entry.Offset : entry.Offset+typ.NumberOfBytes], nil
}
