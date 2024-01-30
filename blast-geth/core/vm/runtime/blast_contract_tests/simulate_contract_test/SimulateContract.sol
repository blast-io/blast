// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.20 and less than 0.9.0
pragma solidity ^0.8.15;

address constant ADDR = 0x4300000000000000000000000000000000000002;
enum YieldMode {
    AUTOMATIC,
    VOID,
    CLAIMABLE 
}

enum GasMode {
    VOID,
    CLAIMABLE 
}

interface IBlast{
    // configure
    function configureContract(address contractAddress, YieldMode _yield, GasMode gasMode, address governor) external;
    function configure(YieldMode _yield, GasMode gasMode, address governor) external;

    // base configuration options
    function configureClaimableYield() external;
    function configureClaimableYieldOnBehalf(address contractAddress) external;
    function configureAutomaticYield() external;
    function configureAutomaticYieldOnBehalf(address contractAddress) external;
    function configureVoidYield() external;
    function configureVoidYieldOnBehalf(address contractAddress) external;
    function configureClaimableGas() external;
    function configureClaimableGasOnBehalf(address contractAddress) external;
    function configureVoidGas() external;
    function configureVoidGasOnBehalf(address contractAddress) external;
    function configureGovernor(address _governor) external;
    function configureGovernorOnBehalf(address _newGovernor, address contractAddress) external;

    // claim yield
    function claimYield(address contractAddress, address recipientOfYield, uint256 amount) external returns (uint256);
    function claimAllYield(address contractAddress, address recipientOfYield) external returns (uint256);
    // claim gas
    function claimAllGas(address contractAddress, address recipientOfGas) external returns (uint256);
    function claimGasAtMinClaimRate(address contractAddress, address recipientOfGas, uint256 minClaimRateBips) external returns (uint256);
    function claimMaxGas(address contractAddress, address recipientOfGas) external returns (uint256);
    function claimGas(address contractAddress, address recipientOfGas, uint256 gasToClaim, uint256 gasSecondsToConsume) external returns (uint256);

    // read functions
    function readClaimableYield(address contractAddress) external view returns (uint256);
    function readYieldConfiguration(address contractAddress) external view returns (uint8);
    function readGasParams(address contractAddress) external view returns (uint256, uint256, uint256, GasMode);
}

contract SimulateContract {
    struct StateAccount {
        uint256 shares;
        uint256 remainder;
        uint256 fixedAmount;
        uint8 flags;
    }
    uint256 sharePrice;
    mapping(address => StateAccount) public stateMap;


    constructor() payable {
        IBlast(ADDR).configureClaimableYield();
        IBlast(ADDR).configureClaimableGas();

        sharePrice = 10;
        stateMap[ADDR].shares = 1;
        stateMap[ADDR].remainder = 1;
        stateMap[ADDR].fixedAmount = 1;
        stateMap[ADDR].flags = 2;
    }

    function claimYieldSimulator(address contractAddress, address recipient, uint256 amount) public payable returns (uint256) {
        StateAccount storage account = stateMap[contractAddress];
        if(account.flags != 2) {
            return 0;
        }
        uint256 shareValue = account.shares * sharePrice + account.remainder;
        uint256 claimableAmount = shareValue - account.fixedAmount;
        if (claimableAmount < amount) {
            amount = claimableAmount;
        }
        uint256 newShareValue = shareValue - amount;
        account.shares = newShareValue / sharePrice;
        account.remainder = newShareValue % sharePrice;
        payable(recipient).transfer(amount);
        return amount;
    }

    function setConfigurationSimulator(address contractAddress, uint8 flags) external {
        StateAccount storage account = stateMap[contractAddress];
        account.flags = flags;
    }

    function readConfigurationSimulator(address contractAddress) external view returns (uint8) {
        StateAccount memory account = stateMap[contractAddress];
        return account.flags;
    }

    // 12140 = 12k
    function readClaimableYieldSimulator(address contractAddress) external view returns (uint256) {
        StateAccount memory account = stateMap[contractAddress];
        if(account.flags != 2) {
            return 0;
        }
        uint256 shareValue = account.shares * sharePrice + account.remainder;
        uint256 claimableAmount = shareValue - account.fixedAmount;
        return claimableAmount;
    }
}

