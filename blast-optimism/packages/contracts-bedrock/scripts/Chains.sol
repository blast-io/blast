// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Chain IDs for the various networks.
library Chains {
    uint256 internal constant Mainnet = 1;
    uint256 internal constant BlastMainnet = 81457;
    uint256 internal constant Sepolia = 11155111;
    uint256 internal constant BlastSepolia = 11155420;
    uint256 internal constant LocalDevnet = 900;
    uint256 internal constant BlastLocalDevnet = 901;
    uint256 internal constant GethDevnet = 1337;
    uint256 internal constant Hardhat = 31337;
}
