// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Script } from "forge-std/Script.sol";

import { console2 as console } from "forge-std/console2.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { Executables } from "scripts/Executables.sol";
import { ERC20Mock } from "@openzeppelin/contracts/mocks/ERC20Mock.sol";

import { Deployer } from "scripts/Deployer.sol";

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

contract E2EInitializer is Deployer {

    function name() public pure override returns (string memory name_) {
        name_ = "E2E";
    }

    address user;
    address alice = vm.addr(0x01);

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


    function setUp() public override {
        super.setUp();

        user = msg.sender;
        if (vm.exists("step.json")) {
            _loadStep();
        }

        if (_chainIsL1()) {
            op = OptimismPortal(payable(mustGetAddress("OptimismPortalProxy")));
            vm.label(address(op), "OptimismPortal");
            eym = ETHYieldManager(payable(mustGetAddress("ETHYieldManagerProxy")));
            vm.label(address(eym), "ETHYieldManager");
            uym = USDYieldManager(payable(mustGetAddress("USDYieldManagerProxy")));
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

            if (vm.exists("l1state.json")) {
                _loadL1State();
            }
        } else {
            if (vm.exists("l2state.json")) {
                _loadL2State();
            }
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

    struct Step {
        uint256 t;
    }

    struct L1State {
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
    }
    struct L2State {
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
    Step step;

    function createL1State(address _user) external {
        user = _user;
        if (_isFork()) {
            l1State = L1State(
                user.balance,
                0,
                lido.balanceOf(user),
                alice.balance,
                0,
                lido.balanceOf(alice),
                address(op).balance,
                address(eym).balance,
                dai.balanceOf(user),
                usdc.balanceOf(user),
                usdt.balanceOf(user),
                dai.balanceOf(alice),
                usdc.balanceOf(alice),
                usdt.balanceOf(alice),
                0,
                dai.balanceOf(address(uym)),
                0,
                dai.balanceOf(address(l1bb))
            );
        } else {
            l1State = L1State(
                user.balance,
                usd.balanceOf(user),
                steth.balanceOf(user),
                alice.balance,
                usd.balanceOf(alice),
                steth.balanceOf(alice),
                address(op).balance,
                address(eym).balance,
                0,
                0,
                0,
                0,
                0,
                0,
                usd.balanceOf(address(uym)),
                0,
                usd.balanceOf(address(l1bb)),
                0
            );
        }
        _storeL1State();
    }

    function createL2State(address _user) external {
        user = _user;
        l2State = L2State(
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

    function createStep(uint256 t) public {
        step = Step(t);
        _storeStep();
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
        json = vm.serializeUint(json, "bridgeDAIBalance", l1State.bridgeDAIBalance);
        vm.writeJson({ json: json, path: "l1state.json" });
    }

    function _storeL2State() internal {
        string memory json = "";
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

    function _storeStep() internal {
        string memory json = "";
        json = vm.serializeUint(json, "t", step.t);
        vm.writeJson({ json: json, path: "step.json" });
    }

    function _loadL1State() internal {
        string memory json = vm.readFile("l1state.json");
        l1State = L1State(
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
            vm.parseJsonUint(json, "$.bridgeDAIBalance")
        );
    }

    function _loadL2State() internal {
        string memory json = vm.readFile("l2state.json");
        l2State = L2State(
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

    function _loadStep() internal {
        string memory json = vm.readFile("step.json");
        step = Step(
            vm.parseJsonUint(json, "$.t")
        );
    }


    /*//////////////////////////////////////
                    HELPERS
    //////////////////////////////////////*/

    function sendETH(address to, uint256 amount) internal {
        (bool success,) = to.call{value: amount}("");
        require(success, "ETH transfer failed");
    }

    function _chainIsL1() internal returns (bool) {
        return _compareStrings(_getDeploymentContext(), "devnetL1") || _compareStrings(_getDeploymentContext(), "sepolia");
    }


    function _isFork() internal view returns (bool) {
        return 0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84.code.length > 0;
    }

    function _compareStrings(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
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
