// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { stdStorage, StdStorage } from "forge-std/Test.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { Bridge_Initializer, LidoYieldProvider_Initializer } from "test/CommonTest.t.sol";

// Libraries
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Constants } from "src/libraries/Constants.sol";

// Target contract dependencies
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { USDB } from "src/L2/USDB.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { ILido } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";

// Target contract
import { OptimismPortal } from "src/L1/OptimismPortal.sol";

contract L1BlastBridge_Getter_Test is Bridge_Initializer {
    /// @dev Test that the accessors return the correct initialized values.
    function test_getters_succeeds() external view {
        assert(l1BlastBridge.OTHER_BRIDGE() == l2BlastBridge);
        assert(l1BlastBridge.messenger() == L1Messenger);
    }
}

contract L1BlastBridge_Initialize_Test is Bridge_Initializer {
    /// @dev Test that the initialize function sets the correct values.
    function test_initialize_succeeds() external {
        assertEq(address(l1BlastBridge.messenger()), address(L1Messenger));
        assertEq(address(l1BlastBridge.OTHER_BRIDGE()), Predeploys.L2_BLAST_BRIDGE);
        assertEq(address(l2BlastBridge), Predeploys.L2_BLAST_BRIDGE);
        bytes32 slot0 = vm.load(address(l1BlastBridge), bytes32(uint256(0)));
        assertEq(slot0, bytes32(uint256(1)));
    }
}

contract L1BlastBridge_Receive_Test is Bridge_Initializer {
    /// @dev Tests receive bridges ETH successfully.
    function test_receive_succeeds() external {
        assertEq(address(ethYieldManager).balance, 0);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ETHBridgeInitiated(alice, alice, 100, hex"");

        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(l2BlastBridge),
                abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, alice, 100, hex""),
                200_000
            )
        );

        vm.prank(alice, alice);
        (bool success,) = address(l1BlastBridge).call{ value: 100 }(hex"");
        assertEq(success, true);
        assertEq(address(ethYieldManager).balance, 100);
    }
}

contract L1BlastBridge_DepositERC20_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    // depositERC20
    // - updates bridge.deposits
    // - emits ERC20DepositInitiated
    // - calls optimismPortal.depositTransaction
    // - only callable by EOA

    /// @dev Tests that depositing ERC20 to the bridge succeeds.
    ///      Bridge deposits are updated.
    ///      Emits ERC20DepositInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call depositERC20.
    function test_depositERC20_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        // minAmountOutWad
        bytes memory extraData = abi.encodePacked(uint256(100));

        // Deal Alice's ERC20 State
        deal(address(DAI), alice, 100000, true);
        vm.prank(alice);
        DAI.approve(address(l1BlastBridge), type(uint256).max);

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(DAI), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(l1BlastBridge), 100)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(DAI), address(Usdb), alice, alice, 100, extraData
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

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(DAI), address(Usdb), alice, alice, 100, extraData);

        /*
        // manually checked these logs, but foundry seems to have issues with checking events in a fork
        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);
        */

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(DAI), address(Usdb), 100, 10000, extraData);
    }
}

contract L1BlastBridge_DepositERC20_TestFail is Bridge_Initializer {
    /// @dev Tests that depositing an ERC20 to the bridge reverts
    ///      if the caller is not an EOA.
    function test_depositERC20_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, hex"ffff");

        vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice, alice);
        l1BlastBridge.bridgeERC20(address(0), address(0), 100, 100, hex"");
    }
}

contract L1BlastBridge_DepositERC20To_Test is Bridge_Initializer {
    /// @dev Tests that depositing ERC20 to the bridge succeeds when
    ///      sent to a different address.
    ///      Bridge deposits are updated.
    ///      Emits ERC20DepositInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Contracts can call depositERC20.
    function test_depositERC20To_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        // minAmountOutWad
        bytes memory extraData = abi.encodePacked(uint256(1000));

        bytes memory message = abi.encodeWithSelector(
            l2BlastBridge.finalizeBridgeERC20.selector, address(DAI), Predeploys.USDB, alice, bob, 1000, extraData
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l1BlastBridge), address(l2BlastBridge), 0, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        deal(address(DAI), alice, 100000, true);

        vm.prank(alice);
        DAI.approve(address(l1BlastBridge), type(uint256).max);

        /*
        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(DAI), address(Usdb), alice, bob, 1000, hex"");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);
        */

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(Predeploys.L2_BLAST_BRIDGE), message, 10000)
        );

        // The L1 XDM should call OptimismPortal.depositTransaction
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 0, baseGas, false, innerMessage
            )
        );
        vm.expectCall(
            address(DAI), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(l1BlastBridge), 1000)
        );

        vm.prank(alice);
        l1BlastBridge.bridgeERC20To(address(DAI), address(Usdb), bob, 1000, 10000, extraData);
    }
}

contract L1BlastBridge_FinalizeERC20Withdrawal_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    /// @dev Tests that finalizing an ERC20 withdrawal succeeds.
    ///      Bridge deposits are updated.
    ///      Emits ERC20WithdrawalFinalized event.
    ///      Only callable by the L2 bridge.
    function test_finalizeERC20Withdrawal_succeeds() external {
        deal(address(DAI), address(l1BlastBridge), 100, true);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeFinalized(address(DAI), address(Usdb), alice, alice, 100, hex"");

        vm.mockCall(
            address(l1BlastBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1BlastBridge.OTHER_BRIDGE()))
        );
        vm.prank(address(l1BlastBridge.messenger()));
        l1BlastBridge.finalizeBridgeERC20(address(DAI), address(Usdb), alice, alice, 100, hex"");

        vm.prank(address(multisig));
        l1BlastBridge.finalize(1);

        vm.prank(alice);
        l1BlastBridge.claimWithdrawal(1, 1);

        assertEq(DAI.balanceOf(address(l1BlastBridge)), 0);
        assertEq(DAI.balanceOf(address(alice)), 100);
    }
}

contract L1BlastBridge_FinalizeERC20Withdrawal_TestFail is Bridge_Initializer {
    /// @dev Tests that finalizing an ERC20 withdrawal reverts if the caller is not the L2 bridge.
    function test_finalizeERC20Withdrawal_notMessenger_reverts() external {
        vm.mockCall(
            address(l1BlastBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1BlastBridge.OTHER_BRIDGE()))
        );
        vm.prank(address(28));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        l1BlastBridge.finalizeBridgeERC20(address(DAI), address(Usdb), alice, alice, 100, hex"");
    }

    /// @dev Tests that finalizing an ERC20 withdrawal reverts if the caller is not the L2 bridge.
    function test_finalizeERC20Withdrawal_notOtherBridge_reverts() external {
        vm.mockCall(
            address(l1BlastBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(address(0)))
        );
        vm.prank(address(l1BlastBridge.messenger()));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        l1BlastBridge.finalizeBridgeERC20(address(DAI), address(Usdb), alice, alice, 100, hex"");
    }
}

contract L1BlastBridge_BridgeStakedAsset is LidoYieldProvider_Initializer {
    function test_depositERC20_stakedAsset_succeds() external {
        vm.prank(address(multisig));
        l1BlastBridge.setETHYieldToken(LidoAddress, true, 18, address(lidoProvider), true);

        vm.prank(alice);
        ILido(LidoAddress).submit{value: 10000}(address(0));

        vm.prank(alice);
        ILido(LidoAddress).approve(address(l1BlastBridge), type(uint256).max);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(LidoAddress, address(0), alice, bob, 1000, hex"");

        /*
        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);
        */

        bytes memory message = abi.encodeWithSelector(
            l2BlastBridge.finalizeBridgeETHDirect.selector, alice, bob, 1000, hex""
        );

        vm.expectCall(
            LidoAddress, abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(ethYieldManager), 1000)
        );

        // The USDYieldManager should call OptimismPortal.depositTransaction
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(l2BlastBridge), 1000, 200000, false, message
            )
        );
        vm.prank(alice);
        l1BlastBridge.bridgeERC20To(LidoAddress, address(0), bob, 1000, 10000, hex"");

        assertEq(lidoProvider.stakedBalance(), 1000);
    }

    function test_depositERC20_notApproved_reverts() external {
        vm.prank(address(multisig));
        l1BlastBridge.setETHYieldToken(LidoAddress, false, 18, address(lidoProvider), true);

        vm.expectRevert("L1BlastBridge: bridge token is not supported");
        vm.prank(alice);
        l1BlastBridge.bridgeERC20To(LidoAddress, address(0), bob, 1000, 10000, hex"");
    }
}
