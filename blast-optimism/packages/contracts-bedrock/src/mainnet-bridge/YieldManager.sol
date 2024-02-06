// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { Ownable2StepUpgradeable } from "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import { SignedMath } from "@openzeppelin/contracts/utils/math/SignedMath.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { WithdrawalQueue } from "src/mainnet-bridge/withdrawal-queue/WithdrawalQueue.sol";
import { YieldProvider } from "src/mainnet-bridge/yield-providers/YieldProvider.sol";
import { Types } from "src/libraries/Types.sol";
import { SafeCall } from "src/libraries/SafeCall.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { SharesBase } from "src/L2/Shares.sol";
import { DelegateCalls } from "src/mainnet-bridge/DelegateCalls.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";
import { Semver } from "src/universal/Semver.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/// @title YieldManager
/// @notice Base contract to centralize accounting, asset management and
///         yield reporting from yield providers of a common base asset.
abstract contract YieldManager is Ownable2StepUpgradeable, WithdrawalQueue, DelegateCalls {
    using EnumerableSet for EnumerableSet.AddressSet;

    /// @notice Maximum gas limit for the yield report call on L2.
    uint32 internal constant REPORT_YIELD_DEFAULT_GAS_LIMIT = 200_000;

    /// @notice Maximum insurance fee the owner is allowed to set.
    uint256 public constant MAX_INSURANCE_FEE_BIPS = 10_000; // 100%

    /// @notice Number of basis points representing 100 percent.
    uint256 internal constant BASIS_POINTS = 10_000;

    /// @notice Set of provider addresses.
    EnumerableSet.AddressSet private _providers;

    /// @notice Address of the admin handling regular tasks such as
    ///         `stake`, `unstake`, `claim`, `commitYieldReport`, and
    ///         `finalize`.
    address public admin;

    /// @notice Address of the insurance module.
    address public insurance;

    /// @notice Address of the L1BlastBridge.
    address public blastBridge;

    /// @notice Sum of negative yields to track the slippage between L2-L1 share price.
    ///         If negative yields accumulate, L1 withdrawals are discounted to cover the
    ///         loss.
    uint256 public accumulatedNegativeYields;

    /// @notice Current insurance fee in bips.
    uint256 public insuranceFeeBips;

    /// @notice Amount of additional funds to withdraw from insurance.
    ///         This buffer addresses the scenario where the transfer of the exact amount of accumulated
    ///         negative yields from insurance does not fully pay off the outstanding amount. In Lido's
    ///         system, the transfer logic is based on shares, which may lead to discrepancies in the
    ///         withdrawal of insurance funds. By including this buffer, the system ensures that when
    ///         insurance funds are withdrawn, the total amount withdrawn is the exact required amount
    ///         plus an additional buffer. This approach guarantees the complete payoff of any negative
    ///         yields, accommodating for any potential rounding discrepancies inherent in the share-based
    ///         transfer logic.
    uint256 public insuranceWithdrawalBuffer;

    /// @notice Address of the OptimismPortal.
    OptimismPortal public portal;

    struct ProviderInfo {
        bytes32 id;
        address providerAddress;
        uint256 stakedValue;
        uint256 pendingBalance;
        uint256 stakedBalance;
        uint256 totalValue;
        int256 yield;
    }

    /// @notice Emitted when the yield report is committed on L1 and
    ///         the yield is communicated to L2.
    /// @param yield                Amount of yield generated at this checkpoint.
    /// @param insurancePremiumPaid Amount paid in insurance.
    /// @param insuranceWithdrawn   Amount withdrawn from insurance.
    event YieldReport(
        int256  yield,
        uint256 insurancePremiumPaid,
        uint256 insuranceWithdrawn
    );

    error CallerIsNotAdmin();
    error FailedToInitializeProvider();
    error ProviderAddressDoesNotMatchIndex();
    error InsufficientInsuranceBalance();
    error NegativeYieldFromInsuredProvider();
    error TotalValueIsZero();
    error CallerIsNotBlastBridge();
    error ProviderNotFound();
    error YieldProviderIsNotMeantForThisManager();
    error UnsupportedToken();

    modifier onlyAdmin() {
        if (msg.sender != admin) {
            revert CallerIsNotAdmin();
        }
        _;
    }

    /// @param _token Address of withdrawal token.
    constructor(address _token) WithdrawalQueue(_token) {}

    /// @notice initializer
    /// @param _portal Address of the OptimismPortal.
    function __YieldManager_init(OptimismPortal _portal) internal onlyInitializing {
        __Ownable2Step_init();
        __WithdrawalQueue_init();

        portal = _portal;
        admin = msg.sender;
    }

    /* ========== OWNER FUNCTIONS ========== */

    /// @notice Set new admin account to handle regular tasks including
    ///         (stake, unstake, claim).
    /// @param _admin Address of new admin
    function setAdmin(address _admin) external onlyOwner {
        require(_admin != address(0));
        admin = _admin;
    }

    /// @notice Set the yield insurance parameters.
    /// @param _insurance        Address of the insurance module.
    /// @param _insuranceFeeBips Insurance fee to take from positive yields.
    /// @param _withdrawalBuffer Amount of additional funds to withdraw from insurance.
    function setInsurance(address _insurance, uint256 _insuranceFeeBips, uint256 _withdrawalBuffer) external onlyOwner {
        require(_insurance != address(0));
        require(_insuranceFeeBips <= MAX_INSURANCE_FEE_BIPS);
        insurance = _insurance;
        insuranceFeeBips = _insuranceFeeBips;
        insuranceWithdrawalBuffer = _withdrawalBuffer;
    }

    /// @notice Set the address of the L1BlastBridge.
    /// @param _blastBridge Address of the L1BlastBridge.
    function setBlastBridge(address _blastBridge) external onlyOwner {
        require(_blastBridge != address(0));
        blastBridge = _blastBridge;
    }

    /// @notice Add a yield provider contract.
    /// @param provider Address of the yield provider.
    function addProvider(address provider) external onlyOwner {
        if (address(YieldProvider(provider).YIELD_MANAGER()) != address(this)) {
            revert YieldProviderIsNotMeantForThisManager();
        }
        _providers.add(provider);
        (bool success,) = provider.delegatecall(abi.encodeWithSignature("initialize()"));
        if (!success) {
            revert FailedToInitializeProvider();
        }
    }

    /// @notice Remove a yield provider contract.
    /// @param provider Address of the yield provider.
    function removeProvider(address provider) external onlyOwner {
        _providers.remove(provider);
    }

    /* ========== ADMIN FUNCTIONS ========== */

    /// @notice Stake funds for a particular yield provider and record the
    ///         staked deposit. The stake call is made via 'delegatecall'
    ///         so the yield provider implementation is executed with the
    ///         yield manager's funds.
    /// @param idx             Index of the provider.
    /// @param providerAddress Address of the provider at index 'idx'.
    /// @param amount          Amount to stake (wad).
    function stake(uint256 idx, address providerAddress, uint256 amount) external onlyAdmin {
        if (_providers.at(idx) != providerAddress) {
            revert ProviderAddressDoesNotMatchIndex();
        }
        _delegatecall_uint256_arg(providerAddress, stake_signature, amount);
        YieldProvider(providerAddress).recordStakedDeposit(amount);
    }

    /// @notice Unstake funds for a particular yield provider and record the
    ///         staked withdraw. The stake call is made via 'delegatecall'
    ///         so the yield provider implementation is executed with the
    ///         yield manager's funds.
    /// @param idx             Index of the provider.
    /// @param providerAddress Address of the provider at index 'idx'.
    /// @param amount          Amount to stake (wad).
    function unstake(uint256 idx, address providerAddress, uint256 amount) external onlyAdmin {
        if (_providers.at(idx) != providerAddress) {
            revert ProviderAddressDoesNotMatchIndex();
        }
        uint256 pending = _delegatecall_uint256_arg_returns_uint256(providerAddress, unstake_signature, amount);
        YieldProvider(providerAddress).recordStakedWithdraw(amount);
        if (pending > 0) {
            YieldProvider(providerAddress).recordPending(pending);
        } else {
            YieldProvider(providerAddress).recordClaimed(amount);
        }
    }

    /// @notice Claim pending funds. Needed only for some yield providers that require two-step
    ///         withdrawal (e.g. Lido).
    /// @param idx             Index of the provider.
    /// @param providerAddress Address of the provider at index 'idx'.
    /// @param requestIds Array of the request ids to claim.
    function claimPending(uint256 idx, address providerAddress, uint256[] calldata requestIds) external onlyAdmin {
        if (_providers.at(idx) != providerAddress) {
            revert ProviderAddressDoesNotMatchIndex();
        }
        uint256 claimed = _delegatecall_uint256_arr_arg_returns_uint256(providerAddress, claim_signature, requestIds);
        YieldProvider(providerAddress).recordClaimed(claimed);
    }

    /// @notice Commit yield report.
    /// @param enableInsurance Whether insurance should be taken from positive yields
    ///        and paid out for negative yields. If false, negative yields will
    ///        accumulate and withdrawals will be discounted. If true (and insurance
    ///        is supported by the provider), it will guarantee that committed yield
    ///        is always non-negative, or else revert.
    function commitYieldReport(bool enableInsurance) external onlyAdmin {
        uint256 providersLength = _providers.length();
        uint256 totalInsurancePremiumPaid;
        uint256 totalInsuranceWithdrawal;
        int256 totalYield;

        // For each provider, commit yield after paying to/from the insurance as necessary
        for (uint256 i; i < providersLength; i++) {
            // read the current yield from the provider
            int256 yield = YieldProvider(_providers.at(i)).yield();
            uint256 insurancePayment;

            // take care of insurance payments and withdrawals
            if (
                enableInsurance &&
                YieldProvider(_providers.at(i)).supportsInsurancePayment() &&
                insurance != address(0)
            ) {
                if (yield > 0) {
                    // pay the insurance premium
                    insurancePayment = uint256(yield) * insuranceFeeBips / BASIS_POINTS;
                    _delegatecall_uint256_arg(_providers.at(i), payInsurancePremium_signature, insurancePayment);
                    totalInsurancePremiumPaid += insurancePayment;
                } else if (yield < 0) {
                    // withdraw from the insurance to cover the loss
                    uint256 insuranceWithdrawal = SignedMath.abs(yield) + insuranceWithdrawalBuffer;
                    uint256 insuranceBalance = YieldProvider(_providers.at(i)).insuranceBalance();
                    if (insuranceBalance < insuranceWithdrawal) {
                        revert InsufficientInsuranceBalance();
                    }
                    _delegatecall_uint256_arg(_providers.at(i), withdrawFromInsurance_signature, insuranceWithdrawal);
                    totalInsuranceWithdrawal += insuranceWithdrawal;
                }
            }

            // Commit the yield for the provider
            int256 committedYield = YieldProvider(_providers.at(i)).commitYield();

            // Sanity check
            if (
                enableInsurance &&
                YieldProvider(_providers.at(i)).supportsInsurancePayment() &&
                insurance != address(0)
            ) {
                if (committedYield < 0) {
                    revert NegativeYieldFromInsuredProvider();
                }
            }

            // update totalYield
            totalYield += committedYield;
        }

        // reflect the accumulated negative yield in totalYield
        if (accumulatedNegativeYields > 0) {
            totalYield -= int256(accumulatedNegativeYields);
        }

        emit YieldReport(totalYield, totalInsurancePremiumPaid, totalInsuranceWithdrawal);

        if (totalYield < 0) {
            accumulatedNegativeYields = uint256(-1 * totalYield);
        } else {
            accumulatedNegativeYields = 0;
            if (totalYield > 0) {
                _reportYield(
                    abi.encodeWithSelector(
                        SharesBase.addValue.selector,
                        totalYield
                    )
                );
            }
        }
    }

    /// @notice Finalize withdrawal requests up to 'requestId'.
    /// @param requestId Last request id to finalize in this batch.
    function finalize(uint256 requestId) external onlyAdmin returns (uint256 checkpointId) {
        uint256 nominalAmount; uint256 realAmount;
        (nominalAmount, realAmount, checkpointId) = _finalize(requestId, getTokenBalance(), sharePrice());
        // nominalAmount - realAmount is the share of the accumulated negative yield
        // that should be paid by the current withdrawal
        if (accumulatedNegativeYields > 0) {
            accumulatedNegativeYields -= (nominalAmount - realAmount);
        }
    }

    /* ========== VIRTUAL FUNCTIONS ========== */

    /// @notice Get the balance of the withdrawal token held by the yield manager
    ///         minus the amount that is locked in the withdrawal queue.
    function getTokenBalance() public view virtual returns (uint256);

    /// @notice Get the amount of the withdrawal token that is held by the yield manager.
    function lockedValue() public view virtual returns (uint256);

    /// @notice Send the yield report to the L2 contract that is responsible for
    ///         updating the L2 share price.
    /// @param data Calldata to send in the message.
    function _reportYield(bytes memory data) internal virtual;

    /* ========== VIEW FUNCTIONS ========== */

    /// @notice Get the total value of all yield providers denominated in the withdrawal token.
    function totalProviderValue() public view returns (uint256 sum) {
        uint256 providersLength = _providers.length();
        for (uint256 i; i < providersLength; i++) {
            sum += YieldProvider(_providers.at(i)).totalProviderValue();
        }
    }

    /// @notice Get the total value of all yield providers plus the locked value.
    function totalValue() public view returns (uint256) {
        return lockedValue() + totalProviderValue();
    }

    /// @notice Get the share price of the withdrawal token with 1e27 precision.
    ///         The share price is capped at 1e27 and can only go down if there
    ///         are accumulated negative yields.
    function sharePrice() public view returns (uint256) {
        uint256 value = totalValue();
        if (value == 0) {
            revert TotalValueIsZero();
        }
        return value * E27_PRECISION_BASE / (value + accumulatedNegativeYields);
    }

    /// @notice Get an accounting report on the current state of a yield provider.
    ///         Due to how EnumerableSet works, 'idx' is not guaranteed to be stable
    ///         across add/remove operations so admin should verify the idx before
    ///         calling state-changing functions (e.g. stake, unstake).
    /// @param idx Index of the provider.
    /// @return info Accounting report on the yield provider.
    function getProviderInfoAt(uint256 idx) external view returns (ProviderInfo memory info) {
        YieldProvider provider = YieldProvider(_providers.at(idx));

        info.id = provider.id();
        info.providerAddress = address(provider);
        info.stakedValue = provider.stakedValue();
        info.pendingBalance = provider.pendingBalance();
        info.stakedBalance = provider.stakedBalance();
        info.totalValue = provider.totalProviderValue();
        info.yield = provider.yield();
    }

    /// @notice Record an increase to the staked funds represented
    ///         by the provider.
    /// @param providerAddress Address of yield provider.
    /// @param amount          Amount of additional staked funds.
    function _recordStakedDeposit(address providerAddress, uint256 amount) internal {
        if (!_providers.contains(providerAddress)) {
            revert ProviderNotFound();
        }
        YieldProvider(providerAddress).recordStakedDeposit(amount);
    }
}

/// @title USDYieldManager
/// @notice Coordinates the accounting, asset management and
///         yield reporting from USD yield providers.
abstract contract USDYieldManager is YieldManager {
    constructor() YieldManager(0x6B175474E89094C44Da98b954EedeAC495271d0F) {} /// DAI address

    /// @notice Reserve extra slots (to a total of 50) in the storage layout for future upgrades.
    ///         A gap size of 41 was chosen here, so that the first slot used in a child contract
    ///         would be a multiple of 50.
    uint256[41] private __gap;

    /// @inheritdoc YieldManager
    function getTokenBalance() public view override returns (uint256) {
        return IERC20(TOKEN).balanceOf(address(this)) - getLockedAmount();
    }

    /// @inheritdoc YieldManager
    function lockedValue() public view override returns (uint256) {
        return IERC20(TOKEN).balanceOf(address(this));
    }

    /// @notice Convert between USDC, USDT, and DAI through Curve's 3Pool and Maker's
    ///         Peg Stability Mechanism.
    /// @param inputToken  Curve 3Pool token index.
    /// @param outputToken Curve 3Pool token index.
    /// @return amountReceived Amount of output token received.
    function convert(int128 inputToken, int128 outputToken, uint256 inputAmountWad, uint256 minOutputAmountWad) external onlyAdmin returns (uint256 amountReceived) {
        return USDConversions._convert(inputToken, outputToken, inputAmountWad, minOutputAmountWad);
    }

    /// @notice Sends the yield report to the USDB contract.
    /// @param data Calldata to send in the message.
    function _reportYield(bytes memory data) internal override {
        portal.depositTransaction(Predeploys.USDB, 0, REPORT_YIELD_DEFAULT_GAS_LIMIT, false, data);
    }
}
