// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Vm } from "forge-std/Test.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { Executables } from "scripts/Executables.sol";

import { Types } from "src/libraries/Types.sol";
import { Hashing } from "src/libraries/Hashing.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";
import { E2EInitializer } from "scripts/Scripts.s.sol";

contract WithdrawE2E is E2EInitializer {
    event WithdrawalRequested(
        uint256 indexed requestId, address indexed requestor, address indexed recipient, uint256 amount
    );

    event MessagePassed(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data,
        bytes32 withdrawalHash
    );

    function _loadWithdrawalTx()
        internal
        view
        returns (Types.WithdrawalTransaction memory _tx, uint32 l2BlockNumber, uint256 requestId)
    {
        string memory json = vm.readFile("withdrawal.json");
        l2BlockNumber = uint32(vm.parseJsonUint(json, "$.l2BlockNumber"));
        requestId = vm.parseJsonUint(json, "$.requestId");
        _tx = Types.WithdrawalTransaction({
            nonce: vm.parseJsonUint(json, "$.nonce"),
            sender: vm.parseJsonAddress(json, "$.sender"),
            target: vm.parseJsonAddress(json, "$.target"),
            value: vm.parseJsonUint(json, "$.value"),
            gasLimit: vm.parseJsonUint(json, "$.gasLimit"),
            data: vm.parseJsonBytes(json, "$.data")
        });
    }

    function _storeWithdrawalTx(
        Types.WithdrawalTransaction memory _tx,
        uint32 l2BlockNumber,
        uint256 requestId
    )
        internal
    {
        string memory json = "";
        vm.serializeUint(json, "l2BlockNumber", l2BlockNumber);
        vm.serializeUint(json, "nonce", _tx.nonce);
        vm.serializeAddress(json, "sender", _tx.sender);
        vm.serializeAddress(json, "target", _tx.target);
        vm.serializeUint(json, "value", _tx.value);
        vm.serializeUint(json, "gasLimit", _tx.gasLimit);
        vm.serializeBytes(json, "data", _tx.data);
        json = vm.serializeUint(json, "requestId", requestId);
        vm.writeJson({ json: json, path: "withdrawal.json" });
    }

    function _proveWithdrawal(Types.WithdrawalTransaction memory _tx, uint32 l2BlockNumber) internal {
        bytes32 messageHash = Hashing.hashWithdrawal(_tx);
        bytes32 slot = keccak256(abi.encode(messageHash, uint256(0)));
        uint256 l2OutputIndex =
            oracle.getL2OutputIndexAfter(l2BlockNumber);
        uint32 l2BlockNumber2 =
            uint32(oracle.getL2Output(l2OutputIndex).l2BlockNumber);
        (bytes32 stateRoot, bytes32 blockHash, bytes32 storageHash, bytes[] memory proof) =
            _getWithdrawalProof(slot, l2BlockNumber2);
        op.proveWithdrawalTransaction(
            _tx,
            l2OutputIndex,
            Types.OutputRootProof({
                version: bytes32(0),
                stateRoot: stateRoot,
                messagePasserStorageRoot: storageHash,
                latestBlockhash: blockHash
            }),
            proof
        );
    }

    function _getWithdrawalProof(
        bytes32 slot,
        uint32 l2BlockNumber
    )
        internal
        returns (bytes32 stateRoot, bytes32 blockHash, bytes32 storageHash, bytes[] memory proof)
    {
        string[] memory cmd = new string[](3);
        cmd[0] = Executables.bash;
        cmd[1] = "-c";
        cmd[2] = string.concat(
            "./scripts/proof.sh ", Strings.toHexString(uint256(slot)), " ", Strings.toString(l2BlockNumber), " | jq"
        );
        string memory res = string(vm.ffi(cmd));
        stateRoot = stdJson.readBytes32(res, "$.stateRoot");
        blockHash = stdJson.readBytes32(res, "$.hash");
        storageHash = stdJson.readBytes32(res, "$.storageHash");
        proof = stdJson.readBytesArray(res, "$.proof");
    }
}

contract WithdrawE2ETest is WithdrawE2E {
    function check() external view override returns (bool) {
        (Types.WithdrawalTransaction memory _tx, uint32 l2BlockNumber,) = _loadWithdrawalTx();
        bytes32 hash = Hashing.hashWithdrawal(_tx);
        (, uint128 timestamp,,) = op.provenWithdrawals(hash);
        if (timestamp == 0) {
            return oracle.latestBlockNumber() >= l2BlockNumber;
        } else {
            uint256 l2OutputIndex = oracle.getL2OutputIndexAfter(l2BlockNumber);
            return op.isOutputFinalized(l2OutputIndex);
        }
    }

    function testWithdrawETH(uint256 amount) public {
        if (_chainIsL1()) {
            (Types.WithdrawalTransaction memory _tx, uint32 l2BlockNumber, uint256 requestId) = _loadWithdrawalTx();
            if (l1State.numCalls == 0) {
                _incrementL1NumCalls();
                vm.startBroadcast();
                vm.recordLogs();
                _proveWithdrawal(_tx, l2BlockNumber);
                Vm.Log memory log = _getLog(vm.getRecordedLogs(), WithdrawalRequested.selector);
                _storeWithdrawalTx(_tx, l2BlockNumber, uint256(log.topics[1]));
                vm.stopBroadcast();
            } else {
                vm.startBroadcast();
                eym.finalize(requestId);
                op.finalizeWithdrawalTransaction{ gas: 1_000_000 }(
                    eym.findCheckpointHint(requestId, 1, eym.getLastCheckpointId()), _tx
                );
                vm.stopBroadcast();
            }
        } else {
            vm.startBroadcast();
            vm.recordLogs();
            l2bb.bridgeETH{ value: amount }(200_000, hex"");
            Vm.Log memory log = _getLog(vm.getRecordedLogs(), MessagePassed.selector);
            (uint256 value, uint256 gasLimit, bytes memory data,) =
                abi.decode(log.data, (uint256, uint256, bytes, bytes32));
            Types.WithdrawalTransaction memory _tx = Types.WithdrawalTransaction({
                nonce: uint256(log.topics[1]),
                sender: address(uint160(uint256(log.topics[2]))),
                target: address(uint160(uint256(log.topics[3]))),
                value: value,
                gasLimit: gasLimit,
                data: data
            });
            _storeWithdrawalTx(_tx, uint32(block.number), 0);
            vm.stopBroadcast();
        }
    }
}

contract WithdrawUSDE2ETest is WithdrawE2E {
    function check() external view override returns (bool) {
        (Types.WithdrawalTransaction memory _tx, uint32 l2BlockNumber,) = _loadWithdrawalTx();
        bytes32 hash = Hashing.hashWithdrawal(_tx);
        (, uint128 timestamp,,) = op.provenWithdrawals(hash);
        if (timestamp == 0) {
            return oracle.latestBlockNumber() >= l2BlockNumber;
        } else {
            uint256 l2OutputIndex = oracle.getL2OutputIndexAfter(l2BlockNumber);
            return op.isOutputFinalized(l2OutputIndex);
        }
    }

    function testWithdrawUSD(uint256 amount) public {
        if (_chainIsL1()) {
            (Types.WithdrawalTransaction memory _tx, uint32 l2BlockNumber,) = _loadWithdrawalTx();
            if (l1State.numCalls == 0) {
                _incrementL1NumCalls();
                vm.startBroadcast();
                _proveWithdrawal(_tx, l2BlockNumber);
                _storeWithdrawalTx(_tx, l2BlockNumber, 0);
                vm.stopBroadcast();
            } else if (l1State.numCalls == 1) {
                vm.startBroadcast();
                vm.recordLogs();
                op.finalizeWithdrawalTransaction{ gas: 1_000_000 }(0, _tx);
                Vm.Log memory log = _getLog(vm.getRecordedLogs(), WithdrawalRequested.selector);
                uint256 requestId = uint256(log.topics[1]);
                uym.finalize(requestId);
                uym.claimWithdrawal(requestId, uym.findCheckpointHint(requestId, 1, uym.getLastCheckpointId()));
                vm.stopBroadcast();
            }
        } else {
            vm.startBroadcast();
            vm.recordLogs();
            l2bb.bridgeERC20(address(usdb), usdb.REMOTE_TOKEN(), amount, 200_000, hex"");
            Vm.Log memory log = _getLog(vm.getRecordedLogs(), MessagePassed.selector);
            (uint256 value, uint256 gasLimit, bytes memory data,) =
                abi.decode(log.data, (uint256, uint256, bytes, bytes32));
            Types.WithdrawalTransaction memory _tx = Types.WithdrawalTransaction({
                nonce: uint256(log.topics[1]),
                sender: address(uint160(uint256(log.topics[2]))),
                target: address(uint160(uint256(log.topics[3]))),
                value: value,
                gasLimit: gasLimit,
                data: data
            });
            _storeWithdrawalTx(_tx, uint32(block.number), 0);
            vm.stopBroadcast();
        }
    }
}

contract DepositE2E is E2EInitializer {
    function depositETHDirect(uint256 amount) public onlyL1 broadcast {
        sendETH(address(op), amount);
    }

    function depositETH(uint256 amount) public onlyL1 broadcast {
        l1bb.bridgeETH{ value: amount }(200_000, hex"");
    }

    function depositETHTo(address to, uint256 amount) public onlyL1 broadcast {
        l1bb.bridgeETHTo{ value: amount }(to, 200_000, hex"");
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
}

contract TestDepositE2E is DepositE2E {
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
            require(
                dai.balanceOf(address(l1bb)) == l1State.bridgeDAIBalance + amount, "L1BlastBridge balance incorrect"
            );
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
            require(
                dai.balanceOf(address(l1bb)) == l1State.bridgeDAIBalance + amountWad, "L1BlastBridge balance incorrect"
            );
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
            require(
                dai.balanceOf(address(l1bb)) >= l1State.bridgeDAIBalance + amountWad, "L1BlastBridge balance incorrect"
            );
        } else {
            uint256 amountWad = USDConversions._usdToWad(amount);
            require(usdb.balanceOf(user) == l2State.userUSDBBalance + amountWad, "User USDB balance incorrect");
        }
    }
}

contract YieldE2E is E2EInitializer {
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
                lido.submit{ value: uint256(amount) }(address(0));
                lido.transfer(address(eym), uint256(amount));
            }
        } else {
            if (amount < 0) {
                eyp.recordYield(amount);
            } else {
                eyp.recordYield{ value: uint256(amount) }(amount);
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
}

contract YieldE2ETest is YieldE2E {
    /*//////////////////////////////////
                ETH Yield
    //////////////////////////////////*/

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
            require(
                (l2State.ethSharePrice - shares.price()) * shares.count() + shares.pending() - l2State.ethPending
                    == yield,
                "ETH share price did not increase"
            );
        }
    }

    function testCommitYieldETH_withInsurance(uint256 yield) public {
        if (_chainIsL1()) {
            fakeETHYield(int256(yield));
            commitETHYieldReport(true);

            if (_isFork()) {
                require(lyp.yield() == 0, "LidoYieldProvider yield is not zero");
                require(lido.balanceOf(address(ei)) == yield * eym.insuranceFeeBips() / 10_000, "Insurance not correct");
            } else {
                require(eyp.yield() == 0, "ETHYieldProvider yield is not zero");
                require(
                    steth.balanceOf(address(ei)) == yield * uym.insuranceFeeBips() / 10_000, "Insurance not correct"
                );
            }
        } else {
            require(
                (l2State.ethSharePrice - shares.price()) * shares.count() + shares.pending() - l2State.ethPending
                    == yield,
                "ETH share price did not increase"
            );
        }
    }

    function testCommitYieldETH_paysOffDebt(uint256 yield) public {
        if (_chainIsL1()) {
            require(eym.accumulatedNegativeYields() > 0, "no debt to pay off");

            fakeETHYield(int256(yield));
            commitETHYieldReport(false);

            if (_isFork()) {
                require(lyp.yield() == 0, "ETHYieldManager yield is not zero");
            } else {
                require(eyp.yield() == 0, "ETHYieldManager yield is not zero");
            }
            require(eym.accumulatedNegativeYields() == 0, "ETHYieldManager negative yield is not zero");
        } else {
            require(
                l2State.ethSharePrice == shares.price() && shares.pending() == l2State.ethPending,
                "ETH share price changed"
            );
        }
    }

    function testCommitYieldETH_negative(uint256 yield) public {
        if (_chainIsL1()) {
            if (_isFork()) {
                require(lyp.stakedPrincipal() > yield, "Insufficient balance to simulate negative yield");
            } else {
                require(eyp.stakedPrincipal() > yield, "Insufficient balance to simulate negative yield");
            }
            fakeETHYield(-1 * int256(yield));
            commitETHYieldReport(false);

            if (_isFork()) {
                require(lyp.yield() == 0, "LidoYieldProvider yield is not zero");
            } else {
                require(eyp.yield() == 0, "ETHYieldProvider yield is not zero");
            }
            require(eym.accumulatedNegativeYields() == yield, "ETHYieldManager negative yields is incorrect");
        } else {
            require(
                l2State.ethSharePrice == shares.price() && shares.pending() == l2State.ethPending,
                "ETH share price changed"
            );
        }
    }

    function testCommitYieldETH_negative_takeFromInsurance(uint256 yield) public {
        if (_chainIsL1()) {
            if (_isFork()) {
                require(lyp.stakedPrincipal() > yield, "Insufficient balance to simulate negative yield");
                require(lyp.insuranceBalance() > yield, "Insufficient insurance to pay negative yield");
            } else {
                require(eyp.stakedPrincipal() > yield, "Insufficient balance to simulate negative yield");
                require(eyp.insuranceBalance() > yield, "Insufficient insurance to pay negative yield");
            }

            fakeETHYield(-1 * int256(yield));
            commitETHYieldReport(true);

            if (_isFork()) {
                require(eyp.yield() == 0, "ETHYieldProvider yield is not zero");
            } else {
                require(lyp.yield() == 0, "LidoYieldProvider yield is not zero");
            }
            require(eym.accumulatedNegativeYields() == 0, "ETHYieldManager negative yield is not zero");
        } else {
            require(
                l2State.ethSharePrice == shares.price() && shares.pending() == l2State.ethPending,
                "ETH share price changed"
            );
        }
    }

    /*//////////////////////////////////
                USD Yield
    //////////////////////////////////*/

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
}
