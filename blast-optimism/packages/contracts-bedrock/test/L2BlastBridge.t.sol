// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
// Target contract is imported by the `Bridge_Initializer`
import { Bridge_Initializer } from "test/CommonTest.t.sol";
import { stdStorage, StdStorage } from "forge-std/Test.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// Libraries
import { Hashing } from "src/libraries/Hashing.sol";
import { Types } from "src/libraries/Types.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";

// Target contract dependencies
import { Predeploys } from "src/libraries/Predeploys.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { OptimismMintableERC20 } from "src/universal/OptimismMintableERC20.sol";
import { YieldMode } from "src/L2/Blast.sol";
import { GasMode } from "src/L2/Gas.sol";

contract L2BlastBridge_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    function test_blastConfig() external {
        assertEq(blast.governorMap(address(l2BlastBridge)), address(0xdead));
        assertTrue(blast.readYieldConfiguration(address(l2BlastBridge)) == uint8(YieldMode.VOID));
        (,,, GasMode gasMode) = blast.readGasParams(address(l2BlastBridge));
        assertTrue(gasMode == GasMode.VOID);
    }

    /// @dev Tests that the bridge is initialized correctly.
    function test_initialize_succeeds() external {
        assertEq(address(l2BlastBridge.messenger()), address(L2Messenger));
        assertEq(address(l1BlastBridge.OTHER_BRIDGE()), address(l2BlastBridge));
        assertEq(address(l2BlastBridge.OTHER_BRIDGE()), address(l1BlastBridge));
    }

    /// @dev Tests that the bridge receives ETH and successfully initiates a withdrawal.
    function test_receive_succeeds() external {
        assertEq(address(messagePasser).balance, 0);
        uint256 nonce = L2Messenger.messageNonce();

        bytes memory message =
            abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, alice, 100, hex"");
        uint64 baseGas = L2Messenger.baseGas(message, 200_000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(l2BlastBridge),
            address(l1BlastBridge),
            100,
            200_000,
            message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(L2Messenger),
                target: address(L1Messenger),
                value: 100,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeInitiated(alice, alice, 100, hex"");

        // L2ToL1MessagePasser will emit a MessagePassed event
        vm.expectEmit(true, true, true, true, address(messagePasser));
        emit MessagePassed(
            nonce, address(L2Messenger), address(L1Messenger), 100, baseGas, withdrawalData, withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessage(address(l1BlastBridge), address(l2BlastBridge), message, nonce, 200_000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessageExtension1(address(l2BlastBridge), 100);

        vm.expectCall(
            address(L2Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(l1BlastBridge),
                message,
                200_000 // StandardBridge's RECEIVE_DEFAULT_GAS_LIMIT
            )
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector, address(L1Messenger), baseGas, withdrawalData
            )
        );

        vm.prank(alice, alice);
        (bool success,) = address(l2BlastBridge).call{ value: 100 }(hex"");
        assertEq(success, true);
        assertEq(address(messagePasser).balance, 100);
    }
}

contract PreBridgeERC20 is Bridge_Initializer {
    function mint(address account, uint256 amount) internal {
        vm.prank(address(l2BlastBridge));
        Usdb.mint(account, amount);
    }

    /// @dev Sets up expected calls and emits for a successful ERC20 withdrawal.
    function _preBridgeERC20(address _l2Token) internal {
        // Alice has 100 Usdb
        mint(alice, 100);
        assertEq(ERC20(_l2Token).balanceOf(alice), 100);
        uint256 nonce = L2Messenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(DAI), _l2Token, alice, alice, 100, hex""
        );
        uint64 baseGas = L2Messenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l2BlastBridge), address(l1BlastBridge), 0, 1000, message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(L2Messenger),
                target: address(L1Messenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectCall(
            address(l2BlastBridge),
            abi.encodeWithSelector(l2BlastBridge.bridgeERC20.selector, _l2Token, address(DAI), 100, 1000, hex"")
        );

        vm.expectCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l1BlastBridge), message, 1000)
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector, address(L1Messenger), baseGas, withdrawalData
            )
        );

        // The l2BlastBridge should burn the tokens
        vm.expectCall(_l2Token, abi.encodeWithSelector(OptimismMintableERC20.burn.selector, alice, 100));

        vm.expectEmit(true, true, true, true);
        emit ERC20BridgeInitiated(_l2Token, address(DAI), alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true);
        emit MessagePassed(
            nonce, address(L2Messenger), address(L1Messenger), 0, baseGas, withdrawalData, withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true);
        emit SentMessage(address(l1BlastBridge), address(l2BlastBridge), message, nonce, 1000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true);
        emit SentMessageExtension1(address(l2BlastBridge), 0);

        vm.prank(alice, alice);
    }
}

contract L2BlastBridge_BridgeERC20_Test is PreBridgeERC20 {
    // BridgeERC20
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeERC20_succeeds() external {
        _preBridgeERC20({ _l2Token: address(Usdb) });
        l2BlastBridge.bridgeERC20(address(Usdb), address(DAI), 100, 1000, hex"");

        assertEq(Usdb.balanceOf(alice), 0);
    }
}

contract PreBridgeERC20To is Bridge_Initializer {
    function mint(address account, uint256 amount) internal {
        vm.prank(address(l2BlastBridge));
        Usdb.mint(account, amount);
    }

    // withdrawTo and BridgeERC20To should behave the same when transferring ERC20 tokens
    // so they should share the same setup and expectEmit calls
    function _preBridgeERC20To(address _l2Token) internal {
        mint(alice, 100);
        assertEq(Usdb.balanceOf(alice), 100);
        uint256 nonce = L2Messenger.messageNonce();
        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(DAI), _l2Token, alice, bob, 100, hex""
        );
        uint64 baseGas = L2Messenger.baseGas(message, 1000);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(l2BlastBridge), address(l1BlastBridge), 0, 1000, message
        );
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(L2Messenger),
                target: address(L1Messenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        vm.expectEmit(true, true, true, true, address(l2BlastBridge));
        emit ERC20BridgeInitiated(_l2Token, address(DAI), alice, bob, 100, hex"");

        vm.expectEmit(true, true, true, true, address(messagePasser));
        emit MessagePassed(
            nonce, address(L2Messenger), address(L1Messenger), 0, baseGas, withdrawalData, withdrawalHash
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessage(address(l1BlastBridge), address(l2BlastBridge), message, nonce, 1000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessageExtension1(address(l2BlastBridge), 0);

        vm.expectCall(
            address(l2BlastBridge),
            abi.encodeWithSelector(
                l2BlastBridge.bridgeERC20To.selector, _l2Token, address(DAI), bob, 100, 1000, hex""
            )
        );

        vm.expectCall(
            address(L2Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(l1BlastBridge), message, 1000)
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector, address(L1Messenger), baseGas, withdrawalData
            )
        );

        // The l2BlastBridge should burn the tokens
        vm.expectCall(address(Usdb), abi.encodeWithSelector(OptimismMintableERC20.burn.selector, alice, 100));

        vm.prank(alice, alice);
    }
}

contract L2BlastBridge_BridgeERC20To_Test is PreBridgeERC20To {
    /// @dev Tests that `bridgeERC20To` burns the tokens, emits `WithdrawalInitiated`,
    ///      and initiates a withdrawal with `Withdrawer.initiateWithdrawal`.
    function test_bridgeERC20To_succeeds() external {
        _preBridgeERC20To({ _l2Token: address(Usdb) });
        l2BlastBridge.bridgeERC20To(address(Usdb), address(DAI), bob, 100, 1000, hex"");
        assertEq(Usdb.balanceOf(alice), 0);
    }
}

contract L2BlastBridge_Bridge_Test is Bridge_Initializer {
    /// @dev Tests that `finalizeDeposit` reverts if the amounts do not match.
    /*
    function test_finalizeBridgeETH_incorrectValue_reverts() external {
        vm.mockCall(
            address(l2BlastBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2BlastBridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: amount sent does not match amount required");
        l2BlastBridge.finalizeBridgeETH{ value: 50 }(alice, alice, 100, hex"");
    }
    */

    /// @dev Tests that `finalizeDeposit` reverts if the receipient is the other bridge.
    function test_finalizeBridgeETH_sendToSelf_reverts() external {
        vm.mockCall(
            address(l2BlastBridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2BlastBridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: cannot send to self");
        l2BlastBridge.finalizeBridgeETH{ value: 100 }(alice, address(l2BlastBridge), 100, hex"");
    }
}

contract L2BlastBridge_FinalizeBridgeETH_Test is Bridge_Initializer {
    /// @dev Tests that `finalizeBridgeETH` succeeds.
    function test_finalizeBridgeETH_succeeds() external {
        address messenger = address(l2BlastBridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(l2BlastBridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        l2BlastBridge.finalizeBridgeETH{ value: 100 }(alice, alice, 100, hex"");
    }

    function test_finalizeBridgeETHDirect_succeeds() external {
        vm.deal(AddressAliasHelper.applyL1ToL2Alias(address(l2BlastBridge.OTHER_BRIDGE())), 100);
        vm.prank(AddressAliasHelper.applyL1ToL2Alias(address(l2BlastBridge.OTHER_BRIDGE())));

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        l2BlastBridge.finalizeBridgeETHDirect{ value: 100 }(alice, alice, 100, hex"");
    }
}
