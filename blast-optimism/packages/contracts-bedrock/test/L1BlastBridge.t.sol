// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { stdStorage, StdStorage } from "forge-std/Test.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Bridge_Initializer, LidoYieldProvider_Initializer } from "test/CommonTest.t.sol";

// Libraries
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Constants } from "src/libraries/Constants.sol";

// Target contract dependencies
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { USDB } from "src/L2/USDB.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { ILido } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";

// Target contract
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import "forge-std/console.sol";

interface IUSDC {
    function mint(address, uint256) external;
    function masterMinter() external view returns (address);
    function configureMinter(address, uint256) external returns (bool);
}

interface IUSDT {
    function owner() external view returns (address);
    function issue(uint256) external;
}

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

    function test_receive_notEoa_reverts() external {
        vm.etch(alice, address(L1Token).code);
        vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice);
        address(l1BlastBridge).call{ value: 1 }("");
    }
}

contract PreBridgeETH is Bridge_Initializer {
    /// @dev Asserts the expected calls and events for bridging ETH depending
    ///      on whether the bridge call is legacy or not.
    function _preBridgeETH(bool isLegacy) internal {
        assertEq(address(ethYieldManager).balance, 0);
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        bytes memory message =
            abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, alice, 500, hex"dead");

        if (isLegacy) {
            vm.expectCall(
                address(l1BlastBridge), 500, abi.encodeWithSelector(l1BlastBridge.bridgeETH.selector, 50000, hex"dead")
            );
        } else {
            vm.expectCall(address(l1BlastBridge), 500, abi.encodeWithSelector(l1BlastBridge.bridgeETH.selector, 50000, hex"dead"));
        }
        vm.expectCall(
            address(L1Messenger),
            500,
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2BlastBridge), message, 50000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l1BlastBridge), address(l2BlastBridge), 500, 50000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 50000);
        vm.expectCall(
            address(op),
            500,
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 500, baseGas, false, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(500), uint256(500), baseGas, false, innerMessage);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ETHBridgeInitiated(alice, alice, 500, hex"dead");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(l2BlastBridge), address(l1BlastBridge), message, nonce, 50000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 500);

        vm.prank(alice, alice);
    }
}

contract L1BlastBridge_BridgeETH_Test is PreBridgeETH {
    /// @dev Tests that bridging ETH succeeds.
    ///      Emits ETHDepositInitiated and ETHBridgeInitiated events.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call bridgeETH.
    ///      ETH ends up in the optimismPortal.
    function test_bridgeETH_succeeds() external {
        _preBridgeETH({ isLegacy: false });
        l1BlastBridge.bridgeETH{ value: 500 }(50000, hex"dead");
        assertEq(address(ethYieldManager).balance, 500);
    }
}

contract PreBridgeETHTo is Bridge_Initializer {
    /// @dev Asserts the expected calls and events for bridging ETH to a different
    ///      address depending on whether the bridge call is legacy or not.
    function _preBridgeETHTo(bool isLegacy) internal {
        assertEq(address(ethYieldManager).balance, 0);
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        if (isLegacy) {
            vm.expectCall(
                address(l1BlastBridge), 600, abi.encodeWithSelector(l1BlastBridge.bridgeETHTo.selector, bob, 60000, hex"dead")
            );
        } else {
            vm.expectCall(
                address(l1BlastBridge), 600, abi.encodeWithSelector(l1BlastBridge.bridgeETHTo.selector, bob, 60000, hex"dead")
            );
        }

        bytes memory message =
            abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, bob, 600, hex"dead");

        // the L1 bridge should call
        // L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l2BlastBridge), message, 60000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l1BlastBridge), address(l2BlastBridge), 600, 60000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 60000);
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 600, baseGas, false, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(600), uint256(600), baseGas, false, innerMessage);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ETHBridgeInitiated(alice, bob, 600, hex"dead");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(l2BlastBridge), address(l1BlastBridge), message, nonce, 60000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 600);

        // deposit eth to bob
        vm.prank(alice, alice);
    }
}

contract L1BlastBridge_BridgeETHTo_Test is PreBridgeETHTo {
    /// @dev Tests that bridging ETH to a different address succeeds.
    ///      Emits ETHDepositInitiated and ETHBridgeInitiated events.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call bridgeETHTo.
    ///      ETH ends up in the optimismPortal.
    function test_bridgeETHTo_succeeds() external {
        _preBridgeETHTo({ isLegacy: false });
        l1BlastBridge.bridgeETHTo{ value: 600 }(bob, 60000, hex"dead");
        assertEq(address(ethYieldManager).balance, 600);
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
        uint256 amount = 1000;

        // minAmountOutWad
        bytes memory extraData = abi.encodePacked(amount);

        // Deal Alice's ERC20 State
        deal(address(DAI), alice, 100000, true);
        vm.prank(alice);
        DAI.approve(address(l1BlastBridge), type(uint256).max);

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(DAI), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(usdYieldManager), amount)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(Usdb), address(DAI), alice, alice, amount, extraData
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

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(l2BlastBridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(DAI), address(Usdb), alice, alice, amount, extraData);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(DAI), address(Usdb), amount, 10000, extraData);
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
            l2BlastBridge.finalizeBridgeERC20.selector, Predeploys.USDB, address(DAI), alice, bob, 1000, extraData
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l1BlastBridge), address(l2BlastBridge), 0, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        deal(address(DAI), alice, 100000, true);

        vm.prank(alice);
        DAI.approve(address(l1BlastBridge), type(uint256).max);

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(l2BlastBridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(DAI), address(Usdb), alice, bob, 1000, extraData);

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
            address(DAI), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(usdYieldManager), 1000)
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
        deal(address(DAI), address(usdYieldManager), 100, true);

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
        usdYieldManager.finalize(1);

        vm.prank(alice);
        usdYieldManager.claimWithdrawal(1, 1);

        assertEq(DAI.balanceOf(address(l1BlastBridge)), 0);
        assertEq(DAI.balanceOf(address(usdYieldManager)), 0);
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
    function test_depositERC20_stakedAsset_succeeds() external {
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1BlastBridgeAliased = AddressAliasHelper.applyL1ToL2Alias(address(l1BlastBridge));
        uint256 amount = 1000;

        vm.prank(address(multisig));
        l1BlastBridge.setETHYieldToken(address(Lido), true, 18, address(lidoProvider), true);

        vm.prank(alice);
        Lido.submit{value: amount}(address(0));

        vm.prank(alice);
        Lido.approve(address(l1BlastBridge), type(uint256).max);

        bytes memory message = abi.encodeWithSelector(
            l2BlastBridge.finalizeBridgeETHDirect.selector, alice, bob, amount, hex""
        );

        bytes memory opaqueData = abi.encodePacked(amount, amount, uint64(200_000), false, message);

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1BlastBridgeAliased, address(l2BlastBridge), version, opaqueData);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(Lido), address(0), alice, bob, amount, hex"");

        vm.expectCall(
            address(Lido), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(ethYieldManager), amount)
        );

        // The USDYieldManager should call OptimismPortal.depositTransaction

        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(l2BlastBridge), amount, 200_000, false, message
            )
        );
        vm.prank(alice);
        l1BlastBridge.bridgeERC20To(address(Lido), address(0), bob, amount, 200_000, hex"");

        assertEq(lidoProvider.stakedPrincipal(), amount);
    }

    function test_depositERC20_notApproved_reverts() external {
        vm.prank(address(multisig));
        l1BlastBridge.setETHYieldToken(address(Lido), false, 18, address(lidoProvider), true);

        vm.expectRevert("L1BlastBridge: bridge token is not supported");
        vm.prank(alice);
        l1BlastBridge.bridgeERC20To(address(Lido), address(0), bob, 1000, 10000, hex"");
    }
}

contract L1BlastBridge_FinalizeBridgeETH_Test is Bridge_Initializer {
    /// @dev Tests that finalizing bridged ETH succeeds.
    function test_finalizeBridgeETH_succeeds() external {
        address messenger = address(l1BlastBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1BlastBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        l1BlastBridge.finalizeBridgeETH{ value: 100 }(alice, alice, 100, hex"");
    }

    /// @dev Tests that finalizing bridged ETH with a discounted withdrawal succeeds.
    function test_finalizeBridgeETH_discounted_succeeds() external {
        address messenger = address(l1BlastBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1BlastBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ETHBridgeFinalized(alice, alice, 50, hex"");

        l1BlastBridge.finalizeBridgeETH{ value: 50 }(alice, alice, 100, hex"");
    }
}

contract L1BlastBridge_FinalizeBridgeETH_TestFail is Bridge_Initializer {
    /// @dev Tests that finalizing bridged ETH reverts if the destination is the L1 bridge.
    function test_finalizeBridgeETH_sendToSelf_reverts() external {
        address messenger = address(l1BlastBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1BlastBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("L1BlastBridge: cannot send to self");
        l1BlastBridge.finalizeBridgeETH{ value: 100 }(alice, address(l1BlastBridge), 100, hex"");
    }

    /// @dev Tests that finalizing bridged ETH reverts if the destination is the messenger.
    function test_finalizeBridgeETH_sendToMessenger_reverts() external {
        address messenger = address(l1BlastBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l1BlastBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("L1BlastBridge: cannot send to messenger");
        l1BlastBridge.finalizeBridgeETH{ value: 100 }(alice, messenger, 100, hex"");
    }
}
contract L1BlastBridge_BridgeFork is Bridge_Initializer {
    using SafeERC20 for ERC20;
    using stdStorage for StdStorage;

    function setUp() public override {
        vm.createSelectFork(vm.envString("ETH_RPC_URL"));
        super.setUp();
    }

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
    function test_depositERC20_DAI_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));
        uint256 amount = 1000;

        // minAmountOutWad
        bytes memory extraData = abi.encodePacked(amount);

        // Deal Alice's ERC20 State
        deal(address(DAI), alice, 100000, true);
        vm.prank(alice);
        DAI.approve(address(l1BlastBridge), type(uint256).max);

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(DAI), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(usdYieldManager), amount)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(Usdb), address(DAI), alice, alice, amount, extraData
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

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(l2BlastBridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(DAI), address(Usdb), alice, alice, amount, extraData);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(DAI), address(Usdb), amount, 10000, extraData);
    }

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
    function test_depositERC20_USDC_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));
        uint256 amount = 1000;

        // minAmountOutWad
        bytes memory extraData = abi.encodePacked(amount);

        // Deal Alice's ERC20 State
        vm.prank(IUSDC(address(USDC)).masterMinter());
        IUSDC(address(USDC)).configureMinter(address(this), type(uint256).max);
        IUSDC(address(USDC)).mint(alice, 100000);

        vm.prank(alice);
        USDC.approve(address(l1BlastBridge), type(uint256).max);

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(USDC), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(usdYieldManager), amount)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(Usdb), address(DAI), alice, alice, amount*1e12, extraData
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

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(l2BlastBridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(USDC), address(Usdb), alice, alice, amount * 1e12, extraData);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(USDC), address(Usdb), amount, 10000, extraData);
    }

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
    function test_depositERC20_USDT_succeeds() external {
        vm.skip(true);
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));
        uint256 amount = 1000;

        // minAmountOutWad
        // bytes memory extraData = abi.encodePacked(amount * 9 / 10);
        bytes memory extraData = abi.encodePacked(uint256(0));

        // Deal Alice's ERC20 State
        vm.startPrank(IUSDT(address(USDT)).owner());
        IUSDT(address(USDT)).issue(amount);
        USDT.safeTransfer(alice, amount);
        vm.stopPrank();

        vm.startPrank(alice);
        USDT.safeApprove(address(l1BlastBridge), type(uint256).max);
        vm.stopPrank();

        // The l1BlastBridge should transfer alice's tokens to itself
        vm.expectCall(
            address(USDT), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(usdYieldManager), amount)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(Usdb), address(DAI), alice, alice, amount*1e12, extraData
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

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(l2BlastBridge), address(l1BlastBridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(l1BlastBridge), 0);

        vm.expectEmit(true, true, true, true, address(l1BlastBridge));
        emit ERC20BridgeInitiated(address(USDT), address(Usdb), alice, alice, amount, extraData);

        vm.prank(alice);
        l1BlastBridge.bridgeERC20(address(USDT), address(Usdb), amount, 10000, extraData);
    }
}
