// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ISemver } from "src/universal/ISemver.sol";
import { Storage } from "src/libraries/Storage.sol";

/// @title StorageSetter
/// @notice A simple contract that allows setting arbitrary storage slots.
///         WARNING: this contract is not safe to be called by untrusted parties.
///         It is only meant as an intermediate step during upgrades.
contract StorageSetter is ISemver {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice Stores a bytes32 `_value` at `_slot`. Any storage slots that
    ///         are packed should be set through this interface.
    function setBytes32(bytes32 _slot, bytes32 _value) public {
        Storage.setBytes32(_slot, _value);
    }

    /// @notice Retrieves a bytes32 value from `_slot`.
    function getBytes32(bytes32 _slot) external view returns (bytes32) {
        return Storage.getBytes32(_slot);
    }

    /// @notice Stores a uint256 `_value` at `_slot`.
    function setUint(bytes32 _slot, uint256 _value) public {
        Storage.setUint(_slot, _value);
    }

    /// @notice Retrieves a uint256 value from `_slot`.
    function getUint(bytes32 _slot) external view returns (uint256) {
        return Storage.getUint(_slot);
    }

    /// @notice Stores an address `_value` at `_slot`.
    function setAddress(bytes32 _slot, address _address) public {
        Storage.setAddress(_slot, _address);
    }

    /// @notice Retrieves an address value from `_slot`.
    function getAddress(bytes32 _slot) external view returns (address) {
        return Storage.getAddress(_slot);
    }
}
