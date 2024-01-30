// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

import { ERC20Rebasing } from "src/L2/ERC20Rebasing.sol";
import { SharesBase } from "src/L2/Shares.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { IOptimismMintableERC20 } from "src/universal/IOptimismMintableERC20.sol";
import { Semver } from "src/universal/Semver.sol";
import { Blast, YieldMode, GasMode } from "src/L2/Blast.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/// @custom:proxied
/// @custom:predeploy 0x4300000000000000000000000000000000000003
/// @title USDB
/// @notice Rebasing ERC20 token with the share price determined by an L1
///         REPORTER. Conforms to OptimismMintableERC20 interface to allow mint/burn
///         interactions from the L1BlastBridge.
contract USDB is ERC20Rebasing, Semver, IOptimismMintableERC20 {
    /// @notice Address of the corresponding version of this token on the remote chain.
    address public immutable REMOTE_TOKEN;

    /// @notice Address of the BlastBridge on this network.
    address public immutable BRIDGE;

    /// @notice Emitted whenever tokens are minted for an account.
    /// @param account Address of the account tokens are being minted for.
    /// @param amount  Amount of tokens minted.
    event Mint(address indexed account, uint256 amount);

    /// @notice Emitted whenever tokens are burned from an account.
    /// @param account Address of the account tokens are being burned from.
    /// @param amount  Amount of tokens burned.
    event Burn(address indexed account, uint256 amount);

    error CallerIsNotBridge();

    /// @notice A modifier that only allows the bridge to call
    modifier onlyBridge() {
        if (msg.sender != BRIDGE) {
            revert CallerIsNotBridge();
        }
        _;
    }

    /// @custom:semver 1.0.0
    /// @param _usdYieldManager Address of the USD Yield Manager. SharesBase yield reporter.
    /// @param _l2Bridge        Address of the L2 Blast bridge.
    /// @param _remoteToken     Address of the corresponding L1 token.
    constructor(address _usdYieldManager, address _l2Bridge, address _remoteToken)
        ERC20Rebasing(_usdYieldManager, 18)
        Semver(1, 0, 0)
    {
        BRIDGE = _l2Bridge;
        REMOTE_TOKEN = _remoteToken;
        _disableInitializers();
    }

    /// @notice Initializer
    function initialize() public initializer {
        __ERC20Rebasing_init("Rebasing USD", "USDB", 1e9);
        Blast(Predeploys.BLAST).configureContract(
            address(this),
            YieldMode.VOID,
            GasMode.VOID,
            address(0xdead) /// don't set a governor
        );
    }

    /// @notice ERC165 interface check function.
    /// @param _interfaceId Interface ID to check.
    /// @return Whether or not the interface is supported by this contract.
    function supportsInterface(bytes4 _interfaceId) external pure returns (bool) {
        bytes4 iface1 = type(IERC165).interfaceId;
        // Interface corresponding to the updated OptimismMintableERC20.
        bytes4 iface2 = type(IOptimismMintableERC20).interfaceId;
        return _interfaceId == iface1 || _interfaceId == iface2;
    }

    /// @custom:legacy
    /// @notice Legacy getter for REMOTE_TOKEN.
    function remoteToken() public view returns (address) {
        return REMOTE_TOKEN;
    }

    /// @custom:legacy
    /// @notice Legacy getter for BRIDGE.
    function bridge() public view returns (address) {
        return BRIDGE;
    }

    /// @notice Allows the StandardBridge on this network to mint tokens.
    /// @param _to     Address to mint tokens to.
    /// @param _amount Amount of tokens to mint.
    function mint(address _to, uint256 _amount)
        external
        virtual
        onlyBridge
    {
        if (_to == address(0)) {
            revert TransferToZeroAddress();
        }

        _deposit(_to, _amount);
        emit Mint(_to, _amount);
    }

    /// @notice Allows the StandardBridge on this network to burn tokens.
    /// @param _from   Address to burn tokens from.
    /// @param _amount Amount of tokens to burn.
    function burn(address _from, uint256 _amount)
        external
        virtual
        onlyBridge
    {
        if (_from == address(0)) {
            revert TransferFromZeroAddress();
        }

        _withdraw(_from, _amount);
        emit Burn(_from, _amount);
    }
}
