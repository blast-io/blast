// SPDX-License-Identifier: MIT
pragma solidity ^0.8.3;

address constant GAS_PREDEPLOY_ADDR = 0x4300000000000000000000000000000000000002;

enum YieldMode {
    AUTOMATIC,
    VOID,
    CLAIMABLE
}

enum GasMode {
    VOID,
    CLAIMABLE
}

interface IBlast {
    // configure
    function configureContract(address contractAddress, YieldMode _yield, GasMode gasMode, address governor) external;
    function configure(YieldMode _yield, GasMode gasMode, address governor) external;

    // base configuration options
    function configureClaimableYield() external;
    function configureClaimableYield(address contractAddress) external;
    function configureAutomaticYield() external;
    function configureAutomaticYield(address contractAddress) external;
    function configureVoidYield() external;
    function configureVoidYield(address contractAddress) external;
    function configureClaimableGas() external;
    function configureClaimableGas(address contractAddress) external;
    function configureVoidGas() external;
    function configureVoidGas(address contractAddress) external;
    function configureGovernor(address _governor) external;
    function configureGovernor(address _newGovernor, address contractAddress) external;

    // claim yield
    function claimYield(address contractAddress, address recipientOfYield, uint256 amount) external returns (uint256);
    function claimAllYield(address contractAddress, address recipientOfYield) external returns (uint256);
    // claim gas
    function claimAllGas(address contractAddress, address recipientOfGas) external returns (uint256);
    function claimGasAtMinClaimRate(
        address contractAddress,
        address recipientOfGas,
        uint256 minClaimRateBips
    ) external returns (uint256);
    function claimMaxGas(address contractAddress, address recipientOfGas) external returns (uint256);
    function claimGas(
        address contractAddress,
        address recipientOfGas,
        uint256 gasToClaim,
        uint256 gasSecondsToConsume
    ) external returns (uint256);

    // read functions
    function readClaimableYield(address contractAddress) external view returns (uint256);
    function readYieldConfiguration(address contractAddress) external view returns (uint8);
    function readGasParams(address contractAddress) external view returns (uint256, uint256, uint256, GasMode);
}

contract Worker {
    constructor() {
        IBlast(GAS_PREDEPLOY_ADDR).configureClaimableYield();
        IBlast(GAS_PREDEPLOY_ADDR).configureClaimableGas();
    }

    function burn_gas() external {
        uint256 i = 0;
        while (gasleft() > 100) {
            ++i;
        }
    }
}
