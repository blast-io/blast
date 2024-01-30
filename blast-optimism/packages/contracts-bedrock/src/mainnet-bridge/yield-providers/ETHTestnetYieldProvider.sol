// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";
import { TestnetYieldProvider } from "src/mainnet-bridge/yield-providers/TestnetYieldProvider.sol";

contract ETHTestnetYieldProvider is TestnetYieldProvider {
    constructor(
        YieldManager _yieldManager,
        address _owner,
        address _token
    ) TestnetYieldProvider(_yieldManager, _owner, _token) {}

    /// @inheritdoc YieldProvider
    function stake(uint256 amount) external override onlyDelegateCall {
        (bool success,) = THIS.call{value: amount}("");
        require(success);
    }

    function stakedBalance() public view override returns (uint256) {
        return uint256(int256(stakedPrincipal) + yield());
    }

    function sendAsset(uint256 amount) external override onlyYieldManager {
        (bool success,) = address(YIELD_MANAGER).call{value: amount}("");
        require(success);
    }

    function recordYield(int256 amount) external payable onlyOwner {
        if (amount > 0) {
            require(msg.value == uint256(amount));
        } else {
            (bool success,) = owner().call{value: uint256(-1 * amount)}("");
            require(success);
        }
        _reportedYield += amount;
    }

    /// @inheritdoc YieldProvider
    function payInsurancePremium(uint256 amount) external override onlyDelegateCall {
        address(YIELD_MANAGER.insurance()).call{value:amount}("");

        emit InsurancePremiumPaid(id(), amount);
    }
}
