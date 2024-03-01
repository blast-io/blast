// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console2 as console } from "forge-std/console2.sol";

import { ScriptInitializer } from "scripts/CommonE2E.s.sol";
import { ETHTestnetYieldProvider } from "src/mainnet-bridge/yield-providers/ETHTestnetYieldProvider.sol";
import { USDTestnetYieldProvider } from "src/mainnet-bridge/yield-providers/USDTestnetYieldProvider.sol";
import { LidoYieldProvider } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { Insurance } from "src/mainnet-bridge/Insurance.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";

contract Sandbox is ScriptInitializer {
    function name() public pure override returns (string memory) {
        return "Sandbox";
    }
}
