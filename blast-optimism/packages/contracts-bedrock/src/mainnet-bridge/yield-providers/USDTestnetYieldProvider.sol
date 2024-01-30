// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";
import { TestnetYieldProvider } from "src/mainnet-bridge/yield-providers/TestnetYieldProvider.sol";

contract USDTestnetYieldProvider is TestnetYieldProvider {
    constructor(
        YieldManager _yieldManager,
        address _owner,
        address _token
    ) TestnetYieldProvider(_yieldManager, _owner, _token) {}

    /// @inheritdoc YieldProvider
    function stake(uint256 amount) external override onlyDelegateCall {
        TOKEN.transfer(THIS, amount);
    }

    function sendAsset(uint256 amount) external override onlyYieldManager {
        TOKEN.transfer(address(YIELD_MANAGER), amount);
    }

    function recordYield(int256 amount) external onlyOwner {
        if (amount > 0) {
            TOKEN.transferFrom(owner(), THIS, uint256(amount));
        } else {
            TOKEN.transfer(owner(), uint256(-1 * amount));
        }
        _reportedYield += amount;
    }

    /// @inheritdoc YieldProvider
    function payInsurancePremium(uint256 amount) external override onlyDelegateCall {
        TOKEN.transfer(YIELD_MANAGER.insurance(), amount);

        emit InsurancePremiumPaid(id(), amount);
    }
}
