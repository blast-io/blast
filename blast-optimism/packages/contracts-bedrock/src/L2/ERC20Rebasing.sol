// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { YieldMode } from "src/L2/Blast.sol";
import { ERC20Permit } from "src/L2/ERC20Permit.sol";

/// @custom:upgradeable
/// @title ERC20Rebasing
/// @notice ERC20 implementation with rebasing token balances. There are 3 yield
/// modes with different rebasing behaviors.
///
/// AUTOMATIC dynamically updates the balance as the share price increases.
///
/// VOID fixes the balance and exempts the account from receiving yields.
///
/// CLAIMABLE fixes the balance and allows the account to claim yields to
/// another account.
///
/// The child implementation is responsible for deciding how the share price is set.
abstract contract ERC20Rebasing is Initializable, ERC20Permit, IERC20 {
    /// @notice Number of decimals.
    uint8 public immutable decimals;

    /// @notice Name of the token.
    string public name;
    /// @notice Symbol of the token.
    string public symbol;

    /// @notice Mapping that stores the number of shares for each account.
    mapping(address => uint256) private _shares;

    /// @notice Total number of shares distributed.
    uint256 internal _totalShares;

    /// @notice Mapping that stores the number of remainder tokens for each account.
    mapping(address => uint256) private _remainders;

    /// @notice Mapping that stores the number of fixed tokens for each account.
    mapping(address => uint256) private _fixed;

    /// @notice Total number of non-rebasing tokens.
    uint256 internal _totalVoidAndRemainders;

    /// @notice Mapping that stores the configured yield mode for each account.
    mapping(address => YieldMode) private _yieldMode;

    /// @notice Mapping that stores the allowance for a given spender and operator pair.
    mapping(address => mapping(address => uint256)) private _allowances;

    /// @notice Reserve extra slots (to a total of 50) in the storage layout for future upgrades.
    ///         A gap size of 41 was chosen here, so that the first slot used in a child contract
    ///         would be a multiple of 50.
    uint256[41] private __gap;

    /// @notice Emitted when an account configures their yield mode.
    /// @param account   Address of the account.
    /// @param yieldMode Yield mode that was configured.
    event Configure(address indexed account, YieldMode yieldMode);

    /// @notice Emitted when a CLAIMABLE account claims their yield.
    /// @param account   Address of the account.
    /// @param recipient Address of the recipient.
    /// @param amount    Amount of yield claimed.
    event Claim(address indexed account, address indexed recipient, uint256 amount);

    error InsufficientBalance();
    error InsufficientAllowance();
    error CannotClaimToSameAccount();
    error TransferFromZeroAddress();
    error TransferToZeroAddress();
    error ApproveFromZeroAddress();
    error ApproveToZeroAddress();
    error NotClaimableAccount();

    /// @param _name     Token name.
    /// @param _symbol   Token symbol.
    /// @param _decimals Number of decimals.
    constructor(string memory _name, string memory _symbol, uint8 _decimals) ERC20Permit(_name) {
        decimals = _decimals;
    }

    /// @param _name     Token name.
    /// @param _symbol   Token symbol.
    function __ERC20Rebasing_init(string memory _name, string memory _symbol) internal onlyInitializing {
        name = _name;
        symbol = _symbol;
    }

    /// @notice --- ERC20 Interface ---

    /// @inheritdoc IERC20
    function totalSupply() external view returns (uint256) {
        return sharePrice() * _totalShares + _totalVoidAndRemainders;
    }

    /// @inheritdoc IERC20
    function balanceOf(address account)
        public
        view
        virtual
        returns (uint256 value)
    {
        YieldMode yieldMode = _yieldMode[account];
        if (yieldMode == YieldMode.AUTOMATIC) {
            value = _computeShareValue(sharePrice(), _shares[account], _remainders[account]);
        } else {
            value = _fixed[account];
        }
    }

    /// @inheritdoc IERC20
    function allowance(address owner, address spender)
        public
        view
        virtual
        returns (uint256)
    {
        return _allowances[owner][spender];
    }

    /// @inheritdoc IERC20
    function transfer(address to, uint256 amount)
        public
        virtual
        returns (bool)
    {
        _transfer(msg.sender, to, amount);
        return true;
    }

    /// @inheritdoc IERC20
    function approve(address spender, uint256 amount)
        public
        virtual
        returns (bool)
    {
        address owner = msg.sender;
        _approve(owner, spender, amount);
        return true;
    }

    /// @inheritdoc IERC20
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public virtual returns (bool) {
        _spendAllowance(from, msg.sender, amount);
        _transfer(from, to, amount);
        return true;
    }

    /// @notice --- Blast Interface ---

    /// @notice Query an account's configured yield mode.
    /// @param account Address to query the configuration.
    /// @return Configured yield mode.
    function getConfiguration(address account) public view returns (YieldMode) {
        return _yieldMode[account];
    }

    /// @notice Get current share price. To be overridden by
    ///         child contract.
    /// @return sharePrice Current share price.
    function sharePrice() public view virtual returns (uint256);

    /// @notice Query an CLAIMABLE account's claimable yield.
    /// @param account Address to query the claimable amount.
    /// @return amount Claimable amount.
    function getClaimableAmount(address account) public view returns (uint256) {
        if (getConfiguration(account) != YieldMode.CLAIMABLE) {
            revert NotClaimableAccount();
        }

        uint256 shareValue = _computeShareValue(sharePrice(), _shares[account], _remainders[account]);
        return shareValue - _fixed[account];
    }

    /// @notice Claim yield from a CLAIMABLE account and send to
    ///         a recipient.
    /// @param recipient Address to receive the claimed balance.
    /// @param amount    Amount to claim.
    /// @return amount Amount claimed.
    function claim(address recipient, uint256 amount) external returns (uint256) {
        address account = msg.sender;

        if (account == recipient) {
            revert CannotClaimToSameAccount();
        }

        if (getConfiguration(account) != YieldMode.CLAIMABLE) {
            revert NotClaimableAccount();
        }

        uint256 currentSharePrice = sharePrice();
        uint256 shareValue = _computeShareValue(currentSharePrice, _shares[account], _remainders[account]);

        uint256 claimableAmount = shareValue - _fixed[account];
        if (amount > claimableAmount) {
            revert InsufficientBalance();
        }

        (uint256 newShares, uint256 newRemainder) = _computeSharesAndRemainder(currentSharePrice, shareValue - amount);

        _deposit(recipient, amount);
        _updateBalance(account, newShares, newRemainder, _fixed[account]);

        emit Claim(msg.sender, recipient, amount);

        return amount;
    }

    /// @notice Change the yield mode of the caller and update the
    ///         balance to reflect the configuration.
    /// @param yieldMode Yield mode to configure
    /// @return Current user balance
    function configure(YieldMode yieldMode) external returns (uint256) {
        _configure(msg.sender, yieldMode);

        emit Configure(msg.sender, yieldMode);

        return balanceOf(msg.sender);
    }

    /// @notice Moves `amount` of tokens from `from` to `to`.
    /// @param from   Address of the sender.
    /// @param to     Address of the recipient.
    /// @param amount Amount of tokens to send.
    function _transfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {
        if (from == address(0)) revert TransferFromZeroAddress();
        if (to == address(0)) revert TransferToZeroAddress();

        uint256 currentSharePrice = sharePrice();
        uint256 fromBalance = balanceOf(from);
        if (amount > fromBalance) {
            revert InsufficientBalance();
        }
        uint256 toBalance = balanceOf(to);
        _setBalance(from, fromBalance - amount, currentSharePrice, false);
        _setBalance(to, toBalance + amount, currentSharePrice, false);

        /// If the sender is configured as VOID, then the amount
        /// is deducted from the total voided funds.
        YieldMode fromYieldMode = getConfiguration(from);
        if (fromYieldMode == YieldMode.VOID) {
            _totalVoidAndRemainders -= amount;
        }

        /// If the receiver is configured as VOID, then the amount
        /// is added to the total voided funds.
        YieldMode toYieldMode = getConfiguration(to);
        if (toYieldMode == YieldMode.VOID) {
            _totalVoidAndRemainders += amount;
        }

        emit Transfer(from, to, amount);
    }

    /// @notice Sets `amount` as the allowance of `spender` over the `owner` s tokens.
    /// @param owner   Address of the owner.
    /// @param spender Address of the spender.
    /// @param amount  Amount of tokens to approve.
    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) internal override {
        if (owner == address(0)) revert ApproveFromZeroAddress();
        if (spender == address(0)) revert ApproveToZeroAddress();

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /// @notice Updates `owner` s allowance for `spender` based on spent `amount`.
    /// @param owner   Address of the owner.
    /// @param spender Address of the spender.
    /// @param amount  Amount of tokens to spender.
    function _spendAllowance(
        address owner,
        address spender,
        uint256 amount
    ) internal virtual {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance != type(uint256).max) {
            if (amount > currentAllowance) revert InsufficientAllowance();
            unchecked {
                _approve(owner, spender, currentAllowance - amount);
            }
        }
    }

    /// @notice Deposit to an account.
    /// @param account Address of the account to deposit to.
    /// @param amount  Amount to deposit to the account.
    function _deposit(address account, uint256 amount) internal {
        uint256 balanceAfter = balanceOf(account) + amount;
        _setBalance(account, balanceAfter, sharePrice(), false);

        /// If the user is configured as VOID, then the amount
        /// is added to the total voided funds.
        YieldMode yieldMode = getConfiguration(account);
        if (yieldMode == YieldMode.VOID) {
            _totalVoidAndRemainders += amount;
        }
    }

    /// @notice Withdraw from an account.
    /// @param account Address of the account to withdraw from.
    /// @param amount  Amount to withdraw to the account.
    function _withdraw(address account, uint256 amount) internal {
        uint256 balance = balanceOf(account);
        if (amount > balance) {
            revert InsufficientBalance();
        }

        unchecked {
            _setBalance(account, balance - amount, sharePrice(), false);
        }

        /// If the user is configured as VOID, then the amount
        /// is deducted from the total voided funds.
        YieldMode yieldMode = getConfiguration(account);
        if (yieldMode == YieldMode.VOID) {
            _totalVoidAndRemainders -= amount;
        }
    }

    /// @notice Configures a new yield mode for an account and updates
    ///         the balance storage to reflect the change.
    /// @param account      Address of the account to configure.
    /// @param newYieldMode New yield mode to configure.
    function _configure(address account, YieldMode newYieldMode) internal {
        uint256 balance = balanceOf(account);

        YieldMode prevYieldMode = getConfiguration(account);
        _yieldMode[account] = newYieldMode;

        uint256 prevFixed = _fixed[account];

        _setBalance(account, balance, sharePrice(), true);

        /// If the previous yield mode was VOID, then the amount
        /// is deducted from the total voided funds.
        if (prevYieldMode == YieldMode.VOID) {
            _totalVoidAndRemainders -= prevFixed;
        }

        /// If the new yield mode is VOID, then the amount
        /// is added to the total voided funds.
        if (newYieldMode == YieldMode.VOID) {
            _totalVoidAndRemainders += balance;
        }
    }

    /// @notice Sets the balance of an account according to it's yield mode
    ///         configuration.
    /// @param account           Address of the account to set the balance of.
    /// @param amount            Balance to set for the account.
    /// @param currentSharePrice Current share price.
    /// @param resetClaimable    If the account is CLAIMABLE, true if the share
    ///                          balance should be set to the amount. Should only be true when
    ///                          configuring the account.
    function _setBalance(address account, uint256 amount, uint256 currentSharePrice, bool resetClaimable) internal {
        uint256 newShares; uint256 newRemainder; uint256 newFixed;
        YieldMode yieldMode = getConfiguration(account);
        if (yieldMode == YieldMode.AUTOMATIC) {
            (newShares, newRemainder) = _computeSharesAndRemainder(currentSharePrice, amount);
        } else if (yieldMode == YieldMode.VOID) {
            newFixed = amount;
        } else if (yieldMode == YieldMode.CLAIMABLE) {
            newFixed = amount;
            uint256 shareValue = amount;
            if (!resetClaimable) {
                /// In order to not reset the claimable balance, we have to compute
                /// the user's current share balance and add or subtract the change in
                /// fixed balance before computing the new shares balance parameters.
                shareValue = _computeShareValue(currentSharePrice, _shares[account], _remainders[account]);
                shareValue = shareValue + amount - _fixed[account];
            }
            (newShares, newRemainder) = _computeSharesAndRemainder(currentSharePrice, shareValue);
        }

        _updateBalance(account, newShares, newRemainder, newFixed);
    }

    /// @notice Update the balance paramters of an account and appropriately refresh the global sums
    ///         to reflect the change of allocation.
    /// @param account      Address of account to update.
    /// @param newShares    New shares value for account.
    /// @param newRemainder New remainder value for account.
    /// @param newFixed     New fixed value for account.
    function _updateBalance(address account, uint256 newShares, uint256 newRemainder, uint256 newFixed) internal {
        _totalShares = _totalShares + newShares - _shares[account];
        _totalVoidAndRemainders = _totalVoidAndRemainders + newRemainder - _remainders[account];

        _shares[account] = newShares;
        _remainders[account] = newRemainder;
        _fixed[account] = newFixed;
    }

    /// @notice Convert nominal value to number of shares with remainder.
    /// @param currentSharePrice Current share price (wad).
    /// @param value             Amount to convert to shares (wad).
    /// @return shares Number of shares (wad), remainder Remainder (wad).
    function _computeSharesAndRemainder(uint256 currentSharePrice, uint256 value) internal pure returns (uint256 shares, uint256 remainder) {
        if (currentSharePrice == 0) {
            remainder = value;
        } else {
            shares = value / currentSharePrice;
            remainder = value % currentSharePrice;
        }
    }

    /// @notice Compute nominal value from number of shares.
    /// @param currentSharePrice Current share price (wad).
    /// @param shares            Number of shares (wad).
    /// @param remainders        (wad).
    /// @return value (wad).
    function _computeShareValue(uint256 currentSharePrice, uint256 shares, uint256 remainders) internal pure returns (uint256) {
        return currentSharePrice * shares + remainders;
    }
}
