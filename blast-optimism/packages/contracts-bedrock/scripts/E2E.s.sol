// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Script } from "forge-std/Script.sol";

import { console2 as console } from "forge-std/console2.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { Executables } from "scripts/Executables.sol";

import { Deployer } from "scripts/Deployer.sol";

import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { AddressManager } from "src/legacy/AddressManager.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { L1StandardBridge } from "src/L1/L1StandardBridge.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { USDB } from "src/L2/USDB.sol";
import { Shares } from "src/L2/Shares.sol";
import { Blast, YieldMode, GasMode } from "src/L2/Blast.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";
import { L1BlastBridge } from "src/mainnet-bridge/L1BlastBridge.sol";
import { ETHYieldManager } from "src/mainnet-bridge/ETHYieldManager.sol";
import { USDYieldManager } from "src/mainnet-bridge/USDYieldManager.sol";
import { L2BlastBridge } from "src/mainnet-bridge/L2BlastBridge.sol";
import { WithdrawalQueue } from "src/mainnet-bridge/withdrawal-queue/WithdrawalQueue.sol";
import { YieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { WETHRebasing } from "src/L2/WETHRebasing.sol";
import { L1ChugSplashProxy } from "src/legacy/L1ChugSplashProxy.sol";
import { ResolvedDelegateProxy } from "src/legacy/ResolvedDelegateProxy.sol";
import { L1CrossDomainMessenger } from "src/L1/L1CrossDomainMessenger.sol";
import { L2OutputOracle } from "src/L1/L2OutputOracle.sol";
import { OptimismMintableERC20Factory } from "src/universal/OptimismMintableERC20Factory.sol";
import { Insurance } from "src/mainnet-bridge/Insurance.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Chains } from "scripts/Chains.sol";
import { LidoYieldProvider } from "src/mainnet-bridge/yield-providers/LidoYieldProvider.sol";
import { DSRYieldProvider } from "src/mainnet-bridge/yield-providers/DSRYieldProvider.sol";
import { ETHTestnetYieldProvider } from "src/mainnet-bridge/yield-providers/ETHTestnetYieldProvider.sol";
import { USDTestnetYieldProvider } from "src/mainnet-bridge/yield-providers/USDTestnetYieldProvider.sol";
import { E2EInitializer } from "scripts/CommonE2E.s.sol";

contract E2E is E2EInitializer {
    // Blast contract stubs
    function configureContract(address contractAddress, YieldMode _yield, GasMode gasMode, address governor) public onlyL2 broadcast {
        b.configureContract(contractAddress, _yield, gasMode, governor);
    }

    function configure(YieldMode _yield, GasMode gasMode, address governor) public onlyL2 broadcast {
        b.configure(_yield, gasMode, governor);
    }

    function configureClaimableYield() public onlyL2 broadcast {
        b.configureClaimableYield();
    }

    function configureClaimableYieldOnBehalf(address contractAddress) public onlyL2 broadcast {
        b.configureClaimableYieldOnBehalf(contractAddress);
    }

    function configureAutomaticYield() public onlyL2 broadcast {
        b.configureAutomaticYield();
    }

    function configureAutomaticYieldOnBehalf(address contractAddress) public onlyL2 broadcast {
        b.configureAutomaticYieldOnBehalf(contractAddress);
    }

    function configureVoidYield() public onlyL2 broadcast {
        b.configureVoidYield();
    }

    function configureVoidYieldOnBehalf(address contractAddress) public onlyL2 broadcast {
        b.configureVoidYieldOnBehalf(contractAddress);
    }

    function configureClaimableGas() public onlyL2 broadcast {
        b.configureClaimableGas();
    }

    function configureClaimableGasOnBehalf(address contractAddress) public onlyL2 broadcast {
        b.configureClaimableGasOnBehalf(contractAddress);
    }

    function configureVoidGas() public onlyL2 broadcast {
        b.configureVoidGas();
    }

    function configureVoidGasOnBehalf(address contractAddress) public onlyL2 broadcast {
        b.configureVoidGasOnBehalf(contractAddress);
    }

    function configureGovernor(address _governor) public onlyL2 broadcast {
        b.configureGovernor(_governor);
    }

    function configureGovernorOnBehalf(address _newGovernor, address contractAddress) public onlyL2 broadcast {
        b.configureGovernorOnBehalf(_newGovernor, contractAddress);
    }

    function claimAllGas(address contractAddress, address recipientOfGas) public onlyL2 broadcast returns (uint256) {
        return b.claimAllGas(contractAddress, recipientOfGas);
    }

    function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) public onlyL2 broadcast returns (uint256) {
        return b.claimGasAtMinClaimRate(contractAddress, recipientOfGas, minClaimRateBips);
    }

    function claimMaxGas(address contractAddress, address recipientOfGas) public onlyL2 broadcast returns (uint256) {
        return b.claimMaxGas(contractAddress, recipientOfGas);
    }

    function claimGas(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) public onlyL2 broadcast returns (uint256) {
        return b.claimGas(contractAddress, recipientOfGas, gasToClaim, gasSecondsToConsume);
    }

    function readClaimableYield(address a) public onlyL2 broadcast returns (uint256) {
        return b.readClaimableYield(a);
    }

    function readYieldConfiguration(address contractAddress) external view returns (uint8) {
        return b.readYieldConfiguration(contractAddress);
    }

    function claimYield(address contractAddress, address beneficiary, uint256 amount) public onlyL2 broadcast returns (uint256) {
        return b.claimYield(contractAddress, beneficiary, amount);
    }

    function claimAllYield(address contractAddress, address beneficiary) public onlyL2 broadcast returns (uint256) {
        return b.claimAllYield(contractAddress, beneficiary);
    }

    function readGasParams(address contractAddress) external view returns (uint256 etherSeconds, uint256 etherBalance, uint256 lastUpdated, GasMode) {
        return b.readGasParams(contractAddress);
    }

    function depositETHDirect(uint256 amount) public onlyL1 broadcast {
        sendETH(address(op), amount);
    }
    function depositETH(uint256 amount) public onlyL1 broadcast {
        l1bb.bridgeETH{value: amount}(200_000, hex"");
    }
    function depositETHTo(address to, uint256 amount) public onlyL1 broadcast {
        l1bb.bridgeETHTo{value: amount}(to, 200_000, hex"");
    }

    function depositUSD(address token, uint256 amount) public onlyL1 broadcast {
        IERC20(token).approve(address(l1bb), type(uint256).max);
        bytes memory extraData = abi.encodePacked(uint256(amount));
        l1bb.bridgeERC20(token, address(usdb), amount, 100_000, extraData);
    }
    function depositUSDTo(address token, address to, uint256 amount) public onlyL1 broadcast {
        IERC20(token).approve(address(l1bb), type(uint256).max);
        bytes memory extraData = abi.encodePacked(uint256(amount));
        l1bb.bridgeERC20To(token, address(usdb), to, amount, 100_000, extraData);
    }

    function depositStakedETH(uint256 amount) public onlyL1 broadcast {
        steth.approve(address(l1bb), type(uint256).max);
        bytes memory extraData = abi.encodePacked(uint256(amount));
        l1bb.bridgeERC20(address(steth), address(0), amount, 100_000, extraData);
    }
    function depositStakedETHTo(address to, uint256 amount) public onlyL1 broadcast {
        steth.approve(address(l1bb), type(uint256).max);
        bytes memory extraData = abi.encodePacked(uint256(amount));
        l1bb.bridgeERC20To(address(steth), address(0), to, amount, 100_000, extraData);
    }

    function fakeUSDYield(int256 amount) public onlyL1 broadcast {
        if (_isFork()) {
            revert("Not supported");
        } else {
            usd.approve(address(uyp), uint256(amount));
            uyp.recordYield(amount);
        }
    }
    function fakeETHYield(int256 amount) public onlyL1 broadcast {
        if (_isFork()) {
            if (amount < 0) {
                lido.transferFrom(address(eym), address(this), uint256(-1 * amount));
            } else {
                lido.submit{value: uint256(amount)}(address(0));
                lido.transfer(address(eym), uint256(amount));
            }
        } else {
            if (amount < 0) {
                eyp.recordYield(amount);
            } else {
                eyp.recordYield{value: uint256(amount)}(amount);
            }
        }
    }

    function commitETHYieldReport(bool insurance) public onlyL1 {
        vm.startBroadcast(eym.admin());
        eym.commitYieldReport(insurance);
        vm.stopBroadcast();
    }
    function commitUSDYieldReport(bool insurance) public onlyL1 {
        vm.startBroadcast(uym.admin());
        uym.commitYieldReport(insurance);
        vm.stopBroadcast();
    }

    /*/////////////////////////////
            DEPOSIT ETH
    /////////////////////////////*/
    function testDepositETHDirect(uint256 amount) public {
        if (_chainIsL1()) {
            depositETHDirect(amount);

            require(user.balance == l1State.userETHBalance - amount, "User L1 ETH balance incorrect");
            require(address(eym).balance == l1State.eymBalance + amount, "ETHYieldManager balance incorrect");
        } else {
            require(user.balance == l2State.userETHBalance + amount, "User L2 ETH balance incorrect");
        }
    }

    function testDepositETH(uint256 amount) public {
        if (_chainIsL1()) {
            depositETH(amount);

            require(user.balance <= l1State.userETHBalance - amount, "User L1 ETH balance incorrect");
            require(address(eym).balance == l1State.eymBalance + amount, "ETHYieldManager balance incorrect");
        } else {
            require(user.balance == l2State.userETHBalance + amount, "User L2 ETH balance incorrect");
        }
    }

    function testDepositETHTo(uint256 amount) public {
        if (_chainIsL1()) {
            depositETHTo(alice, amount);

            require(user.balance <= l1State.userETHBalance - amount, "User L1 ETH balance incorrect");
            require(address(eym).balance == l1State.eymBalance + amount, "ETHYieldManager balance incorrect");
        } else {
            require(alice.balance == l2State.aliceETHBalance + amount, "Alice L2 ETH balance incorrect");
        }
    }

    function testDepositETHYield(uint256 amount) public {
        if (_chainIsL1()) {
            uint256 stakedPrincipal = eyp.stakedPrincipal();
            depositStakedETH(amount);

            require(steth.balanceOf(user) == l1State.userStETHBalance - amount, "User L1 ETH balance incorrect");
            require(eyp.stakedPrincipal() == stakedPrincipal + amount, "Staked principal is incorrect");
        } else {
            require(user.balance == l2State.userETHBalance + amount, "User L2 ETH balance incorrect");
        }
    }

    function testDepositETHYieldTo(uint256 amount) public {
        if (_chainIsL1()) {
            uint256 stakedPrincipal = eyp.stakedPrincipal();
            depositStakedETHTo(alice, amount);

            require(steth.balanceOf(user) == l1State.userStETHBalance - amount, "User L1 ETH balance incorrect");
            require(eyp.stakedPrincipal() == stakedPrincipal + amount, "Staked principal is incorrect");
        } else {
            require(alice.balance == l2State.aliceETHBalance + amount, "Alice L2 ETH balance incorrect");
        }
    }

    /*/////////////////////////////
            DEPOSIT USD
    /////////////////////////////*/
    function testDepositUSD(uint256 amount) public {
        if (_chainIsL1()) {
            vm.startBroadcast();
            usd.mint(msg.sender, amount);
            vm.stopBroadcast();
            depositUSD(address(usd), amount);

            require(usd.balanceOf(user) == l1State.userUSDBalance, "User USD balance incorrect");
            require(usd.balanceOf(address(l1bb)) == l1State.bridgeUSDBalance, "L1BlastBridge balance incorrect");
            require(usd.balanceOf(address(uym)) == l1State.uymUSDBalance + amount, "USDYieldManager balance incorrect");
        } else {
            require(usdb.balanceOf(user) == l2State.userUSDBBalance + amount, "User USDB balance incorrect");
        }
    }

    function testDepositUSDTo(uint256 amount) public {
        if (_chainIsL1()) {
            vm.startBroadcast();
            usd.mint(msg.sender, amount);
            vm.stopBroadcast();
            depositUSDTo(address(usd), alice, amount);

            require(usd.balanceOf(user) == l1State.userUSDBalance, "User USD balance incorrect");
            require(usd.balanceOf(address(l1bb)) == l1State.bridgeUSDBalance, "L1BlastBridge balance incorrect");
            require(usd.balanceOf(address(uym)) == l1State.uymUSDBalance + amount, "USDYieldManager balance incorrect");
        } else {
            require(usdb.balanceOf(alice) == l2State.aliceUSDBBalance + amount, "Alice USDB balance incorrect");
        }
    }

    function testDepositDAI(uint256 amount) public {
        if (_chainIsL1()) {
            depositUSD(address(dai), amount);

            require(dai.balanceOf(user) == l1State.userDAIBalance - amount, "User DAI balance incorrect");
            require(dai.balanceOf(address(l1bb)) == l1State.bridgeDAIBalance + amount, "L1BlastBridge balance incorrect");
        } else {
            require(usdb.balanceOf(user) == l2State.userUSDBBalance + amount, "User USDB balance incorrect");
        }
    }

    function testDepositUSDC(uint256 amount) public {
        if (_chainIsL1()) {
            usdc.approve(0x0A59649758aa4d66E25f08Dd01271e891fe52199, type(uint256).max);
            depositUSD(address(usdc), amount);

            uint256 amountWad = USDConversions._usdToWad(amount);
            require(usdc.balanceOf(user) == l1State.userUSDCBalance - amount, "User USDC balance incorrect");
            require(dai.balanceOf(address(l1bb)) == l1State.bridgeDAIBalance + amountWad, "L1BlastBridge balance incorrect");
        } else {
            uint256 amountWad = USDConversions._usdToWad(amount);
            require(usdb.balanceOf(user) == l2State.userUSDBBalance + amountWad, "User USDB balance incorrect");
        }
    }

    function testDepositUSDT(uint256 amount) public {
        if (_chainIsL1()) {
            // usdt.approve(0x0A59649758aa4d66E25f08Dd01271e891fe52199, type(uint256).max);
            depositUSD(address(usdt), amount);

            uint256 amountWad = USDConversions._usdToWad(amount);
            require(usdt.balanceOf(user) == l1State.userUSDTBalance - amount, "User USDT balance incorrect");
            require(dai.balanceOf(address(l1bb)) >= l1State.bridgeDAIBalance + amountWad, "L1BlastBridge balance incorrect");
        } else {
            uint256 amountWad = USDConversions._usdToWad(amount);
            require(usdb.balanceOf(user) == l2State.userUSDBBalance + amountWad, "User USDB balance incorrect");
        }
    }


    /*///////////////////////////////////////////////////////
                            ETH Yield
    ///////////////////////////////////////////////////////*/

    function testCommitYieldETH(uint256 yield) public {
        if (_chainIsL1()) {
            fakeETHYield(int256(yield));
            commitETHYieldReport(false);

            if (_isFork()) {
                require(lyp.yield() == 0, "LidoYieldProvider yield is not zero");
            } else {
                require(eyp.yield() == 0, "ETHYieldProvider yield is not zero");
            }
        } else {
            require((l2State.ethSharePrice - shares.price()) * shares.count() + shares.pending() - l2State.ethPending == yield, "ETH share price did not increase");
        }
    }

    function testCommitYieldETH_withInsurance(uint256 yield) public {
        if (_chainIsL1()) {
            fakeETHYield(int256(yield));
            commitETHYieldReport(true);

            if (_isFork()) {
                require(lyp.yield() == 0, "LidoYieldProvider yield is not zero");
                require(lido.balanceOf(address(ei)) == yield / 10, "Insurance not correct");
            } else {
                require(eyp.yield() == 0, "ETHYieldProvider yield is not zero");
                // require(uyt.balanceOf(address(ei)) == yield / 10, "Insurance not correct");
            }
        } else {
            require((l2State.ethSharePrice - shares.price()) * shares.count() + shares.pending() - l2State.ethPending == yield, "ETH share price did not increase");
        }
    }

    function testCommitYieldETH_paysOffDebt(uint256 yield) public {
        if (_chainIsL1()) {
            fakeETHYield(int256(yield));
            commitETHYieldReport(false);

            fakeETHYield(-1 * int256(yield));
            commitETHYieldReport(false);

            fakeETHYield(int256(yield));
            commitETHYieldReport(false);

            require(lyp.yield() == 0, "ETHYieldManager yield is not zero");
            require(eym.accumulatedNegativeYields() == 0, "ETHYieldManager negative yield is not zero");
        } else {
            require(l2State.ethSharePrice == shares.price() && shares.pending() == l2State.ethPending, "ETH share price changed");
        }
    }

    function testCommitYieldETH_negative(uint256 yield) public {
        if (_chainIsL1()) {
            fakeETHYield(-1 * int256(yield));
            commitETHYieldReport(false);

            if (_isFork()) {
                require(lyp.yield() == 0, "LidoYieldProvider yield is not zero");
            } else {
                require(eyp.yield() == 0, "ETHYieldProvider yield is not zero");
            }
            require(eym.accumulatedNegativeYields() == yield, "ETHYieldManager negative yields is incorrect");
        } else {
            require(l2State.ethSharePrice == shares.price() && shares.pending() == l2State.ethPending, "ETH share price changed");
        }
    }

    function testCommitYieldETH_negative_takeFromInsurance(uint256 yield) public {
        if (_chainIsL1()) {
            fakeETHYield(int256(yield));
            commitETHYieldReport(true);

            fakeETHYield(-1 * int256(yield)/10);
            commitETHYieldReport(true);

            if (_isFork()) {
                require(eyp.yield() == 0, "ETHYieldProvider yield is not zero");
            } else {
                require(lyp.yield() == 0, "LidoYieldProvider yield is not zero");
            }
            require(eym.accumulatedNegativeYields() == 0, "ETHYieldManager negative yield is not zero");
        } else {
            require(l2State.ethSharePrice == shares.price() && shares.pending() == l2State.ethPending, "ETH share price changed");
        }
    }

    /*/////////////////////
          USD Yield
    /////////////////////*/
    function testCommitYieldUSD(uint256 yield) public {
        if (_chainIsL1()) {
            fakeUSDYield(int256(yield));
            commitUSDYieldReport(false);
            if (_isFork()) {
                require(dyp.yield() == 0, "DSRYieldManager yield is not zero");
            } else {
                require(uyp.yield() == 0, "USDYieldManager yield is not zero");
            }
        } else {
            require(l2State.usdSharePrice < usdb.price(), "USD share price did not increase");
        }
    }

    /*/////////////////////////////////////
          Claiming Yield on BLAST
    /////////////////////////////////////*/

    // expects l2 -> l1 -> l2
    // blast e2e devnet 212 "function testIncreaseSharePriceAndClaimYield()"
    function testIncreaseSharePriceAndClaimYield() public {
        if (_chainIsL1()) {
            fakeETHYield(int256(100*1e18));
            commitETHYieldReport(false);
        } else if (step.t == 0) {
            configureClaimableYield();
            require(user.balance > 0, "user balance must be > 0");
        } else {
            uint256 yield = readClaimableYield(msg.sender);
            require(yield > 0, "yield must be non-negative");
            uint256 bal = user.balance;
            uint256 yieldClaimed = claimAllYield(user, user);
            require(yieldClaimed == yield, "yield claimed must be identical");
            require(user.balance - bal == yieldClaimed, "yield was claimed");
        }
    }
}
