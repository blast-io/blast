// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { ERC20Rebasing } from "src/L2/ERC20Rebasing.sol";
import { WETHRebasing } from "src/L2/WETHRebasing.sol";
import { NrERC20 } from "src/L2/nrERC20.sol";

/// @custom:proxied
/// @title NrETH wrapper to accept ETH deposits.
contract NrETH is NrERC20 {
    constructor(ERC20Rebasing _token, uint8 _decimals) NrERC20(_token, _decimals) {}

    /// @notice Shortcut to deposit native ETH, wrap into WETH, and
    ///         then wrap WETH into nrETH.
    receive() external payable {
        WETHRebasing(payable(address(TOKEN))).deposit{value: msg.value}();
        uint256 shares = getNrERC20ByStERC20(msg.value);
        uint256 realAmount = getStERC20ByNrERC20(shares);
        _mint(msg.sender, shares);
        address(msg.sender).call{value: msg.value - realAmount}("");
    }
}
