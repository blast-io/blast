// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

import { Predeploys } from "src/libraries/Predeploys.sol";
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { L2BlastBridge } from "src/mainnet-bridge/L2BlastBridge.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { OptimismPortal } from "src/L1/OptimismPortal.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { USDYieldManager } from "src/mainnet-bridge/YieldManager.sol";
import { ETHYieldManager } from "src/mainnet-bridge/ETHYieldManager.sol";
import { USDB } from "src/L2/USDB.sol";
import { USDConversions } from "src/mainnet-bridge/USDConversions.sol";

/// @custom:proxied
/// @title L1BlastBridge
/// @notice The L1BlastBridge is responsible for transfering ETH and yield-bearing ERC20 tokens between L1 and L2.
///         It is also responsible for managing and coordinating the staking, unstaking, and yield reporting
///         of bridged USD tokens that are eligible for yield.
///
///         The current implementation converts all deposited USD tokens to DAI before bridging them to L2 to mint USDB.
///         Hence, the amount of USDB that is minted on L2 will be equal to the amount of DAI that is received on L1.
///         This is done to simplify the yield management, as DSR is the only yield provider that is currently supported.
///         When non-DAI USD tokens are deposited, the user is expected to provide the minimum amount of DAI that should
///         be received (i.e. the minimum amount of USDB that should be minted on L2). This amount must be specified
///         in the extraData field of the deposit transaction (uin256 minAmountInWad).
contract L1BlastBridge is StandardBridge, USDYieldManager, ISemver {
    struct YieldToken {
        bool approved;
        uint8 decimals;
        address provider;
        bool reportStakedBalance;
    }

    /// @notice Mapping of potential deposit tokens to whether they're
    ///         approved as ETH yield tokens and additional metadata.
    mapping(address => YieldToken) public ethYieldTokens;

    /// @notice Mapping of potential deposit tokens to whether they're
    ///         approved as USD yield tokens and additional metadata.
    mapping(address => YieldToken) public usdYieldTokens;

    /// @notice Address of the ETH Yield Manager.
    ETHYieldManager public ethYieldManager;

    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @custom:semver 1.0.0
    /// @notice Constructs the L1BlastBridge contract.
    constructor() StandardBridge(StandardBridge(payable(Predeploys.L2_BLAST_BRIDGE))) {
        initialize({
            _portal: OptimismPortal(payable(address(0))),
            _messenger: CrossDomainMessenger(address(0)),
            _yieldManager: ETHYieldManager(payable(address(0)))
        });
    }

    /// @notice Initializer
    /// @param _portal       Address of the OptimismPortal.
    /// @param _messenger    Address of the L1CrossDomainMessenger.
    /// @param _yieldManager Address of the ETHYieldManager.
    function initialize(OptimismPortal _portal, CrossDomainMessenger _messenger, ETHYieldManager _yieldManager) public initializer {
        ethYieldManager = _yieldManager;
        __StandardBridge_init(_messenger);
        __YieldManager_init(_portal);
        USDConversions._init();
    }

    /// @notice Add/remove an approved USD yield token.
    /// @param token               Address of token.
    /// @param approved            Whether the token is an approved yield token.
    /// @param decimals            Number of token decimals.
    /// @param provider            Address of the yield provider for the token.
    /// @param reportStakedBalance Whether a deposit needs to be reported to the yield provider.
    function setUSDYieldToken(address token, bool approved, uint8 decimals, address provider, bool reportStakedBalance) external onlyOwner {
        usdYieldTokens[token] = YieldToken({
            approved: approved,
            decimals: decimals,
            provider: provider,
            reportStakedBalance: reportStakedBalance
        });
    }

    /// @notice Add/remove an approved ETH yield token.
    /// @param token               Address of token.
    /// @param approved            Whether the token is an approved yield token.
    /// @param decimals            Number of token decimals.
    /// @param provider            Address of the yield provider for the token.
    /// @param reportStakedBalance Whether a deposit needs to be reported to the yield provider.
    function setETHYieldToken(address token, bool approved, uint8 decimals, address provider, bool reportStakedBalance) external onlyOwner {
        require(token != address(0));
        ethYieldTokens[token] = YieldToken({
            approved: approved,
            decimals: decimals,
            provider: provider,
            reportStakedBalance: reportStakedBalance
        });
    }

    /// @notice Allows EOAs to bridge ETH by sending directly to the bridge.
    receive() external payable override {
        _initiateBridgeETH(msg.sender, msg.sender, msg.value, RECEIVE_DEFAULT_GAS_LIMIT, hex"");
    }

    /// @notice Finalizes an ERC20 bridge on this chain. Can only be triggered by the other
    ///         StandardBridge contract on the remote chain.
    /// @param _localToken  Address of the ERC20 on this chain.
    /// @param _remoteToken Address of the ERC20 on this chain.
    /// @param _from        Address of the sender.
    /// @param _to          Address of the receiver.
    /// @param _amount      Amount of the ERC20 being bridged.
    /// @param _extraData   Extra data to be sent with the transaction. Note that the recipient will
    ///                     not be triggered with this data, but it will be emitted and can be used
    ///                     to identify the transaction.
    function finalizeBridgeERC20(
        address _localToken,
        address _remoteToken,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _extraData
    )
        public
        override
        onlyOtherBridge
    {
        require(_to != address(this), "StandardBridge: cannot send to self");
        require(_to != address(messenger), "StandardBridge: cannot send to messenger");

        require(_localToken == TOKEN, "L1BlastBridge: only DAI can be received through this bridge");
        require(_remoteToken == Predeploys.USDB, "L1BlastBridge: only USDB can be withdrawn through this bridge");

        // Emit the correct events. By default this will be ERC20BridgeFinalized, but child
        // contracts may override this function in order to emit legacy events as well.
        _emitERC20BridgeFinalized(_localToken, _remoteToken, _from, _to, _amount, _extraData);

        _requestWithdrawal(_to, _amount);
    }

    /// @notice Sends approved yield-bearing ERC20 tokens to a receiver's address on the other chain.
    ///         Only USDB or ETH are accepted as _remoteToken. ETH-based tokens are sent to the
    ///         Optimism Portal.
    /// @param _localToken  Address of the ERC20 on this chain.
    /// @param _remoteToken Address of the corresponding token on the remote chain.
    /// @param _to          Address of the receiver.
    /// @param _amount      Amount of local tokens to deposit.
    /// @param _minGasLimit Minimum amount of gas that the bridge can be relayed with.
    /// @param _extraData   Extra data to be sent with the transaction. Note that the recipient will
    ///                     not be triggered with this data, but it will be emitted and can be used
    ///                     to identify the transaction. For bridging yield-bearing USD tokens (except DAI),
    ///                     the extra data should contain the minimum amount of tokens in WAD to be minted.
    ///                     When the deposited USD tokens are converted to DAI, it ensures that the amount
    ///                     of DAI received (and hence the amount of USDB minted) is at least the minimum
    ///                     amount specified.
    function _initiateBridgeERC20(
        address _localToken,
        address _remoteToken,
        address _from,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes memory _extraData
    )
        internal
        override
    {
        YieldToken memory usdYieldToken = usdYieldTokens[_localToken];
        YieldToken memory ethYieldToken = ethYieldTokens[_localToken];
        if (usdYieldToken.approved) {
            require(_remoteToken == Predeploys.USDB, "L1BlastBridge: this token can only be bridged to USDB");
            IERC20(_localToken).transferFrom(_from, address(this), _amount);

            uint256 amountWad = USDConversions._convertDecimals(_amount, usdYieldToken.decimals, USDConversions.WAD_DECIMALS);
            uint256 amountToMintWad = USDConversions._convertToDAI(
                _localToken,
                amountWad,
                _extraData
            );

            // Update the yield provider with the staked deposit.
            if (usdYieldToken.reportStakedBalance) {
                require(usdYieldToken.provider != address(0));
                _recordStakedDeposit(usdYieldToken.provider, amountToMintWad);
            }

            messenger.sendMessage(
                Predeploys.L2_BLAST_BRIDGE,
                abi.encodeWithSelector(
                    StandardBridge.finalizeBridgeERC20.selector,
                    address(USDConversions.DAI),
                    Predeploys.USDB,
                    _from,
                    _to,
                    amountToMintWad,
                    _extraData
                ),
                _minGasLimit
            );

            // Emit the correct events. By default this will be ERC20BridgeInitiated, but child
            // contracts may override this function in order to emit legacy events as well.
            _emitERC20BridgeInitiated(address(USDConversions.DAI), _remoteToken, _from, _to, amountToMintWad, _extraData);
        } else if (ethYieldToken.approved) {
            require(_remoteToken == address(0), "L1BlastBridge: this token can only be bridge to ETH");
            IERC20(_localToken).transferFrom(_from, address(ethYieldManager), _amount);

            // Update the yield provider with the staked deposit.
            if (ethYieldToken.reportStakedBalance) {
                require(ethYieldToken.provider != address(0));
                ethYieldManager.recordStakedDeposit(ethYieldToken.provider, _amount);
            }

            // Message has to be sent to the OptimismPortal directly because we have to
            // request the L2 message has value without sending ETH.
            portal.depositTransaction(
                Predeploys.L2_BLAST_BRIDGE,
                _amount,
                RECEIVE_DEFAULT_GAS_LIMIT,
                false,
                abi.encodeWithSelector(
                    L2BlastBridge.finalizeBridgeETHDirect.selector,
                    _from,
                    _to,
                    USDConversions._convertDecimals(_amount, ethYieldToken.decimals, USDConversions.WAD_DECIMALS),
                    _extraData
                )
            );

            // Emit the correct events. By default this will be ERC20BridgeInitiated, but child
            // contracts may override this function in order to emit legacy events as well.
            _emitERC20BridgeInitiated(_localToken, _remoteToken, _from, _to, _amount, _extraData);
        } else {
            revert("L1BlastBridge: bridge token is not supported");
        }
    }
}
