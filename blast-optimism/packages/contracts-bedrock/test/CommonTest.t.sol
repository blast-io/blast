// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { Test, StdUtils } from "forge-std/Test.sol";
import { Vm } from "forge-std/Vm.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { L1StandardBridge } from "src/L1/L1StandardBridge.sol";
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { WithdrawalQueue } from "src/mainnet-bridge/withdrawal-queue/WithdrawalQueue.sol";
import { L1ERC721Bridge } from "src/L1/L1ERC721Bridge.sol";
import { L2ERC721Bridge } from "src/L2/L2ERC721Bridge.sol";
import { OptimismMintableERC20Factory } from "src/universal/OptimismMintableERC20Factory.sol";
import { OptimismMintableERC721Factory } from "src/universal/OptimismMintableERC721Factory.sol";
import { OptimismMintableERC20 } from "src/universal/OptimismMintableERC20.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { L1CrossDomainMessenger } from "src/L1/L1CrossDomainMessenger.sol";
import { L2CrossDomainMessenger } from "src/L2/L2CrossDomainMessenger.sol";
import { USDB } from "src/L2/USDB.sol";
import { WETHRebasing } from "src/L2/WETHRebasing.sol";
import { Shares } from "src/L2/Shares.sol";
import { L1BlastBridge } from "src/mainnet-bridge/L1BlastBridge.sol";
import { L2BlastBridge } from "src/mainnet-bridge/L2BlastBridge.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { Blast } from "src/L2/Blast.sol";
import { Gas } from "src/L2/Gas.sol";
import { SequencerFeeVault } from "src/L2/SequencerFeeVault.sol";
import { FeeVault } from "src/universal/FeeVault.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { LegacyERC20ETH } from "src/legacy/LegacyERC20ETH.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Types } from "src/libraries/Types.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { ERC20Mock } from "@openzeppelin/contracts/mocks/ERC20Mock.sol";
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { ResolvedDelegateProxy } from "src/legacy/ResolvedDelegateProxy.sol";
import { AddressManager } from "src/legacy/AddressManager.sol";
import { L1ChugSplashProxy } from "src/legacy/L1ChugSplashProxy.sol";
import { IL1ChugSplashDeployer } from "src/legacy/L1ChugSplashProxy.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { LegacyMintableERC20 } from "src/legacy/LegacyMintableERC20.sol";
import { SystemConfig } from "src/L1/SystemConfig.sol";
import { ResourceMetering } from "src/L1/ResourceMetering.sol";
import { Constants } from "src/libraries/Constants.sol";
import { LidoYieldProvider, ILido } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";
import { DSRYieldProvider } from "src/mainnet-bridge/yield-providers/DSRYieldProvider.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { DSRYieldProvider, IDsrManager } from "src/mainnet-bridge/yield-providers/DSRYieldProvider.sol";
import { ETHYieldManager } from "src/mainnet-bridge/ETHYieldManager.sol";
import { USDYieldManager } from "src/mainnet-bridge/USDYieldManager.sol";
import { USDConversions, IDssPsm } from "src/mainnet-bridge/USDConversions.sol";
import { Insurance } from "src/mainnet-bridge/Insurance.sol";
import { YieldMode } from "src/L2/Blast.sol";
import { GasMode } from "src/L2/Gas.sol";

interface IPot {
    function chi() external view returns (uint256);
    function rho() external view returns (uint256);
    function dsr() external view returns (uint256);
    function drip() external;
}

contract MockGas is Gas {
    constructor(
        address _admin,
        address _blastConfigurationContract,
        address _blastFeeVault,
        uint256 _zeroClaimRate,
        uint256 _baseGasSeconds,
        uint256 _baseClaimRate,
        uint256 _ceilGasSeconds,
        uint256 _ceilClaimRate
    ) Gas(_admin, _blastConfigurationContract, _blastFeeVault, _zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate) {
    }

    function updateGasParams(address contractAddress, uint256 etherSeconds, uint256 etherBalance, GasMode mode) external {
        _updateGasParams(contractAddress, etherSeconds, etherBalance, mode);
    }
}

contract CommonTest is Test {
    address alice = address(128);
    address bob = address(257);
    address charlie = address(512);
    address multisig = address(1024);

    address immutable ZERO_ADDRESS = address(0);
    address immutable NON_ZERO_ADDRESS = address(1);
    uint256 immutable NON_ZERO_VALUE = 100;
    uint256 immutable ZERO_VALUE = 0;
    uint64 immutable NON_ZERO_GASLIMIT = 50000;
    bytes32 nonZeroHash = keccak256(abi.encode("NON_ZERO"));
    bytes NON_ZERO_DATA = hex"0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff0000";

    event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData);

    /// @dev OpenZeppelin Ownable.sol transferOwnership event
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    FFIInterface ffi;

    function setUp() public virtual {
        // Give alice and bob some ETH
        vm.deal(alice, 1 << 16);
        vm.deal(bob, 1 << 16);
        vm.deal(charlie, 1 << 16);
        vm.deal(multisig, 1 << 16);

        vm.label(alice, "alice");
        vm.label(bob, "bob");
        vm.label(charlie, "charlie");
        vm.label(multisig, "multisig");

        // Make sure we have a non-zero base fee
        vm.fee(1000000000);

        ffi = new FFIInterface();
    }

    function emitTransactionDeposited(
        address _from,
        address _to,
        uint256 _mint,
        uint256 _value,
        uint64 _gasLimit,
        bool _isCreation,
        bytes memory _data
    )
        internal
    {
        emit TransactionDeposited(_from, _to, 0, abi.encodePacked(_mint, _value, _gasLimit, _isCreation, _data));
    }
}

contract L2OutputOracle_Initializer is CommonTest {
    // Test target
    L2OutputOracle oracle;
    L2OutputOracle oracleImpl;

    Blast blast;
    MockGas gas;

    L2ToL1MessagePasser messagePasser = L2ToL1MessagePasser(payable(Predeploys.L2_TO_L1_MESSAGE_PASSER));

    // Constructor arguments
    address internal proposer = 0x000000000000000000000000000000000000AbBa;
    address internal owner = 0x000000000000000000000000000000000000ACDC;
    uint256 internal submissionInterval = 1800;
    uint256 internal l2BlockTime = 2;
    uint256 internal startingBlockNumber = 200;
    uint256 internal startingTimestamp = 1000;
    uint256 internal finalizationPeriodSeconds = 7 days;
    address guardian;

    // Test data
    uint256 initL1Time;

    event OutputProposed(
        bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp
    );

    event OutputsDeleted(uint256 indexed prevNextOutputIndex, uint256 indexed newNextOutputIndex);

    // Advance the evm's time to meet the L2OutputOracle's requirements for proposeL2Output
    function warpToProposeTime(uint256 _nextBlockNumber) public {
        vm.warp(oracle.computeL2Timestamp(_nextBlockNumber) + 1);
    }

    /// @dev Helper function to propose an output.
    function proposeAnotherOutput() public {
        bytes32 proposedOutput2 = keccak256(abi.encode());
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        uint256 nextOutputIndex = oracle.nextOutputIndex();
        warpToProposeTime(nextBlockNumber);
        uint256 proposedNumber = oracle.latestBlockNumber();

        // Ensure the submissionInterval is enforced
        assertEq(nextBlockNumber, proposedNumber + submissionInterval);

        vm.roll(nextBlockNumber + 1);

        vm.expectEmit(true, true, true, true);
        emit OutputProposed(proposedOutput2, nextOutputIndex, nextBlockNumber, block.timestamp);

        vm.prank(proposer);
        oracle.proposeL2Output(proposedOutput2, nextBlockNumber, 0, 0);
    }

    function checkBlastConfig(
        address contractAddress,
        address governor,
        YieldMode yieldMode,
        GasMode gasMode
    ) public {
        assertEq(blast.governorMap(contractAddress), governor);
        assertTrue(blast.readYieldConfiguration(contractAddress) == uint8(yieldMode));
        (,,, GasMode _gasMode) = blast.readGasParams(contractAddress);
        assertTrue(_gasMode == gasMode);
    }

    function setUp() public virtual override {
        super.setUp();
        vm.deal(address(messagePasser), 0);
        guardian = makeAddr("guardian");

        // By default the first block has timestamp and number zero, which will cause underflows in the
        // tests, so we'll move forward to these block values.
        initL1Time = startingTimestamp + 1;
        vm.warp(initL1Time);
        vm.roll(startingBlockNumber);
        // Deploy the L2OutputOracle and transfer owernship to the proposer
        oracleImpl = new L2OutputOracle({
            _submissionInterval: submissionInterval,
            _l2BlockTime: l2BlockTime,
            _finalizationPeriodSeconds: finalizationPeriodSeconds
        });
        Proxy proxy = new Proxy(multisig);
        vm.prank(multisig);
        proxy.upgradeToAndCall(
            address(oracleImpl),
            abi.encodeCall(L2OutputOracle.initialize, (startingBlockNumber, startingTimestamp, proposer, owner))
        );
        oracle = L2OutputOracle(address(proxy));
        vm.label(address(oracle), "L2OutputOracle");

        // Set the L2ToL1MessagePasser at the correct address
        vm.etch(Predeploys.L2_TO_L1_MESSAGE_PASSER, address(new L2ToL1MessagePasser()).code);
        vm.label(Predeploys.L2_TO_L1_MESSAGE_PASSER, "L2ToL1MessagePasser");

        MockGas mockGas = new MockGas({ _admin: address(0x0), 
                            _blastConfigurationContract: address(Predeploys.BLAST),
                            _blastFeeVault: address(Predeploys.BASE_FEE_VAULT),
                            _zeroClaimRate: 0x1F4, // 5000 = 5%
                            _baseClaimRate: 0x7D0, // 20000 = 20%
                            _ceilClaimRate: 0x2710, // 10_000 = 100%
                            _baseGasSeconds: 0x1, // 1s
                            _ceilGasSeconds: 0xC8 // 200s
                            });
        vm.etch(Predeploys.GAS, address(mockGas).code);
        gas = MockGas(Predeploys.GAS);
        vm.label(address(gas), "GAS");

        MockYield mockYield = new MockYield();
        Blast mockBlast = new Blast({ _gasContract: address(Predeploys.GAS), _yieldContract: address(mockYield) });
        vm.etch(Predeploys.BLAST, address(mockBlast).code);
        blast = Blast(Predeploys.BLAST);
        vm.label(address(blast), "BLAST");
    }
}

contract MockYield {
    mapping(address => uint8) internal configurations;

    function getConfiguration(address x) external view returns (uint8) {
        return configurations[x];
    }
    function configure(address x, uint8 mode) external returns (uint256) {
        configurations[x] = mode;
    }
    function claim(address, address, uint256) external returns (uint256) {}
}

contract Portal_Initializer is L2OutputOracle_Initializer {
    // Test target
    OptimismPortal internal opImpl;
    OptimismPortal internal op;
    ETHYieldManager internal ethYieldManager;
    USDYieldManager internal usdYieldManager;
    SystemConfig systemConfig;

    ERC20 USDC;
    ERC20 USDT;
    ERC20 DAI;

    event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success);
    event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to, uint256 requestId);

    function setUp() public virtual override {
        super.setUp();

        Proxy systemConfigProxy = new Proxy(multisig);

        SystemConfig systemConfigImpl = new SystemConfig();

        vm.prank(multisig);
        systemConfigProxy.upgradeToAndCall(
            address(systemConfigImpl),
            abi.encodeCall(
                SystemConfig.initialize,
                (
                    address(1), //_owner,
                    0, //_overhead,
                    10000, //_scalar,
                    bytes32(0), //_batcherHash,
                    30_000_000, //_gasLimit,
                    address(0), //_unsafeBlockSigner,
                    Constants.DEFAULT_RESOURCE_CONFIG(), //_config,
                    0, //_startBlock
                    address(0xff), // _batchInbox
                    SystemConfig.Addresses({ // _addresses
                        l1CrossDomainMessenger: address(0),
                        l1ERC721Bridge: address(0),
                        l1StandardBridge: address(0),
                        l2OutputOracle: address(oracle),
                        optimismPortal: address(op),
                        optimismMintableERC20Factory: address(0)
                    })
                )
            )
        );

        systemConfig = SystemConfig(address(systemConfigProxy));

        opImpl = new OptimismPortal();

        Proxy portalProxy = new Proxy(multisig);

        if (address(USDConversions.DAI).code.length == 0) {
            vm.etch(address(USDConversions.DAI), address(new ERC20("DAI", "DAI")).code);
        }
        DAI = ERC20(address(USDConversions.DAI));
        if (address(USDConversions.USDC).code.length == 0) {
            vm.etch(address(USDConversions.USDC), address(new ERC20("USDC", "USDC")).code);
        }
        USDC = ERC20(address(USDConversions.USDC));
        if (address(USDConversions.USDT).code.length == 0) {
            vm.etch(address(USDConversions.USDT), address(new ERC20("USDT", "USDT")).code);
        }
        USDT = ERC20(address(USDConversions.USDT));

        L1ChugSplashProxy eymProxy = new L1ChugSplashProxy(multisig);
        vm.mockCall(multisig, abi.encodeWithSelector(IL1ChugSplashDeployer.isUpgrading.selector), abi.encode(true));

        vm.prank(multisig);
        portalProxy.upgradeToAndCall(
            address(opImpl), abi.encodeCall(OptimismPortal.initialize, (oracle, guardian, systemConfig, false, ETHYieldManager(payable(address(eymProxy)))))
        );
        op = OptimismPortal(payable(address(portalProxy)));
        vm.label(address(op), "OptimismPortal");

        vm.startPrank(multisig);
        eymProxy.setCode(address(new ETHYieldManager()).code);
        vm.clearMockedCalls();
        address ETHYieldManager_Impl = eymProxy.getImplementation();
        ethYieldManager = ETHYieldManager(payable(address(eymProxy)));
        vm.stopPrank();
        ethYieldManager.initialize(op, multisig);

        L1ChugSplashProxy uymProxy = new L1ChugSplashProxy(multisig);
        vm.mockCall(multisig, abi.encodeWithSelector(IL1ChugSplashDeployer.isUpgrading.selector), abi.encode(true));

        vm.startPrank(multisig);
        uymProxy.setCode(address(new USDYieldManager(address(DAI))).code); // TODO
        vm.clearMockedCalls();
        address USDYieldManager_Impl = uymProxy.getImplementation();
        usdYieldManager = USDYieldManager(payable(address(uymProxy)));
        vm.stopPrank();
        usdYieldManager.initialize(op, multisig);

        vm.label(address(eymProxy), "ETHYieldManager_Proxy");
        vm.label(address(ETHYieldManager_Impl), "ETHYieldManager_Impl");

        vm.label(address(uymProxy), "USDYieldManager_Proxy");
        vm.label(address(USDYieldManager_Impl), "USDYieldManager_Impl");
    }
}

contract Messenger_Initializer is Portal_Initializer {
    AddressManager internal addressManager;
    L1CrossDomainMessenger internal L1Messenger;
    L2CrossDomainMessenger internal L2Messenger = L2CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER);

    event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit);

    event SentMessageExtension1(address indexed sender, uint256 value);

    event MessagePassed(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data,
        bytes32 withdrawalHash
    );

    event RelayedMessage(bytes32 indexed msgHash);
    event FailedRelayedMessage(bytes32 indexed msgHash);

    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 mint,
        uint256 value,
        uint64 gasLimit,
        bool isCreation,
        bytes data
    );

    event WhatHappened(bool success, bytes returndata);

    function setUp() public virtual override {
        super.setUp();

        // Deploy the address manager
        vm.prank(multisig);
        addressManager = new AddressManager();

        // Setup implementation
        L1CrossDomainMessenger L1MessengerImpl = new L1CrossDomainMessenger();

        // Setup the address manager and proxy
        vm.prank(multisig);
        addressManager.setAddress("OVM_L1CrossDomainMessenger", address(L1MessengerImpl));
        ResolvedDelegateProxy proxy = new ResolvedDelegateProxy(
            addressManager,
            "OVM_L1CrossDomainMessenger"
        );
        L1Messenger = L1CrossDomainMessenger(address(proxy));
        L1Messenger.initialize(op);

        vm.etch(Predeploys.L2_CROSS_DOMAIN_MESSENGER, address(new L2CrossDomainMessenger(address(L1Messenger))).code);

        L2Messenger.initialize();

        // Label addresses
        vm.label(address(addressManager), "AddressManager");
        vm.label(address(L1MessengerImpl), "L1CrossDomainMessenger_Impl");
        vm.label(address(L1Messenger), "L1CrossDomainMessenger_Proxy");
        vm.label(Predeploys.LEGACY_ERC20_ETH, "LegacyERC20ETH");
        vm.label(Predeploys.L2_CROSS_DOMAIN_MESSENGER, "L2CrossDomainMessenger");

        vm.label(AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)), "L1CrossDomainMessenger_aliased");
    }
}

contract Bridge_Initializer is Messenger_Initializer {
    L1StandardBridge L1Bridge;
    L2StandardBridge L2Bridge;
    L1BlastBridge l1BlastBridge;
    L2BlastBridge l2BlastBridge;
    OptimismMintableERC20Factory L2TokenFactory;
    OptimismMintableERC20Factory L1TokenFactory;
    ERC20 L1Token;
    ERC20 BadL1Token;
    OptimismMintableERC20 L2Token;
    LegacyMintableERC20 LegacyL2Token;
    USDB Usdb;
    Shares SHARES;
    WETHRebasing WETH;
    ERC20 NativeL2Token;
    ERC20 BadL2Token;
    OptimismMintableERC20 RemoteL1Token;

    event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes data);

    event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes data);

    event ERC20DepositInitiated(
        address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data
    );

    event ERC20WithdrawalFinalized(
        address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data
    );

    event WithdrawalInitiated(
        address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data
    );

    event DepositFinalized(
        address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data
    );

    event DepositFailed(
        address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data
    );

    event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes data);

    event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes data);

    event ERC20BridgeInitiated(
        address indexed localToken,
        address indexed remoteToken,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    event ERC20BridgeFinalized(
        address indexed localToken,
        address indexed remoteToken,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    function setUp() public virtual override {
        super.setUp();

        L1Token = new ERC20("Native L1 Token", "L1T");

        vm.label(Predeploys.L2_STANDARD_BRIDGE, "L2StandardBridge");
        vm.label(Predeploys.OPTIMISM_MINTABLE_ERC20_FACTORY, "OptimismMintableERC20Factory");

        // Deploy the L1 bridge and initialize it with the address of the
        // L1CrossDomainMessenger
        L1ChugSplashProxy proxy = new L1ChugSplashProxy(multisig);
        vm.mockCall(multisig, abi.encodeWithSelector(IL1ChugSplashDeployer.isUpgrading.selector), abi.encode(true));
        vm.startPrank(multisig);
        proxy.setCode(address(new L1StandardBridge()).code);
        vm.clearMockedCalls();
        address L1Bridge_Impl = proxy.getImplementation();
        vm.stopPrank();

        L1Bridge = L1StandardBridge(payable(address(proxy)));
        L1Bridge.initialize({ _messenger: L1Messenger });

        vm.label(address(proxy), "L1StandardBridge_Proxy");
        vm.label(address(L1Bridge_Impl), "L1StandardBridge_Impl");

        // Deploy the L2StandardBridge, move it to the correct predeploy
        // address and then initialize it. It is safe to call initialize directly
        // on the proxy because the bytecode was set in state with `etch`.
        vm.etch(Predeploys.L2_STANDARD_BRIDGE, address(new L2StandardBridge(StandardBridge(payable(proxy)))).code);
        L2Bridge = L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE));
        L2Bridge.initialize();

        proxy = new L1ChugSplashProxy(multisig);
        vm.mockCall(multisig, abi.encodeWithSelector(IL1ChugSplashDeployer.isUpgrading.selector), abi.encode(true));
        vm.startPrank(multisig);
        proxy.setCode(address(new L1BlastBridge()).code);
        vm.clearMockedCalls();
        address L1BlastBridge_Impl = proxy.getImplementation();
        l1BlastBridge = L1BlastBridge(payable(address(proxy)));
        l1BlastBridge.initialize({
          _portal: op,
          _messenger: L1Messenger,
          _usdYieldManager: usdYieldManager,
          _ethYieldManager: ethYieldManager
        });
        ethYieldManager.setBlastBridge(address(l1BlastBridge));
        usdYieldManager.setBlastBridge(address(l1BlastBridge));
        ethYieldManager.setAdmin(multisig);
        usdYieldManager.setAdmin(multisig);
        vm.stopPrank();

        vm.label(address(proxy), "L1BlastBridge_Proxy");
        vm.label(address(L1BlastBridge_Impl), "L1BlastBridge_Impl");

        vm.etch(Predeploys.L2_BLAST_BRIDGE, address(new L2BlastBridge(StandardBridge(payable(proxy)))).code);
        l2BlastBridge = L2BlastBridge(payable(Predeploys.L2_BLAST_BRIDGE));
        l2BlastBridge.initialize();

        // Set up the L2 mintable token factory
        OptimismMintableERC20Factory factory = new OptimismMintableERC20Factory();
        vm.etch(Predeploys.OPTIMISM_MINTABLE_ERC20_FACTORY, address(factory).code);
        L2TokenFactory = OptimismMintableERC20Factory(Predeploys.OPTIMISM_MINTABLE_ERC20_FACTORY);
        L2TokenFactory.initialize(Predeploys.L2_STANDARD_BRIDGE);

        vm.etch(Predeploys.LEGACY_ERC20_ETH, address(new LegacyERC20ETH()).code);

        vm.startPrank(multisig);
        l1BlastBridge.setUSDYieldToken(address(L1Token), true, 18, address(0), false);
        vm.stopPrank();
        vm.startPrank(multisig);
        l1BlastBridge.setUSDYieldToken(address(USDC), true, 6, address(0), false);
        l1BlastBridge.setUSDYieldToken(address(USDT), true, 6, address(0), false);
        l1BlastBridge.setUSDYieldToken(address(DAI), true, 18, address(0), false);
        vm.stopPrank();

        USDB usdb = new USDB({
            _usdYieldManager: address(usdYieldManager),
            _l2Bridge: address(l2BlastBridge),
            _remoteToken: address(DAI)
        });
        vm.etch(Predeploys.USDB, address(usdb).code);
        Usdb = USDB(Predeploys.USDB);
        Usdb.initialize();
        vm.label(address(Usdb), "USDB");

        Shares shares = new Shares({ _reporter: address(ethYieldManager) });
        vm.etch(Predeploys.SHARES, address(shares).code);
        SHARES = Shares(Predeploys.SHARES);
        SHARES.initialize(1e9);
        vm.label(address(SHARES), "SHARES");


        vm.deal(address(this), 100 ether);
        WETHRebasing weth = new WETHRebasing();
        vm.etch(Predeploys.WETH_REBASING, address(weth).code);
        WETH = WETHRebasing(payable(Predeploys.WETH_REBASING));
        WETH.initialize();
        vm.label(address(WETH), "WETH");

        LegacyL2Token = new LegacyMintableERC20({
            _l2Bridge: address(L2Bridge),
            _l1Token: address(L1Token),
            _name: string.concat("LegacyL2-", L1Token.name()),
            _symbol: string.concat("LegacyL2-", L1Token.symbol())
        });
        vm.label(address(LegacyL2Token), "LegacyMintableERC20");

        // Deploy the L2 ERC20 now
        L2Token = OptimismMintableERC20(
            L2TokenFactory.createStandardL2Token(
                address(L1Token),
                string(abi.encodePacked("L2-", L1Token.name())),
                string(abi.encodePacked("L2-", L1Token.symbol()))
            )
        );

        BadL2Token = OptimismMintableERC20(
            L2TokenFactory.createStandardL2Token(
                address(1),
                string(abi.encodePacked("L2-", L1Token.name())),
                string(abi.encodePacked("L2-", L1Token.symbol()))
            )
        );

        NativeL2Token = new ERC20("Native L2 Token", "L2T");
        Proxy factoryProxy = new Proxy(multisig);
        OptimismMintableERC20Factory L1TokenFactoryImpl = new OptimismMintableERC20Factory();

        vm.prank(multisig);
        factoryProxy.upgradeToAndCall(
            address(L1TokenFactoryImpl), abi.encodeCall(OptimismMintableERC20Factory.initialize, address(L1Bridge))
        );

        L1TokenFactory = OptimismMintableERC20Factory(address(factoryProxy));

        RemoteL1Token = OptimismMintableERC20(
            L1TokenFactory.createStandardL2Token(
                address(NativeL2Token),
                string(abi.encodePacked("L1-", NativeL2Token.name())),
                string(abi.encodePacked("L1-", NativeL2Token.symbol()))
            )
        );

        BadL1Token = OptimismMintableERC20(
            L1TokenFactory.createStandardL2Token(
                address(1),
                string(abi.encodePacked("L1-", NativeL2Token.name())),
                string(abi.encodePacked("L1-", NativeL2Token.symbol()))
            )
        );
    }
}

contract ERC721Bridge_Initializer is Bridge_Initializer {
    L1ERC721Bridge L1NFTBridge;
    L2ERC721Bridge L2NFTBridge;

    function setUp() public virtual override {
        super.setUp();

        // Deploy the L1ERC721Bridge.
        L1ERC721Bridge l1BridgeImpl = new L1ERC721Bridge();
        Proxy l1BridgeProxy = new Proxy(multisig);

        vm.prank(multisig);
        l1BridgeProxy.upgradeToAndCall(
            address(l1BridgeImpl), abi.encodeCall(L1ERC721Bridge.initialize, (CrossDomainMessenger(L1Messenger)))
        );

        L1NFTBridge = L1ERC721Bridge(address(l1BridgeProxy));

        // Deploy the implementation for the L2ERC721Bridge and etch it into the predeploy address.
        L2ERC721Bridge l2BridgeImpl = new L2ERC721Bridge(address(L1NFTBridge));
        Proxy l2BridgeProxy = new Proxy(multisig);
        vm.etch(Predeploys.L2_ERC721_BRIDGE, address(l2BridgeProxy).code);

        // set the storage slot for admin
        bytes32 OWNER_KEY = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        vm.store(Predeploys.L2_ERC721_BRIDGE, OWNER_KEY, bytes32(uint256(uint160(multisig))));

        vm.prank(multisig);
        Proxy(payable(Predeploys.L2_ERC721_BRIDGE)).upgradeToAndCall(
            address(l2BridgeImpl), abi.encodeCall(L2ERC721Bridge.initialize, ())
        );

        // Set up a reference to the L2ERC721Bridge.
        L2NFTBridge = L2ERC721Bridge(Predeploys.L2_ERC721_BRIDGE);

        // Label the L1 and L2 bridges.
        vm.label(address(L1NFTBridge), "L1ERC721Bridge");
        vm.label(address(L2NFTBridge), "L2ERC721Bridge");
    }
}

contract FeeVault_Initializer is Bridge_Initializer {
    SequencerFeeVault vault = SequencerFeeVault(payable(Predeploys.SEQUENCER_FEE_WALLET));
    address constant recipient = address(2048);

    event Withdrawal(uint256 value, address to, address from);

    event Withdrawal(uint256 value, address to, address from, FeeVault.WithdrawalNetwork withdrawalNetwork);
}

interface IWithdrawalQueue {
    function finalize(uint256 _lastRequestIdToBeFinalized, uint256 _maxShareRate) external payable;
    function getWithdrawalRequests(address _owner) external view returns (uint256[] memory requestsIds);
}

contract LidoYieldProvider_Initializer is Bridge_Initializer {
    LidoYieldProvider internal lidoProvider;
    Insurance internal insurance;

    ILido Lido = ILido(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
    IWithdrawalQueue WithdrawalQueue = IWithdrawalQueue(0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1);

    function finalizeWithdrawals() internal {
        uint256[] memory _requestIds = WithdrawalQueue.getWithdrawalRequests(address(ethYieldManager));
        uint256 shareRate = Lido.getPooledEthByShares(1e27);
        vm.prank(address(Lido));
        WithdrawalQueue.finalize(_requestIds[_requestIds.length-1], shareRate);
    }

    function setUp() public virtual override {
        vm.createSelectFork(vm.envString("ETH_RPC_URL")); // 2023-12-20
        super.setUp();

        lidoProvider = new LidoYieldProvider(YieldManager(address(ethYieldManager)));
        vm.label(address(lidoProvider), "LidoYieldProvider");

        vm.label(address(Lido), "Lido");
        vm.label(address(WithdrawalQueue), "LidoWithdrawalQueue");

        L1ChugSplashProxy proxy = new L1ChugSplashProxy(multisig);
        vm.mockCall(multisig, abi.encodeWithSelector(IL1ChugSplashDeployer.isUpgrading.selector), abi.encode(true));
        vm.startPrank(multisig);
        proxy.setCode(address(new Insurance(ethYieldManager)).code);
        vm.clearMockedCalls();
        insurance = Insurance(payable(address(proxy)));
        insurance.initialize();
        vm.label(address(insurance), "Insurance");
        vm.stopPrank();

        vm.startPrank(multisig);
        ethYieldManager.addProvider(address(lidoProvider));
        ethYieldManager.setInsurance(address(insurance), 1000, 10); // rate = 10% and buffer = 10 wei
        vm.stopPrank();

        // faucet stETH
        vm.deal(address(this), 100 ether);
        Lido.submit{value: 100 ether}(address(0));
    }
}

interface IDAI is IERC20 {
    function mint(address, uint256) external;
}

interface IUSDC {
    function configureMinter(address, uint256) external;
    function mint(address, uint256) external;
    function owner() external view returns (address);
    function masterMinter() external view returns (address);
    function balanceOf(address) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
}

contract DSRYieldProvider_Initializer is Bridge_Initializer {
    IDsrManager public constant DSR_MANAGER = IDsrManager(0x373238337Bfe1146fb49989fc222523f83081dDb);
    IUSDC public constant RealUSDC = IUSDC(0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48);
    IDAI public constant RealDAI = IDAI(0x6B175474E89094C44Da98b954EedeAC495271d0F);

    DSRYieldProvider internal dsrProvider;
    Insurance internal insurance;

    function setUp() public virtual override {
        vm.createSelectFork(vm.envString("ETH_RPC_URL"), 18878218);
        bytes memory realUSDCImpl = address(0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48).code;
        bytes memory realDAIImpl = address(0x6B175474E89094C44Da98b954EedeAC495271d0F).code;

        super.setUp();

        vm.etch(0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48, realUSDCImpl);
        vm.etch(0x6B175474E89094C44Da98b954EedeAC495271d0F, realDAIImpl);

        // re-ran USDConversions._init() for RealUSDC and RealDAI
        address CURVE_3POOL = 0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7;
        address GEM_JOIN = 0x0A59649758aa4d66E25f08Dd01271e891fe52199;
        vm.startPrank(address(l1BlastBridge));
        RealUSDC.approve(address(CURVE_3POOL), type(uint256).max);
        RealUSDC.approve(GEM_JOIN, type(uint256).max);
        RealDAI.approve(address(CURVE_3POOL), type(uint256).max);
        RealDAI.approve(GEM_JOIN, type(uint256).max);
        vm.stopPrank();

        dsrProvider = new DSRYieldProvider(usdYieldManager);
        vm.label(address(dsrProvider), "DSRYieldProvider");

        vm.prank(multisig);
        insurance = new Insurance(usdYieldManager);
        vm.label(address(insurance), "Insurance");

        vm.startPrank(multisig);
        usdYieldManager.addProvider(address(dsrProvider));
        usdYieldManager.setInsurance(address(insurance), 1000, 0); // rate = 10% and buffer = 0 wei
        vm.stopPrank();

        IPot pot = IPot(DSR_MANAGER.pot());
        uint256 rho = pot.rho();
        if (rho > block.timestamp) {
            skip(rho - block.timestamp);
        }
        pot.drip();
    }
}

contract FFIInterface is Test {
    function getProveWithdrawalTransactionInputs(Types.WithdrawalTransaction memory _tx)
        external
        returns (bytes32, bytes32, bytes32, bytes32, bytes[] memory)
    {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "getProveWithdrawalTransactionInputs";
        cmds[3] = vm.toString(_tx.nonce);
        cmds[4] = vm.toString(_tx.sender);
        cmds[5] = vm.toString(_tx.target);
        cmds[6] = vm.toString(_tx.value);
        cmds[7] = vm.toString(_tx.gasLimit);
        cmds[8] = vm.toString(_tx.data);

        bytes memory result = vm.ffi(cmds);
        (
            bytes32 stateRoot,
            bytes32 storageRoot,
            bytes32 outputRoot,
            bytes32 withdrawalHash,
            bytes[] memory withdrawalProof
        ) = abi.decode(result, (bytes32, bytes32, bytes32, bytes32, bytes[]));

        return (stateRoot, storageRoot, outputRoot, withdrawalHash, withdrawalProof);
    }

    function hashCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    )
        external
        returns (bytes32)
    {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashCrossDomainMessage";
        cmds[3] = vm.toString(_nonce);
        cmds[4] = vm.toString(_sender);
        cmds[5] = vm.toString(_target);
        cmds[6] = vm.toString(_value);
        cmds[7] = vm.toString(_gasLimit);
        cmds[8] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashWithdrawal(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    )
        external
        returns (bytes32)
    {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashWithdrawal";
        cmds[3] = vm.toString(_nonce);
        cmds[4] = vm.toString(_sender);
        cmds[5] = vm.toString(_target);
        cmds[6] = vm.toString(_value);
        cmds[7] = vm.toString(_gasLimit);
        cmds[8] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashOutputRootProof(
        bytes32 _version,
        bytes32 _stateRoot,
        bytes32 _messagePasserStorageRoot,
        bytes32 _latestBlockhash
    )
        external
        returns (bytes32)
    {
        string[] memory cmds = new string[](7);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashOutputRootProof";
        cmds[3] = Strings.toHexString(uint256(_version));
        cmds[4] = Strings.toHexString(uint256(_stateRoot));
        cmds[5] = Strings.toHexString(uint256(_messagePasserStorageRoot));
        cmds[6] = Strings.toHexString(uint256(_latestBlockhash));

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashDepositTransaction(
        address _from,
        address _to,
        uint256 _mint,
        uint256 _value,
        uint64 _gas,
        bytes memory _data,
        uint64 _logIndex
    )
        external
        returns (bytes32)
    {
        string[] memory cmds = new string[](11);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "hashDepositTransaction";
        cmds[3] = "0x0000000000000000000000000000000000000000000000000000000000000000";
        cmds[4] = vm.toString(_logIndex);
        cmds[5] = vm.toString(_from);
        cmds[6] = vm.toString(_to);
        cmds[7] = vm.toString(_mint);
        cmds[8] = vm.toString(_value);
        cmds[9] = vm.toString(_gas);
        cmds[10] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function encodeDepositTransaction(Types.UserDepositTransaction calldata txn) external returns (bytes memory) {
        string[] memory cmds = new string[](12);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "encodeDepositTransaction";
        cmds[3] = vm.toString(txn.from);
        cmds[4] = vm.toString(txn.to);
        cmds[5] = vm.toString(txn.value);
        cmds[6] = vm.toString(txn.mint);
        cmds[7] = vm.toString(txn.gasLimit);
        cmds[8] = vm.toString(txn.isCreation);
        cmds[9] = vm.toString(txn.data);
        cmds[10] = vm.toString(txn.l1BlockHash);
        cmds[11] = vm.toString(txn.logIndex);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes));
    }

    function encodeCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    )
        external
        returns (bytes memory)
    {
        string[] memory cmds = new string[](9);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "encodeCrossDomainMessage";
        cmds[3] = vm.toString(_nonce);
        cmds[4] = vm.toString(_sender);
        cmds[5] = vm.toString(_target);
        cmds[6] = vm.toString(_value);
        cmds[7] = vm.toString(_gasLimit);
        cmds[8] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes));
    }

    function decodeVersionedNonce(uint256 nonce) external returns (uint256, uint256) {
        string[] memory cmds = new string[](4);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "decodeVersionedNonce";
        cmds[3] = vm.toString(nonce);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (uint256, uint256));
    }

    function getMerkleTrieFuzzCase(string memory variant)
        external
        returns (bytes32, bytes memory, bytes memory, bytes[] memory)
    {
        string[] memory cmds = new string[](6);
        cmds[0] = "./scripts/go-ffi/go-ffi";
        cmds[1] = "trie";
        cmds[2] = variant;

        return abi.decode(vm.ffi(cmds), (bytes32, bytes, bytes, bytes[]));
    }

    function getCannonMemoryProof(uint32 pc, uint32 insn) external returns (bytes32, bytes memory) {
        string[] memory cmds = new string[](5);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "cannonMemoryProof";
        cmds[3] = vm.toString(pc);
        cmds[4] = vm.toString(insn);
        bytes memory result = vm.ffi(cmds);
        (bytes32 memRoot, bytes memory proof) = abi.decode(result, (bytes32, bytes));
        return (memRoot, proof);
    }

    function getCannonMemoryProof(
        uint32 pc,
        uint32 insn,
        uint32 memAddr,
        uint32 memVal
    )
        external
        returns (bytes32, bytes memory)
    {
        string[] memory cmds = new string[](7);
        cmds[0] = "scripts/go-ffi/go-ffi";
        cmds[1] = "diff";
        cmds[2] = "cannonMemoryProof";
        cmds[3] = vm.toString(pc);
        cmds[4] = vm.toString(insn);
        cmds[5] = vm.toString(memAddr);
        cmds[6] = vm.toString(memVal);
        bytes memory result = vm.ffi(cmds);
        (bytes32 memRoot, bytes memory proof) = abi.decode(result, (bytes32, bytes));
        return (memRoot, proof);
    }
}

library EIP1967Helper {
    Vm internal constant vm = Vm(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);

    function getAdmin(address _proxy) internal view returns (address) {
        return address(uint160(uint256(vm.load(address(_proxy), Constants.PROXY_OWNER_ADDRESS))));
    }

    function getImplementation(address _proxy) internal view returns (address) {
        return address(uint160(uint256(vm.load(address(_proxy), Constants.PROXY_IMPLEMENTATION_ADDRESS))));
    }
}

// Used for testing a future upgrade beyond the current implementations.
// We include some variables so that we can sanity check accessing storage values after an upgrade.
contract NextImpl is Initializable {
    // Initializable occupies the zero-th slot.
    bytes32 slot1;
    bytes32[19] __gap;
    bytes32 slot21;
    bytes32 public constant slot21Init = bytes32(hex"1337");

    function initialize(uint8 _init) public reinitializer(_init) {
        // Slot21 is unused by an of our upgradeable contracts.
        // This is used to verify that we can access this value after an upgrade.
        slot21 = slot21Init;
    }
}

contract Reverter {
    fallback() external {
        revert();
    }
}

// Useful for testing reentrancy guards
contract CallerCaller {
    event WhatHappened(bool success, bytes returndata);

    fallback() external {
        (bool success, bytes memory returndata) = msg.sender.call(msg.data);
        emit WhatHappened(success, returndata);
        assembly {
            switch success
            case 0 { revert(add(returndata, 0x20), mload(returndata)) }
            default { return(add(returndata, 0x20), mload(returndata)) }
        }
    }
}

// Used for testing the `CrossDomainMessenger`'s per-message reentrancy guard.
contract ConfigurableCaller {
    bool doRevert = true;
    address target;
    bytes payload;

    event WhatHappened(bool success, bytes returndata);

    /// @notice Call the configured target with the configured payload OR revert.
    function call() external {
        if (doRevert) {
            revert("ConfigurableCaller: revert");
        } else {
            (bool success, bytes memory returndata) = address(target).call(payload);
            emit WhatHappened(success, returndata);
            assembly {
                switch success
                case 0 { revert(add(returndata, 0x20), mload(returndata)) }
                default { return(add(returndata, 0x20), mload(returndata)) }
            }
        }
    }

    /// @notice Set whether or not to have `call` revert.
    function setDoRevert(bool _doRevert) external {
        doRevert = _doRevert;
    }

    /// @notice Set the target for the call made in `call`.
    function setTarget(address _target) external {
        target = _target;
    }

    /// @notice Set the payload for the call made in `call`.
    function setPayload(bytes calldata _payload) external {
        payload = _payload;
    }

    /// @notice Fallback function that reverts if `doRevert` is true.
    ///        Otherwise, it does nothing.
    fallback() external {
        if (doRevert) {
            revert("ConfigurableCaller: revert");
        }
    }
}
