// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";

/// @title TestnetYieldProvider
/// @notice Provider for simulating a yield source on testnet.
abstract contract TestnetYieldProvider is YieldProvider, Ownable {
    IERC20 immutable TOKEN;
    address immutable THIS;

    int256 internal _reportedYield;

    /// @param _yieldManager Address of the yield manager for the underlying
    ///        yield asset of this provider.
    constructor(YieldManager _yieldManager, address _owner, address _token) YieldProvider(_yieldManager) {
        _transferOwnership(_owner);
        TOKEN = IERC20(_token);
        THIS = address(this);
    }

    /// @inheritdoc YieldProvider
    function initialize() external override onlyDelegateCall {}

    /// @inheritdoc YieldProvider
    function name() public pure override returns (string memory) {
        return "TestnetYieldProvider";
    }

    /// @inheritdoc YieldProvider
    function isStakingEnabled(address token) public view override returns (bool) {
        return token == address(TOKEN);
    }

    /// @inheritdoc YieldProvider
    function stakedBalance() public view virtual override returns (uint256) {
        return TOKEN.balanceOf(address(YIELD_MANAGER));
    }

    /// @inheritdoc YieldProvider
    function yield() public view override returns (int256) {
        return _reportedYield;
    }

    /// @inheritdoc YieldProvider
    function supportsInsurancePayment() public pure override returns (bool) {
        return true;
    }

    /// @inheritdoc YieldProvider
    function unstake(uint256 amount) external override onlyDelegateCall returns (uint256, uint256) {
        TestnetYieldProvider(THIS).sendAsset(amount);
        return (0, 0);
    }

    function sendAsset(uint256 amount) external virtual;

    function _afterCommitYield() internal override {
        _reportedYield = 0;
    }
}
