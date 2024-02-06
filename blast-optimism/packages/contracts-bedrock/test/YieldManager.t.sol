pragma solidity 0.8.15;
// Testing utilities
import { Test, StdUtils, Vm, StdStorage, stdStorage } from "forge-std/Test.sol";
import {
    Portal_Initializer,
    LidoYieldProvider_Initializer,
    DSRYieldProvider_Initializer,
    CommonTest,
    NextImpl
} from "test/CommonTest.t.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { console } from "forge-std/console.sol";

// Target contract dependencies
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";
import { LidoYieldProvider, ILido, IERC20, IWithdrawalQueue } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";
import { DSRYieldProvider, IDsrManager, IPot } from "src/mainnet-bridge/yield-providers/DSRYieldProvider.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";

// Target contract
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";

contract Util is Test {
    function assertClose(int256 a, int256 b) internal {
        int256 min = a < b ? a : b;
        if (min >= 0) {
            assertApproxEqRel(uint256(a), uint256(b), 1.0e12);
        } else {
            assertApproxEqRel(uint256(a + min), uint256(b + min), 1.0e12);
        }
    }

    function assertClose(uint256 a, uint256 b) internal {
        assertApproxEqRel(a, b, 1.0e12);
    }

    function addressToBytes32(address a) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(uint256(uint160(a))));
    }
}

contract ETH_YieldManager_Test is LidoYieldProvider_Initializer, Util {
    using stdStorage for StdStorage;

    event LidoUnstakeInitiated(uint256 indexed requestId, uint256 amount);
    event YieldReport(
        int256  yield,
        uint256 insurancePremiumPaid,
        uint256 insuranceWithdrawn
    );

    error InsufficientInsuranceBalance();

    function emulatePositiveYield(uint256 amount) internal {
        ILido(LidoAddress).transfer(address(ethYieldManager), amount);
    }

    function test_getProviderInfo_succeeds() external {
        emulatePositiveYield(1 ether);

        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.providerAddress, address(lidoProvider));
        assertEq(info.id, keccak256(abi.encodePacked("LidoYieldProvider", string(abi.encodePacked("1.0.0")))));
        assertEq(info.stakedBalance, uint256(0));
        assertClose(info.stakedValue, 1 ether);
        assertClose(uint256(info.yield), 1 ether);
    }

    function test_stake_Lido_succeeds() external {
        vm.deal(address(ethYieldManager), 1 ether);

        uint256 lockedValue = ethYieldManager.lockedValue();
        assertEq(lockedValue, 1 ether);

        vm.prank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 0.4 ether);
        assertEq(address(ethYieldManager).balance, 0.6 ether);

        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.stakedBalance, uint256(0.4 ether));
        assertClose(info.stakedValue, uint256(0.4 ether));
        assertClose(info.yield, int256(0 ether));
        assertEq(ethYieldManager.lockedValue(), uint256(0.6 ether));
        uint256 totalEth = ethYieldManager.totalValue();
        assertClose(totalEth, 1 ether);
    }

    function test_unstake_Lido_succeeds() external {
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 1 ether);

        uint256 lockedValue = ethYieldManager.lockedValue();
        assertEq(lockedValue, 1 ether);

        vm.prank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 0.4 ether);

        vm.expectEmit(false, false, false, true, address(ethYieldManager));
        emit LidoUnstakeInitiated(1234, 0.2 ether);

        vm.prank(multisig);
        ethYieldManager.unstake(0, address(lidoProvider), 0.2 ether);

        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.pendingBalance, uint256(0.2 ether));
        assertEq(info.stakedBalance, uint256(0.2 ether));
    }

    function test_claimPending_Lido_succeeds() external {
        vm.pauseGasMetering();
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 1 ether);

        uint256 lockedValue = ethYieldManager.lockedValue();
        assertEq(lockedValue, 1 ether);

        vm.startPrank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 0.4 ether);

        ethYieldManager.unstake(0, address(lidoProvider), 0.3 ether);

        setUpMockLidoClaim(1, 1, 0.2 ether, address(ethYieldManager));
        uint256[] memory requestIds = new uint256[](1);
        requestIds[0] = 1;
        vm.recordLogs();
        ethYieldManager.claimPending(0, address(lidoProvider), requestIds);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // ensure there's no DepositedTransaction event
        assertEq(entries.length, 1);
        assertEq(entries[0].topics[0], keccak256("Claimed(bytes32,uint256)"));
        assertEq(entries[0].topics[1], YieldProvider(lidoProvider).id());
        uint256 claimed = abi.decode(entries[0].data, (uint256));
        assertEq(claimed, uint256(0.2 ether));

        assertEq(ethYieldManager.lockedValue(), 0.8 ether);
        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.pendingBalance, uint256(0.1 ether));  // 0.3 - 0.2
        assertEq(info.stakedBalance, uint256(0.1 ether));   // 0.4 - 0.3
        assertClose(info.stakedValue, uint256(0.1 ether));  // 0.4 - 0.3
        assertClose(info.totalValue, uint256(0.2 ether));   // 0.4 - 0.2
    }

    function test_recordStakedDeposit_Lido_succeeds() external {
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 0.4 ether);
        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.stakedBalance, uint256(0.4 ether));
    }

    function test_commitYieldReport_positive_yield_no_insurance_succeeds() external {
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 0.4 ether);
        emulatePositiveYield(0.5 ether);

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        assertEq(entries.length, 3);

        assertEq(entries[0].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[0].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[0].data, (int256));
        assertClose(yield, int256(0.1 ether));

        assertEq(entries[1].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, , ) = abi.decode(entries[1].data, (int256, uint256, uint256));
        assertEq(totalYield, yield);

        assertEq(entries[2].topics[0], keccak256("TransactionDeposited(address,address,uint256,bytes)"));
        assertEq(entries[2].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager))));
        assertEq(entries[2].topics[2], addressToBytes32(0x4300000000000000000000000000000000000000));
    }

    function test_commitYieldReport_positive_yield_insurance_premium_succeeds() external {
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 0.4 ether);
        emulatePositiveYield(0.5 ether);

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // first two are Lido transfer events
        assertEq(entries.length, 6);

        assertEq(entries[2].topics[0], keccak256("InsurancePremiumPaid(bytes32,uint256)"));
        uint256 insuranceAmount = abi.decode(entries[2].data, (uint256));
        assertClose(insuranceAmount, uint256(0.01 ether));

        assertEq(entries[3].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[3].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[3].data, (int256));
        assertClose(yield, int256(0.09 ether));

        assertEq(entries[4].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[4].data, (int256, uint256, uint256));
        assertClose(totalYield, int256(0.09 ether));
        assertClose(insurancePremiumPaid, uint256(0.01 ether));
        assertEq(insuranceWithdrawn, 0);

        assertEq(entries[5].topics[0], keccak256("TransactionDeposited(address,address,uint256,bytes)"));
        assertEq(entries[5].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager))));
        assertEq(entries[5].topics[2], addressToBytes32(0x4300000000000000000000000000000000000000));

        uint256 insuranceBalance = ILido(LidoAddress).balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(0.01 ether));
    }

    function test_commitYieldReport_payoff_and_positive_yield_succeeds() external {
        vm.prank(address(l1BlastBridge));

        // 1 ETH yield
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 4 ether);
        emulatePositiveYield(5 ether);

        // 0.2 ETH accumulated negative yield
        uint256 slot = stdstore.target(address(ethYieldManager)).sig("accumulatedNegativeYields()").find();
        vm.store(address(ethYieldManager), bytes32(slot), bytes32(uint256(0.2 ether)));

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // first two are Lido transfer events
        assertEq(entries.length, 6);

        // expecting
        //  - 0.1 ETH insurance premium
        //  - 0.9 ETH committed yield from Lido
        //  - 0.7 ETH total yield in the report (0.2 ETH negative yields paid off)
        assertEq(entries[2].topics[0], keccak256("InsurancePremiumPaid(bytes32,uint256)"));
        uint256 insuranceAmount = abi.decode(entries[2].data, (uint256));
        assertClose(insuranceAmount, uint256(0.1 ether));

        assertEq(entries[3].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[3].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[3].data, (int256));
        assertClose(yield, int256(0.9 ether));

        assertEq(entries[4].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[4].data, (int256, uint256, uint256));
        assertClose(totalYield, int256(0.7 ether));
        assertClose(insurancePremiumPaid, uint256(0.1 ether));
        assertEq(insuranceWithdrawn, 0);

        assertEq(entries[5].topics[0], keccak256("TransactionDeposited(address,address,uint256,bytes)"));
        assertEq(entries[5].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager))));
        assertEq(entries[5].topics[2], addressToBytes32(0x4300000000000000000000000000000000000000));

        uint256 insuranceBalance = ILido(LidoAddress).balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(0.1 ether));

        assertEq(ethYieldManager.accumulatedNegativeYields(), 0);
    }

    function test_commitYieldReport_payoff_only_succeeds() external {
        vm.prank(address(l1BlastBridge));

        // 1 ETH yield
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 4 ether);
        emulatePositiveYield(5 ether);

        // 1.2 ETH accumulated negative yield
        uint256 slot = stdstore.target(address(ethYieldManager)).sig("accumulatedNegativeYields()").find();
        vm.store(address(ethYieldManager), bytes32(slot), bytes32(uint256(1.2 ether)));

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // first two are Lido transfer events, no transaction deposited
        assertEq(entries.length, 5);

        // expecting
        //  - 0.1 ETH insurance premium
        //  - 0.9 ETH committed yield from Lido
        //  - -0.3 ETH total yield in the report (0.9 negative yields paid off)
        assertEq(entries[2].topics[0], keccak256("InsurancePremiumPaid(bytes32,uint256)"));
        uint256 insuranceAmount = abi.decode(entries[2].data, (uint256));
        assertClose(insuranceAmount, uint256(0.1 ether));

        assertEq(entries[3].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[3].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[3].data, (int256));
        assertClose(yield, int256(0.9 ether));

        assertEq(entries[4].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[4].data, (int256, uint256, uint256));
        assertClose(totalYield, int256(-0.3 ether));
        assertClose(insurancePremiumPaid, uint256(0.1 ether));
        assertEq(insuranceWithdrawn, 0);

        uint256 insuranceBalance = ILido(LidoAddress).balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(0.1 ether));

        assertClose(ethYieldManager.accumulatedNegativeYields(), 0.3 ether);
    }

    function test_commitYieldReport_negative_yield_no_insurance_succeeds() external {
        vm.prank(address(l1BlastBridge));

        // -1 ETH yield
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 4 ether);
        emulatePositiveYield(3 ether);

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        assertEq(entries.length, 2);

        assertEq(entries[0].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[0].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[0].data, (int256));
        assertClose(yield, int256(-1 ether));

        assertEq(entries[1].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[1].data, (int256, uint256, uint256));
        assertClose(totalYield, int256(-1 ether));
        assertClose(insurancePremiumPaid, uint256(0));
        assertEq(insuranceWithdrawn, 0);


        uint256 insuranceBalance = ILido(LidoAddress).balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(0 ether));

        assertClose(ethYieldManager.accumulatedNegativeYields(), 1 ether);
    }

    function test_commitYieldReport_negative_yield_with_insurance_succeeds() external {
        vm.prank(address(l1BlastBridge));

        // -1 ETH yield
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 4 ether);
        emulatePositiveYield(3 ether);

        // send 2 ETH to insurance
        ILido(LidoAddress).transfer(ethYieldManager.insurance(), 2 ether);

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        assertEq(entries.length, 6);

        assertEq(entries[2].topics[0], keccak256("InsuranceWithdrawn(bytes32,uint256)"));
        uint256 insuranceAmount = abi.decode(entries[2].data, (uint256));
        assertClose(insuranceAmount, uint256(1 ether));

        assertEq(entries[3].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[3].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[3].data, (int256));
        uint256 buffer = ethYieldManager.insuranceWithdrawalBuffer();
        assertGe(yield, 0);
        assertLe(yield, int256(buffer));

        assertEq(entries[4].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[4].data, (int256, uint256, uint256));

        assertGe(totalYield, 0);
        assertLe(totalYield, int256(buffer));

        assertClose(insurancePremiumPaid, uint256(0));
        assertClose(insuranceWithdrawn, 1 ether + buffer);


        uint256 insuranceBalance = ILido(LidoAddress).balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(1 ether));

        assertEq(ethYieldManager.accumulatedNegativeYields(), 0 ether);
    }

    function test_commitYieldReport_negative_yield_insufficient_insurance_reverts() external {
        vm.prank(address(l1BlastBridge));

        // -1 ETH yield
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 4 ether);
        emulatePositiveYield(3 ether);

        // send 0.5 ETH to insurance
        ILido(LidoAddress).transfer(ethYieldManager.insurance(), 0.5 ether);

        vm.expectRevert(InsufficientInsuranceBalance.selector);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
    }

    // TODO: withdrawals
    // TODO: fuzz test: L1 share price only changes upon yield report
    // TODO: set admin
    // TODO: set insurance
    // TODO: accounting e2e

}

contract USD_YieldManager_Test is DSRYieldProvider_Initializer, Util {
    error InvalidExtraData();

    function test_getProviderInfo_DSR_empty_succeeds() external {
        IPot pot = IPot(DSR_MANAGER.pot());
        vm.mockCall(
            address(pot),
            abi.encodeWithSelector(
                pot.rho.selector
            ),
            abi.encode(block.timestamp)
        );

        YieldManager.ProviderInfo memory info = l1BlastBridge.getProviderInfoAt(0);
        assertEq(info.providerAddress, address(dsrProvider));
        assertEq(info.id, keccak256(abi.encodePacked("DSRYieldProvider", string(abi.encodePacked("1.0.0")))));
        assertEq(info.stakedBalance, uint256(0));
        assertClose(info.stakedValue, 0 ether);
        assertClose(uint256(info.yield), 0 ether);
    }

    /*
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

    function _testConvert(int128 inputToken, int128 outputToken, uint256 inputAmountWad, uint256 minOutputAmountWad) internal {
        uint256 inputTokenBalance = IERC20(USDConversions._token(inputToken)).balanceOf(address(l1BlastBridge));
        uint256 outputTokenBalance = IERC20(USDConversions._token(outputToken)).balanceOf(address(l1BlastBridge));
        vm.prank(address(multisig));
        uint256 amountReceived = l1BlastBridge.convert(inputToken, outputToken, inputAmountWad, minOutputAmountWad);
        assertGe(amountReceived, USDConversions._convertDecimals(minOutputAmountWad, outputToken));
        assertEq(IERC20(USDConversions._token(inputToken)).balanceOf(address(l1BlastBridge)), inputTokenBalance - USDConversions._convertDecimals(inputAmountWad, inputToken));
        assertEq(IERC20(USDConversions._token(outputToken)).balanceOf(address(l1BlastBridge)), outputTokenBalance + amountReceived);
    }

    function test_usdConversions_USDCToDAI() external {
        deal(address(USDC), address(l1BlastBridge), 2 ether);
        _testConvert(USDC_INDEX, DAI_INDEX, 2 ether, 2 ether);
    }

    function test_usdConversions_DAIToUSDC() external {
        deal(address(DAI), address(l1BlastBridge), 2 ether);
        _testConvert(DAI_INDEX, USDC_INDEX, 2 ether, 2 ether);
    }

    // TODO: for some reason Curve isn't working on a fork so these tests aren't working, the call to Curve looks correct though. Just a random EVM error that makes it hard to debug.
    function test_usdConversions_DAIToUSDT() external {
        deal(address(DAI), address(l1BlastBridge), 2 ether);
        _testConvert(DAI_INDEX, USDT_INDEX, 2 ether, 1.9 ether);
    }

    function test_usdConversions_USDTToDAI() external {
        deal(address(USDT), address(l1BlastBridge), 2 ether);
        _testConvert(USDT_INDEX, DAI_INDEX, 2 ether, 1.9 ether);
    }

    function test_usdConversions_USDCToUSDT() external {
        deal(address(USDC), address(l1BlastBridge), 2 ether);
        _testConvert(USDC_INDEX, USDT_INDEX, 2 ether, 1.9 ether);
    }

    function test_usdConversions_USDTToUSDC() external {
        deal(address(USDT), address(l1BlastBridge), 100 ether);
        _testConvert(USDT_INDEX, USDC_INDEX, 100 ether, 50 ether);
    }

    function test_usdConversions_reverts() external {
        deal(address(DAI), address(l1BlastBridge), 2 ether);
        deal(address(USDC), address(l1BlastBridge), 2 ether);
        deal(address(USDT), address(l1BlastBridge), 2 ether);

        vm.expectRevert(MinimumAmountNotMet.selector);
        vm.prank(address(multisig));
        l1BlastBridge.convert(0, 1, 2 ether, 2 ether + 1);

        vm.expectRevert();
        vm.prank(address(multisig));
        l1BlastBridge.convert(3, 0, 2 ether, 2 ether);

        vm.expectRevert();
        vm.prank(address(multisig));
        l1BlastBridge.convert(0, 3, 2 ether, 2 ether);

        vm.expectRevert();
        vm.prank(address(multisig));
        l1BlastBridge.convert(0, 0, 2 ether, 2 ether);
    }
    */

    function test_stake_DSR_succeeds() external {
        vm.startPrank(0x9759A6Ac90977b93B58547b4A71c78317f391A28);
        RealDAI.mint(address(l1BlastBridge), 10 ether);
        vm.stopPrank();

        vm.prank(multisig);
        l1BlastBridge.stake(0, address(dsrProvider), 10 ether);

        IPot pot = IPot(DSR_MANAGER.pot());
        vm.mockCall(
            address(pot),
            abi.encodeWithSelector(
                pot.rho.selector
            ),
            abi.encode(block.timestamp - 30 days)
        );

        YieldManager.ProviderInfo memory info = l1BlastBridge.getProviderInfoAt(0);
        assertEq(info.providerAddress, address(dsrProvider));
        assertEq(info.id, keccak256(abi.encodePacked("DSRYieldProvider", string(abi.encodePacked("1.0.0")))));
        assertEq(info.stakedBalance, uint256(10 ether));
        assertGt(info.stakedValue, 10 ether);
        uint256 yield = info.stakedValue - 10 ether;
        assertEq(uint256(info.yield), yield);
        uint256 totalValue = l1BlastBridge.totalValue();
        assertEq(totalValue, 10 ether + yield);
    }

    function test_convert_USDC_TO_DAI_succeeds() external {
        vm.prank(RealUSDC.masterMinter());
        RealUSDC.configureMinter(address(this), type(uint256).max);
        RealUSDC.mint(address(l1BlastBridge), 10e6); // 10 USDC

        vm.prank(address(l1BlastBridge));
        RealUSDC.approve(0x0A59649758aa4d66E25f08Dd01271e891fe52199, type(uint256).max);

        vm.prank(multisig);
        l1BlastBridge.convert(1, 0, 10 ether, 10 ether); // convert USDC to DAI
        assertEq(RealDAI.balanceOf(address(l1BlastBridge)), 10 ether);
        assertEq(RealUSDC.balanceOf(address(l1BlastBridge)), 0);
    }

    function test_DAI_deposit_conversion_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();

        // Deal Alice's USDC State
        deal(address(RealDAI), alice, 10 ether, true);
        vm.prank(alice);
        RealDAI.approve(address(l1BlastBridge), type(uint256).max);

        // no need to specify extraData for DAI
        bytes memory extraData = hex"";

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(RealDAI), abi.encodeWithSelector(DAI.transferFrom.selector, alice, address(l1BlastBridge), 10 ether)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(DAI), address(Usdb), alice, alice, 10 ether, extraData
        );

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2BlastBridge), message, 10000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l1BlastBridge), address(l2BlastBridge), 0, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 0, baseGas, false, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(RealDAI), address(Usdb), 10 ether, 10000, extraData);
        assertEq(RealDAI.balanceOf(address(l1BlastBridge)), 10 ether);
    }

    function test_USDC_deposit_conversion_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();

        // Deal Alice's USDC State
        deal(address(RealUSDC), alice, 10 * 1.0e6, true);
        vm.prank(alice);
        RealUSDC.approve(address(l1BlastBridge), type(uint256).max);

        // specify minAmountOut to be 10 DAI
        bytes memory extraData = abi.encodePacked(uint256(10 ether)); // this has to be in wad

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(RealUSDC), abi.encodeWithSelector(USDC.transferFrom.selector, alice, address(l1BlastBridge), 10 * 1.0e6)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(DAI), address(Usdb), alice, alice, 10 ether, extraData
        );

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2BlastBridge), message, 10000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l1BlastBridge), address(l2BlastBridge), 0, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 0, baseGas, false, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(RealUSDC), address(Usdb), 10 * 1.0e6, 10000, extraData);
        assertEq(RealUSDC.balanceOf(address(l1BlastBridge)), 0);
        assertEq(RealDAI.balanceOf(address(l1BlastBridge)), 10 ether);
    }

    function test_USDC_deposit_no_extra_data_reverts() external {
        // Deal Alice's USDC State
        deal(address(RealUSDC), alice, 10 * 1.0e6, true);
        vm.prank(alice);
        RealUSDC.approve(address(l1BlastBridge), type(uint256).max);

        vm.expectRevert(InvalidExtraData.selector);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(RealUSDC), address(Usdb), 10 * 1.0e6, 10000, hex"");
    }

    function test_USDT_deposit_conversion_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();

        // Deal Alice's USDT State
        deal(address(USDT), alice, 10 * 1.0e6, true);
        vm.prank(alice);
        USDT.approve(address(l1BlastBridge), type(uint256).max);

        // specify minAmountOut to be 10 DAI
        bytes memory extraData = abi.encodePacked(uint256(9.9 ether)); // this has to be in wad

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(USDT), abi.encodeWithSelector(USDT.transferFrom.selector, alice, address(l1BlastBridge), 10 * 1.0e6)
        );

        // TODO: check the relayed message content


        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(USDT), address(Usdb), 10 * 1.0e6, 10000, extraData);
        assertEq(USDT.balanceOf(address(l1BlastBridge)), 0);
        assertGe(RealDAI.balanceOf(address(l1BlastBridge)), 9.9 ether);
    }

    function test_USDT_deposit_no_extra_data_reverts() external {
        // Deal Alice's USDT State
        deal(address(USDT), alice, 10 * 1.0e6, true);
        vm.prank(alice);
        USDT.approve(address(l1BlastBridge), type(uint256).max);

        vm.expectRevert(InvalidExtraData.selector);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(USDT), address(Usdb), 10 * 1.0e6, 10000, hex"");
    }

    function test_unstake_DSR_succeeds() external {
        vm.startPrank(0x9759A6Ac90977b93B58547b4A71c78317f391A28);
        RealDAI.mint(address(l1BlastBridge), 10 ether);
        vm.stopPrank();

        vm.startPrank(multisig);
        l1BlastBridge.stake(0, address(dsrProvider), 10 ether);

        assertEq(RealDAI.balanceOf(address(l1BlastBridge)), 0);

        skip(3600);

        vm.recordLogs();
        l1BlastBridge.unstake(0, address(dsrProvider), 10 ether);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        uint256 logLength = entries.length;
        uint256 claimedLogIndex = logLength - 1;
        uint256 unstakedLogIndex = logLength - 2;

        assertEq(entries[claimedLogIndex].topics[0], keccak256("Claimed(bytes32,uint256)"));
        assertEq(entries[claimedLogIndex].topics[1], YieldProvider(dsrProvider).id());
        uint256 claimed = abi.decode(entries[claimedLogIndex].data, (uint256));
        assertEq(claimed, uint256(10 ether));

        assertEq(entries[unstakedLogIndex].topics[0], keccak256("Unstaked(bytes32,uint256)"));
        assertEq(entries[unstakedLogIndex].topics[1], YieldProvider(dsrProvider).id());
        uint256 unstaked = abi.decode(entries[unstakedLogIndex].data, (uint256));
        assertEq(unstaked, uint256(10 ether));

        assertEq(RealDAI.balanceOf(address(l1BlastBridge)), 10 ether);
    }

    /*
    function test_recordStakedDeposit_DSR_succeeds() external {
        vm.prank(address(l1BlastBridge));
        l1BlastBridge.recordStakedDeposit(address(dsrProvider), 10 ether);
        YieldManager.ProviderInfo memory info = l1BlastBridge.getProviderInfoAt(0);
        assertEq(info.stakedBalance, uint256(10 ether));
    }
    */

    function test_commitYieldReport_DSR_positive_yield_no_insurance_succeeds() external {
        uint256 yield = emulatePositiveYield(10 ether, 30 days);

        vm.recordLogs();
        vm.prank(multisig);
        l1BlastBridge.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        assertEq(entries.length, 3);

        assertEq(entries[0].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[0].topics[1], YieldProvider(dsrProvider).id());
        int256 reportedYield = abi.decode(entries[0].data, (int256));
        assertClose(reportedYield, int256(yield));

        assertEq(entries[1].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, , ) = abi.decode(entries[1].data, (int256, uint256, uint256));
        assertEq(totalYield, int256(yield));

        assertEq(entries[2].topics[0], keccak256("TransactionDeposited(address,address,uint256,bytes)"));
        assertEq(entries[2].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(l1BlastBridge))));
        assertEq(entries[2].topics[2], addressToBytes32(Predeploys.USDB));
        // TODO: verify the tx data
    }

    function test_commitYieldReport_DSR_positive_yield_insurance_premium_succeeds() external {
        uint256 yield = emulatePositiveYield(10 ether, 30 days);

        vm.recordLogs();
        vm.prank(multisig);
        l1BlastBridge.commitYieldReport(true);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        uint256 logLength = entries.length;
        uint256 depositLogIndex = logLength - 1;
        uint256 yieldReportLogIndex = logLength - 2;
        uint256 yieldCommitLogIndex = logLength - 3;
        uint256 insurancePaidLogIndex = logLength - 4;

        assertEq(entries[insurancePaidLogIndex].topics[0], keccak256("InsurancePremiumPaid(bytes32,uint256)"));
        uint256 insuranceAmount = abi.decode(entries[insurancePaidLogIndex].data, (uint256));
        assertClose(insuranceAmount, uint256(yield / 10)); // 10% of yield

        assertEq(entries[yieldCommitLogIndex].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[yieldCommitLogIndex].topics[1], YieldProvider(dsrProvider).id());
        int256 committedYield = abi.decode(entries[yieldCommitLogIndex].data, (int256));
        assertClose(committedYield, int256(yield - insuranceAmount));

        assertEq(entries[yieldReportLogIndex].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[yieldReportLogIndex].data, (int256, uint256, uint256));
        assertClose(totalYield, int256(yield - insuranceAmount));
        assertClose(insurancePremiumPaid, insuranceAmount);
        assertEq(insuranceWithdrawn, 0);

        assertEq(entries[depositLogIndex].topics[0], keccak256("TransactionDeposited(address,address,uint256,bytes)"));
        assertEq(entries[depositLogIndex].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(l1BlastBridge))));
        assertEq(entries[depositLogIndex].topics[2], addressToBytes32(Predeploys.USDB)); // USDB

        uint256 insuranceBalance = RealDAI.balanceOf(l1BlastBridge.insurance());
        assertClose(insuranceBalance, uint256(yield / 10)); // 10% of yield
    }

    function emulatePositiveYield(uint256 initialAmount, uint256 elapsed) internal returns (uint256 yield) {
        vm.startPrank(0x9759A6Ac90977b93B58547b4A71c78317f391A28);
        RealDAI.mint(address(l1BlastBridge), initialAmount);
        vm.stopPrank();

        vm.startPrank(multisig);
        l1BlastBridge.stake(0, address(dsrProvider), initialAmount);

        skip(elapsed);

        YieldManager.ProviderInfo memory info = l1BlastBridge.getProviderInfoAt(0);
        yield = info.stakedValue - initialAmount;
    }
}
