// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
// Target contract is imported by the `Bridge_Initializer`
import { Bridge_Initializer } from "test/CommonTest.t.sol";
import { ILegacyMintableERC20, IOptimismMintableERC20 } from "src/universal/IOptimismMintableERC20.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

// Libraries
import { Predeploys } from "src/libraries/Predeploys.sol";

// Target contract dependencies
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { YieldMode } from "src/L2/Blast.sol";
import { GasMode } from "src/L2/Gas.sol";

contract Shares_Test is Bridge_Initializer {
    event NewPrice(uint256 price);

    function setUp() public virtual override {
        super.setUp();

        setShareCount(1);
    }

    function test_blastConfig() external {
        checkBlastConfig(address(SHARES), address(0xdead), YieldMode.VOID, GasMode.VOID);
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

    function test_blastConfig() external {
        checkBlastConfig(address(WETH), address(0xdead), YieldMode.AUTOMATIC, GasMode.VOID);
    }

    function totalETH() internal view returns (uint256) {
        return address(WETH).balance + alice.balance + bob.balance + charlie.balance;
    }

    function addYield(uint256 yield) internal {
        uint256 perToken = 1e18 + (yield * 1e18) / totalETH();
        vm.deal(address(WETH), (perToken * address(WETH).balance) / 1e18);
        vm.deal(alice, (perToken*alice.balance) / 1e18);
        vm.deal(bob, (perToken*bob.balance) / 1e18);
        vm.deal(charlie, (perToken*charlie.balance) / 1e18);
        vm.prank(Predeploys.SHARES);
        WETH.addValue(0);
    }

    function addBalance(address account, uint256 amount) internal {
        vm.deal(account, account.balance + amount);
    }

    function testConfigure(YieldMode yieldMode, uint256 expectedBalance, uint256 expectedSharePrice) internal {
        vm.prank(alice);
        WETH.configure(yieldMode);
        assertEq(WETH.price(), expectedSharePrice);
        assertEq(WETH.balanceOf(alice), expectedBalance);
        assertTrue(WETH.getConfiguration(alice) == yieldMode);
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
        uint256 price = WETH.price();
        vm.deal(alice, price);
        vm.prank(alice);
        WETH.deposit{value: price}();

        addYield(totalETH());
        uint256 aliceBalance = alice.balance;

        vm.expectEmit(true, true, true, true);
        emit Withdrawal(alice, 2*price);

        vm.prank(alice);
        WETH.withdraw(2*price);
        assertEq(WETH.balanceOf(alice), 0);
        assertEq(alice.balance, aliceBalance+2*price);
    }

    function test_withdraw_insufficientBalance_fails() external {
        uint256 price = WETH.price();
        vm.deal(alice, price);
        vm.prank(alice);
        WETH.deposit{value: price}();

        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        WETH.withdraw(2*price);

        assertEq(WETH.balanceOf(alice), price);
    }
}
