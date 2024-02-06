// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Bridge_Initializer } from "test/CommonTest.t.sol";
import { ILegacyMintableERC20, IOptimismMintableERC20 } from "src/universal/IOptimismMintableERC20.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import { YieldMode } from "src/L2/ERC20Rebasing.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";

contract Shares_Test is Bridge_Initializer {
    event NewPrice(uint256 price);

    function setUp() public virtual override {
        super.setUp();

        setShareCount(1);
    }

    function totalETH() internal view returns (uint256) {
        return address(WETH).balance + alice.balance + bob.balance + charlie.balance;
    }

    function setShareCount(uint256 count) internal {
        vm.store(address(SHARES), bytes32(uint256(51)), bytes32(count));
    }

    function test_addValue_succeeds() external {
        uint256 initialPrice = SHARES.price();

        setShareCount(10);

        vm.expectEmit(true, true, true, true);
        emit NewPrice(initialPrice + 1);

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager));
        vm.prank(caller);
        SHARES.addValue(10);

        assertEq(SHARES.price(), initialPrice + 1);
        assertEq(SHARES.pending(), 0);
    }

    function test_addValue_pending_succeeds() external {
        setShareCount(2e18);

        uint256 initialPrice = SHARES.price();

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager));
        vm.prank(caller);
        SHARES.addValue(1);

        assertEq(SHARES.price(), initialPrice);
        assertEq(SHARES.pending(), 1);
    }

    function testFuzz_addValue(uint256 shares, uint256 yield) external {
        vm.assume(shares < 1e30);
        vm.assume(yield <= shares);

        uint256 initialPrice = SHARES.price();
        setShareCount(shares);

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(ethYieldManager));
        vm.prank(caller);
        SHARES.addValue(yield);

        assertEq(SHARES.price() * SHARES.count() + SHARES.pending(), shares * initialPrice + yield);
    }
}

contract WETHRebasing_Test is Bridge_Initializer {
    event Deposit(address indexed dst, uint256 wad);
    event Withdrawal(address indexed src, uint256 wad);

    error InsufficientBalance();
    error InsufficientAllowance();
    error NotClaimableAccount();

    function totalETH() internal view returns (uint256) {
        return address(WETH).balance + alice.balance + bob.balance + charlie.balance;
    }

    function addYield(uint256 yield) internal {
        uint256 perToken = 1e18 + (yield * 1e18) / totalETH();
        vm.deal(address(WETH), (perToken * address(WETH).balance) / 1e18);
        vm.deal(alice, (perToken*alice.balance) / 1e18);
        vm.deal(bob, (perToken*bob.balance) / 1e18);
        vm.deal(charlie, (perToken*charlie.balance) / 1e18);
    }

    function addBalance(address account, uint256 amount) internal {
        vm.deal(account, account.balance + amount);
    }

    function testConfigure(YieldMode yieldMode, uint256 expectedBalance, uint256 expectedSharePrice) internal {
        vm.prank(alice);
        WETH.configure(yieldMode);
        assertEq(WETH.sharePrice(), expectedSharePrice);
        assertEq(WETH.balanceOf(alice), expectedBalance);
        assertTrue(WETH.getConfiguration(alice) == yieldMode);
    }

    function test_configure_succeeds() external {
        vm.prank(alice);
        WETH.deposit{value: WETH.sharePrice()}();

        // defaults to Automatic
        assertTrue(WETH.getConfiguration(alice) == YieldMode.AUTOMATIC);

        uint256 aliceBalance = WETH.balanceOf(alice);
        uint256 sharePrice = WETH.sharePrice();

        // AUTOMATIC -> CLAIMABLE
        testConfigure(YieldMode.CLAIMABLE, aliceBalance, sharePrice);
        // CLAIMABLE -> AUTOMATIC
        testConfigure(YieldMode.AUTOMATIC, aliceBalance, sharePrice);
        // AUTOMATIC -> VOID
        testConfigure(YieldMode.VOID, aliceBalance, sharePrice);
        // VOID -> CLAIMABLE
        testConfigure(YieldMode.CLAIMABLE, aliceBalance, sharePrice);
        // CLAIMABLE -> VOID
        testConfigure(YieldMode.VOID, aliceBalance, sharePrice);
        // VOID -> AUTOMATIC
        testConfigure(YieldMode.AUTOMATIC, aliceBalance, sharePrice);
    }

    function test_transfer_succeeds() external {
        vm.prank(alice);
        WETH.deposit{value: 100}();
        assertEq(WETH.balanceOf(alice), 100);
        vm.prank(bob);
        WETH.deposit{value: 100}();
        assertEq(WETH.balanceOf(bob), 100);
        vm.prank(charlie);
        WETH.deposit{value: 100}();
        assertEq(WETH.balanceOf(charlie), 100);

        vm.prank(bob);
        WETH.configure(YieldMode.VOID);

        uint256 sharePrice = WETH.sharePrice();

        vm.prank(charlie);
        WETH.configure(YieldMode.CLAIMABLE);

        vm.startPrank(alice);
        WETH.transfer(bob, 10);
        assertEq(WETH.balanceOf(alice), 90);
        assertEq(WETH.balanceOf(bob), 110);
        assertEq(WETH.sharePrice(), sharePrice);
        WETH.transfer(charlie, 10);
        assertEq(WETH.balanceOf(alice), 80);
        assertEq(WETH.balanceOf(charlie), 110);
        assertEq(WETH.sharePrice(), sharePrice);
        vm.stopPrank();

        vm.startPrank(bob);
        WETH.transfer(alice, 10);
        assertEq(WETH.balanceOf(bob), 100);
        assertEq(WETH.balanceOf(alice), 90);
        assertEq(WETH.sharePrice(), sharePrice);
        WETH.transfer(charlie, 10);
        assertEq(WETH.balanceOf(bob), 90);
        assertEq(WETH.balanceOf(charlie), 120);
        assertEq(WETH.sharePrice(), sharePrice);
        vm.stopPrank();

        vm.startPrank(charlie);
        WETH.transfer(alice, 10);
        assertEq(WETH.balanceOf(charlie), 110);
        assertEq(WETH.balanceOf(alice), 100);
        assertEq(WETH.sharePrice(), sharePrice);
        WETH.transfer(bob, 10);
        assertEq(WETH.balanceOf(charlie), 100);
        assertEq(WETH.balanceOf(bob), 100);
        assertEq(WETH.sharePrice(), sharePrice);
        vm.stopPrank();
    }

    function test_transfer_insufficientBalance_reverts() external {
        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        WETH.transfer(bob, 10);
    }

    function test_claim_succeeds() external {
        vm.prank(alice);
        WETH.configure(YieldMode.CLAIMABLE);

        uint256 sharePrice = WETH.sharePrice();
        vm.prank(alice);
        WETH.deposit{value: sharePrice}();

        assertEq(WETH.balanceOf(alice), sharePrice);

        addYield(totalETH());

        WETH.balanceOf(alice);
        assertEq(WETH.getClaimableAmount(alice), sharePrice);
        vm.prank(alice);
        WETH.claim(bob, sharePrice/2);
        assertEq(WETH.balanceOf(alice), sharePrice);
        assertEq(WETH.balanceOf(bob), sharePrice/2);
        assertEq(WETH.getClaimableAmount(alice), sharePrice/2);
    }

    function test_claim_insufficientFunds_reverts() external {
        vm.prank(alice);
        WETH.configure(YieldMode.CLAIMABLE);

        vm.prank(alice);
        WETH.deposit{value: 100}();

        assertEq(WETH.balanceOf(alice), 100);

        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        WETH.claim(bob, 1);
    }

    function test_claim_notClaimable_reverts() external {
        vm.prank(alice);
        WETH.configure(YieldMode.AUTOMATIC);

        vm.prank(alice);
        WETH.deposit{value: 100}();

        vm.expectRevert(NotClaimableAccount.selector);
        vm.prank(alice);
        Usdb.claim(bob, 1);
    }

    function test_approval_succeeds() external {
        vm.prank(alice);
        WETH.approve(bob, 100);

        assertEq(WETH.allowance(alice, bob), 100);
    }

    function test_transferFrom_withApproval_succeeds() external {
        vm.prank(alice);
        WETH.deposit{value: 100}();

        vm.prank(alice);
        WETH.approve(bob, 100);

        vm.prank(bob);
        WETH.transferFrom(alice, bob, 100);

        assertEq(WETH.balanceOf(alice), 0);
        assertEq(WETH.balanceOf(bob), 100);
    }

    function test_transferFrom_withoutApproval_reverts() external {
        vm.prank(alice);
        WETH.deposit{value: 100}();

        vm.expectRevert(InsufficientAllowance.selector);
        vm.prank(bob);
        WETH.transferFrom(alice, bob, 100);
    }

    function test_rebasing() external {
        uint256 sharePrice = WETH.sharePrice();
        vm.prank(alice);
        WETH.deposit{value: sharePrice/2}();
        vm.prank(bob);
        WETH.deposit{value: sharePrice/2}();
        vm.prank(charlie);
        WETH.deposit{value: sharePrice}();

        addBalance(address(WETH), 1);

        assertEq(WETH.balanceOf(alice), sharePrice/2);
        assertEq(WETH.balanceOf(bob), sharePrice/2);
        assertEq(WETH.balanceOf(charlie), sharePrice);
        assertEq(WETH.sharePrice(), sharePrice);

        addBalance(address(WETH), 2);

        assertEq(WETH.balanceOf(alice), sharePrice/2);
        assertEq(WETH.balanceOf(bob), sharePrice/2);
        assertEq(WETH.balanceOf(charlie), sharePrice + 1);
        assertEq(WETH.sharePrice(), sharePrice + 1);
    }

    function test_deposit_succeeds() external {
        vm.prank(alice);
        WETH.configure(YieldMode.AUTOMATIC);
        vm.prank(bob);
        WETH.configure(YieldMode.VOID);
        vm.prank(charlie);
        WETH.configure(YieldMode.CLAIMABLE);

        vm.expectEmit(true, true, true, true);
        emit Deposit(alice, 100);
        vm.prank(alice);
        WETH.deposit{value: 100}();
        assertEq(WETH.balanceOf(alice), 100);

        vm.expectEmit(true, true, true, true);
        emit Deposit(bob, 100);
        vm.prank(bob);
        WETH.deposit{value: 100}();
        assertEq(WETH.balanceOf(bob), 100);

        vm.expectEmit(true, true, true, true);
        emit Deposit(charlie, 100);
        vm.prank(charlie);
        WETH.deposit{value: 100}();
        assertEq(WETH.balanceOf(charlie), 100);
    }

    function test_withdraw_succeeds() external {
        vm.deal(alice, 1e6);
        vm.prank(alice);
        WETH.deposit{value: 1e6}();

        addYield(totalETH());
        uint256 aliceBalance = alice.balance;

        vm.expectEmit(true, true, true, true);
        emit Withdrawal(alice, 2e6);

        vm.prank(alice);
        WETH.withdraw(2e6);
        assertEq(WETH.balanceOf(alice), 0);
        assertEq(alice.balance, aliceBalance+2e6);
    }

    function test_withdraw_insufficientBalance_fails() external {
        vm.prank(alice);
        WETH.deposit{value: 100}();

        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        WETH.withdraw(200);
        assertEq(WETH.balanceOf(alice), 100);
    }

    function testFuzz_transfer(uint8 aliceConfiguration, uint8 bobConfiguration, uint256 transferAmount) external {
        uint256 aliceBalance = alice.balance;
        vm.assume(transferAmount <= aliceBalance);
        vm.assume(aliceConfiguration <= 2);
        vm.assume(bobConfiguration <= 2);

        vm.prank(alice);
        WETH.deposit{value: aliceBalance}();

        vm.prank(alice);
        WETH.configure(YieldMode(aliceConfiguration));

        vm.prank(bob);
        WETH.configure(YieldMode(bobConfiguration));

        vm.prank(alice);
        WETH.transfer(bob, transferAmount);

        assertEq(WETH.balanceOf(alice), aliceBalance - transferAmount);
        assertEq(WETH.balanceOf(bob), transferAmount);
    }

    /*
    function testFuzz_yield(
        uint256 aliceDeposit,
        uint256 bobDeposit,
        uint256 charlieDeposit,
        uint256 yield
    ) external {
        vm.assume(aliceDeposit <= alice.balance);
        vm.assume(bobDeposit <= bob.balance);
        vm.assume(charlieDeposit <= charlie.balance);
        vm.assume(yield < 100e6);

        vm.prank(alice);
        WETH.configure(YieldMode.AUTOMATIC);
        vm.prank(bob);
        WETH.configure(YieldMode.VOID);
        vm.prank(charlie);
        WETH.configure(YieldMode.CLAIMABLE);

        vm.prank(alice);
        WETH.deposit{value: aliceDeposit}();
        vm.prank(bob);
        WETH.deposit{value: bobDeposit}();
        vm.prank(charlie);
        WETH.deposit{value: charlieDeposit}();

        uint256 sharePrice = WETH.sharePrice();
        uint256 aliceYield = (yield * (aliceDeposit - aliceDeposit % sharePrice)) / (address(WETH).balance - bobDeposit) + aliceDeposit;
        uint256 charlieYield = (yield * (charlieDeposit - charlieDeposit % sharePrice)) / (address(WETH).balance - bobDeposit);
        addBalance(address(WETH), yield);

        assertEq(WETH.balanceOf(alice), aliceYield);
        assertEq(WETH.balanceOf(bob), bobDeposit);
        assertEq(WETH.balanceOf(charlie), charlieDeposit);
        assertEq(WETH.getClaimableAmount(charlie), charlieYield);
    }
    */
}
