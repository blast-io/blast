// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Bridge_Initializer } from "test/CommonTest.t.sol";
import { ILegacyMintableERC20, IOptimismMintableERC20 } from "src/universal/IOptimismMintableERC20.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { YieldMode } from "src/L2/Blast.sol";
import { GasMode } from "src/L2/Gas.sol";

contract USDB_Test is Bridge_Initializer {
    event Mint(address indexed account, uint256 amount);
    event Burn(address indexed account, uint256 amount);
    event NewPrice(uint256 price);

    error CallerIsNotBridge();

    function test_blastConfig() external {
        checkBlastConfig(address(Usdb), address(0xdead), YieldMode.VOID, GasMode.VOID);
    }

    function test_addValue_succeeds() external {
        uint256 price = Usdb.price();

        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, price);
        vm.stopPrank();

        uint256 shares = Usdb.count();

        vm.expectEmit(true, true, true, true);
        emit NewPrice(price + 1);

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(usdYieldManager));
        vm.prank(caller);
        Usdb.addValue(shares); // add 1 token per share

        assertEq(Usdb.balanceOf(alice), price + 1);
        assertEq(Usdb.price(), price + 1);
    }

    function test_addValue_pending_succeeds() external {
        vm.mockCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(Usdb.BRIDGE()))
        );
        uint256 price = Usdb.price();
        vm.startPrank(address(Usdb.BRIDGE()));
        Usdb.mint(alice, 10*price);
        vm.stopPrank();

        uint256 shares = Usdb.count();

        address caller = AddressAliasHelper.applyL1ToL2Alias(address(usdYieldManager));
        vm.prank(caller);
        Usdb.addValue(shares - 1);

        assertEq(Usdb.balanceOf(alice), 10*price);
        assertEq(Usdb.price(), price);
        assertEq(Usdb.pending(), shares - 1);

        vm.prank(caller);
        Usdb.addValue(1);

        assertEq(Usdb.balanceOf(alice), 10*price + 10);
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
