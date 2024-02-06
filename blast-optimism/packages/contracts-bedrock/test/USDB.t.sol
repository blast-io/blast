// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Bridge_Initializer } from "test/CommonTest.t.sol";
import { ILegacyMintableERC20, IOptimismMintableERC20 } from "src/universal/IOptimismMintableERC20.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import { YieldMode } from "src/L2/ERC20Rebasing.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";

contract USDB_Test is Bridge_Initializer {
    event Mint(address indexed account, uint256 amount);
    event Burn(address indexed account, uint256 amount);
    event NewPrice(uint256 price);

    error InsufficientBalance();
    error InsufficientAllowance();
    error CallerIsNotBridge();
    error NotClaimableAccount();

    function test_configure_succeeds() external {
        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.prank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 100);

        // defaults to Automatic
        assertTrue(Usdb.getConfiguration(alice) == YieldMode.AUTOMATIC);
        Usdb.balanceOf(alice);

        vm.prank(alice);
        Usdb.configure(YieldMode.CLAIMABLE);
        assertEq(Usdb.balanceOf(alice), 100);
        assertTrue(Usdb.getConfiguration(alice) == YieldMode.CLAIMABLE);

        vm.prank(alice);
        Usdb.configure(YieldMode.AUTOMATIC);
        assertEq(Usdb.balanceOf(alice), 100);
        assertTrue(Usdb.getConfiguration(alice) == YieldMode.AUTOMATIC);

        vm.prank(alice);
        Usdb.configure(YieldMode.VOID);
        assertEq(Usdb.balanceOf(alice), 100);
        assertTrue(Usdb.getConfiguration(alice) == YieldMode.VOID);
    }

    function test_transfer_succeeds() external {
        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 100);
        Usdb.mint(bob, 100);
        Usdb.mint(charlie, 100);
        vm.stopPrank();

        vm.prank(bob);
        Usdb.configure(YieldMode.VOID);

        vm.prank(charlie);
        Usdb.configure(YieldMode.CLAIMABLE);

        vm.startPrank(alice);
        Usdb.transfer(bob, 10);
        assertEq(Usdb.balanceOf(alice), 90);
        assertEq(Usdb.balanceOf(bob), 110);
        Usdb.transfer(charlie, 10);
        assertEq(Usdb.balanceOf(alice), 80);
        assertEq(Usdb.balanceOf(charlie), 110);
        vm.stopPrank();

        vm.startPrank(bob);
        Usdb.transfer(alice, 10);
        assertEq(Usdb.balanceOf(bob), 100);
        assertEq(Usdb.balanceOf(alice), 90);
        Usdb.transfer(charlie, 10);
        assertEq(Usdb.balanceOf(bob), 90);
        assertEq(Usdb.balanceOf(charlie), 120);
        vm.stopPrank();

        vm.startPrank(charlie);
        Usdb.transfer(alice, 10);
        assertEq(Usdb.balanceOf(charlie), 110);
        assertEq(Usdb.balanceOf(alice), 100);
        Usdb.transfer(bob, 10);
        assertEq(Usdb.balanceOf(charlie), 100);
        assertEq(Usdb.balanceOf(bob), 100);
        vm.stopPrank();
    }

    function test_transfer_insufficientBalance_reverts() external {
        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        Usdb.transfer(bob, 10);
    }

    function test_claim_succeeds() external {
        vm.prank(alice);
        Usdb.configure(YieldMode.CLAIMABLE);

        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 1e8);
        vm.stopPrank();

        assertEq(Usdb.balanceOf(alice), 1e8);

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(l1BlastBridge));
        vm.startPrank(caller);
        Usdb.addValue(Usdb.count() * 100);
        vm.stopPrank();

        Usdb.balanceOf(alice);
        assertEq(Usdb.getClaimableAmount(alice), 100);
        vm.prank(alice);
        Usdb.claim(bob, 50);
        assertEq(Usdb.balanceOf(alice), 1e8);
        assertEq(Usdb.balanceOf(bob), 50);
        assertEq(Usdb.getClaimableAmount(alice), 50);
    }

    function test_claim_insufficientFunds_reverts() external {
        vm.prank(alice);
        Usdb.configure(YieldMode.CLAIMABLE);

        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 100);
        vm.stopPrank();

        assertEq(Usdb.balanceOf(alice), 100);

        vm.expectRevert(InsufficientBalance.selector);
        vm.prank(alice);
        Usdb.claim(bob, 1);
    }

    function test_claim_notClaimable_reverts() external {
        vm.prank(alice);
        Usdb.configure(YieldMode.AUTOMATIC);

        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 100);
        vm.stopPrank();

        vm.expectRevert(NotClaimableAccount.selector);
        vm.prank(alice);
        Usdb.claim(bob, 1);
    }

    function test_approval_succeeds() external {
        vm.prank(alice);
        Usdb.approve(bob, 100);

        assertEq(Usdb.allowance(alice, bob), 100);
    }

    function test_transferFrom_withApproval_succeeds() external {
        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 100);
        vm.stopPrank();

        vm.prank(alice);
        Usdb.approve(bob, 100);

        vm.prank(bob);
        Usdb.transferFrom(alice, bob, 100);

        assertEq(Usdb.balanceOf(alice), 0);
        assertEq(Usdb.balanceOf(bob), 100);
    }

    function test_transferFrom_withoutApproval_reverts() external {
        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 100);
        vm.stopPrank();

        vm.expectRevert(InsufficientAllowance.selector);
        vm.prank(bob);
        Usdb.transferFrom(alice, bob, 100);
    }

    function signPermit(address user, uint256 pk, uint256 value, uint256 deadline) internal view returns (uint8, bytes32, bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(
                Usdb.PERMIT_TYPEHASH(),
                user,
                address(alice),
                value,
                Usdb.nonces(user),
                deadline
            )
        );

        bytes32 digest = keccak256(
            abi.encodePacked("\x19\x01", Usdb.DOMAIN_SEPARATOR(), structHash)
        );

        return vm.sign(pk, digest);
    }

    function test_permitApproval_succeeds() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(user, userPk, 100, block.timestamp+1);
        Usdb.permit(user, alice, 100, block.timestamp+1, v, r, s);

        assertEq(Usdb.allowance(user, alice), 100);
    }

    function test_permitApproval_pastDeadline_reverts() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(user, userPk, 100, block.timestamp-1);
        vm.expectRevert("ERC20Permit: expired deadline");
        Usdb.permit(user, alice, 100, block.timestamp-1, v, r, s);
    }

    function test_permitApproval_invalidSignature_reverts() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(alice, userPk, 100, block.timestamp+1);
        vm.expectRevert("ERC20Permit: invalid signature");
        Usdb.permit(user, alice, 100, block.timestamp+1, v, r, s);
    }

    function test_permitApproval_invalidNonce_reverts() external {
        uint256 userPk = 0x01;
        address user = vm.addr(userPk);
        (uint8 v, bytes32 r, bytes32 s) = signPermit(user, userPk, 100, block.timestamp+1);

        Usdb.permit(user, alice, 100, block.timestamp+1, v, r, s);

        vm.expectRevert("ERC20Permit: invalid signature");
        Usdb.permit(user, alice, 100, block.timestamp+1, v, r, s);
    }

    function test_addValue_succeeds() external {
        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 1e8);
        vm.stopPrank();

        uint256 shares = Usdb.count();
        uint256 price = Usdb.price();

        vm.expectEmit(true, true, true, true);
        emit NewPrice(price + 1);

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(l1BlastBridge));
        vm.prank(caller);
        Usdb.addValue(shares); // add 1 token per share

        assertEq(Usdb.balanceOf(alice), 1e8 + 1);
        assertEq(Usdb.price(), price + 1);
    }

    function test_addValue_pending_succeeds() external {
        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 10e8);
        vm.stopPrank();

        uint256 shares = Usdb.count();
        uint256 price = Usdb.price();

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(l1BlastBridge));
        vm.prank(caller);
        Usdb.addValue(shares - 1);

        assertEq(Usdb.balanceOf(alice), 10e8);
        assertEq(Usdb.price(), price);
        assertEq(Usdb.pending(), shares - 1);

        vm.prank(caller);
        Usdb.addValue(1);

        assertEq(Usdb.balanceOf(alice), 10e8+10);
        assertEq(Usdb.price(), price + 1);
        assertEq(Usdb.pending(), 0);
    }

    function test_remoteToken_succeeds() external {
        assertEq(Usdb.REMOTE_TOKEN(), address(DAI));
    }

    function test_bridge_succeeds() external {
        assertEq(Usdb.BRIDGE(), address(l2BlastBridge));
    }

    function test_mint_succeeds() external {
        vm.expectEmit(true, true, true, true);
        emit Mint(alice, 100);

        vm.prank(address(l2BlastBridge));
        Usdb.mint(alice, 100);

        assertEq(Usdb.balanceOf(alice), 100);
    }

    function test_mint_notBridge_reverts() external {
        // NOT the bridge
        vm.expectRevert(CallerIsNotBridge.selector);
        vm.prank(address(alice));
        Usdb.mint(alice, 100);
    }

    function test_burn_succeeds() external {
        vm.prank(address(l2BlastBridge));
        Usdb.mint(alice, 100);

        vm.expectEmit(true, true, true, true);
        emit Burn(alice, 100);

        vm.prank(address(l2BlastBridge));
        Usdb.burn(alice, 100);

        assertEq(Usdb.balanceOf(alice), 0);
    }

    function test_burn_notBridge_reverts() external {
        // NOT the bridge
        vm.expectRevert(CallerIsNotBridge.selector);
        vm.prank(address(alice));
        Usdb.burn(alice, 100);
    }

    function test_getClaimableBalance_notAffectedByTransfer() external {
        vm.prank(alice);
        Usdb.configure(YieldMode.CLAIMABLE);

        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.bridge()))
        );
        vm.startPrank(address(Usdb.bridge()));
        Usdb.mint(alice, 100);
        Usdb.mint(bob, 100);
        vm.stopPrank();

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(l1BlastBridge));
        vm.startPrank(caller);
        Usdb.addValue(Usdb.count());
        vm.stopPrank();

        uint256 claimableAmount = Usdb.getClaimableAmount(alice);

        vm.prank(bob);
        Usdb.transfer(alice, 100);

        assertEq(Usdb.getClaimableAmount(alice), claimableAmount);
    }

    function test_erc165_supportsInterface_succeeds() external {
        // The assertEq calls in this test are comparing the manual calculation of the iface,
        // with what is returned by the solidity's type().interfaceId, just to be safe.
        bytes4 iface1 = bytes4(keccak256("supportsInterface(bytes4)"));
        assertEq(iface1, type(IERC165).interfaceId);
        assert(Usdb.supportsInterface(iface1));

        bytes4 iface2 =
            Usdb.remoteToken.selector ^ Usdb.bridge.selector ^ Usdb.mint.selector ^ Usdb.burn.selector;
        assertEq(iface2, type(IOptimismMintableERC20).interfaceId);
        assert(Usdb.supportsInterface(iface2));
    }

    function testFuzz_transfer(uint8 aliceConfiguration, uint8 bobConfiguration, uint256 transferAmount) external {
        vm.assume(transferAmount <= 100 ether);
        vm.assume(aliceConfiguration <= 2);
        vm.assume(bobConfiguration <= 2);

        vm.prank(address(l2BlastBridge));
        Usdb.mint(alice, 100 ether);

        vm.prank(alice);
        Usdb.configure(YieldMode(aliceConfiguration));

        vm.prank(bob);
        Usdb.configure(YieldMode(bobConfiguration));

        vm.prank(alice);
        Usdb.transfer(bob, transferAmount);

        assertEq(Usdb.balanceOf(alice), 100 ether - transferAmount);
        assertEq(Usdb.balanceOf(bob), transferAmount);
    }
}
