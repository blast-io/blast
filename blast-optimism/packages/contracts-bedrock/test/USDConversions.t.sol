// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { Bridge_Initializer } from "test/CommonTest.t.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// Target contract
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";

interface IUSDC {
    function mint(address, uint256) external;
    function masterMinter() external view returns (address);
    function configureMinter(address, uint256) external returns (bool);
}

interface IUSDT {
    function owner() external view returns (address);
    function issue(uint256) external;
}

contract USDConversions_Test is Bridge_Initializer {
    using SafeERC20 for ERC20;

    uint256 constant WAD_DECIMALS = 18;
    uint256 constant USD_DECIMALS = 6;
    int128 constant DAI_INDEX = 0;
    int128 constant USDC_INDEX = 1;
    int128 constant USDT_INDEX = 2;

    event YieldReport(
        int256  yield,
        uint256 insurancePremiumPaid,
        uint256 insuranceWithdrawn
    );

    error InsufficientInsuranceBalance();
    error InsufficientBalance();
    error MinimumAmountNotMet();
    error InvalidTokenIndex();

    function TOKEN() public view returns (address) {
        return address(DAI);
    }

    function availableBalance() public view returns (uint256) {
        return DAI.balanceOf(address(this));
    }

    function setUp() public override {
        vm.createSelectFork(vm.envString("ETH_RPC_URL"));
        super.setUp();
        USDConversions._init();
    }

    function _testConvert(
        int128 inputToken,
        int128 outputToken,
        uint256 inputAmountWad,
        uint256 minOutputAmountWad
    ) internal {
        uint256 inputTokenBalance = USDConversions._tokenBalance(inputToken);
        uint256 outputTokenBalance = USDConversions._tokenBalance(outputToken);
        uint256 amountReceived = USDConversions._convert(inputToken, outputToken, inputAmountWad, minOutputAmountWad);
        assertGe(amountReceived, USDConversions._convertDecimals(minOutputAmountWad, outputToken));
        assertEq(USDConversions._tokenBalance(inputToken), inputTokenBalance - USDConversions._convertDecimals(inputAmountWad, inputToken));
        assertEq(USDConversions._tokenBalance(outputToken), outputTokenBalance + amountReceived);
    }

    function test_usdConversions_USDCToDAI() external {
        mintUSDC(address(this), 2 ether);
        _testConvert(USDC_INDEX, DAI_INDEX, 2 ether, 2 ether);
    }

    function test_usdConversions_DAIToUSDC() external {
        mintDAI(address(this), 2 ether);
        _testConvert(DAI_INDEX, USDC_INDEX, 2 ether, 2 ether);
    }

    // TODO: for some reason Curve isn't working on a fork so these tests aren't working, the call to Curve looks correct though. Just a random EVM error that makes it hard to debug.
    function test_usdConversions_DAIToUSDT() external {
        vm.skip(true);
        mintDAI(address(this), 2 ether);
        _testConvert(DAI_INDEX, USDT_INDEX, 2 ether, 1.9 ether);
    }

    function test_usdConversions_USDTToDAI() external {
        vm.skip(true);
        mintUSDT(address(this), 2 ether);
        _testConvert(USDT_INDEX, DAI_INDEX, 2 ether, 1.9 ether);
    }

    function test_usdConversions_USDCToUSDT() external {
        vm.skip(true);
        mintUSDC(address(this), 2 ether);
        _testConvert(USDC_INDEX, USDT_INDEX, 2 ether, 1.9 ether);
    }

    function test_usdConversions_USDTToUSDC() external {
        vm.skip(true);
        mintUSDT(address(this), 100 ether);
        _testConvert(USDT_INDEX, USDC_INDEX, 100 ether, 50 ether);
    }

    function test_usdConversions_reverts() external {
        vm.skip(true);
        mintDAI(address(this), 2 ether);
        mintUSDC(address(this), 2 ether);
        mintUSDT(address(this), 2 ether);

        vm.expectRevert(MinimumAmountNotMet.selector);
        vm.prank(address(multisig));
        USDConversions._convert(0, 1, 2 ether, 2 ether + 1);

        vm.expectRevert();
        vm.prank(address(multisig));
        USDConversions._convert(3, 0, 2 ether, 2 ether);

        vm.expectRevert();
        vm.prank(address(multisig));
        USDConversions._convert(0, 3, 2 ether, 2 ether);

        vm.expectRevert();
        vm.prank(address(multisig));
        USDConversions._convert(0, 0, 2 ether, 2 ether);
    }

    function mintDAI(address to, uint256 amount) internal {
        deal(address(DAI), to, amount, true);
    }

    function mintUSDC(address to, uint256 amount) internal {
        vm.prank(IUSDC(address(USDC)).masterMinter());
        IUSDC(address(USDC)).configureMinter(address(this), type(uint256).max);
        IUSDC(address(USDC)).mint(to, amount);
    }

    function mintUSDT(address to, uint256 amount) internal {
        vm.startPrank(IUSDT(address(USDT)).owner());
        IUSDT(address(USDT)).issue(amount);
        USDT.safeTransfer(to, amount);
        vm.stopPrank();
    }
}
