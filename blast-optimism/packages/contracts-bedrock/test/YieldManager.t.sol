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
import { Types } from "src/libraries/Types.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";
import { LidoYieldProvider, ILido, IERC20, IWithdrawalQueue } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";
import { DSRYieldProvider, IDsrManager, IPot } from "src/mainnet-bridge/yield-providers/DSRYieldProvider.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { Insurance } from "src/mainnet-bridge/Insurance.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";

// Target contract
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { ETHYieldManager } from "src/mainnet-bridge/ETHYieldManager.sol";

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
        Lido.transfer(address(ethYieldManager), amount);
    }

    function emulateNegativeYield(uint256 amount) internal {
        vm.prank(address(ethYieldManager));
        Lido.transfer(address(this), amount);
    }

    function test_getProviderInfo_succeeds() external {
        emulatePositiveYield(1 ether);

        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.providerAddress, address(lidoProvider));
        assertEq(info.id, keccak256(abi.encodePacked("LidoYieldProvider", string(abi.encodePacked("1.0.0")))));
        assertEq(info.stakedPrincipal, uint256(0));
        assertClose(info.stakedBalance, 1 ether);
        assertClose(uint256(info.yield), 1 ether);
    }

    function test_stake_Lido_succeeds() external {
        vm.deal(address(ethYieldManager), 1 ether);

        uint256 availableBalance = ethYieldManager.availableBalance();
        assertEq(availableBalance, 1 ether);

        vm.prank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 0.4 ether);
        assertEq(address(ethYieldManager).balance, 0.6 ether);

        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.stakedPrincipal, uint256(0.4 ether));
        assertClose(info.stakedBalance, uint256(0.4 ether));
        assertClose(info.yield, int256(0 ether));
        assertEq(ethYieldManager.availableBalance(), uint256(0.6 ether));
        uint256 totalEth = ethYieldManager.totalValue();
        assertClose(totalEth, 1 ether);
    }

    function test_unstake_Lido_succeeds() external { // ensure withdrawal queue is not paused skip(1700000000);
        vm.deal(address(ethYieldManager), 1 ether);

        uint256 availableBalance = ethYieldManager.availableBalance();
        assertEq(availableBalance, 1 ether);

        // avoid paused error
        vm.warp(1706000000);

        vm.prank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 0.4 ether);

        vm.expectEmit(false, false, false, true, address(ethYieldManager));
        emit LidoUnstakeInitiated(1234, 0.2 ether);

        vm.prank(multisig);
        ethYieldManager.unstake(0, address(lidoProvider), 0.2 ether);

        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.pendingBalance, uint256(0.2 ether));
        assertEq(info.stakedPrincipal, uint256(0.2 ether));
    }

    function test_recordStakedDeposit_Lido_succeeds() external {
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 0.4 ether);
        YieldManager.ProviderInfo memory info = ethYieldManager.getProviderInfoAt(0);
        assertEq(info.stakedPrincipal, uint256(0.4 ether));
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

        uint256 insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
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

        uint256 insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
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

        uint256 insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
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


        uint256 insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(0 ether));

        assertClose(ethYieldManager.accumulatedNegativeYields(), 1 ether);
    }

    function test_commitYieldReport_negative_yield_with_insurance_succeeds() external {
        vm.prank(address(l1BlastBridge));

        // -1 ETH yield
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 4 ether);
        emulatePositiveYield(3 ether);

        // send 2 ETH to insurance
        Lido.transfer(ethYieldManager.insurance(), 2 ether);

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


        uint256 insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(1 ether));

        assertEq(ethYieldManager.accumulatedNegativeYields(), 0 ether);
    }

    function test_commitYieldReport_negative_yield_insufficient_insurance_reverts() external {
        vm.prank(address(l1BlastBridge));

        // -1 ETH yield
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 4 ether);
        emulatePositiveYield(3 ether);

        // send 0.5 ETH to insurance
        Lido.transfer(ethYieldManager.insurance(), 0.5 ether);

        vm.expectRevert(InsufficientInsuranceBalance.selector);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
    }

    /// @notice This test ensures that the yield manager claims the correct amount
    ///         of claimable withdrawals from the Lido withdrawal queue before
    ///         committing the yield report. It also ensures that any realized
    ///         negative yield from the claim is reflected in the yield report.
    function test_commitYieldReport_claim_succeeds() external {
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 1 ether);

        vm.startPrank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 0.4 ether);
        ethYieldManager.unstake(0, address(lidoProvider), 0.3 ether);
        vm.stopPrank();

        finalizeWithdrawals();

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // verify the last three events (Claimed, YieldCommit, YieldReport)
        // there should be no TransactionDeposited event
        uint256 claimedIndex = entries.length - 3;
        uint256 yieldCommitIndex = entries.length - 2;
        uint256 yieldReportIndex = entries.length - 1;

        assertEq(entries[claimedIndex].topics[0], keccak256("Claimed(bytes32,uint256,uint256)"));
        assertEq(entries[claimedIndex].topics[1], YieldProvider(lidoProvider).id());
        (uint256 claimed, uint256 expected) = abi.decode(entries[claimedIndex].data, (uint256,uint256));

        assertEq(expected, 0.3 ether);
        assertClose(claimed, 0.3 ether);
        uint256 negativeYield = expected - claimed;

        assertEq(entries[yieldCommitIndex].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[yieldCommitIndex].topics[1], YieldProvider(lidoProvider).id());

        // yield could be slightly negative
        int256 yield = abi.decode(entries[yieldCommitIndex].data, (int256));
        assertClose(yield + 1 ether, int256(0 ether + 1 ether));

        int256 expectedTotalYield = yield + int256(negativeYield) * -1;

        assertEq(entries[yieldReportIndex].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[yieldReportIndex].data, (int256, uint256, uint256));
        assertEq(totalYield, expectedTotalYield);
        assertEq(insurancePremiumPaid, 0);
        assertEq(insuranceWithdrawn, 0);

        uint256 expectedAccumulatedNegativeYield = totalYield >= 0 ?
            uint256(0) :
            uint256(totalYield * -1);

        assertEq(ethYieldManager.accumulatedNegativeYields(), expectedAccumulatedNegativeYield);
    }

    /// @notice This test ensures that any negative yield that is realized from
    ///         the claim is paid off by the positive yield.
    function test_commitYieldReport_claim_positive_yield_succeeds() external {
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 1 ether);

        vm.startPrank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 1.0 ether);
        ethYieldManager.unstake(0, address(lidoProvider), 0.4 ether);
        vm.stopPrank();

        emulatePositiveYield(0.1 ether);

        finalizeWithdrawals();

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // verify the last four events (Claimed, YieldCommit, YieldReport, TransactionDeposited)
        uint256 claimedIndex = entries.length - 4;
        uint256 yieldCommitIndex = entries.length - 3;
        uint256 yieldReportIndex = entries.length - 2;
        uint256 transactionDepositedIndex = entries.length - 1;

        assertEq(entries[claimedIndex].topics[0], keccak256("Claimed(bytes32,uint256,uint256)"));
        assertEq(entries[claimedIndex].topics[1], YieldProvider(lidoProvider).id());
        (uint256 claimed, uint256 expected) = abi.decode(entries[claimedIndex].data, (uint256,uint256));

        assertEq(expected, 0.4 ether);
        assertClose(claimed, 0.4 ether);
        uint256 negativeYield = expected - claimed;

        assertEq(entries[yieldCommitIndex].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[yieldCommitIndex].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[yieldCommitIndex].data, (int256));
        assertClose(yield, int256(0.1 ether));

        int256 totalExpectedYield = yield + int256(negativeYield) * -1;

        assertEq(entries[yieldReportIndex].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[yieldReportIndex].data, (int256, uint256, uint256));
        assertEq(totalYield, totalExpectedYield);
        assertEq(insurancePremiumPaid, 0);
        assertEq(insuranceWithdrawn, 0);

        assertEq(entries[transactionDepositedIndex].topics[0], keccak256("TransactionDeposited(address,address,uint256,bytes)"));
        assertEq(entries[transactionDepositedIndex].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager))));
        assertEq(entries[transactionDepositedIndex].topics[2], addressToBytes32(0x4300000000000000000000000000000000000000));

        assertEq(ethYieldManager.accumulatedNegativeYields(), 0);
    }

    /// @notice This test ensures that when there are multiple claimable withdrawals
    ///         the yield manager claims the correct amount of claimable withdrawals
    function test_commitYieldReport_claim_multiple_succeeds() external {
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 1 ether);

        vm.startPrank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 1.0 ether);

        // total of 0.6 ETH unstaked in three withdrawals
        ethYieldManager.unstake(0, address(lidoProvider), 0.1 ether);
        ethYieldManager.unstake(0, address(lidoProvider), 0.2 ether);
        ethYieldManager.unstake(0, address(lidoProvider), 0.3 ether);
        vm.stopPrank();

        finalizeWithdrawals();

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // verify the last three events (Claimed, YieldCommit, YieldReport)
        // there should be no TransactionDeposited event
        uint256 claimedIndex = entries.length - 3;
        uint256 yieldCommitIndex = entries.length - 2;
        uint256 yieldReportIndex = entries.length - 1;

        assertEq(entries[claimedIndex].topics[0], keccak256("Claimed(bytes32,uint256,uint256)"));
        assertEq(entries[claimedIndex].topics[1], YieldProvider(lidoProvider).id());
        (uint256 claimed, uint256 expected) = abi.decode(entries[claimedIndex].data, (uint256,uint256));

        assertEq(expected, 0.6 ether);
        assertClose(claimed, 0.6 ether);
        uint256 negativeYield = expected - claimed;

        assertEq(entries[yieldCommitIndex].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[yieldCommitIndex].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[yieldCommitIndex].data, (int256));
        // due to transfers during unstaking, the yield is not exactly 0
        // adding 1 ether to both sides to avoid comparison with 0
        assertClose(yield + int256(1 ether), int256(1 ether));

        int256 expectedTotalYield = yield + int256(negativeYield) * -1;

        assertEq(entries[yieldReportIndex].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[yieldReportIndex].data, (int256, uint256, uint256));
        assertEq(totalYield, expectedTotalYield);
        assertEq(insurancePremiumPaid, 0);
        assertEq(insuranceWithdrawn, 0);

        uint256 expectedAccumulatedNegativeYields = totalYield < 0 ? uint256(totalYield * -1) : 0;

        assertEq(ethYieldManager.accumulatedNegativeYields(), expectedAccumulatedNegativeYields);
    }

    /// @notice This test ensures that when the number of claimable withdrawals
    ///         exceeds the batch size, the yield manager claims the correct amount
    function test_commitYieldReport_claim_batch_size_succeeds() external {
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 2 ether);

        vm.startPrank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 2.0 ether);

        // unstake 11 times, but because the batch size is 10, only 10 should be claimed
        for (uint256 i = 0; i < 11; i++) {
            ethYieldManager.unstake(0, address(lidoProvider), 0.1 ether);
        }
        vm.stopPrank();

        finalizeWithdrawals();

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // verify the last three events (Claimed, YieldCommit, YieldReport)
        // there should be no TransactionDeposited event
        uint256 claimedIndex = entries.length - 3;
        uint256 yieldCommitIndex = entries.length - 2;
        uint256 yieldReportIndex = entries.length - 1;

        assertEq(entries[claimedIndex].topics[0], keccak256("Claimed(bytes32,uint256,uint256)"));
        assertEq(entries[claimedIndex].topics[1], YieldProvider(lidoProvider).id());
        (uint256 claimed, uint256 expected) = abi.decode(entries[claimedIndex].data, (uint256,uint256));

        assertEq(expected, 1.0 ether);
        assertClose(claimed, 1.0 ether);
        uint256 negativeYield = expected - claimed;

        assertEq(entries[yieldCommitIndex].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[yieldCommitIndex].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[yieldCommitIndex].data, (int256));

        // due to multiple transfers from unstaking, the yield might not be exactly 0
        assertClose(yield + int256(1 ether), int256(1 ether));

        int256 totalExpectedYield = yield + int256(negativeYield) * -1;

        assertEq(entries[yieldReportIndex].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[yieldReportIndex].data, (int256, uint256, uint256));
        assertEq(totalYield, totalExpectedYield);
        assertEq(insurancePremiumPaid, 0);
        assertEq(insuranceWithdrawn, 0);

        uint256 expectedAccumulatedNegativeYield = totalExpectedYield >= 0 ?
            uint256(totalExpectedYield) :
            uint256(totalExpectedYield * -1);

        assertEq(ethYieldManager.accumulatedNegativeYields(), expectedAccumulatedNegativeYield);

        assertEq(lidoProvider.lastClaimedIndex(), 10);
        assertEq(lidoProvider.lastUnstakeRequestIndex(), 11);
    }

    /// @notice This test ensures that when there are no claimable withdrawals
    ///         the yield manager does not claim any withdrawals
    function test_commitYieldReport_not_claimable_succeeds() external {
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 2 ether);

        vm.startPrank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 2.0 ether);

        // unstake 11 times, but because the batch size is 10, only 10 should be claimed
        for (uint256 i = 0; i < 11; i++) {
            ethYieldManager.unstake(0, address(lidoProvider), 0.1 ether);
        }
        vm.stopPrank();

        emulatePositiveYield(0.1 ether);

        // do not finalize withdrawals
        //finalizeWithdrawals();

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // verify the last three events (YieldCommit, YieldReport, TransactionDeposited)
        uint256 yieldCommitIndex = entries.length - 3;
        uint256 yieldReportIndex = entries.length - 2;
        uint256 transactionDepositedIndex = entries.length - 1;

        assertEq(entries[yieldCommitIndex].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[yieldCommitIndex].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[yieldCommitIndex].data, (int256));

        assertClose(yield, int256(0.1 ether));

        assertEq(entries[yieldReportIndex].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[yieldReportIndex].data, (int256, uint256, uint256));
        assertEq(totalYield, yield);
        assertEq(insurancePremiumPaid, 0);
        assertEq(insuranceWithdrawn, 0);

        assertEq(entries[transactionDepositedIndex].topics[0], keccak256("TransactionDeposited(address,address,uint256,bytes)"));
        assertEq(entries[transactionDepositedIndex].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager))));
        assertEq(entries[transactionDepositedIndex].topics[2], addressToBytes32(0x4300000000000000000000000000000000000000));

        assertEq(ethYieldManager.accumulatedNegativeYields(), 0);

    }

    /// @notice This test ensures that negative yield from claim and negative yield
    ///         from staked balance are combined
    function test_commitYieldReport_claim_negative_yield_combined_succeeds() external {
        // ensure withdrawal queue is not paused
        skip(1700000000);
        vm.deal(address(ethYieldManager), 1 ether);

        vm.startPrank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), 1.0 ether);
        ethYieldManager.unstake(0, address(lidoProvider), 0.2 ether);
        vm.stopPrank();

        emulateNegativeYield(0.1 ether);

        finalizeWithdrawals();

        vm.recordLogs();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        // verify the last three events (Claimed, YieldCommit, YieldReport)
        uint256 claimedIndex = entries.length - 3;
        uint256 yieldCommitIndex = entries.length - 2;
        uint256 yieldReportIndex = entries.length - 1;

        assertEq(entries[claimedIndex].topics[0], keccak256("Claimed(bytes32,uint256,uint256)"));
        assertEq(entries[claimedIndex].topics[1], YieldProvider(lidoProvider).id());
        (uint256 claimed, uint256 expected) = abi.decode(entries[claimedIndex].data, (uint256,uint256));

        assertEq(expected, 0.2 ether);
        assertClose(claimed, 0.2 ether);
        uint256 negativeYieldFromClaim = expected - claimed;

        assertEq(entries[yieldCommitIndex].topics[0], keccak256("YieldCommit(bytes32,int256)"));
        assertEq(entries[yieldCommitIndex].topics[1], YieldProvider(lidoProvider).id());
        int256 yield = abi.decode(entries[yieldCommitIndex].data, (int256));

        assertClose(yield, int256(-0.1 ether));

        int256 totalExpectedYield = yield + int256(negativeYieldFromClaim) * -1;

        assertEq(entries[yieldReportIndex].topics[0], keccak256("YieldReport(int256,uint256,uint256)"));
        (int256 totalYield, uint256 insurancePremiumPaid, uint256 insuranceWithdrawn) = abi.decode(entries[yieldReportIndex].data, (int256, uint256, uint256));
        assertEq(totalYield, totalExpectedYield);
        assertEq(insurancePremiumPaid, 0);
        assertEq(insuranceWithdrawn, 0);

        assertEq(ethYieldManager.accumulatedNegativeYields(), uint256(totalExpectedYield * -1));

    }

    function test_requestWithdrawal_portal_succeeds() external {
        vm.prank(address(op));
        uint256 requestId = ethYieldManager.requestWithdrawal(1 ether);
        assertEq(requestId, 1);
    }

    function test_requestWithdrawal_not_portal_reverts() external {
        vm.prank(address(0x4));
        vm.expectRevert(ETHYieldManager.CallerIsNotPortal.selector);
        ethYieldManager.requestWithdrawal(1 ether);
    }

    function test_setInsurance_succeeds() external {
        Insurance newInsurance = new Insurance(ethYieldManager);

        vm.prank(multisig);
        ethYieldManager.setInsurance(address(newInsurance), 5000, 1000);

        assertEq(ethYieldManager.insurance(), address(newInsurance));
        assertEq(ethYieldManager.insuranceFeeBips(), 5000);
        assertEq(ethYieldManager.insuranceWithdrawalBuffer(), 1000);
    }

    function test_setInsurance_non_owner_reverts() external {
        Insurance newInsurance = new Insurance(ethYieldManager);

        vm.expectRevert();
        vm.prank(address(0x4));
        ethYieldManager.setInsurance(address(newInsurance), 5000, 1000);
    }
}

contract ETH_Insurance_Test is LidoYieldProvider_Initializer, Util {
    function test_setAdmin_succeeds() external {
        assertEq(insurance.admin(), multisig);
        vm.prank(multisig);
        insurance.setAdmin(address(0x4));
        assertEq(insurance.admin(), address(0x4));
    }

    function test_coverLoss_admin_succeeds() external {

        Lido.transfer(address(insurance), 2 ether);

        assertEq(ethYieldManager.totalValue(), 0 ether);

        vm.prank(multisig);
        insurance.coverLoss(address(Lido), 1 ether);

        assertClose(ethYieldManager.totalValue(), 1 ether);
    }

    function test_coverLoss_ym_succeeds() external {

        Lido.transfer(address(insurance), 2 ether);

        assertEq(ethYieldManager.totalValue(), 0 ether);

        vm.prank(address(ethYieldManager));
        insurance.coverLoss(address(Lido), 1 ether);

        assertClose(ethYieldManager.totalValue(), 1 ether);
    }

    function test_coverLoss_non_admin_reverts() external {

        Lido.transfer(address(insurance), 2 ether);

        assertEq(ethYieldManager.totalValue(), 0 ether);

        vm.expectRevert(abi.encodeWithSelector(Insurance.OnlyAdminOrYieldManager.selector));

        vm.prank(address(0x4));
        insurance.coverLoss(address(Lido), 1 ether);
    }

    function test_different_insurance_rates_succeeds() external {
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 0.1 ether);
        Lido.transfer(address(ethYieldManager), 0.2 ether);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);

        uint256 insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(0.01 ether));

        // change rate to 0%
        vm.startPrank(multisig);
        ethYieldManager.setInsurance(address(insurance), 0, ethYieldManager.insuranceWithdrawalBuffer());
        vm.stopPrank();

        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 0.1 ether);
        Lido.transfer(address(ethYieldManager), 0.2 ether);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);

        insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
        assertClose(insuranceBalance, uint256(0.01 ether));

        // change rate to 50%
        vm.startPrank(multisig);
        ethYieldManager.setInsurance(address(insurance), 5000, ethYieldManager.insuranceWithdrawalBuffer());
        vm.stopPrank();

        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), 0.1 ether);
        Lido.transfer(address(ethYieldManager), 0.2 ether);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);

        insuranceBalance = Lido.balanceOf(ethYieldManager.insurance());
        // 0.01 + 0.05
        assertClose(insuranceBalance, uint256(0.06 ether));
    }
}

/// @notice Tests specifically related to accounting (incl. negative yields and withdrawal discounts)
///         This test does not use the insurance contract (i.e. commitYieldReport(false))
///
/// Formulas:
///    - `totalValue()`: tokenBalance() - getLockedBalance() + totalProviderValue()
///    - `sharePrice()`: totalValue() / (totalValue() + accumulatedNegativeYields)
///
/// Exhaustive list of `accumulatedNegativeYields` mutations:
///    - `commitYieldReport()`
///        - via LidoYieldProvider._claim() (via `_delegatecall_preCommitYieldReportDelegateCallHook` -> yp._claim() -> `ym.recordNegativeYield()`)
///        - via `totalYield` update
///    - `finalize()`: negative yields are paid off from discounted withdrawals
///
/// Exhaustive list of `totalValue()` mutations:
///    - YM.tokenBalance() = address(ETHYieldManager).balance
///    - YM.getLockedBalance() (via WithdrawalQueue.lockedBalance)
///        - increase via YM.finalize() (-> WithdrawalQueue._finalize())
///        - decrease via YM.claimWithdrawal()
///    - YM.totalProviderValue() = YP.stakedBalance() + YP.pendingBalance
///        - YP.stakedBalance() = LIDO.balanceOf(ETHYieldManager)
///        - YP.pendingBalance
///            - increase via YM.unstake()
///            - decrease via YM.commitYieldReport() (-> YP.preCommitYieldReportDelegateCallHook() -> YP._claim())
contract ETH_YieldManager_Accounting_No_Insurance_Test is LidoYieldProvider_Initializer, Util {
    // finalize decreases accumulatedNegativeYields by correct amount
    // i.e. share price never changes after finalize
    function testFuzz_sharePrice_does_not_change_after_single_finalize(
        uint256 balance,
        uint256 withdrawalAmount
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(withdrawalAmount > 0);
        vm.assume(balance > withdrawalAmount);

        // set up balance
        vm.deal(address(ethYieldManager), balance);

        // set up withdrawal
        vm.prank(address(op));
        uint256 requestId = ethYieldManager.requestWithdrawal(withdrawalAmount);

        assertEq(ethYieldManager.accumulatedNegativeYields(), 0);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();

        vm.prank(multisig);
        ethYieldManager.finalize(requestId);

        assertEq(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    // TODO: fuzz: multiple finalizes

    function testFuzz_sharePrice_does_not_decrease_with_negative_yields_after_single_finalize(
        uint256 balance,
        uint256 withdrawalAmount,
        uint256 originalNegativeYield
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(withdrawalAmount > 0);
        vm.assume(balance > withdrawalAmount);
        vm.assume(balance > originalNegativeYield);

        // set up balance
        vm.deal(address(ethYieldManager), balance);

        // set up withdrawal
        vm.prank(address(op));
        uint256 requestId = ethYieldManager.requestWithdrawal(withdrawalAmount);

        // set up negative yield
        vm.prank(address(ethYieldManager));
        ethYieldManager.recordNegativeYield(originalNegativeYield);
        assertEq(ethYieldManager.accumulatedNegativeYields(), originalNegativeYield);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();

        vm.prank(multisig);
        ethYieldManager.finalize(requestId);

        // TODO: figure out the precision range
        assertGe(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    // see OptimismPortal_FinalizeWithdrawal_Test for testFuzz_sharePrice_does_not_change_after_finalizeWithdrawalTransaction

    // share price should stay the same after ETH balance increase
    function testFuzz_sharePrice_does_not_change_after_ETH_balance_increase(
        uint256 balance,
        uint256 increaseAmount
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > increaseAmount);
        vm.assume(increaseAmount > 0);

        // set up balance
        vm.deal(address(ethYieldManager), balance);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();

        vm.deal(address(ethYieldManager), increaseAmount + balance);

        assertEq(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    function testFuzz_sharePrice_does_not_decrease_with_negative_yields_after_ETH_balance_increase(
        uint256 balance,
        uint256 increaseAmount,
        uint256 originalNegativeYield
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > increaseAmount);
        vm.assume(balance > originalNegativeYield);
        vm.assume(originalNegativeYield > 0);
        vm.assume(increaseAmount > 0);

        // set up balance
        vm.deal(address(ethYieldManager), balance);

        // set up negative yield
        vm.prank(address(ethYieldManager));
        ethYieldManager.recordNegativeYield(originalNegativeYield);
        assertEq(ethYieldManager.accumulatedNegativeYields(), originalNegativeYield);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();

        vm.deal(address(ethYieldManager), increaseAmount + balance);

        assertGe(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    function testFuzz_sharePrice_does_not_decrease_after_stETH_balance_increase(
        uint256 balance,
        uint256 increaseAmount,
        uint256 originalNegativeYield
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > increaseAmount);
        vm.assume(balance > originalNegativeYield);
        //vm.assume(originalNegativeYield > 0);
        vm.assume(increaseAmount > 0);

        // set up balance
        vm.deal(address(this), balance);
        Lido.submit{value: balance}(address(0));
        Lido.transfer(address(ethYieldManager), balance);

        // set up negative yield
        vm.prank(address(ethYieldManager));
        ethYieldManager.recordNegativeYield(originalNegativeYield);
        assertEq(ethYieldManager.accumulatedNegativeYields(), originalNegativeYield);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();

        vm.deal(address(ethYieldManager), increaseAmount + balance);

        // can only increase (by a very small amount) but never decrease
        assertGe(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    function testFuzz_sharePrice_does_not_change_after_pendingBalance_increase_via_unstake(
        uint256 balance,
        uint256 unstakeAmount
    ) external {
        // share price = 1
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(1000 ether > unstakeAmount);
        vm.assume(unstakeAmount > 100);

        // avoid paused error
        vm.warp(1706000000);

        // set up balance
        vm.deal(address(this), balance + unstakeAmount);
        Lido.submit{value: balance + unstakeAmount}(address(0));
        Lido.transfer(address(ethYieldManager), balance + unstakeAmount);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + unstakeAmount);

        assertEq(ethYieldManager.sharePrice(), 1e27);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        vm.prank(multisig);
        ethYieldManager.unstake(0, address(lidoProvider), unstakeAmount);
        assertGe(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    function testFuzz_sharePrice_decreases_after_negative_yield_from_provider(
        uint256 balance,
        uint256 negativeYield,
        uint256 originalNegativeYield
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > negativeYield);
        vm.assume(100000 ether > originalNegativeYield);
        vm.assume(originalNegativeYield > 0);
        vm.assume(balance > originalNegativeYield + negativeYield);
        // using a sufficiently large negative yield to avoid precision errors
        vm.assume(negativeYield > 1e9);

        // set up balance and original negative yield
        vm.deal(address(this), balance);
        Lido.submit{value: balance}(address(0));
        Lido.transfer(address(ethYieldManager), balance);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + originalNegativeYield);

        // first commit yield report
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        assertLt(sharePriceBefore, 1e27);

        // more negative yield
        vm.prank(address(ethYieldManager));
        Lido.transfer(address(this), negativeYield);

        // commit yield report again
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        assertLt(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    function testFuzz_sharePrice_decreases_after_negative_yield_from_claim(
        uint256 balance,
        uint256 unstakeAmount
    ) external {
        // share price = 1
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(1000 ether > unstakeAmount);
        vm.assume(unstakeAmount > 100);

        // make sure withdrawal queue is not paused
        vm.warp(1706000000);

        // set up balance and stake
        vm.deal(address(this), balance + unstakeAmount);
        Lido.submit{value: balance + unstakeAmount}(address(0));
        Lido.transfer(address(ethYieldManager), balance + unstakeAmount);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + unstakeAmount);

        // unstake and finalize
        vm.prank(multisig);
        ethYieldManager.unstake(0, address(lidoProvider), unstakeAmount);

        finalizeWithdrawals();

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        assertEq(sharePriceBefore, 1e27);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);
        assertLt(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    // guarantees only when share price is 1
    function testFuzz_sharePrice_does_not_change_after_stake(
        uint256 balance,
        uint256 stakeAmount
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > stakeAmount);
        vm.assume(stakeAmount > 0);

        // set up balance
        vm.deal(address(ethYieldManager), balance + stakeAmount);

        // set up negative yield

        uint256 sharePriceBefore = ethYieldManager.sharePrice();

        vm.prank(multisig);
        ethYieldManager.stake(0, address(lidoProvider), stakeAmount);

        assertEq(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    // when share price is less than 1, stake() can inifinitesimally influence the share price
    //function testFuzz_sharePrice_does_not_change_with_negative_yields_after_stake() external {
    //    // TODO
    //}

    function testFuzz_sharePrice_does_not_change_after_positive_yield(
        uint256 balance,
        uint256 yield
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > yield);
        vm.assume(yield > 100);

        // set up balance and stake
        vm.deal(address(this), balance + yield);
        Lido.submit{value: balance + yield}(address(0));
        Lido.transfer(address(ethYieldManager), balance + yield);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        assertEq(sharePriceBefore, 1e27);

        // commit yield report
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        assertEq(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    function testFuzz_sharePrice_increases_after_positive_yield(
        uint256 balance,
        uint256 positiveYield,
        uint256 originalNegativeYield
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > positiveYield);
        vm.assume(100000 ether > originalNegativeYield);
        vm.assume(originalNegativeYield > 0);
        vm.assume(balance > originalNegativeYield);
        // using a sufficiently large negative yield to avoid precision errors
        vm.assume(positiveYield > 1e9);

        // set up balance and original negative yield
        vm.deal(address(this), balance + positiveYield + 1 ether);
        Lido.submit{value: balance + positiveYield + 1 ether}(address(0));
        Lido.transfer(address(ethYieldManager), balance);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + originalNegativeYield);

        // first commit yield report
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        assertLt(sharePriceBefore, 1e27);

        // provide positive yield
        Lido.transfer(address(ethYieldManager), positiveYield);

        // commit yield report again
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        assertGt(ethYieldManager.sharePrice(), sharePriceBefore);
    }

    function testFuzz_sharePrice_recovers_to_1_after_positive_yield(
        uint256 balance,
        uint256 positiveYield,
        uint256 originalNegativeYield
    ) external {
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(100000 ether > positiveYield);
        vm.assume(100000 ether > originalNegativeYield);
        vm.assume(originalNegativeYield > 0);
        vm.assume(balance > originalNegativeYield);
        // using a sufficiently large negative yield to avoid precision errors
        vm.assume(positiveYield > 1e9);
        // 100 as a buffer to avoid precision issues
        vm.assume(positiveYield > originalNegativeYield + 100);

        // set up balance and original negative yield
        vm.deal(address(this), balance + positiveYield + 1 ether);
        Lido.submit{value: balance + positiveYield + 1 ether}(address(0));
        Lido.transfer(address(ethYieldManager), balance);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + originalNegativeYield);

        // first commit yield report
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        assertLt(sharePriceBefore, 1e27);

        // provide positive yield
        Lido.transfer(address(ethYieldManager), positiveYield);

        // commit yield report again
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        assertEq(ethYieldManager.sharePrice(), 1e27);
    }

    function testFuzz_negative_yields_accumulate_after_commitYieldReport(
        uint256 balance,
        uint256 negativeYield
    ) external {
        vm.assume(balance > 1 ether);
        vm.assume(balance < 100000 ether);
        vm.assume(negativeYield < 100000 ether);
        vm.assume(negativeYield > 0);

        // emulate negative yield
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + negativeYield);
        vm.deal(address(this), balance);
        Lido.submit{value: balance}(address(0));
        Lido.transfer(address(ethYieldManager), balance);


        uint256 accumulatedBefore = ethYieldManager.accumulatedNegativeYields();
        assertEq(accumulatedBefore, 0);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(false);

        assertGt(ethYieldManager.accumulatedNegativeYields(), accumulatedBefore);
        //assertClose(accumulatedBefore + negativeYield, ethYieldManager.accumulatedNegativeYields());
    }

    // TODO claim should be always 0 for LidoYP.recordUnstaked
}

/// @notice Tests related to accounting (incl. negative yields and withdrawal discounts)
///         This test uses the insurance contract (i.e. commitYieldReport(false))
contract ETH_YieldManager_Accounting_With_Insurance_Test is LidoYieldProvider_Initializer, Util {
    error InsufficientInsuranceBalance();
    error NegativeYieldIncrease();
    // withdrawal buffer

    // when insurance is enabled and it doesn't have enough balance to cover
    // the negative yield, it should revert
    function testFuzz_commitYieldReport_insufficient_insurance_to_cover_negative_yields_reverts(
        uint256 balance,
        uint256 negativeYield,
        uint256 insuranceBalance
    ) external {
        vm.assume(balance > 1 ether);
        vm.assume(balance < 100000 ether);
        vm.assume(insuranceBalance < 100000 ether);
        vm.assume(negativeYield < 100000 ether);
        vm.assume(negativeYield > 0);
        vm.assume(insuranceBalance < negativeYield);

        // emulate negative yield
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + negativeYield);
        // 1 ether is a buffer
        vm.deal(address(this), balance + insuranceBalance + 1 ether);
        Lido.submit{value: balance}(address(0));
        Lido.transfer(address(ethYieldManager), balance);

        // set up insurance balance
        Lido.transfer(address(insurance), insuranceBalance);

        vm.expectRevert(InsufficientInsuranceBalance.selector);
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
    }

    // commitYieldReport(true) should revert upon claim loss if insurance doesn't have enough balance
    function testFuzz_commitYieldReport_insufficient_insurance_to_cover_claim_loss_reverts(
        uint256 balance,
        uint256 unstakeAmount
    ) external {
        // no negative yields
        vm.assume(balance > 0);
        vm.assume(1000 ether > balance);
        vm.assume(balance > unstakeAmount);
        vm.assume(unstakeAmount > 100);

        // make sure withdrawal queue is not paused
        vm.warp(1706000000);

        // set up balance and stake and give 1 ether to insurance
        vm.deal(address(this), balance + unstakeAmount + 1 ether);
        Lido.submit{value: balance + unstakeAmount}(address(0));
        Lido.transfer(address(ethYieldManager), balance + unstakeAmount);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + unstakeAmount);

        assertEq(Lido.balanceOf(address(insurance)), 0);

        // unstake and finalize
        vm.prank(multisig);
        ethYieldManager.unstake(0, address(lidoProvider), unstakeAmount);

        finalizeWithdrawals();

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        assertEq(sharePriceBefore, 1e27);

        // revert with InsufficientInsuranceBalance or NegativeYieldIncrease
        vm.expectRevert();
        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);
    }

    // commitYieldReport(true) should pay off negative yields from insurance
    function testFuzz_commitYieldReport_sufficient_insurance_to_cover_negative_yields_succeeds(
        uint256 balance,
        uint256 negativeYield,
        uint256 insuranceBalance
    ) external {
        vm.assume(balance > 1 ether);
        vm.assume(balance < 100000 ether);
        vm.assume(insuranceBalance < 100000 ether);
        vm.assume(negativeYield < 100000 ether);
        vm.assume(negativeYield > 0);
        vm.assume(insuranceBalance > negativeYield);
        // to avoid precision errors
        vm.assume(insuranceBalance > 100);

        // emulate negative yield
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + negativeYield);
        // 1 ether is a buffer
        vm.deal(address(this), balance + insuranceBalance + 1 ether);
        Lido.submit{value: balance + insuranceBalance + 1 ether}(address(0));
        Lido.transfer(address(ethYieldManager), balance);

        // set up insurance balance
        Lido.transfer(address(insurance), insuranceBalance);

        vm.prank(multisig);
        ethYieldManager.commitYieldReport(true);

        assertEq(ethYieldManager.accumulatedNegativeYields(), 0);
    }

    // ensure that share price remains 1 when pending balance decreases and a loss can be realized as a result of LidoYP._claim
    function testFuzz_commitYieldReportAfterInsuranceWithdrawal_succeeds_after_claim_loss(
        uint256 balance,
        uint256 unstakeAmount
    ) external {
        // share price = 1
        vm.assume(balance > 0);
        vm.assume(100000 ether > balance);
        vm.assume(balance > unstakeAmount);
        vm.assume(1000 ether > unstakeAmount);
        vm.assume(unstakeAmount > 100);

        // make sure withdrawal queue is not paused
        vm.warp(1706000000);

        // set up balance and stake and give 1 ether to insurance
        vm.deal(address(this), balance + unstakeAmount + 1 ether);
        Lido.submit{value: balance + unstakeAmount}(address(0));
        Lido.transfer(address(ethYieldManager), balance + unstakeAmount);
        Lido.transfer(address(insurance), 1 ether);
        vm.prank(address(l1BlastBridge));
        ethYieldManager.recordStakedDeposit(address(lidoProvider), balance + unstakeAmount);


        // unstake and finalize
        vm.prank(multisig);
        ethYieldManager.unstake(0, address(lidoProvider), unstakeAmount);

        finalizeWithdrawals();

        uint256 sharePriceBefore = ethYieldManager.sharePrice();
        assertEq(sharePriceBefore, 1e27);

        vm.prank(multisig);
        ethYieldManager.commitYieldReportAfterInsuranceWithdrawal(address(Lido), 0.1 ether);
        assertEq(ethYieldManager.sharePrice(), sharePriceBefore);
    }
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

        YieldManager.ProviderInfo memory info = usdYieldManager.getProviderInfoAt(0);
        assertEq(info.providerAddress, address(dsrProvider));
        assertEq(info.id, keccak256(abi.encodePacked("DSRYieldProvider", string(abi.encodePacked("1.0.0")))));
        assertEq(info.stakedPrincipal, uint256(0));
        assertClose(info.stakedBalance, 0 ether);
        assertClose(uint256(info.yield), 0 ether);
    }

    function test_stake_DSR_succeeds() external {
        vm.startPrank(0x9759A6Ac90977b93B58547b4A71c78317f391A28);
        RealDAI.mint(address(usdYieldManager), 10 ether);
        vm.stopPrank();

        vm.prank(multisig);
        usdYieldManager.stake(0, address(dsrProvider), 10 ether);

        IPot pot = IPot(DSR_MANAGER.pot());
        vm.mockCall(
            address(pot),
            abi.encodeWithSelector(
                pot.rho.selector
            ),
            abi.encode(block.timestamp - 30 days)
        );

        YieldManager.ProviderInfo memory info = usdYieldManager.getProviderInfoAt(0);
        assertEq(info.providerAddress, address(dsrProvider));
        assertEq(info.id, keccak256(abi.encodePacked("DSRYieldProvider", string(abi.encodePacked("1.0.0")))));
        assertEq(info.stakedPrincipal, uint256(10 ether));
        assertGt(info.stakedBalance, 10 ether);
        uint256 yield = info.stakedBalance - 10 ether;
        assertEq(uint256(info.yield), yield);
        uint256 totalValue = usdYieldManager.totalValue();
        assertEq(totalValue, 10 ether + yield);
    }

    function test_unstake_DSR_succeeds() external {
        vm.startPrank(0x9759A6Ac90977b93B58547b4A71c78317f391A28);
        RealDAI.mint(address(usdYieldManager), 10 ether);
        vm.stopPrank();

        vm.startPrank(multisig);
        usdYieldManager.stake(0, address(dsrProvider), 10 ether);

        assertEq(RealDAI.balanceOf(address(usdYieldManager)), 0);

        skip(3600);

        vm.recordLogs();
        usdYieldManager.unstake(0, address(dsrProvider), 10 ether);
        vm.stopPrank();
        Vm.Log[] memory entries = vm.getRecordedLogs();

        uint256 logLength = entries.length;
        uint256 claimedLogIndex = logLength - 1;
        uint256 unstakedLogIndex = logLength - 2;

        assertEq(entries[claimedLogIndex].topics[0], keccak256("Claimed(bytes32,uint256,uint256)"));
        assertEq(entries[claimedLogIndex].topics[1], YieldProvider(dsrProvider).id());
        (uint256 claimed, uint256 expected) = abi.decode(entries[claimedLogIndex].data, (uint256,uint256));
        assertEq(claimed, uint256(10 ether));
        assertEq(expected, uint256(10 ether));

        assertEq(entries[unstakedLogIndex].topics[0], keccak256("Unstaked(bytes32,uint256)"));
        assertEq(entries[unstakedLogIndex].topics[1], YieldProvider(dsrProvider).id());
        uint256 unstaked = abi.decode(entries[unstakedLogIndex].data, (uint256));
        assertEq(unstaked, uint256(10 ether));

        assertEq(RealDAI.balanceOf(address(usdYieldManager)), 10 ether);
    }

    /*
    function test_recordStakedDeposit_DSR_succeeds() external {
        vm.prank(address(l1BlastBridge));
        l1BlastBridge.recordStakedDeposit(address(dsrProvider), 10 ether);
        YieldManager.ProviderInfo memory info = l1BlastBridge.getProviderInfoAt(0);
        assertEq(info.stakedPrincipal, uint256(10 ether));
    }
    */

    function test_commitYieldReport_DSR_positive_yield_no_insurance_succeeds() external {
        uint256 yield = emulatePositiveYield(10 ether, 30 days);

        vm.recordLogs();
        vm.startPrank(multisig);
        usdYieldManager.commitYieldReport(false);
        vm.stopPrank();
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
        assertEq(entries[2].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(usdYieldManager))));
        assertEq(entries[2].topics[2], addressToBytes32(Predeploys.USDB));
        // TODO: verify the tx data
    }

    function test_commitYieldReport_DSR_positive_yield_insurance_premium_succeeds() external {
        uint256 yield = emulatePositiveYield(10 ether, 30 days);

        vm.recordLogs();
        vm.startPrank(multisig);
        usdYieldManager.commitYieldReport(true);
        vm.stopPrank();
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
        assertEq(entries[depositLogIndex].topics[1], addressToBytes32(AddressAliasHelper.applyL1ToL2Alias(address(usdYieldManager))));
        assertEq(entries[depositLogIndex].topics[2], addressToBytes32(Predeploys.USDB)); // USDB

        uint256 insuranceBalance = RealDAI.balanceOf(usdYieldManager.insurance());
        assertClose(insuranceBalance, uint256(yield / 10)); // 10% of yield
    }

    function emulatePositiveYield(uint256 initialAmount, uint256 elapsed) internal returns (uint256 yield) {
        vm.startPrank(0x9759A6Ac90977b93B58547b4A71c78317f391A28);
        RealDAI.mint(address(usdYieldManager), initialAmount);
        vm.stopPrank();

        vm.prank(multisig);
        usdYieldManager.stake(0, address(dsrProvider), initialAmount);

        skip(elapsed);

        YieldManager.ProviderInfo memory info = usdYieldManager.getProviderInfoAt(0);
        yield = info.stakedBalance - initialAmount;
    }
}

contract USD_Insurance_Test is DSRYieldProvider_Initializer, Util {

    function setUp() public override {
        super.setUp();

        vm.startPrank(0x9759A6Ac90977b93B58547b4A71c78317f391A28);
        RealDAI.mint(address(insurance), 10 ether);
        vm.stopPrank();
    }

    function test_setAdmin_succeeds() external {
        assertEq(insurance.admin(), multisig);
        vm.prank(multisig);
        insurance.setAdmin(address(0x4));
        assertEq(insurance.admin(), address(0x4));
    }

    function test_coverLoss_admin_succeeds() external {

        assertEq(insurance.admin(), multisig);

        assertEq(usdYieldManager.totalValue(), 0 ether);

        vm.prank(multisig);
        insurance.coverLoss(address(RealDAI), 1 ether);

        assertEq(usdYieldManager.totalValue(), 1 ether);
    }

    function test_coverLoss_ym_succeeds() external {

        assertEq(usdYieldManager.totalValue(), 0 ether);

        vm.prank(address(usdYieldManager));
        insurance.coverLoss(address(RealDAI), 1 ether);

        assertClose(usdYieldManager.totalValue(), 1 ether);
    }

    function test_coverloss_non_admin_reverts() external {

        assertEq(usdYieldManager.totalValue(), 0 ether);

        vm.expectRevert(abi.encodeWithSelector(Insurance.OnlyAdminOrYieldManager.selector));

        vm.prank(address(0x4));
        insurance.coverLoss(address(RealDAI), 1 ether);
    }
}
