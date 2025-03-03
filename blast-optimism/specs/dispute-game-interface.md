# Dispute Game Interface

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [Types](#types)
- [`DisputeGameFactory` Interface](#disputegamefactory-interface)
- [`DisputeGame` Interface](#disputegame-interface)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

A dispute game is played between multiple parties when contesting the truthiness
of a claim. In the context of an optimistic rollup, claims are made about the
state of the layer two network to enable withdrawals to the layer one. A proposer
makes a claim about the layer two state such that they can withdraw and a
challenger can dispute the validity of the claim. The security of the layer two
comes from the ability of fraudulent withdrawals being able to be disputed.

A dispute game interface is defined to allow for multiple implementations of
dispute games to exist. If multiple dispute games run in production, it gives
a similar security model as having multiple protocol clients, as a bug in a
single dispute game will not result in the bug becoming consensus.

## Types

For added context, we define a few types that are used in the following snippets.

```solidity

/// @notice A custom type for a generic hash.
type Hash is bytes32;

/// @notice The type of proof system being used.
enum GameType {
    /// @dev The game will use a `IDisputeGame` implementation that utilizes fault proofs.
    FAULT,
    /// @dev The game will use a `IDisputeGame` implementation that utilizes validity proofs.
    VALIDITY,
    /// @dev The game will use a `IDisputeGame` implementation that utilizes attestation proofs.
    ATTESTATION
}

/// @notice The current status of the dispute game.
enum GameStatus {
    /// @dev The game is currently in progress, and has not been resolved.
    IN_PROGRESS,
    /// @dev The game has concluded, and the `rootClaim` was challenged successfully.
    CHALLENGER_WINS,
    /// @dev The game has concluded, and the `rootClaim` could not be contested.
    DEFENDER_WINS
}

/// @notice A `Claim` type represents a 32 byte hash or other unique identifier for a claim about
///         a certain piece of information.
/// @dev For the `FAULT` `GameType`, this will be a root of the merklized state of the fault proof
///      program at the end of the state transition.
///      For the `ATTESTATION` `GameType`, this will be an output root.
type Claim is bytes32;
```

## `DisputeGameFactory` Interface

The dispute game factory is responsible for creating new `DisputeGame` contracts
given a `GameType` and a root `Claim`. Challenger agents will listen to the
`DisputeGameCreated` events that are emitted by the factory as well as other events
that pertain to detecting fault (i.e. `OutputProposed(bytes32,uint256,uint256,uint256)`) in order to keep up
with on-going disputes in the protocol.

A [`clones-with-immutable-args`](https://github.com/Saw-mon-and-Natalie/clones-with-immutable-args) factory
(originally by @wighawag, but forked by @Saw-mon-and-Natalie) is used to create Clones. Each `GameType` has
a corresponding implementation within the factory, and when a new game is created, the factory creates a
clone of the `GameType`'s pre-deployed implementation contract.

The `rootClaim` of created dispute games can either be a claim that the creator agrees or disagrees with.
This is an implementation detail that is left up to the `IDisputeGame` to handle within its `resolve` function.

When the `DisputeGameFactory` creates a new `DisputeGame` contract, it calls `initialize()` on the clone to
set up the game.

```solidity
/// @title IDisputeGameFactory
/// @notice The interface for a DisputeGameFactory contract.
interface IDisputeGameFactory {
    /// @notice Emitted when a new dispute game is created
    /// @param disputeProxy The address of the dispute game proxy
    /// @param gameType The type of the dispute game proxy's implementation
    /// @param rootClaim The root claim of the dispute game
    event DisputeGameCreated(address indexed disputeProxy, GameType indexed gameType, Claim indexed rootClaim);

    /// @notice Emitted when a new game implementation added to the factory
    /// @param impl The implementation contract for the given `GameType`.
    /// @param gameType The type of the DisputeGame.
    event ImplementationSet(address indexed impl, GameType indexed gameType);

    /// @notice The total number of dispute games created by this factory.
    /// @return gameCount_ The total number of dispute games created by this factory.
    function gameCount() external view returns (uint256 gameCount_);

    /// @notice `games` queries an internal mapping that maps the hash of
    ///         `gameType ++ rootClaim ++ extraData` to the deployed `DisputeGame` clone.
    /// @dev `++` equates to concatenation.
    /// @param _gameType The type of the DisputeGame - used to decide the proxy implementation
    /// @param _rootClaim The root claim of the DisputeGame.
    /// @param _extraData Any extra data that should be provided to the created dispute game.
    /// @return proxy_ The clone of the `DisputeGame` created with the given parameters.
    ///         Returns `address(0)` if nonexistent.
    /// @return timestamp_ The timestamp of the creation of the dispute game.
    function games(GameType _gameType, Claim _rootClaim, bytes calldata _extraData)
        external
        view
        returns (IDisputeGame proxy_, Timestamp timestamp_);

    /// @notice `gameAtIndex` returns the dispute game contract address and its creation timestamp
    ///          at the given index. Each created dispute game increments the underlying index.
    /// @param _index The index of the dispute game.
    /// @return gameType_ The type of the DisputeGame - used to decide the proxy implementation.
    /// @return timestamp_ The timestamp of the creation of the dispute game.
    /// @return proxy_ The clone of the `DisputeGame` created with the given parameters.
    ///         Returns `address(0)` if nonexistent.
    function gameAtIndex(uint256 _index)
        external
        view
        returns (GameType gameType_, Timestamp timestamp_, IDisputeGame proxy_);

    /// @notice `gameImpls` is a mapping that maps `GameType`s to their respective
    ///         `IDisputeGame` implementations.
    /// @param _gameType The type of the dispute game.
    /// @return impl_ The address of the implementation of the game type.
    ///         Will be cloned on creation of a new dispute game with the given `gameType`.
    function gameImpls(GameType _gameType) external view returns (IDisputeGame impl_);

    /// @notice Creates a new DisputeGame proxy contract.
    /// @param _gameType The type of the DisputeGame - used to decide the proxy implementation.
    /// @param _rootClaim The root claim of the DisputeGame.
    /// @param _extraData Any extra data that should be provided to the created dispute game.
    /// @return proxy_ The address of the created DisputeGame proxy.
    function create(GameType _gameType, Claim _rootClaim, bytes calldata _extraData)
        external
        returns (IDisputeGame proxy_);

    /// @notice Sets the implementation contract for a specific `GameType`.
    /// @dev May only be called by the `owner`.
    /// @param _gameType The type of the DisputeGame.
    /// @param _impl The implementation contract for the given `GameType`.
    function setImplementation(GameType _gameType, IDisputeGame _impl) external;

    /// @notice Returns a unique identifier for the given dispute game parameters.
    /// @dev Hashes the concatenation of `gameType . rootClaim . extraData`
    ///      without expanding memory.
    /// @param _gameType The type of the DisputeGame.
    /// @param _rootClaim The root claim of the DisputeGame.
    /// @param _extraData Any extra data that should be provided to the created dispute game.
    /// @return uuid_ The unique identifier for the given dispute game parameters.
    function getGameUUID(GameType _gameType, Claim _rootClaim, bytes memory _extraData)
        external
        pure
        returns (Hash uuid_);
}
```

## `DisputeGame` Interface

The dispute game interface should be generic enough to allow it to work with any
proof system. This means that it should work fault proofs, validity proofs,
an attestation based proof system, or any other source of truth that adheres to
the interface.

Clones of the `IDisputeGame`'s `initialize` functions will be called by the `DisputeGameFactory` upon creation.

```solidity
////////////////////////////////////////////////////////////////
//                    GENERIC DISPUTE GAME                    //
////////////////////////////////////////////////////////////////

/// @title IDisputeGame
/// @notice The generic interface for a DisputeGame contract.
interface IDisputeGame {
    /// @notice Initializes the DisputeGame contract.
    /// @custom:invariant The `initialize` function may only be called once.
    function initialize() external;

    /// @notice Emitted when the game is resolved.
    /// @param status The status of the game after resolution.
    event Resolved(GameStatus indexed status);

    /// @notice Returns the timestamp that the DisputeGame contract was created at.
    function createdAt() external pure returns (Timestamp createdAt_);

    /// @notice Returns the current status of the game.
    function status() external view returns (GameStatus status_);

    /// @notice Getter for the game type.
    /// @dev `clones-with-immutable-args` argument #1
    /// @dev The reference impl should be entirely different depending on the type (fault, validity)
    ///      i.e. The game type should indicate the security model.
    /// @return gameType_ The type of proof system being used.
    function gameType() external view returns (GameType gameType_);

    /// @notice Getter for the root claim.
    /// @return rootClaim_ The root claim of the DisputeGame.
    /// @dev `clones-with-immutable-args` argument #2
    function rootClaim() external view returns (Claim rootClaim_);

    /// @notice Getter for the extra data.
    /// @dev `clones-with-immutable-args` argument #3
    /// @return extraData_ Any extra data supplied to the dispute game contract by the creator.
    function extraData() external view returns (bytes memory extraData_);

    /// @notice Returns the address of the `BondManager` used
    function bondManager() public view returns (IBondManager bondManager_);

    /// @notice If all necessary information has been gathered, this function should mark the game
    ///         status as either `CHALLENGER_WINS` or `DEFENDER_WINS` and return the status of
    ///         the resolved game. It is at this stage that the bonds should be awarded to the
    ///         necessary parties.
    /// @dev May only be called if the `status` is `IN_PROGRESS`.
    /// @return status_ The status of the game after resolution.
    function resolve() public returns (GameStatus status_);

    /// @notice A compliant implementation of this interface should return the components of the
    ///         game UUID's preimage provided in the cwia payload. The preimage of the UUID is
    ///         constructed as `keccak256(gameType . rootClaim . extraData)` where `.` denotes
    ///         concatenation.
    /// @return gameType_ The type of proof system being used.
    /// @return rootClaim_ The root claim of the DisputeGame.
    /// @return extraData_ Any extra data supplied to the dispute game contract by the creator.
    function gameData() external view returns (GameType gameType_, Claim rootClaim_, bytes memory extraData_);
}
```
