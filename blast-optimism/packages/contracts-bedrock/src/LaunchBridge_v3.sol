// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { Ownable2StepUpgradeable } from "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IERC20Permit } from "@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

interface ILido is IERC20, IERC20Permit {
    function submit(address user) external payable;
}

interface IDAI is IERC20 {
    function permit(
        address holder,
        address spender,
        uint256 nonce,
        uint256 expiry,
        bool allowed,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external;
    function nonces(address user) external view returns (uint256);
}

interface IUSDC is IERC20, IERC20Permit {
    function transferWithAuthorization(address, address, uint256, uint256, uint256, bytes32, uint8, bytes32, bytes32) external;
}

interface IUSDT {
    function transfer(address to, uint256 amount) external;
    function transferFrom(address from, address to, uint256 amount) external;
    function approve(address spender, uint256 amount) external;
    function basisPointsRate() external view returns (uint256);
    function balanceOf(address) external view returns (uint256);
}

interface IDsrManager {
    function join(address dst, uint256 wad) external;
    function exit(address dst, uint256 wad) external;
    function exitAll(address dst) external;
    function daiBalance(address usr) external returns (uint256 wad);
    function pot() external view returns (address);
    function pieOf(address) external view returns (uint256);
}

interface IDssPsm {
    function sellGem(address usr, uint256 gemAmt) external;
    function buyGem(address usr, uint256 gemAmt) external;
    function dai() external view returns (address);
    function gemJoin() external view returns (address);
    function tin() external view returns (uint256);
    function tout() external view returns (uint256);
}

interface IPot {
    function chi() external view returns (uint256);
    function rho() external view returns (uint256);
    function dsr() external view returns (uint256);
}

interface IMainnetBridge {
    function bridgeETHTo(address _to, uint32 _minGasLimit, bytes calldata _extraData) external payable;
    function bridgeERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes calldata _extraData) external payable;
}

interface ICurve3Pool {
    function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) external;
}

contract LaunchBridge is UUPSUpgradeable, Ownable2StepUpgradeable, PausableUpgradeable {
    mapping(address => uint256) public ethShares;
    uint256 public totalETHShares;

    mapping(address => uint256) public usdShares;
    uint256 public totalUSDShares;

    mapping(address => bool) public transitioned;
    bool public isTransitionEnabled;

    address public staker;

    IMainnetBridge internal _mainnetBridge;

    uint256 constant EMERGENCY_WITHDRAW_TIMESTAMP = 1717200000; // June 1, 2024

    ILido public constant LIDO = ILido(0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84);
    IUSDC public constant USDC = IUSDC(0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48);
    IUSDT public constant USDT = IUSDT(0xdAC17F958D2ee523a2206206994597C13D831ec7);
    IDsrManager public constant DSR_MANAGER = IDsrManager(0x373238337Bfe1146fb49989fc222523f83081dDb);
    IDssPsm public constant PSM = IDssPsm(0x89B78CfA322F6C5dE0aBcEecab66Aee45393cC5A);
    ICurve3Pool public constant CURVE_3POOL = ICurve3Pool(0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7);
    IDAI public constant DAI = IDAI(0x6B175474E89094C44Da98b954EedeAC495271d0F);

    uint256 internal constant _BASIS_POINTS = 10_000;
    address internal constant _INITIAL_TOKEN_HOLDER = 0x000000000000000000000000000000000000dEaD;
    uint256 internal constant _USD_DECIMALS = 6;
    uint256 internal constant _WAD_DECIMALS = 18;
    int128 internal constant _CURVE_USDT_INDEX = 2;
    int128 internal constant _CURVE_DAI_INDEX = 0;
    uint256 internal constant _WAD = 10 ** 18;
    uint256 internal constant _RAY = 10 ** 27;
    uint256 internal constant _INITIAL_DEPOSIT_AMOUNT = 1000;

    event ETHDeposited(address indexed user, uint256 shares, uint256 amount);
    event USDDeposited(address indexed user, uint256 shares, uint256 amount, uint256 daiAmount);
    event Withdraw(address indexed user, uint256 ethAmount, uint256 stETHAmount, uint256 daiAmount);

    error CallerIsNotStaker();
    error TransitionNotEnabled();
    error TransitionIsEnabled();
    error UserAlreadyTransitioned();
    error InsufficientFunds();
    error BridgeIsNotSet();
    error ZeroDeposit();
    error ZeroSharesIssued();
    error SharesNotInitiated();
    error InvalidRecipientSignature();
    error InvalidRecipient();
    error OnlyEOA();

    modifier onlyEOA() {
        if (msg.sender != tx.origin) {
            revert OnlyEOA();
        }
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(address _staker) external initializer {
        __UUPSUpgradeable_init();
        __Ownable2Step_init();
        __Pausable_init();

        _pause();

        staker = _staker;

        USDC.approve(PSM.gemJoin(), type(uint256).max);
        USDT.approve(address(CURVE_3POOL), type(uint256).max);
        DAI.approve(address(DSR_MANAGER), type(uint256).max);
    }

    function _authorizeUpgrade(address target) internal override onlyOwner {}

    /**
     * @notice Pause deposits to the bridge (admin only)
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause deposits to the bridge (admin only)
     */
    function unpause() external onlyOwner {
        if (totalETHShares == 0 && totalUSDShares == 0) {
            revert SharesNotInitiated();
        }
        if (isTransitionEnabled) {
            revert TransitionIsEnabled();
        }
        _unpause();
    }

    /**
     * @notice Set approved staker (admin only)
     * @param _staker New staker address
     */
    function setStaker(address _staker) public onlyOwner {
        staker = _staker;
    }

    /**
     * @notice Open bridge to accept deposits; accept initial ETH and DAI deposit to initiate the shares accounting (admin only)
     * @param from Initial depositor
     * @param nonce Permit signature nonce
     * @param v Permit signature v parameter
     * @param r Permit signature r parameter
     * @param s Permit signature s parameter
     */
    function open(address from, uint256 nonce, uint8 v, bytes32 r, bytes32 s) external payable onlyOwner {
        DAI.permit(
            from,
            address(this),
            nonce,
            type(uint256).max,
            true,
            v,
            r,
            s
        );
        DAI.transferFrom(
            from,
            address(this),
            _INITIAL_DEPOSIT_AMOUNT
        );

        uint256 ethBalance = address(this).balance;
        uint256 daiBalance = DAI.balanceOf(address(this));

        assert(totalETHShares == 0 && totalUSDShares == 0);
        assert(ethBalance >= _INITIAL_DEPOSIT_AMOUNT && daiBalance >= _INITIAL_DEPOSIT_AMOUNT);
        _mintETHShares(_INITIAL_TOKEN_HOLDER, ethBalance);
        _mintUSDShares(_INITIAL_TOKEN_HOLDER, daiBalance);

        _unpause();
    }

    /**
     * @notice Wrapper to get mainnet bridge
     */
    function getMainnetBridge() public view returns (IMainnetBridge mainnetBridge) {
        mainnetBridge = _mainnetBridge;
        if (address(mainnetBridge) == address(0)) {
            revert BridgeIsNotSet();
        }
    }

    /**
     * @notice Wrapper to set mainnet bridge
     */
    function _setMainnetBridge(address mainnetBridge) internal {
        assert(mainnetBridge.code.length > 0);
        _mainnetBridge = IMainnetBridge(mainnetBridge);
    }

    /**
     * @notice Get the user balance in ETH and USD pool
     * @dev Does not update DSR yield
     * @param user User address
     * @return ethBalance User's ETH balance, usdBalance User's USD balance
     */
    function balanceOf(address user) external view returns (uint256 ethBalance, uint256 usdBalance) {
        ethBalance = _ethByShares(ethShares[user]);
        usdBalance = _usdBySharesNoUpdate(usdShares[user]);
    }

    /**
     * @notice Get the current ETH pool balance
     * @return Pooled ETH balance between buffered balance and deposited Lido balance
     */
    function totalETHBalance() public view returns (uint256) {
        return address(this).balance + LIDO.balanceOf(address(this));
    }

    /**
     * @notice Get the current USD pool balance
     * @dev Does not update DSR yield
     * @return Pooled USD balance between buffered balance and deposited DSR balance
     */
    function totalUSDBalanceNoUpdate() public view returns (uint256) {
        IPot pot = IPot(DSR_MANAGER.pot());
        uint256 chi = _rmul(_rpow(pot.dsr(), block.timestamp - pot.rho(), _RAY), pot.chi());
        return DAI.balanceOf(address(this)) + _rmul(DSR_MANAGER.pieOf(address(this)), chi);
    }

    /**
     * @notice Get the current USD pool balance
     * @return Pooled USD balance between buffered balance and deposited DSR balance
     */
    function totalUSDBalance() public returns (uint256) {
        return DAI.balanceOf(address(this)) + DSR_MANAGER.daiBalance(address(this));
    }

    /*/////////////////////////
             DEPOSITS
    /////////////////////////*/

    receive() external payable {
        depositETH();
    }

    /**
     * @notice Deposit ETH to the ETH pool
     */
    function depositETH() public payable {
        if (msg.value == 0) {
            revert ZeroDeposit();
        }
        _recordDepositETHAfterTransfer(msg.value);
    }

    /**
     * @notice Deposit StETH to the ETH pool
     * @param stETHAmount Amount to deposit in StETH (wad)
     */
    function depositStETH(uint256 stETHAmount) public {
        if (stETHAmount == 0) {
            revert ZeroDeposit();
        }
        _recordDepositETHBeforeTransfer(stETHAmount);
        LIDO.transferFrom(msg.sender, address(this), stETHAmount);
    }

    /**
     * @notice Deposit StETH to the ETH pool with a permit signature
     * @param stETHAmount Amount to deposit in StETH (wad)
     * @param allowance Allowance amount
     * @param deadline Permit signature deadline
     * @param v Permit signature v parameter
     * @param r Permit signature r parameter
     * @param s Permit signature s parameter
     */
    function depositStETHWithPermit(uint256 stETHAmount, uint256 allowance, uint256 deadline, uint8 v, bytes32 r, bytes32 s) external {
        LIDO.permit(msg.sender, address(this), allowance, deadline, v, r, s);
        depositStETH(stETHAmount);
    }

    /**
     * @notice Deposit USDC to the USD pool
     * @dev USDC is converted to DAI using Maker DssPsm
     * @param usdcAmount Amount to deposit in USDC
     */
    function depositUSDC(uint256 usdcAmount) public {
        if (usdcAmount == 0) {
            revert ZeroDeposit();
        }
        uint256 wadAmount = _usdToWad(usdcAmount);
        uint256 conversionFee = PSM.tin() * wadAmount / _WAD;
        _recordDepositUSDBeforeTransfer(wadAmount, wadAmount - conversionFee);

        USDC.transferFrom(msg.sender, address(this), usdcAmount);

        /* Convert USDC to DAI through MakerDAO Peg Stability Mechanism. */
        PSM.sellGem(address(this), usdcAmount);
    }

    /**
     * @notice Deposit USDC to the USD pool with a permit signature
     * @dev USDC is converted to DAI using Maker DssPsm
     * @param usdcAmount Amount to deposit in USDC (usd)
     * @param allowance Allowance amount
     * @param deadline Permit signature deadline timestamp
     * @param v Permit signature v parameter
     * @param r Permit signature r parameter
     * @param s Permit signature s parameter
     */
    function depositUSDCWithPermit(uint256 usdcAmount, uint256 allowance, uint256 deadline, uint8 v, bytes32 r, bytes32 s) external {
        USDC.permit(msg.sender, address(this), allowance, deadline, v, r, s);
        depositUSDC(usdcAmount);
    }

    /**
     * @notice Deposit DAI to the USD pool
     * @param daiAmount Amount to deposit in DAI (wad)
     */
    function depositDAI(uint256 daiAmount) public {
        if (daiAmount == 0) {
            revert ZeroDeposit();
        }
        _recordDepositUSDBeforeTransfer(daiAmount, daiAmount);

        DAI.transferFrom(msg.sender, address(this), daiAmount);
    }

    /**
     * @notice Deposit DAI to the USD pool with a permit signature
     * @param daiAmount Amount to deposit in DAI (wad)
     * @param nonce Permit signature nonce
     * @param expiry Permit signature expiry timestamp
     * @param v Permit signature v parameter
     * @param r Permit signature r parameter
     * @param s Permit signature s parameter
     */
    function depositDAIWithPermit(uint256 daiAmount, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) external {
        DAI.permit(msg.sender, address(this), nonce, expiry, true, v, r, s);
        depositDAI(daiAmount);
    }

    /**
     * @notice Deposit USDT to the USD pool
     * @dev USDT is converted to DAI using Curve 3Pool
     * @param usdtAmount Amount to deposit in USDT (usd)
     * @param minDAIAmount Minimum DAI amount to accept when exchanging through Curve (wad)
     */
    function depositUSDT(uint256 usdtAmount, uint256 minDAIAmount) external {
        if (usdtAmount == 0) {
            revert ZeroDeposit();
        }

        uint256 usdtBalance = USDT.balanceOf(address(this));
        USDT.transferFrom(msg.sender, address(this), usdtAmount);
        uint256 receivedUSDT = USDT.balanceOf(address(this)) - usdtBalance;

        /* Exchange USDT to DAI through the Curve 3Pool. */
        uint256 daiBalance = DAI.balanceOf(address(this));
        CURVE_3POOL.exchange(
            _CURVE_USDT_INDEX,
            _CURVE_DAI_INDEX,
            receivedUSDT,
            minDAIAmount
        );

        /* The amount of DAI received in the exchange is uncertain due to slippage, so we must record the deposit after the exchange. */
        uint256 receivedDAI = DAI.balanceOf(address(this)) - daiBalance;
        _recordDepositUSDAfterTransfer(_usdToWad(usdtAmount), receivedDAI);
    }

    /**
     * @notice Mint new ETH shares from new deposit before deposit has been made
     * @param amount Amount deposited in ETH
     */
    function _recordDepositETHBeforeTransfer(uint256 amount) internal {
        _recordDepositETH(amount, false);
    }

    /**
     * @notice Mint new ETH shares from new deposit after deposit has been made
     * @param amount Amount deposited in ETH
     */
    function _recordDepositETHAfterTransfer(uint256 amount) internal {
        _recordDepositETH(amount, true);
    }

    /**
     * @notice Mint new USD shares from new deposit before deposit has been made
     * @param depositedAmount Amount deposited in USD (wad)
     * @param daiAmount Amount of DAI obtained after conversion (wad)
     */
    function _recordDepositUSDBeforeTransfer(uint256 depositedAmount, uint256 daiAmount) internal {
        _recordDepositUSD(depositedAmount, daiAmount, false);
    }

    /**
     * @notice Mint new USD shares from new deposit after deposit has been made
     * @param depositedAmount Amount deposited in USD (wad)
     * @param daiAmount Amount of DAI obtained after conversion (wad)
     */
    function _recordDepositUSDAfterTransfer(uint256 depositedAmount, uint256 daiAmount) internal {
        _recordDepositUSD(depositedAmount, daiAmount, true);
    }

    /**
     * @notice Mint new ETH shares from new deposit
     * @param depositedAmount Amount deposited in ETH (wad)
     * @param alreadyDeposited The amount has already been deposited to the contract
     */
    function _recordDepositETH(uint256 depositedAmount, bool alreadyDeposited) internal whenNotPaused {
        uint256 _totalETHBalance = totalETHBalance();
        if (alreadyDeposited) {
            _totalETHBalance = _totalETHBalance - depositedAmount;
        }
        uint256 sharesToIssue = depositedAmount * totalETHShares / _totalETHBalance;
        if (sharesToIssue == 0) {
            revert ZeroSharesIssued();
        }

        _mintETHShares(msg.sender, sharesToIssue);

        emit ETHDeposited(msg.sender, sharesToIssue, depositedAmount);
    }

    /**
     * @notice Mint new USD shares from new deposit
     * @param depositedAmount Amount deposited in USD (wad)
     * @param daiAmount Amount of DAI obtained after conversion (wad)
     * @param alreadyDeposited The amount has already been deposited to the contract
     */
    function _recordDepositUSD(uint256 depositedAmount, uint256 daiAmount, bool alreadyDeposited) internal whenNotPaused {
        uint256 _totalUSDBalance = totalUSDBalance();
        if (alreadyDeposited) {
            _totalUSDBalance = _totalUSDBalance - daiAmount;
        }
        uint256 sharesToIssue = daiAmount * totalUSDShares / _totalUSDBalance; // user only gets shares for the obtained DAI
        if (sharesToIssue == 0) {
            revert ZeroSharesIssued();
        }

        _mintUSDShares(msg.sender, sharesToIssue);

        emit USDDeposited(msg.sender, sharesToIssue, depositedAmount, daiAmount);
    }

    /**
     * @notice Mint ETH shares
     * @param user User address
     * @param shares Number of ETH shares to mint
     */
    function _mintETHShares(address user, uint256 shares) internal {
        ethShares[user] += shares;
        totalETHShares += shares;
    }

    /**
     * @notice Mint USD shares
     * @param user User address
     * @param shares Number of USD shares to mint
     */
    function _mintUSDShares(address user, uint256 shares) internal {
        usdShares[user] += shares;
        totalUSDShares += shares;
    }

    /**
     * @notice Burn ETH shares
     * @param user User address
     * @param shares Number of ETH shares to burn
     */
    function _burnETHShares(address user, uint256 shares) internal {
        ethShares[user] -= shares;
        totalETHShares -= shares;
    }

    /**
     * @notice Burn USD shares
     * @param user User address
     * @param shares Number of USD shares to burn
     */
    function _burnUSDShares(address user, uint256 shares) internal {
        usdShares[user] -= shares;
        totalUSDShares -= shares;
    }

    /*/////////////////////////
              STAKING
    /////////////////////////*/

    /**
     * @notice Stake pooled ETH funds by submiting ETH to Lido
     * @param amount Amount in ETH to stake (wad)
     */
    function stakeETH(uint256 amount) external {
        if (msg.sender != staker) {
            revert CallerIsNotStaker();
        }
        if (amount > address(this).balance) {
            revert InsufficientFunds();
        }

        LIDO.submit{value: amount}(address(0));
    }

    /**
     * @notice Stake pooled USD funds by depositing DAI into the Maker DSR
     * @param amount Amount in DAI to stake (usd)
     */
    function stakeUSD(uint256 amount) external {
        if (msg.sender != staker) {
            revert CallerIsNotStaker();
        }
        if (amount > DAI.balanceOf(address(this))) {
            revert InsufficientFunds();
        }

        DSR_MANAGER.join(address(this), amount);
    }

    /*/////////////////////////
            TRANSITION
    /////////////////////////*/

    /**
     * @notice Start the transition to the mainnet bridge (admin only)
     * @param mainnetBridge Mainnet bridge address
     */
    function enableTransition(address mainnetBridge) external onlyOwner {
        if (isTransitionEnabled) {
            revert TransitionIsEnabled();
        }

        _pause();
        _setMainnetBridge(mainnetBridge);
        isTransitionEnabled = true;

        LIDO.approve(mainnetBridge, type(uint256).max);
        DAI.approve(mainnetBridge, type(uint256).max);
    }

    /**
     * @notice Transition the caller's portion of the pooled funds to the mainnet bridge
     */
    // NB: This function is now responsible for moving the assets to the new bridge,
    // this was before in the `_moveETH`-related functionality.
    function transition(uint32 minGasLimit) external onlyEOA {
        _transition(msg.sender, msg.sender, minGasLimit);
    }

    /**
     * @notice Transition the caller's portion of the pooled funds to the mainnet bridge
     */
    // NB: This function is now responsible for moving the assets to the new bridge,
    // this was before in the `_moveETH`-related functionality.
    function transition(address recipient, uint8 v, bytes32 r, bytes32 s, uint32 minGasLimit) external {
        {
            if (recipient == address(0)) {
                revert InvalidRecipient();
            }

            /// Verify signature of the recipient address by the recipient address.
            /// This is just a safety check for the user that they own the wallet
            /// they are sending funds to.
            bytes memory prefix = "\x19Ethereum Signed Message:\n32";
            bytes32 prefixedHashMessage = keccak256(abi.encodePacked(prefix, recipient));
            address signer = ecrecover(prefixedHashMessage, v, r, s);
            if (signer != recipient) {
                revert InvalidRecipientSignature();
            }
        }

        _transition(msg.sender, recipient, minGasLimit);
    }

    /**
     * @notice Transition the caller's portion of the pooled funds to the mainnet bridge
     */
    // NB: This function is now responsible for moving the assets to the new bridge,
    // this was before in the `_moveETH`-related functionality.
    function _transition(address user, address recipient, uint32 minGasLimit) internal {
        if (!isTransitionEnabled) {
            revert TransitionNotEnabled();
        }

        if (transitioned[user]) {
            revert UserAlreadyTransitioned();
        }
        transitioned[user] = true;

        (uint ethAmountToMove, uint stETHAmountToMove) = _moveETH(user);
        uint daiAmountToMove = _moveUSD(user);

        IMainnetBridge mainnetBridge = getMainnetBridge();
        if (ethAmountToMove > 0) {
            mainnetBridge.bridgeETHTo{value: ethAmountToMove}(recipient, minGasLimit, bytes(""));
        }
        if (stETHAmountToMove > 0) {
            mainnetBridge.bridgeERC20To(address(LIDO), address(0), recipient, stETHAmountToMove, minGasLimit, hex"");
        }
        if (daiAmountToMove > 0) {
            mainnetBridge.bridgeERC20To(address(DAI), Predeploys.USDB, recipient, daiAmountToMove, minGasLimit, hex"");
        }
    }

    /// In the event multisig keys are lost, users can reclaim their funds after the contract
    /// has expired.
    function emergencyWithdraw() external {
        require(block.timestamp > EMERGENCY_WITHDRAW_TIMESTAMP, "Emergency timestamp not reached");
        _withdraw();
    }

    function _withdraw() internal {
        (uint ethAmountToMove, uint stETHAmountToMove) = _moveETH(msg.sender);
        uint daiAmountToMove = _moveUSD(msg.sender);

        if (stETHAmountToMove > 0) {
            LIDO.transfer(msg.sender, stETHAmountToMove);
        }
        if (daiAmountToMove > 0) {
            DAI.transfer(msg.sender, daiAmountToMove);
        }
        if (ethAmountToMove > 0) {
            payable(msg.sender).transfer(ethAmountToMove);
        }

        emit Withdraw(msg.sender, ethAmountToMove, stETHAmountToMove, daiAmountToMove);
    }

    /**
     * @notice Move user's portion of pooled ETH by the amount of shares
     * @param user User address
     */
    // NB: This function was refactored to return the assets it'd move around, and
    // the caller is responsible for executing the actual transfer.
    function _moveETH(address user) internal returns (uint ethAmountToMove, uint stETHAmountToMove) {
        uint256 userETHShares = ethShares[user];
        if (userETHShares > 0) {
            ethAmountToMove = _ethByShares(userETHShares);
            _burnETHShares(user, userETHShares);

            /*
               If there are insufficient ETH funds in the bridge to cover the user's share,
               then we need to start moving StETH.
            */
            uint256 contractETHBalance = address(this).balance;
            if (ethAmountToMove > contractETHBalance) {
                stETHAmountToMove = ethAmountToMove - contractETHBalance;
                ethAmountToMove = contractETHBalance;
            }
        }
    }

    /**
     * @notice Move user's portion of pooled USD by the amount of shares
     * @param user User address
     */
    function _moveUSD(address user) internal returns (uint daiAmountToMove) {
        uint256 userUSDShares = usdShares[user];
        if (userUSDShares > 0) {
            daiAmountToMove = _usdByShares(userUSDShares);
            _burnUSDShares(user, userUSDShares);

            /*
               If there are insufficient DAI funds in the bridge to cover the user's share,
               then we need to start withdrawing DAI from the DSR.
            */
            uint256 contractDAIBalance = DAI.balanceOf(address(this));
            if (daiAmountToMove > contractDAIBalance) {
                DSR_MANAGER.exit(address(this), daiAmountToMove - contractDAIBalance);
            }
        }
    }

    /*/////////////////////////
              HELPERS
    /////////////////////////*/

    /**
     * @notice Convert ETH to equivalent shares
     * @param shares Number of ETH shares
     * @return Amount of ETH
     */
    function _ethByShares(uint256 shares) internal view returns (uint256) {
        return shares * totalETHBalance() / totalETHShares;
    }

    /**
     * @notice Current shares to equivalent USD
     * @dev Does not update DSR yield
     * @param shares Number of USD shares
     * @return Amount of USD
     */
    function _usdBySharesNoUpdate(uint256 shares) internal view returns (uint256) {
        return shares * totalUSDBalanceNoUpdate() / totalUSDShares;
    }

    /**
     * @notice Current shares to equivalent USD
     * @param shares Number of USD shares
     * @return Amount of USD
     */
    function _usdByShares(uint256 shares) internal returns (uint256) {
        return shares * totalUSDBalance() / totalUSDShares;
    }

    /**
     * @notice Convert from wad (18 decimals) to USD (6 decimals) denomination
     * @param wad Amount in wad
     * @return Amount in USD
     */
    function _wadToUSD(uint256 wad) internal pure returns (uint256) {
        return wad / (10**(_WAD_DECIMALS - _USD_DECIMALS));
    }

    /**
     * @notice Convert from USD (6 decimals) to wad (18 decimals) denomination
     * @param usd Amount in USD
     * @return Amount in wad
     */
    function _usdToWad(uint256 usd) internal pure returns (uint256) {
        return usd * (10**(_WAD_DECIMALS - _USD_DECIMALS));
    }

    /**
     * @dev Based on _rpow from MakerDAO pot.sol contract (https://github.com/makerdao/dss/blob/fa4f6630afb0624d04a003e920b0d71a00331d98/src/pot.sol#L87-L105)
     */
    function _rpow(uint x, uint n, uint base) internal pure returns (uint z) {
        assembly {
            switch x case 0 {switch n case 0 {z := base} default {z := 0}}
            default {
                switch mod(n, 2) case 0 { z := base } default { z := x }
                let half := div(base, 2)  // for rounding.
                for { n := div(n, 2) } n { n := div(n,2) } {
                    let xx := mul(x, x)
                    if iszero(eq(div(xx, x), x)) { revert(0,0) }
                    let xxRound := add(xx, half)
                    if lt(xxRound, xx) { revert(0,0) }
                    x := div(xxRound, base)
                    if mod(n,2) {
                        let zx := mul(z, x)
                        if and(iszero(iszero(x)), iszero(eq(div(zx, x), z))) { revert(0,0) }
                        let zxRound := add(zx, half)
                        if lt(zxRound, zx) { revert(0,0) }
                        z := div(zxRound, base)
                    }
                }
            }
        }
    }

    /**
     * @dev Based on _rmul in MakerDAO pot.sol contract (https://github.com/makerdao/dss/blob/fa4f6630afb0624d04a003e920b0d71a00331d98/src/pot.sol#L109-L111)
     */
    function _rmul(uint x, uint y) internal pure returns (uint z) {
        z = x * y / _RAY;
    }
}
