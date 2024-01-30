// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { Semver } from "src/universal/Semver.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/// @custom:proxied
/// @title ETHYieldManager
/// @notice Coordinates the accounting, asset management and
///         yield reporting from ETH yield providers.
contract ETHYieldManager is YieldManager, Semver {
    error CallerIsNotPortal();

    constructor() YieldManager(address(0)) Semver(1, 0, 0) {
        initialize(OptimismPortal(payable(address(0))), address(0));
    }

    receive() external payable {}

    /// @notice initializer
    /// @param _portal Address of the OptimismPortal.
    /// @param _owner  Address of the YieldManager owner.
    function initialize(OptimismPortal _portal, address _owner) public initializer {
        __YieldManager_init(_portal, _owner);
    }

    /// @inheritdoc YieldManager
    function tokenBalance() public view override returns (uint256) {
        return address(this).balance;
    }

    /// @notice Wrapper for WithdrawalQueue._requestWithdrawal
    function requestWithdrawal(uint256 amount)
        external
        returns (uint256)
    {
        if (msg.sender != address(portal)) {
            revert CallerIsNotPortal();
        }
        return _requestWithdrawal(address(portal), amount);
    }

    /// @notice Sends the yield report to the Shares contract.
    /// @param data Calldata to send in the message.
    function _reportYield(bytes memory data) internal override {
        portal.depositTransaction(Predeploys.SHARES, 0, REPORT_YIELD_DEFAULT_GAS_LIMIT, false, data);
    }
}
