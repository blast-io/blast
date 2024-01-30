// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { Predeploys } from "src/libraries/Predeploys.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { SafeCall } from "src/libraries/SafeCall.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { Blast, YieldMode, GasMode } from "src/L2/Blast.sol";

/// @custom:proxied
/// @custom:predeploy 0x4300000000000000000000000000000000000005
/// @title L2BlastBridge
/// @notice The L2BlastBridge is responsible for transfering ETH and USDB tokens between L1 and
///         L2. In the case that an ERC20 token is native to L2, it will be escrowed within this
///         contract.
contract L2BlastBridge is StandardBridge, ISemver {
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Constructs the L2BlastBridge contract.
    /// @param _otherBridge Address of the L1BlastBridge.
    constructor(StandardBridge _otherBridge) StandardBridge(_otherBridge) {
        _disableInitializers();
    }

    /// @notice Initializer
    function initialize() public initializer {
        __StandardBridge_init({ _messenger: CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER) });
        Blast(Predeploys.BLAST).configureContract(
            address(this),
            YieldMode.VOID,
            GasMode.VOID,
            address(0xdead) /// don't set a governor
        );
    }

    /// @notice Allows EOAs to bridge ETH by sending directly to the bridge.
    receive() external payable override onlyEOA {
        _initiateBridgeETH(msg.sender, msg.sender, msg.value, RECEIVE_DEFAULT_GAS_LIMIT, hex"");
    }

    /// @notice Modified StandardBridge.finalizeBridgeETH function to allow calls directly from
    ///         the L1BlastBridge without going through a messenger.
    /// @notice See { StandardBridge-finalizeBridgeETH }
    function finalizeBridgeETHDirect(
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _extraData
    )
        public
        payable
    {
        require(AddressAliasHelper.undoL1ToL2Alias(msg.sender) == address(OTHER_BRIDGE), "L2BlastBridge: function can only be called from the other bridge");
        require(msg.value == _amount, "L2BlastBridge: amount sent does not match amount required");
        require(_to != address(this), "L2BlastBridge: cannot send to self");
        require(_to != address(messenger), "L2BlastBridge: cannot send to messenger");

        // Emit the correct events. By default this will be _amount, but child
        // contracts may override this function in order to emit legacy events as well.
        _emitETHBridgeFinalized(_from, _to, _amount, _extraData);

        bool success = SafeCall.call(_to, gasleft(), _amount, hex"");
        require(success, "L2BlastBridge: ETH transfer failed");
    }

    /// @notice Wrapper to only accept USDB withdrawals.
    /// @notice See { StandardBridge-_initiateBridgeERC20 }
    function _initiateBridgeERC20(
        address _localToken,
        address _remoteToken,
        address _from,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes memory _extraData
    )
        internal
        override
    {
        require(_localToken == Predeploys.USDB, "L2BlastBridge: only USDB can be withdrawn from this bridge.");
        require(_isCorrectTokenPair(Predeploys.USDB, _remoteToken), "L2BlastBridge: wrong remote token for USDB.");
        super._initiateBridgeERC20(_localToken, _remoteToken, _from, _to, _amount, _minGasLimit, _extraData);
    }
}
