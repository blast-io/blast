// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Script } from "forge-std/Script.sol";
import { Vm } from "forge-std/Test.sol";

import { Safe } from "safe-contracts/Safe.sol";
import { SafeProxyFactory } from "safe-contracts/proxies/SafeProxyFactory.sol";
import { Enum as SafeOps } from "safe-contracts/common/Enum.sol";

import { console2 as console } from "forge-std/console2.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { Executables } from "scripts/Executables.sol";
import { ERC20Mock } from "@openzeppelin/contracts/mocks/ERC20Mock.sol";

import { Deployer } from "scripts/Deployer.sol";

import { DisputeGameFactory } from "src/dispute/DisputeGameFactory.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { AddressManager } from "src/legacy/AddressManager.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { L1StandardBridge } from "src/L1/L1StandardBridge.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { USDB } from "src/L2/USDB.sol";
import { Shares } from "src/L2/Shares.sol";
import { Gas } from "src/L2/Gas.sol";
import { Blast } from "src/L2/Blast.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";
import { L1BlastBridge } from "src/mainnet-bridge/L1BlastBridge.sol";
import { ETHYieldManager } from "src/mainnet-bridge/ETHYieldManager.sol";
import { USDYieldManager } from "src/mainnet-bridge/USDYieldManager.sol";
import { L2BlastBridge } from "src/mainnet-bridge/L2BlastBridge.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";
import { WithdrawalQueue } from "src/mainnet-bridge/withdrawal-queue/WithdrawalQueue.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { USDYieldManager } from "src/mainnet-bridge/USDYieldManager.sol";
import { WETHRebasing } from "src/L2/WETHRebasing.sol";
import { L1ChugSplashProxy } from "src/legacy/L1ChugSplashProxy.sol";
import { ResolvedDelegateProxy } from "src/legacy/ResolvedDelegateProxy.sol";
import { L1CrossDomainMessenger } from "src/L1/L1CrossDomainMessenger.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";
import { OptimismMintableERC20Factory } from "src/universal/OptimismMintableERC20Factory.sol";
import { Insurance } from "src/mainnet-bridge/Insurance.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Chains } from "scripts/Chains.sol";
import { LidoYieldProvider } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";
import { DSRYieldProvider } from "src/mainnet-bridge/yield-providers/DSRYieldProvider.sol";
import { ETHTestnetYieldProvider } from "src/mainnet-bridge/yield-providers/ETHTestnetYieldProvider.sol";
import { USDTestnetYieldProvider } from "src/mainnet-bridge/yield-providers/USDTestnetYieldProvider.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";

interface IToken is IERC20 {
    function mint(address to, uint256 amount) external returns (bool);
}

interface IUSDT {
    function transfer(address to, uint256 amount) external;
    function balanceOf(address) external view returns (uint256);
    function approve(address, uint256) external;
}

interface ILido is IERC20 {
    function submit(address referral) external payable;
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool);
    function isStakingPaused() external view returns (bool);
}

contract ScriptInitializer is Deployer {
    ProxyAdmin proxyAdmin;
    OptimismPortal op;
    ETHYieldManager eym;
    USDYieldManager uym;
    L1BlastBridge l1bb;
    L2BlastBridge l2bb;
    IToken dai;
    IToken usdc;
    IUSDT usdt;
    ILido lido;
    USDB usdb;
    Shares shares;
    L2OutputOracle oracle;
    LidoYieldProvider lyp;
    DSRYieldProvider dyp;
    ETHTestnetYieldProvider eyp;
    USDTestnetYieldProvider uyp;
    ERC20Mock usd;
    ERC20Mock steth;
    Insurance ei;
    Insurance ui;
    Blast b;
    Gas gas;

    function name() public pure virtual override returns (string memory name_) {
        name_ = "Script";
    }

    function setUp() public virtual override {
        super.setUp();

        if (_chainIsL1()) {
            proxyAdmin = ProxyAdmin(mustGetAddress("ProxyAdmin"));
            vm.label(address(proxyAdmin), "ProxyAdmin");
            op = OptimismPortal(payable(mustGetAddress("OptimismPortalProxy")));
            vm.label(address(op), "OptimismPortal");
            eym = ETHYieldManager(payable(mustGetAddress("ETHYieldManagerProxy")));
            vm.label(address(eym), "ETHYieldManager");
            uym = USDYieldManager(payable(getAddress("USDYieldManagerProxy")));
            vm.label(address(uym), "USDYieldManager");
            l1bb = L1BlastBridge(payable(mustGetAddress("L1BlastBridgeProxy")));
            vm.label(address(l1bb), "L1BlastBridge");
            oracle = L2OutputOracle(mustGetAddress("L2OutputOracleProxy"));
            vm.label(address(oracle), "L2OutputOracle");
            ei = Insurance(getAddress("ETHInsuranceProxy"));
            vm.label(address(ei), "ETHInsurance");
            ui = Insurance(getAddress("USDInsuranceProxy"));
            vm.label(address(ui), "USDInsurance");
            usd = ERC20Mock(getAddress("USDToken"));
            vm.label(address(usd), "USDToken");
            steth = ERC20Mock(getAddress("ETHYieldToken"));
            vm.label(address(steth), "ETHYieldToken");
            eyp = ETHTestnetYieldProvider(getAddress("ETHTestnetYieldProvider"));
            vm.label(address(eyp), "ETHTestnetYieldProvider");
            uyp = USDTestnetYieldProvider(getAddress("USDTestnetYieldProvider"));
            vm.label(address(uyp), "USDTestnetYieldProvider");
        }

        if (_isFork()) {
            lyp = LidoYieldProvider(mustGetAddress("LidoYieldProvider"));
            vm.label(address(lyp), "LidoYieldProvider");
            dyp = DSRYieldProvider(mustGetAddress("DSRYieldProvider"));
            vm.label(address(dyp), "DSRYieldProvider");
            lido = ILido(payable(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84));
            vm.label(address(lido), "Lido");
            dai = IToken(address(USDConversions.DAI));
            vm.label(address(dai), "DAI");
            usd = ERC20Mock(address(dai));
            vm.label(address(usd), "USDToken");
            usdc = IToken(address(USDConversions.USDC));
            vm.label(address(usdc), "USDC");
            usdt = IUSDT(address(USDConversions.USDT));
            vm.label(address(usdt), "USDT");
        }

        l2bb = L2BlastBridge(payable(Predeploys.L2_BLAST_BRIDGE));
        vm.label(address(l2bb), "L2BlastBridge");
        usdb = USDB(Predeploys.USDB);
        vm.label(address(usdb), "USDB");
        shares = Shares(Predeploys.SHARES);
        vm.label(address(shares), "Shares");
        gas = Gas(Predeploys.GAS);
        vm.label(address(gas), "Gas");
        b = Blast(Predeploys.BLAST);
        vm.label(address(b), "Blast");
    }

    /// @notice Call from the Safe contract to the Proxy Admin's upgrade and call method
    function _upgradeAndCallViaSafe(address _proxy, address _implementation, bytes memory _innerCallData) internal {
        address proxyAdmin = mustGetAddress("ProxyAdmin");

        bytes memory data =
            abi.encodeCall(ProxyAdmin.upgradeAndCall, (payable(_proxy), _implementation, _innerCallData));

        _callViaSafe({ _target: proxyAdmin, _data: data });
    }

    // @notice Make a call from the Safe contract to an arbitrary address with arbitrary data
    function _callViaSafe(address _target, bytes memory _data) internal {
        Safe safe = Safe(mustGetAddress("SystemOwnerSafe"));

        // This is the signature format used the caller is also the signer.
        bytes memory signature = abi.encodePacked(uint256(uint160(msg.sender)), bytes32(0), uint8(1));

        safe.execTransaction({
            to: _target,
            value: 0,
            data: _data,
            operation: SafeOps.Operation.Call,
            safeTxGas: 0,
            baseGas: 0,
            gasPrice: 0,
            gasToken: address(0),
            refundReceiver: payable(address(0)),
            signatures: signature
        });
    }

    function _deployBehindProxy(address implementation, bytes memory initializeCalldata, string memory name) internal {
        address proxyAdmin = mustGetAddress("ProxyAdmin");
        L1ChugSplashProxy proxy = new L1ChugSplashProxy(proxyAdmin);

        address admin = address(uint160(uint256(vm.load(address(proxy), OWNER_KEY))));
        require(admin == proxyAdmin);

        _upgradeAndCallViaSafe({
            _proxy: payable(address(proxy)),
            _implementation: implementation,
            _innerCallData: initializeCalldata
        });

        string memory version = DisputeGameFactory(mustGetAddress("DisputeGameFactoryProxy")).version();
        save(name, address(implementation));
        save(string.concat(name, "Proxy"), address(proxy));
    }

    /*//////////////////////////////////////
                    HELPERS
    //////////////////////////////////////*/

    function sendETH(address to, uint256 amount) internal {
        (bool success,) = to.call{value: amount}("");
        require(success, "ETH transfer failed");
    }

    function _chainIsL1() internal returns (bool) {
        return _compareStrings(_getDeploymentContext(), "devnetL1") || _compareStrings(_getDeploymentContext(), "sepolia") || _compareStrings(_getDeploymentContext(), "mainnet");
    }

    function _isFork() internal view returns (bool) {
        return 0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84.code.length > 0;
    }

    function _compareStrings(string memory _a, string memory _b) internal pure returns (bool) {
        return keccak256(abi.encodePacked(_a)) == keccak256(abi.encodePacked(_b));
    }

    modifier onlyL1() {
        require(_chainIsL1(), "Function can only be called on L1");
        _;
    }

    modifier onlyL2() {
        require(_compareStrings(_getDeploymentContext(), "901"), "Function can only be called on L2");
        _;
    }

    modifier broadcast() {
        vm.startBroadcast();
        _;
        vm.stopBroadcast();
    }

    modifier broadcastAddr(address wallet) {
        vm.startBroadcast(wallet);
        _;
        vm.stopBroadcast();
    }
}

contract E2EInitializer is ScriptInitializer {

    function name() public pure override returns (string memory name_) {
        name_ = "E2E";
    }

    address user;
    address alice = vm.addr(0x01);

    function setUp() public virtual override {
        super.setUp();

        user = msg.sender;
        if (_chainIsL1()) {
            if (vm.exists("l1state.json")) {
                _loadL1State();
            }
        } else {
            if (vm.exists("l2state.json")) {
                _loadL2State();
            }
        }
    }

    function _incrementL1NumCalls() internal {
        l1State.numCalls += 1;
        _storeL1State();
    }

    function _incrementL2NumCalls() internal {
        l2State.numCalls += 1;
        _storeL1State();
    }

    function check() external view virtual returns (bool) {
        return true;
    }

    struct L1State {
        uint256 numCalls;
        uint256 userETHBalance;
        uint256 userUSDBalance;
        uint256 userStETHBalance;
        uint256 aliceETHBalance;
        uint256 aliceUSDBalance;
        uint256 aliceStETHBalance;
        uint256 portalBalance;
        uint256 eymBalance;
        uint256 userDAIBalance;
        uint256 userUSDCBalance;
        uint256 userUSDTBalance;
        uint256 aliceDAIBalance;
        uint256 aliceUSDCBalance;
        uint256 aliceUSDTBalance;
        uint256 uymUSDBalance;
        uint256 uymDAIBalance;
        uint256 bridgeUSDBalance;
        uint256 bridgeDAIBalance;
        uint256 eymNegativeYields;
        uint256 uymNegativeYields;
    }
    struct L2State {
        uint256 numCalls;
        uint256 userETHBalance;
        uint256 userUSDBBalance;
        uint256 aliceETHBalance;
        uint256 aliceUSDBBalance;
        uint256 ethSharePrice;
        uint256 ethPending;
        uint256 usdSharePrice;
        uint256 usdPending;
    }

    L1State l1State;
    L2State l2State;

    function createL1State(address _user) external {
        user = _user;
        if (_isFork()) {
            l1State = L1State({
                numCalls: 0,
                userETHBalance: user.balance,
                userUSDBalance: 0,
                userStETHBalance: lido.balanceOf(user),
                aliceETHBalance: alice.balance,
                aliceUSDBalance: 0,
                aliceStETHBalance: lido.balanceOf(alice),
                portalBalance: address(op).balance,
                eymBalance: address(eym).balance,
                userDAIBalance: dai.balanceOf(user),
                userUSDCBalance: usdc.balanceOf(user),
                userUSDTBalance: usdt.balanceOf(user),
                aliceDAIBalance: dai.balanceOf(alice),
                aliceUSDCBalance: usdc.balanceOf(alice),
                aliceUSDTBalance: usdt.balanceOf(alice),
                uymUSDBalance: 0,
                uymDAIBalance: dai.balanceOf(address(uym)),
                bridgeUSDBalance: 0,
                bridgeDAIBalance: dai.balanceOf(address(l1bb)),
                eymNegativeYields: eym.accumulatedNegativeYields(),
                uymNegativeYields: uym.accumulatedNegativeYields()
            });
        } else {
            l1State = L1State({
                numCalls: 0,
                userETHBalance: user.balance,
                userUSDBalance: usd.balanceOf(user),
                userStETHBalance: steth.balanceOf(user),
                aliceETHBalance: alice.balance,
                aliceUSDBalance: usd.balanceOf(alice),
                aliceStETHBalance: steth.balanceOf(alice),
                portalBalance: address(op).balance,
                eymBalance: address(eym).balance,
                userDAIBalance: 0,
                userUSDCBalance: 0,
                userUSDTBalance: 0,
                aliceDAIBalance: 0,
                aliceUSDCBalance: 0,
                aliceUSDTBalance: 0,
                uymUSDBalance: usd.balanceOf(address(uym)),
                uymDAIBalance: 0,
                bridgeUSDBalance: usd.balanceOf(address(l1bb)),
                bridgeDAIBalance: 0,
                eymNegativeYields: eym.accumulatedNegativeYields(),
                uymNegativeYields: uym.accumulatedNegativeYields()
            });
        }
        _storeL1State();
    }

    function createL2State(address _user) external {
        user = _user;
        l2State = L2State(
            0,
            user.balance,
            usdb.balanceOf(user),
            alice.balance,
            usdb.balanceOf(alice),
            shares.price(),
            shares.pending(),
            usdb.price(),
            usdb.pending()
        );
        _storeL2State();
    }

    function fundAccount(address _user) public broadcast {
        sendETH(_user, 2 ether);
        if (_isFork()) {
            dai.transfer(_user, 100 ether);
            usdc.transfer(_user, 100e6);
            usdt.transfer(_user, 100e6);
        } else {
            usd.mint(_user, 100 ether);
            steth.mint(_user, 100 ether);
        }
    }

    function _storeL1State() internal {
        string memory json = "";
        vm.serializeUint(json, "numCalls", l1State.numCalls);
        vm.serializeUint(json, "userETHBalance", l1State.userETHBalance);
        vm.serializeUint(json, "userUSDBalance", l1State.userUSDBalance);
        vm.serializeUint(json, "userStETHBalance", l1State.userStETHBalance);
        vm.serializeUint(json, "userDAIBalance", l1State.userDAIBalance);
        vm.serializeUint(json, "userUSDCBalance", l1State.userUSDCBalance);
        vm.serializeUint(json, "userUSDTBalance", l1State.userUSDTBalance);
        vm.serializeUint(json, "aliceETHBalance", l1State.aliceETHBalance);
        vm.serializeUint(json, "aliceUSDBalance", l1State.aliceUSDBalance);
        vm.serializeUint(json, "aliceStETHBalance", l1State.aliceStETHBalance);
        vm.serializeUint(json, "aliceDAIBalance", l1State.aliceDAIBalance);
        vm.serializeUint(json, "aliceUSDCBalance", l1State.aliceUSDCBalance);
        vm.serializeUint(json, "aliceUSDTBalance", l1State.aliceUSDTBalance);
        vm.serializeUint(json, "portalBalance", l1State.portalBalance);
        vm.serializeUint(json, "eymBalance", l1State.eymBalance);
        vm.serializeUint(json, "uymDAIBalance", l1State.uymDAIBalance);
        vm.serializeUint(json, "uymUSDBalance", l1State.uymUSDBalance);
        vm.serializeUint(json, "bridgeUSDBalance", l1State.bridgeUSDBalance);
        vm.serializeUint(json, "eymNegativeYields", l1State.eymNegativeYields);
        vm.serializeUint(json, "uymNegativeYields", l1State.uymNegativeYields);
        json = vm.serializeUint(json, "bridgeDAIBalance", l1State.bridgeDAIBalance);
        vm.writeJson({ json: json, path: "l1state.json" });
    }

    function _storeL2State() internal {
        string memory json = "";
        vm.serializeUint(json, "numCalls", l2State.numCalls);
        vm.serializeUint(json, "userETHBalance", l2State.userETHBalance);
        vm.serializeUint(json, "userUSDBBalance", l2State.userUSDBBalance);
        vm.serializeUint(json, "aliceETHBalance", l2State.aliceETHBalance);
        vm.serializeUint(json, "aliceUSDBBalance", l2State.aliceUSDBBalance);
        vm.serializeUint(json, "ethSharePrice", l2State.ethSharePrice);
        vm.serializeUint(json, "ethPending", l2State.ethPending);
        vm.serializeUint(json, "usdSharePrice", l2State.usdSharePrice);
        json = vm.serializeUint(json, "usdPending", l2State.usdPending);
        vm.writeJson({ json: json, path: "l2state.json" });
    }

    function _loadL1State() internal {
        string memory json = vm.readFile("l1state.json");
        l1State = L1State(
            vm.parseJsonUint(json, "$.numCalls"),
            vm.parseJsonUint(json, "$.userETHBalance"),
            vm.parseJsonUint(json, "$.userUSDBalance"),
            vm.parseJsonUint(json, "$.userStETHBalance"),
            vm.parseJsonUint(json, "$.aliceETHBalance"),
            vm.parseJsonUint(json, "$.aliceUSDBalance"),
            vm.parseJsonUint(json, "$.aliceStETHBalance"),
            vm.parseJsonUint(json, "$.portalBalance"),
            vm.parseJsonUint(json, "$.eymBalance"),
            vm.parseJsonUint(json, "$.userDAIBalance"),
            vm.parseJsonUint(json, "$.userUSDCBalance"),
            vm.parseJsonUint(json, "$.userUSDTBalance"),
            vm.parseJsonUint(json, "$.aliceDAIBalance"),
            vm.parseJsonUint(json, "$.aliceUSDCBalance"),
            vm.parseJsonUint(json, "$.aliceUSDTBalance"),
            vm.parseJsonUint(json, "$.uymUSDBalance"),
            vm.parseJsonUint(json, "$.uymDAIBalance"),
            vm.parseJsonUint(json, "$.bridgeUSDBalance"),
            vm.parseJsonUint(json, "$.bridgeDAIBalance"),
            vm.parseJsonUint(json, "$.eymNegativeYields"),
            vm.parseJsonUint(json, "$.uymNegativeYields")
        );
    }

    function _loadL2State() internal {
        string memory json = vm.readFile("l2state.json");
        l2State = L2State(
            vm.parseJsonUint(json, "$.numCalls"),
            vm.parseJsonUint(json, "$.userETHBalance"),
            vm.parseJsonUint(json, "$.userUSDBBalance"),
            vm.parseJsonUint(json, "$.aliceETHBalance"),
            vm.parseJsonUint(json, "$.aliceUSDBBalance"),
            vm.parseJsonUint(json, "$.ethSharePrice"),
            vm.parseJsonUint(json, "$.ethPending"),
            vm.parseJsonUint(json, "$.usdSharePrice"),
            vm.parseJsonUint(json, "$.ethPending")
        );
    }

    function _getLog(Vm.Log[] memory logs, bytes32 selector) internal pure returns (Vm.Log memory log) {
        for (uint i; i < logs.length; i++) {
            if (logs[i].topics[0] == selector) {
                log = logs[i];
                break;
            }
        }
    }
}
