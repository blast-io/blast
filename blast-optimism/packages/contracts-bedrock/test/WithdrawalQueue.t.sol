pragma solidity 0.8.15;

// Testing utilities
import { Test, StdUtils, Vm, StdStorage, stdStorage } from "forge-std/Test.sol";
import { Portal_Initializer, LidoYieldProvider_Initializer, CommonTest, NextImpl } from "test/CommonTest.t.sol";
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

// Target contracts
import { WithdrawalQueue } from "src/mainnet-bridge/withdrawal-queue/WithdrawalQueue.sol";

contract MockWithdrawalQueue is WithdrawalQueue {
    constructor(address token) WithdrawalQueue(token) {}

    function initialize() external initializer {
        __WithdrawalQueue_init();
    }

    function requestWithdrawal(address recipient, uint256 amount) external returns (uint256) {
        return _requestWithdrawal(recipient, amount);
    }

    function finalize_external(uint256 _lastRequestIdToBeFinalized, uint256 availableBalance, uint256 sharePrice) external returns (uint256 nominalAmountToFinalize, uint256 realAmountToFinalize, uint256 checkpointId) {
        return _finalize(_lastRequestIdToBeFinalized, availableBalance, sharePrice);
    }
}

contract WithdrawalQueue_Test is Test {
    MockWithdrawalQueue ethQueue;
    address alice = address(128);
    address bob = address(257);
    address charlie = address(512);
    address dave = address(1024);

    function setUp() external {
        ethQueue = new MockWithdrawalQueue(address(0));
        ethQueue.initialize();
    }

    function test_requestWithdrawal_succeeds() external {
        uint256 firstRequestId = ethQueue.requestWithdrawal(address(0x1), 1 ether);
        assertEq(firstRequestId, 1);
        assertEq(ethQueue.getLastRequestId(), 1);

        uint256 secondRequestId = ethQueue.requestWithdrawal(address(0x2), 2 ether);
        assertEq(secondRequestId, 2);
        assertEq(ethQueue.getLastRequestId(), 2);

        uint256[] memory requestIds = new uint256[](2);
        requestIds[0] = firstRequestId;
        requestIds[1] = secondRequestId;
        WithdrawalQueue.WithdrawalRequestStatus[] memory statuses = ethQueue.getWithdrawalStatus(
            requestIds
        );

        assertEq(statuses.length, 2);
        assertEq(statuses[0].amount, 1 ether);
        assertEq(statuses[0].recipient, address(0x1));
        assertEq(statuses[0].isFinalized, false);
        assertEq(statuses[0].isClaimed, false);

        assertEq(statuses[1].amount, 2 ether);
        assertEq(statuses[1].recipient, address(0x2));
        assertEq(statuses[1].isFinalized, false);
        assertEq(statuses[1].isClaimed, false);
    }

    function test_finalize_succeeds() external {
        ethQueue.requestWithdrawal(address(0x1), 1 ether);
        ethQueue.requestWithdrawal(address(0x2), 2 ether);
        ethQueue.requestWithdrawal(address(0x3), 3 ether);

        vm.recordLogs();
        (uint256 nominal, uint256 real,) = ethQueue.finalize_external(2, 10 ether, 1.0e27);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        assertEq(nominal, 3 ether);
        assertEq(real, 3 ether);

        assertEq(ethQueue.getLastCheckpointId(), 1);
        assertEq(ethQueue.getLastFinalizedRequestId(), 2);
        assertEq(ethQueue.getLockedBalance(), 3 ether);

        assertEq(entries.length, 1);
        assertEq(entries[0].topics[0], keccak256("WithdrawalsFinalized(uint256,uint256,uint256,uint256,uint256,uint256)"));
        assertEq(entries[0].topics[1], bytes32(uint256(1)));  // from
        assertEq(entries[0].topics[2], bytes32(uint256(2)));  // to
        assertEq(entries[0].topics[3], bytes32(uint256(1)));  // checkpointId
        (uint256 lockedAmount, , uint256 sharePrice) = abi.decode(entries[0].data, (uint256, uint256, uint256));
        assertEq(lockedAmount, 3 ether);
        assertEq(sharePrice, 1.0e27);

        ethQueue.requestWithdrawal(address(0x4), 4 ether);

        // second batch
        vm.recordLogs();
        (nominal, real,) = ethQueue.finalize_external(4, 10 ether, 1.0e27);
        entries = vm.getRecordedLogs();

        assertEq(nominal, 7 ether);
        assertEq(real, 7 ether);

        assertEq(ethQueue.getLastCheckpointId(), 2);
        assertEq(ethQueue.getLastFinalizedRequestId(), 4);
        assertEq(ethQueue.getLockedBalance(), 10 ether);

        assertEq(entries.length, 1);
        assertEq(entries[0].topics[0], keccak256("WithdrawalsFinalized(uint256,uint256,uint256,uint256,uint256,uint256)"));
        assertEq(entries[0].topics[1], bytes32(uint256(3)));  // from
        assertEq(entries[0].topics[2], bytes32(uint256(4)));  // to
        assertEq(entries[0].topics[3], bytes32(uint256(2)));  // checkpointId
        (lockedAmount, , sharePrice) = abi.decode(entries[0].data, (uint256, uint256, uint256));
        assertEq(lockedAmount, 7 ether);
        assertEq(sharePrice, 1.0e27);
    }

    function test_finalize_lower_share_price_succeeds() external {
        ethQueue.requestWithdrawal(address(0x1), 1 ether);
        ethQueue.requestWithdrawal(address(0x2), 2 ether);
        ethQueue.requestWithdrawal(address(0x3), 3 ether);

        vm.recordLogs();
        (uint256 nominal, uint256 real,) = ethQueue.finalize_external(2, 10 ether, 0.9e27);
        Vm.Log[] memory entries = vm.getRecordedLogs();

        assertEq(nominal, 3 ether);
        assertEq(real, 2.7 ether);

        assertEq(ethQueue.getLastCheckpointId(), 1);
        assertEq(ethQueue.getLastFinalizedRequestId(), 2);
        assertEq(ethQueue.getLockedBalance(), 2.7 ether);

        assertEq(entries.length, 1);
        assertEq(entries[0].topics[0], keccak256("WithdrawalsFinalized(uint256,uint256,uint256,uint256,uint256,uint256)"));
        assertEq(entries[0].topics[1], bytes32(uint256(1)));  // from
        assertEq(entries[0].topics[2], bytes32(uint256(2)));  // to
        assertEq(entries[0].topics[3], bytes32(uint256(1)));  // checkpointId
        (uint256 lockedAmount, , uint256 sharePrice) = abi.decode(entries[0].data, (uint256, uint256, uint256));
        assertEq(lockedAmount, 2.7 ether);
        assertEq(sharePrice, 0.9e27);

        ethQueue.requestWithdrawal(address(0x4), 4 ether);

        // second batch
        vm.recordLogs();
        (nominal, real,) = ethQueue.finalize_external(4, 10 ether, 0.8e27);
        entries = vm.getRecordedLogs();

        assertEq(nominal, 7 ether);
        assertEq(real, 5.6 ether);

        assertEq(ethQueue.getLastCheckpointId(), 2);
        assertEq(ethQueue.getLastFinalizedRequestId(), 4);
        assertEq(ethQueue.getLockedBalance(), 8.3 ether);  // 2.7 + 5.6

        assertEq(entries.length, 1);
        assertEq(entries[0].topics[0], keccak256("WithdrawalsFinalized(uint256,uint256,uint256,uint256,uint256,uint256)"));
        assertEq(entries[0].topics[1], bytes32(uint256(3)));  // from
        assertEq(entries[0].topics[2], bytes32(uint256(4)));  // to
        assertEq(entries[0].topics[3], bytes32(uint256(2)));  // checkpointId
        (lockedAmount, , sharePrice) = abi.decode(entries[0].data, (uint256, uint256, uint256));
        assertEq(lockedAmount, 5.6 ether);
        assertEq(sharePrice, 0.8e27);
    }

    // TODO: test more revert cases for finalize

    function test_finalize_invalid_share_price_reverts() external {
        vm.expectRevert(WithdrawalQueue.InvalidSharePrice.selector);
        ethQueue.finalize_external(1, 10 ether, 1.1e27);
    }

    function test_finalize_insufficient_balance_reverts() external {
        ethQueue.requestWithdrawal(alice, 1 ether);
        ethQueue.requestWithdrawal(bob, 2 ether);
        ethQueue.requestWithdrawal(charlie, 3 ether);
        ethQueue.requestWithdrawal(dave, 4 ether);

        ethQueue.finalize_external(2, 2.7 ether, 0.9e27); // (1 + 2) * 0.9 == 2.7 should not revert
        assertEq(ethQueue.getLockedBalance(), 2.7 ether);

        vm.expectRevert(WithdrawalQueue.InsufficientBalance.selector);
        ethQueue.finalize_external(4, 4.8 ether, 0.7e27); // (3 + 4) * 0.7 > 4.8 should revert
    }

    function test_finalize_insufficient_balance_2_reverts() external {
        ethQueue.requestWithdrawal(alice, 1 ether);
        ethQueue.requestWithdrawal(bob, 2 ether);
        ethQueue.requestWithdrawal(charlie, 3 ether);

        vm.expectRevert(WithdrawalQueue.InsufficientBalance.selector);
        ethQueue.finalize_external(2, 1 ether, 1.0e27);
    }

    function test_claim_eth_succeeds() external {
        vm.deal(address(ethQueue), 10 ether);

        ethQueue.requestWithdrawal(bob, 1 ether);
        ethQueue.requestWithdrawal(alice, 2 ether);
        ethQueue.requestWithdrawal(charlie, 3 ether);
        ethQueue.finalize_external(2, address(ethQueue).balance, 1.0e27);

        vm.prank(alice);
        bool success = ethQueue.claimWithdrawal(2, 1);
        assertTrue(success);
        assertEq(alice.balance, 2 ether);
        assertEq(ethQueue.getLockedBalance(), 1 ether);
        assertEq(address(ethQueue).balance, 8 ether);

        ethQueue.requestWithdrawal(dave, 4 ether);
        ethQueue.finalize_external(4, address(ethQueue).balance, 1.0e27);

        vm.prank(bob);
        success = ethQueue.claimWithdrawal(1, 1);
        assertTrue(success);
        assertEq(bob.balance, 1 ether);
        assertEq(ethQueue.getLockedBalance(), 7 ether);
        assertEq(address(ethQueue).balance, 7 ether);

        vm.prank(dave);
        success = ethQueue.claimWithdrawal(4, 2);
        assertTrue(success);
        assertEq(dave.balance, 4 ether);
        assertEq(ethQueue.getLockedBalance(), 3 ether);
        assertEq(address(ethQueue).balance, 3 ether);
    }

    function test_claim_eth_lower_share_price_succeeds() external {
        vm.deal(address(ethQueue), 10 ether);

        ethQueue.requestWithdrawal(bob, 1 ether);
        ethQueue.requestWithdrawal(alice, 2 ether);
        ethQueue.requestWithdrawal(charlie, 3 ether);
        ethQueue.finalize_external(2, address(ethQueue).balance, 0.9e27);

        vm.prank(alice);
        bool success = ethQueue.claimWithdrawal(2, 1);
        assertTrue(success);
        assertEq(alice.balance, 1.8 ether);
        assertEq(ethQueue.getLockedBalance(), 0.9 ether);
        assertEq(address(ethQueue).balance, 8.2 ether);

        ethQueue.requestWithdrawal(dave, 4 ether);
        ethQueue.finalize_external(4, address(ethQueue).balance, 0.8e27);
        // locked amount = (1 + 2) * 0.9 + (3 + 4) * 0.8 = 8.3

        vm.prank(bob);
        success = ethQueue.claimWithdrawal(1, 1);
        assertTrue(success);
        assertEq(bob.balance, 0.9 ether);
        assertEq(ethQueue.getLockedBalance(), 5.6 ether);
        assertEq(address(ethQueue).balance, 7.3 ether);

        vm.prank(dave);
        success = ethQueue.claimWithdrawal(4, 2);
        assertTrue(success);
        assertEq(dave.balance, 3.2 ether);
        assertEq(ethQueue.getLockedBalance(), 2.4 ether);
        assertEq(address(ethQueue).balance, 4.1 ether);
    }

    function test_findCheckpointHint_succeeds() external {
        uint8[5] memory finalizedRequestIds = [2, 5, 7, 12, 19];

        for (uint256 i = 0; i < 19; i++) {
            ethQueue.requestWithdrawal(alice, 1 ether);
        }

        for (uint256 i = 0; i < finalizedRequestIds.length; i++) {
            ethQueue.finalize_external(uint256(finalizedRequestIds[i]), 100 ether, 1.0e27);
        }

        uint8[5] memory testRequestIds = [1, 3, 6, 8, 13];
        uint8[5] memory expectedCheckpointIds = [1, 2, 3, 4, 5];

        for (uint256 i = 0; i < testRequestIds.length; i++) {
            uint256 checkpoint = ethQueue.findCheckpointHint(uint256(testRequestIds[i]), 1, 5);
            assertEq(checkpoint, uint256(expectedCheckpointIds[i]));
        }
    }

    function test_findCheckpointHints_succeeds() external {
        uint8[5] memory finalizedRequestIds = [2, 5, 7, 12, 19];

        for (uint256 i = 0; i < 19; i++) {
            ethQueue.requestWithdrawal(alice, 1 ether);
        }

        for (uint256 i = 0; i < finalizedRequestIds.length; i++) {
            ethQueue.finalize_external(uint256(finalizedRequestIds[i]), 100 ether, 1.0e27);
        }

        uint256[] memory testRequestIds = new uint256[](5);
        testRequestIds[0] = 1;
        testRequestIds[1] = 3;
        testRequestIds[2] = 6;
        testRequestIds[3] = 8;
        testRequestIds[4] = 13;

        uint256[] memory hintIds = ethQueue.findCheckpointHints(testRequestIds, 1, 5);

        assertEq(hintIds.length, 5);
        assertEq(hintIds[0], 1);
        assertEq(hintIds[1], 2);
        assertEq(hintIds[2], 3);
        assertEq(hintIds[3], 4);
        assertEq(hintIds[4], 5);
    }
}
