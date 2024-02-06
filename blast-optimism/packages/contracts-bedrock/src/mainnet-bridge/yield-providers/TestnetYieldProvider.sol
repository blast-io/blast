// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";

/// @title TestnetYieldProvider
/// @notice Provider for simulating a yield source on testnet.
contract TestnetYieldProvider is YieldProvider, Ownable {
    uint256 internal _reportedYield;

    /// @param _yieldManager Address of the yield manager for the underlying
    ///        yield asset of this provider.
    constructor(YieldManager _yieldManager, address _owner) YieldProvider(_yieldManager) {
        _transferOwnership(_owner);
    }

    /// @inheritdoc YieldProvider
    function initialize() external override onlyDelegateCall {}

    /// @inheritdoc YieldProvider
    function name() public pure override returns (string memory) {
        return "TestnetYieldProvider";
    }

    /// @inheritdoc YieldProvider
    function isStakingEnabled(address token) public view override returns (bool) {
        return false;
    }

    /// @inheritdoc YieldProvider
    function stakedValue() public view override returns (uint256) {
        return address(YIELD_MANAGER).balance;
    }

    /// @inheritdoc YieldProvider
    function yield() public view override returns (int256) {
        return int256(_reportedYield);
    }

    /// @inheritdoc YieldProvider
    function supportsInsurancePayment() public view override returns (bool) {
        return false;
    }

    /// @inheritdoc YieldProvider
    function stake(uint256 amount) external override onlyDelegateCall {}

    /// @inheritdoc YieldProvider
    function unstake(uint256 amount) external override onlyDelegateCall returns (uint256 pending) {}

    function recordYield(uint256 amount) external payable onlyOwner {
        require(msg.value == amount);
        (bool success,) = address(YIELD_MANAGER).call{value: amount}("");
        require(success);
        _reportedYield = amount;
    }

    function _afterCommitYield() internal override {
        _reportedYield = 0;
    }
}
