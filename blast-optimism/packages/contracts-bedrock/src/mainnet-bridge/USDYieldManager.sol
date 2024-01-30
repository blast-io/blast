// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";
import { Semver } from "src/universal/Semver.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/// @custom:proxied
/// @title USDYieldManager
/// @notice Coordinates the accounting, asset management and
///         yield reporting from USD yield providers.
contract USDYieldManager is YieldManager, Semver {
    /// @param _token Address of withdrawal token. It is assumed that the token
    ///               has 18 decimals.
    constructor(address _token) YieldManager(_token) Semver(1, 0, 0) {}

    /// @notice initializer
    /// @param _portal Address of the OptimismPortal.
    /// @param _owner  Address of the YieldManager owner.
    function initialize(OptimismPortal _portal, address _owner) public initializer {
        __YieldManager_init(_portal, _owner);
        if (TOKEN == address(USDConversions.DAI)) {
            USDConversions._init();
        }
    }

    /// @inheritdoc YieldManager
    function tokenBalance() public view override returns (uint256) {
        return IERC20(TOKEN).balanceOf(address(this));
    }

    /// @notice Wrapper for WithdrawalQueue._requestWithdrawal
    function requestWithdrawal(address recipient, uint256 amount)
        external
        onlyBlastBridge
        returns (uint256)
    {
        return _requestWithdrawal(address(recipient), amount);
    }

    /// @notice Wrapper for USDConversions._convertTo
    function convert(
        address inputTokenAddress,
        uint256 inputAmountWad,
        bytes memory _extraData
    ) external onlyBlastBridge returns (uint256) {
        return USDConversions._convertTo(
            inputTokenAddress,
            TOKEN,
            inputAmountWad,
            _extraData
        );
    }

    /// @notice Sends the yield report to the USDB contract.
    /// @param data Calldata to send in the message.
    function _reportYield(bytes memory data) internal override {
        portal.depositTransaction(Predeploys.USDB, 0, REPORT_YIELD_DEFAULT_GAS_LIMIT, false, data);
    }
}
