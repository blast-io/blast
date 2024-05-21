// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { ERC20PermitUpgradeable } from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";

import { ERC20Rebasing } from "src/L2/ERC20Rebasing.sol";
import { Blast, YieldMode } from "src/L2/Blast.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { GasMode } from "src/L2/Gas.sol";

/// @custom:proxied
/// @title StERC20 token wrapper with static balances.
/// @dev It's an ERC20 token that represents the account's share of the total
/// supply of stERC20 tokens. NrERC20 token's balance only changes on transfers,
/// unlike StERC20 that is also changed when oracles report staking rewards and
/// penalties. It's a "power user" token for DeFi protocols which don't
/// support rebasable tokens.
///
/// The contract is also a trustless wrapper that accepts stERC20 tokens and mints
/// nrETH in return. Then the user unwraps, the contract burns user's nrERC20
/// and sends user locked stERC20 in return.
///
/// The contract provides the staking shortcut: user can send ERC20 with regular
/// transfer and get nrERC20 in return. The contract will send ERC20 to Lido submit
/// method, staking it and wrapping the received stERC20.
contract NrERC20 is ERC20PermitUpgradeable {
    ERC20Rebasing public immutable TOKEN;
    uint8 internal immutable DECIMALS;

    /// @notice _token Underlying stERC20 token.
    constructor(ERC20Rebasing _token, uint8 _decimals) {
        _disableInitializers();
        TOKEN = _token;
        DECIMALS = _decimals;
    }

    /// @notice Initializer
    function initialize(string memory name, string memory symbol) public initializer {
        __ERC20_init(name, symbol);
        __ERC20Permit_init(name);
        TOKEN.configure(YieldMode.AUTOMATIC);
        Blast(Predeploys.BLAST).configureContract(
            address(this),
            YieldMode.VOID,
            GasMode.VOID,
            address(0xdead) /// don't set a governor
        );
    }

    function decimals() public view override returns (uint8) {
        return DECIMALS;
    }

    /// @notice Exchanges stERC20 to nrERC20.
    /// @param _amount Unwrapped amount to deposit.
    /// @return Amount of nrERC20 user receives after wrap.
    function wrap(uint256 _amount) public returns (uint256) {
        require(_amount > 0, "nrERC20: can't wrap zero stERC20");
        uint256 shares = getNrERC20ByStERC20(_amount);
        uint256 realAmount = getStERC20ByNrERC20(shares);
        _mint(msg.sender, shares);
        TOKEN.transferFrom(msg.sender, address(this), realAmount);
        return shares;
    }

    /// @notice Exchanges nrERC20 to stERC20.
    /// @param _shares Amount of shares to burn.
    /// @return Amount of stERC20 user receives after unwrap.
    function unwrap(uint256 _shares) external returns (uint256) {
        require(_shares > 0, "nrERC20: zero amount unwrap not allowed");
        uint256 amount = getStERC20ByNrERC20(_shares);
        _burn(msg.sender, _shares);
        TOKEN.transfer(msg.sender, amount);
        return amount;
    }

    /// @notice Get amount of nrERC20 for a given amount of stERC20.
    /// @param _amount Amount
    /// @return Amount of nrERC20.
    function getNrERC20ByStERC20(uint256 _amount) public view returns (uint256) {
        return _amount / TOKEN.price();
    }

    /// @notice Get amount of stERC20 for a given amount of nrERC20.
    /// @param _shares Shares
    /// @return Amount of stERC20.
    function getStERC20ByNrERC20(uint256 _shares) public view returns (uint256) {
        return _shares * TOKEN.price();
    }

    /// @notice Get amount of stERC20 for one nrERC20.
    /// @return Amount of stERC20.
    function stERC20PerToken() external view returns (uint256) {
        return getStERC20ByNrERC20(10 ** decimals());
    }

    /// @notice Get amount of nrERC20 for one stERC20.
    /// @return Amount of tokens.
    function tokensPerStERC20() external view returns (uint256) {
        return getNrERC20ByStERC20(10 ** TOKEN.decimals());
    }
}
