// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { CommonTest } from "test/CommonTest.t.sol";
import { ILegacyMintableERC20, IOptimismMintableERC20 } from "src/universal/IOptimismMintableERC20.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import { YieldMode } from "src/L2/ERC20Rebasing.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { ERC20Rebasing } from "src/L2/ERC20Rebasing.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

contract MockERC20Rebasing is ERC20Rebasing {
    constructor(address _reporter, uint8 _decimals) ERC20Rebasing(_reporter, _decimals) {}

    function initialize() public initializer {
        __ERC20Rebasing_init(
            "Test",
            "Test",
            1e9
        );
    }

    function mint(address to, uint256 amount) external {
        _deposit(to, amount);
    }
}

contract ERC20Rebasing_Test is CommonTest {
    MockERC20Rebasing Token;

    event Deposit(address indexed dst, uint256 wad);
    event Withdrawal(address indexed src, uint256 wad);

    error InsufficientBalance();
    error InsufficientAllowance();
    error NotClaimableAccount();

    function setUp() public override {
        super.setUp();

        Token = new MockERC20Rebasing(address(this), 18);
        Token.initialize();
    }

    function addYield(uint256 amount) internal {
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(Token.REPORTER()));
        Token.addValue(amount);
    }

    function testConfigure(YieldMode yieldMode, uint256 expectedBalance, uint256 expectedSharePrice) internal {
        vm.prank(alice);
        Token.configure(yieldMode);
        assertEq(Token.price(), expectedSharePrice);
        assertEq(Token.balanceOf(alice), expectedBalance);
        assertTrue(Token.getConfiguration(alice) == yieldMode);
    }

    function test_configure_succeeds() external {
        Token.mint(alice, Token.price());

        // defaults to Automatic
        assertTrue(Token.getConfiguration(alice) == YieldMode.AUTOMATIC);

        uint256 aliceBalance = Token.balanceOf(alice);
        uint256 price = Token.price();

        addYield(Token.count());
        aliceBalance += 1;
        price += 1;

        // AUTOMATIC -> CLAIMABLE
        testConfigure(YieldMode.CLAIMABLE, aliceBalance, price);
        assertEq(Token.getClaimableAmount(alice), 0);

        addYield(Token.count());
        aliceBalance += 1;
        price += 1;

        // CLAIMABLE -> AUTOMATIC
        testConfigure(YieldMode.AUTOMATIC, aliceBalance, price);

        addYield(Token.count());
        aliceBalance += 1;
        price += 1;

        // AUTOMATIC -> VOID
        testConfigure(YieldMode.VOID, aliceBalance, price);

        addYield(Token.count());

        // VOID -> CLAIMABLE
        testConfigure(YieldMode.CLAIMABLE, aliceBalance, price);
        assertEq(Token.getClaimableAmount(alice), 0);

        addYield(Token.count());
        aliceBalance += 1;
        price += 1;

        // CLAIMABLE -> VOID
        testConfigure(YieldMode.VOID, aliceBalance, price);

        addYield(Token.count());

        // VOID -> AUTOMATIC
        testConfigure(YieldMode.AUTOMATIC, aliceBalance, price);

        addYield(Token.count());
        aliceBalance += 1;
        price += 1;

        // AUTOMATIC -> AUTOMATIC
        testConfigure(YieldMode.AUTOMATIC, aliceBalance, price);

        vm.prank(alice);
        Token.configure(YieldMode.VOID);

        addYield(Token.count());

        // VOID -> VOID
        testConfigure(YieldMode.VOID, aliceBalance, price);

        vm.prank(alice);
        Token.configure(YieldMode.CLAIMABLE);

        addYield(Token.count());
        aliceBalance += 1;
        price += 1;

        // CLAIMABLE -> CLAIMABLE
        testConfigure(YieldMode.CLAIMABLE, aliceBalance, price);
        assertEq(Token.getClaimableAmount(alice), 0);
    }

    function test_transfer_succeeds() external {
        uint256 price = Token.price();
        Token.mint(alice, 10*price);
        assertEq(Token.balanceOf(alice), 10*price);
        Token.mint(bob, 10*price);
        assertEq(Token.balanceOf(bob), 10*price);
        Token.mint(charlie, 10*price);
        assertEq(Token.balanceOf(charlie), 10*price);

        vm.prank(bob);
        Token.configure(YieldMode.VOID);

        vm.prank(charlie);
        Token.configure(YieldMode.CLAIMABLE);

        vm.startPrank(alice);
        Token.transfer(bob, price);
        assertEq(Token.balanceOf(alice), 9*price);
        assertEq(Token.balanceOf(bob), 11*price);
        assertEq(Token.price(), price);
        Token.transfer(charlie, price);
        assertEq(Token.balanceOf(alice), 8*price);
        assertEq(Token.balanceOf(charlie), 11*price);
        assertEq(Token.price(), price);
        vm.stopPrank();

        vm.startPrank(bob);
        Token.transfer(alice, price);
        assertEq(Token.balanceOf(bob), 10*price);
        assertEq(Token.balanceOf(alice), 9*price);
        assertEq(Token.price(), price);
        Token.transfer(charlie, price);
        assertEq(Token.balanceOf(bob), 9*price);
        assertEq(Token.balanceOf(charlie), 12*price);
        assertEq(Token.price(), price);
        vm.stopPrank();

        vm.startPrank(charlie);
        Token.transfer(alice, price);
        assertEq(Token.balanceOf(charlie), 11*price);
        assertEq(Token.balanceOf(alice), 10*price);
        assertEq(Token.price(), price);
        Token.transfer(bob, price);
        assertEq(Token.balanceOf(charlie), 10*price);
        assertEq(Token.balanceOf(bob), 10*price);
        assertEq(Token.price(), price);
        vm.stopPrank();
    }

    function test_transfer_insufficientBalance_reverts() external {
        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        Token.transfer(bob, 10);
    }

    function test_claim_succeeds() external {
        vm.prank(alice);
        Token.configure(YieldMode.CLAIMABLE);

        uint256 price = Token.price();
        Token.mint(alice, price);

        assertEq(Token.balanceOf(alice), price);

        addYield(10*Token.count());

        assertEq(Token.getClaimableAmount(alice), 10);
        vm.prank(alice);
        Token.claim(bob, 5);
        assertEq(Token.balanceOf(alice), price);
        assertEq(Token.balanceOf(bob), 5);
        assertEq(Token.getClaimableAmount(alice), 5);
    }

    function test_claim_toSelf_succeeds() external {
        vm.prank(alice);
        Token.configure(YieldMode.CLAIMABLE);

        uint256 price = Token.price();
        Token.mint(alice, price);

        addYield(2*Token.count());
        assertEq(Token.balanceOf(alice), price);
        assertEq(Token.getClaimableAmount(alice), 2);

        vm.prank(alice);
        Token.claim(alice, 1);
        assertEq(Token.balanceOf(alice), price+1);
        assertEq(Token.getClaimableAmount(alice), 1);
    }

    function test_claim_insufficientFunds_reverts() external {
        vm.prank(alice);
        Token.configure(YieldMode.CLAIMABLE);

        uint256 price = Token.price();
        Token.mint(alice, price);

        assertEq(Token.balanceOf(alice), price);

        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        Token.claim(bob, 1);
    }

    function test_claim_notClaimable_reverts() external {
        vm.prank(alice);
        Token.configure(YieldMode.AUTOMATIC);

        uint256 price = Token.price();
        Token.mint(alice, price);

        vm.expectRevert(NotClaimableAccount.selector);
        vm.prank(alice);
        Token.claim(bob, 1);
    }

    function test_getClaimableBalance_notAffectedByTransfer() external {
        vm.prank(alice);
        Token.configure(YieldMode.CLAIMABLE);

        Token.mint(alice, 100);
        Token.mint(bob, 100);

        addYield(Token.count());

        uint256 claimableAmount = Token.getClaimableAmount(alice);

        vm.prank(bob);
        Token.transfer(alice, 100);

        assertEq(Token.getClaimableAmount(alice), claimableAmount);
    }

    function test_approval_succeeds() external {
        vm.prank(alice);
        Token.approve(bob, 100);

        assertEq(Token.allowance(alice, bob), 100);
    }

    function test_transferFrom_withApproval_succeeds() external {
        uint256 price = Token.price();
        Token.mint(alice, price);

        vm.prank(alice);
        Token.approve(bob, 100);

        vm.prank(bob);
        Token.transferFrom(alice, bob, 100);

        assertEq(Token.balanceOf(alice), price-100);
        assertEq(Token.balanceOf(bob), 100);
    }

    function test_transferFrom_withoutApproval_reverts() external {
        Token.mint(alice, Token.price());

        vm.expectRevert(InsufficientAllowance.selector);
        vm.prank(bob);
        Token.transferFrom(alice, bob, 100);
    }

    function test_transferFrom_selfTransfer() external {
        uint256 price = Token.price();
        Token.mint(alice, price);

        vm.prank(alice);
        Token.transfer(alice, 100);

        assertEq(Token.balanceOf(alice), price);
    }

    function signPermit(address user, uint256 pk, uint256 value, uint256 deadline) internal view returns (uint8, bytes32, bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(
                Token.PERMIT_TYPEHASH(),
                user,
                address(alice),
                value,
                Token.nonces(user),
                deadline
            )
        );

        bytes32 digest = keccak256(
            abi.encodePacked("\x19\x01", Token.DOMAIN_SEPARATOR(), structHash)
        );

        return vm.sign(pk, digest);
    }

    function test_permitApproval_succeeds() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(user, userPk, 100, block.timestamp+1);
        Token.permit(user, alice, 100, block.timestamp+1, v, r, s);

        assertEq(Token.allowance(user, alice), 100);
    }

    function test_permitApproval_pastDeadline_reverts() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(user, userPk, 100, block.timestamp-1);
        vm.expectRevert("ERC20Permit: expired deadline");
        Token.permit(user, alice, 100, block.timestamp-1, v, r, s);
    }

    function test_permitApproval_invalidSignature_reverts() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(alice, userPk, 100, block.timestamp+1);
        vm.expectRevert("ERC20Permit: invalid signature");
        Token.permit(user, alice, 100, block.timestamp+1, v, r, s);
    }

    function test_permitApproval_invalidNonce_reverts() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(user, userPk, 100, block.timestamp+1);

        Token.permit(user, alice, 100, block.timestamp+1, v, r, s);

        vm.expectRevert("ERC20Permit: invalid signature");
        Token.permit(user, alice, 100, block.timestamp+1, v, r, s);
    }

    function test_rebasing() external {
        uint256 price = Token.price();
        Token.mint(alice, price/2);
        Token.mint(bob, price/2);
        Token.mint(charlie, price);

        addYield(Token.count());

        assertEq(Token.balanceOf(alice), price/2);
        assertEq(Token.balanceOf(bob), price/2);
        assertEq(Token.balanceOf(charlie), price + 1);
        assertEq(Token.price(), price + 1);
    }

    function testFuzz_transfer(uint8 aliceConfiguration, uint8 bobConfiguration, uint256 transferAmount) external {
        uint256 aliceBalance = alice.balance;
        vm.assume(transferAmount <= aliceBalance);
        vm.assume(aliceConfiguration <= 2);
        vm.assume(bobConfiguration <= 2);

        Token.mint(alice, aliceBalance);

        vm.prank(alice);
        Token.configure(YieldMode(aliceConfiguration));

        vm.prank(bob);
        Token.configure(YieldMode(bobConfiguration));

        vm.prank(alice);
        Token.transfer(bob, transferAmount);

        assertEq(Token.balanceOf(alice), aliceBalance - transferAmount);
        assertEq(Token.balanceOf(bob), transferAmount);
    }

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
        Token.configure(YieldMode.AUTOMATIC);
        vm.prank(bob);
        Token.configure(YieldMode.VOID);
        vm.prank(charlie);
        Token.configure(YieldMode.CLAIMABLE);

        Token.mint(alice, aliceDeposit);
        Token.mint(bob, bobDeposit);
        Token.mint(charlie, charlieDeposit);

        uint256 price = Token.price();
        uint256 aliceYield = (yield * (aliceDeposit - aliceDeposit % price)) / (Token.totalSupply() - bobDeposit) + aliceDeposit;
        uint256 charlieYield = (yield * (charlieDeposit - charlieDeposit % price)) / (Token.totalSupply() - bobDeposit);

        assertEq(Token.balanceOf(alice), aliceYield);
        assertEq(Token.balanceOf(bob), bobDeposit);
        assertEq(Token.balanceOf(charlie), charlieDeposit);
        assertEq(Token.getClaimableAmount(charlie), charlieYield);
    }
}
