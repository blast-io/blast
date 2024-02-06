// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";
import { WithdrawalQueue } from "src/mainnet-bridge/withdrawal-queue/WithdrawalQueue.sol";

interface ILido is IERC20 {
    function submit(address referral) external payable;
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool);
    function isStakingPaused() external view returns (bool);
}

interface IWithdrawalQueue {
    function getLastCheckpointIndex() external view returns (uint256);
    function findCheckpointHints(uint256[] calldata _requestIds, uint256 _firstIndex, uint256 _lastIndex) external view returns (uint256[] memory hintIds);
    function requestWithdrawals(uint256[] calldata _amounts, address _owner) external returns (uint256[] memory requestIds);
    function claimWithdrawals(uint256[] calldata _requestIds, uint256[] calldata _hints) external;
}

interface IInsurance {
    function coverLoss(address token, uint256 amount) external;
}

/// @title LidoYieldProvider
/// @notice Provider for the Lido (ETH) yield source.
contract LidoYieldProvider is YieldProvider {
    ILido public constant LIDO = ILido(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
    IWithdrawalQueue public constant WITHDRAWAL_QUEUE = IWithdrawalQueue(0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1);

    /// @notice Emitted when a withdrawal is requested from Lido.
    /// @param requestId Lido WitdrawalQueue requestId.
    /// @param amount    Amount requested for withdrawal.
    event LidoUnstakeInitiated(uint256 indexed requestId, uint256 amount);

    /// @param _yieldManager Address of the yield manager for the underlying
    ///        yield asset of this provider.
    constructor(YieldManager _yieldManager) YieldProvider(_yieldManager) {}

    /// @inheritdoc YieldProvider
    function initialize() external override onlyDelegateCall {
        // LIDO.approve(address(WITHDRAWAL_QUEUE), type(uint256).max);
        // LIDO.approve(address(YIELD_MANAGER), type(uint256).max);
    }

    /// @inheritdoc YieldProvider
    function name() public pure override returns (string memory) {
        return "LidoYieldProvider";
    }

    /// @inheritdoc YieldProvider
    function isStakingEnabled(address token) public view override returns (bool) {
        return token == address(LIDO) && LIDO.isStakingPaused();
    }

    /// @inheritdoc YieldProvider
    function stakedValue() public view override returns (uint256) {
        return LIDO.balanceOf(address(YIELD_MANAGER));
    }

    /// @inheritdoc YieldProvider
    function yield() public view override returns (int256) {
        return int256(stakedValue()) - int256(stakedBalance);
    }

    /// @inheritdoc YieldProvider
    function supportsInsurancePayment() public view override returns (bool) {
        return YIELD_MANAGER.insurance() != address(0);
    }

    /// @inheritdoc YieldProvider
    function stake(uint256 amount) external override onlyDelegateCall {
        if (amount > YIELD_MANAGER.lockedValue()) {
            revert InsufficientStakableFunds();
        }
        LIDO.submit{value: amount}(address(0));
    }

    /// @inheritdoc YieldProvider
    function unstake(uint256 amount) external override onlyDelegateCall returns (uint256 pending) {
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = amount;
        uint256 requestId = WITHDRAWAL_QUEUE.requestWithdrawals(amounts, address(YIELD_MANAGER))[0];
        emit LidoUnstakeInitiated(requestId, amount);
        pending = amount;
    }

    /// @inheritdoc YieldProvider
    function claim(uint256[] calldata requestIds) external override onlyDelegateCall returns (uint256 claimed) {
        uint256 balanceBefore = address(YIELD_MANAGER).balance;
        uint256[] memory hintIds = WITHDRAWAL_QUEUE.findCheckpointHints(
            requestIds,
            0,
            WITHDRAWAL_QUEUE.getLastCheckpointIndex()
        );
        WITHDRAWAL_QUEUE.claimWithdrawals(requestIds, hintIds);
        claimed = address(YIELD_MANAGER).balance - balanceBefore;
    }

    /// @inheritdoc YieldProvider
    function payInsurancePremium(uint256 amount) external override onlyDelegateCall {
        require(supportsInsurancePayment(), "insurance not supported");

        // send stETH to insurance
        LIDO.transfer(YIELD_MANAGER.insurance(), amount);
        // there is no need to update staked balance as insurance premium is expected
        // to come from the yield

        emit InsurancePremiumPaid(id(), amount);
    }

    /// @inheritdoc YieldProvider
    function withdrawFromInsurance(uint256 amount) external override onlyDelegateCall {
        require(supportsInsurancePayment(), "insurance not supported");
        IInsurance(YIELD_MANAGER.insurance()).coverLoss(address(LIDO), amount);

        emit InsuranceWithdrawn(id(), amount);
    }

    /// @inheritdoc YieldProvider
    function insuranceBalance() public view override returns (uint256) {
        return LIDO.balanceOf(YIELD_MANAGER.insurance());
    }
}
