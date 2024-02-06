// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { Semver } from "src/universal/Semver.sol";

/// @title YieldProvider
/// @notice Base contract for interacting and accounting for a
///         specific yield source.
abstract contract YieldProvider is Semver {
    YieldManager public immutable YIELD_MANAGER;

    uint256 public stakedBalance;
    uint256 public pendingBalance;

    event YieldCommit(bytes32 indexed provider, int256 yield);
    event Staked(bytes32 indexed provider, uint256 amount);
    event Unstaked(bytes32 indexed provider, uint256 amount);
    event Pending(bytes32 indexed provider, uint256 amount);
    event Claimed(bytes32 indexed provider, uint256 amount);
    event InsurancePremiumPaid(bytes32 indexed provider, uint256 amount);
    event InsuranceWithdrawn(bytes32 indexed provider, uint256 amount);

    error InsufficientStakableFunds();
    error CallerIsNotYieldManager();
    error ContextIsNotYieldManager();
    error NotSupported();

    modifier onlyYieldManager() {
        if (msg.sender != address(YIELD_MANAGER)) {
            revert CallerIsNotYieldManager();
        }
        _;
    }

    modifier onlyDelegateCall() {
        if (address(this) != address(YIELD_MANAGER)) {
            revert ContextIsNotYieldManager();
        }
        _;
    }

    /// @param _yieldManager Address of the yield manager for the underlying
    ///        yield asset of this provider.
    constructor(YieldManager _yieldManager) Semver(1, 0, 0) {
        require(address(_yieldManager) != address(this));
        YIELD_MANAGER = _yieldManager;
    }

    /// @notice initialize
    function initialize() external onlyDelegateCall virtual {}

    function name() public pure virtual returns (string memory);

    function id() public view returns (bytes32) {
        return keccak256(abi.encodePacked(name(), version()));
    }

    /// @notice Whether staking is enabled for the given asset.
    function isStakingEnabled(address token) external view virtual returns (bool);

    /// @notice Current balance of the provider's staked funds.
    function stakedValue() public view virtual returns (uint256);

    /// @notice Total value in the provider's yield method/protocol.
    function totalProviderValue() public view returns (uint256) {
        return stakedValue() + pendingBalance;
    }

    /// @notice Current amount of yield gained since the previous commit.
    function yield() public view virtual returns (int256);

    /// @notice Whether the provider supports yield insurance.
    function supportsInsurancePayment() public view virtual returns (bool) {
        return false;
    }

    /// @notice Gets insurance balance available for the provider's assets.
    function insuranceBalance() public view virtual returns (uint256) {
        revert("not supported");
    }

    /// @notice Commit the current amount of yield and checkpoint the accounting
    ///         variables.
    /// @return Amount of yield at this checkpoint.
    function commitYield() external onlyYieldManager returns (int256) {
        _beforeCommitYield();

        int256 _yield = yield();
        stakedBalance = stakedValue();

        _afterCommitYield();

        emit YieldCommit(id(), _yield);
        return _yield;
    }

    /// @notice Stake YieldManager funds using the provider's yield method/protocol.
    ///         Must be called via `delegatecall` from the YieldManager.
    function stake(uint256) external virtual;

    /// @notice Unstake YieldManager funds from the provider's yield method/protocol.
    ///         Must be called via `delegatecall` from the YieldManager.
    function unstake(uint256) external virtual returns (uint256 pending);

    /// @notice Claim funds pending in an unstaking delay from the provider's yield method/protocol.
    ///         Must be called via `delegatecall` from the YieldManager.
    function claim(uint256[] calldata requestIds) external virtual onlyDelegateCall returns (uint256 claimed) {
        revert NotSupported();
    }

    /// @notice Pay insurance premium during a yield report. Must be called via
    ///         `delegatecall` from the YieldManager.
    function payInsurancePremium(uint256 amount) external virtual onlyDelegateCall {
        revert NotSupported();
    }

    /// @notice Withdraw insurance funds to cover yield losses during a yield report.
    ///         Must be called via `delegatecall` from the YieldManager.
    function withdrawFromInsurance(uint256 amount) external virtual onlyDelegateCall {
        revert NotSupported();
    }

    /// @notice Record a deposit to the stake balance of the provider to track the
    ///         principal balance.
    /// @param amount Amount of new staked balance to record.
    function recordStakedDeposit(uint256 amount) external virtual onlyYieldManager {
        stakedBalance += amount;
        emit Staked(id(), amount);
    }

    /// @notice Record a withdraw the stake balance of the provider.
    /// @param amount Amount of staked balance to remove.
    function recordStakedWithdraw(uint256 amount) external virtual onlyYieldManager {
        stakedBalance -= amount;
        emit Unstaked(id(), amount);
    }

    /// @notice Record a pending balance to the provider. Needed only for providers
    ///         that use two-step withdrawals (e.g. Lido).
    function recordPending(uint256 amount) external virtual onlyYieldManager {
        pendingBalance += amount;
        emit Pending(id(), amount);
    }

    /// @notice Record a claimed balance to the provider. For providers with one-step
    ///         withdrawals, this method should be overriden to just emit the event
    ///         to avoid integer underflow.
    function recordClaimed(uint256 amount) external virtual onlyYieldManager {
        pendingBalance -= amount;
        emit Claimed(id(), amount);
    }

    function _beforeCommitYield() internal virtual {}
    function _afterCommitYield() internal virtual {}
}
