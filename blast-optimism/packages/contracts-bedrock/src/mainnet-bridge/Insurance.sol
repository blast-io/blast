// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { Semver } from "src/mainnet-bridge/YieldManager.sol";

/// @custom:proxied
/// @title Insurace
/// @notice Holds the yield insurance funds and allows yield managers to
///         withdraw to cover losses.
contract Insurance is Initializable, Semver {
    address public admin;
    YieldManager immutable YIELD_MANAGER;

    error OnlyAdmin();
    error OnlyAdminOrYieldManager();
    error InsufficientBalance();

    modifier onlyAdmin() {
        if (msg.sender != admin) {
            revert OnlyAdmin();
        }
        _;
    }

    modifier onlyAdminOrYieldManager() {
        if (msg.sender != admin && msg.sender != address(YIELD_MANAGER)) {
            revert OnlyAdminOrYieldManager();
        }
        _;
    }

    constructor(YieldManager _yieldManager) Semver(1, 0, 0) {
        YIELD_MANAGER = _yieldManager;
        initialize();
    }

    function initialize() public initializer {}

    function setAdmin(address _admin) external onlyAdmin {
        admin = _admin;
    }

    function coverLoss(address token, uint256 amount) external onlyAdminOrYieldManager {
        if (IERC20(token).balanceOf(address(this)) < amount) {
            revert InsufficientBalance();
        }

        IERC20(token).transfer(address(YIELD_MANAGER), amount);
    }
}
