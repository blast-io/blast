// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
// Target contract is imported by the `Bridge_Initializer`
import { Bridge_Initializer } from "test/CommonTest.t.sol";
import { ILegacyMintableERC20, IOptimismMintableERC20 } from "src/universal/IOptimismMintableERC20.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

// Libraries
import { Predeploys } from "src/libraries/Predeploys.sol";

// Target contract dependencies
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { YieldMode } from "src/L2/Blast.sol";
import { GasMode } from "src/L2/Gas.sol";
import 'forge-std/console.sol';

contract Gas_Test is Bridge_Initializer {

    event ClaimRateLogged(uint256 claimRate);

    function setUp() public virtual override {
        super.setUp();
    }

    function test_updateAdminParameters() external {
        updateAdminParameters(0x1, 0x2, 0x3, 0x4, 0x5);
        assertEq(gas.ceilClaimRate(), 0x5);
    }

    function test_minClaimRate() external {
        updateAdminParameters(2500, 60, 5000, 100, 8000);
        vm.warp(0);
        vm.deal(address(gas), 1 ether);
        gas.updateGasParams(alice, 0, 1 ether, GasMode.CLAIMABLE);
        uint256 feeVaultBalance = address(Predeploys.BASE_FEE_VAULT).balance;

        vm.warp(80);
        (uint256 etherSeconds, uint256 etherBalance,,) = gas.readGasParams(alice);
        assertEq(etherBalance, 1 ether);
        assertEq(etherSeconds, 80*1 ether);

        vm.prank(Predeploys.BLAST);
        uint256 userEther = gas.claimGasAtMinClaimRate(alice, alice, 6000);
        uint256 feeVaultBalanceConsumed = address(Predeploys.BASE_FEE_VAULT).balance - feeVaultBalance;
        uint256 totalBalance = feeVaultBalanceConsumed + userEther;
        uint256 claimRate = 10_000 * userEther / totalBalance; 
        console.log("claimRate", claimRate);

        assertTrue(claimRate >= 6000);
        assertTrue(claimRate >= 6010);
    }


    function updateAdminParameters(
        uint256 _zeroClaimRate,
        uint256 _baseGasSeconds,
        uint256 _baseClaimRate,
        uint256 _ceilGasSeconds,
        uint256 _ceilClaimRate
    ) internal {
        vm.prank(address(0x0));
        gas.updateAdminParameters(_zeroClaimRate, _baseGasSeconds, _baseClaimRate, _ceilGasSeconds, _ceilClaimRate);
    }
}
