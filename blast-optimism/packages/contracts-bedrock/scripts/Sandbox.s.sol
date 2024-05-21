// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "scripts/Scripts.s.sol";

contract Sandbox is ScriptInitializer {
    function name() public pure override returns (string memory) {
        return "Sandbox";
    }
}
