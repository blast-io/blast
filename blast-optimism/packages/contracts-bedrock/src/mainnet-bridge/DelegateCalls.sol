// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

abstract contract DelegateCalls {
    // @notice state-changing, uint256 arg, no return
    uint256 constant payInsurancePremium_signature = (
        0x66785d9300000000000000000000000000000000000000000000000000000000
    );  // payInsurancePremium(uint256)
    uint256 constant withdrawFromInsurance_signature = (
        0x9b75759400000000000000000000000000000000000000000000000000000000
    );  // withdrawFromInsurance(uint256)
    uint256 constant stake_signature = (
        0xa694fc3a00000000000000000000000000000000000000000000000000000000
    );  // stake(uint256)

    // @notice state-changing, uint256 arg, returns uint256
    uint256 constant unstake_signature = (
        0x2e17de7800000000000000000000000000000000000000000000000000000000
    );  // unstake(uint256) returns (uint256)

    // @notice state-changing, uint256[] arg, returns uint256
    uint256 constant claim_signature = (
        0x6ba4c13800000000000000000000000000000000000000000000000000000000
    );  // claim(uint256[]) returns (uint256)

    function _delegatecall_no_arg_no_return(address provider, uint256 selector) internal {
        (bool success,) = provider.delegatecall(abi.encodePacked(selector));
        require(success, "delegatecall failed");
    }

    function _delegatecall_no_arg_returns_uint256(address provider, uint256 selector) internal returns (uint256) {
        (bool success, bytes memory res) = provider.delegatecall(abi.encodePacked(selector));
        require(success, "delegatecall failed");
        return abi.decode(res, (uint256));
    }

    function _delegatecall_uint256_arr_arg_returns_uint256(address provider, uint256 selector, uint256[] memory arg) internal returns (uint256) {
        (bool success, bytes memory res) = provider.delegatecall(abi.encodePacked(selector, arg));
        require(success, "delegatecall failed");
        return abi.decode(res, (uint256));
    }

    function _delegatecall_no_arg_returns_int256(address provider, uint256 selector) internal returns (int256) {
        (bool success, bytes memory res) = provider.delegatecall(abi.encodePacked(selector));
        require(success, "delegatecall failed");
        return abi.decode(res, (int256));
    }

    function _delegatecall_uint256_arg_returns_uint256(address provider, uint256 selector, uint256 arg) internal returns (uint256 result) {
        assembly {
            let ptr := mload(0x40)
            mstore(0x40, add(ptr, 0x24))

            mstore(ptr, selector)
            mstore(add(ptr, 0x4), arg)

            let success := delegatecall(gas(), provider, ptr, 0x24, 0, 0)
            let size := returndatasize()

            returndatacopy(ptr, 0, size)

            switch success
            case 0 { revert(ptr, size) }
            default { 
                result := mload(ptr)
            }
        }
    }

    function _delegatecall_uint256_arg(address provider, uint256 selector, uint256 arg) internal {
        assembly {
            let ptr := mload(0x40)
            mstore(0x40, add(ptr, 0x24))

            mstore(ptr, selector)
            mstore(add(ptr, 0x4), arg)

            let success := delegatecall(gas(), provider, ptr, 0x24, 0, 0)

            if eq(success, 0) {
                revert(0, 0)
            }
        }
    }
}
