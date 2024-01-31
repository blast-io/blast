// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { IDisputeGame } from "./IDisputeGame.sol";

import "src/libraries/DisputeTypes.sol";

/// @title IFaultDisputeGame
/// @notice The interface for a fault proof backed dispute game.
interface IFaultDisputeGame is IDisputeGame {
    /// @notice The `ClaimData` struct represents the data associated with a Claim.
    /// @dev TODO(clabby): Add bond ID information.
    struct ClaimData {
        uint32 parentIndex;
        bool countered;
        Claim claim;
        Position position;
        Clock clock;
    }

    /// @notice The `OutputProposal` struct contains information about an output proposal in
    ///         the `L2OutputOracle` at a given index.
    struct OutputProposal {
        uint128 index;
        uint128 l2BlockNumber;
        Hash outputRoot;
    }

    /// @notice A container for two consecutive `OutputProposal`s, used to store the starting
    ///         and disputed output proposals for a given dispute game. The starting output
    ///         proposal will be used to determine where the off chain agents should begin
    ///         running their fault proof program, and the disputed output proposal will be
    ///         fed into the program and treated as disputed state. The program's exit code
    ///         expresses its opinion on the validity of the state transition from the starting,
    ///         trusted output proposal to the disputed output proposal, and ultimately resolves
    ///         the dispute.
    struct OutputProposals {
        OutputProposal starting;
        OutputProposal disputed;
    }

    /// @notice Emitted when a new claim is added to the DAG by `claimant`
    /// @param parentIndex The index within the `claimData` array of the parent claim
    /// @param claim The claim being added
    /// @param claimant The address of the claimant
    event Move(uint256 indexed parentIndex, Claim indexed claim, address indexed claimant);

    /// @notice Attack a disagreed upon `Claim`.
    /// @param _parentIndex Index of the `Claim` to attack in the `claimData` array.
    /// @param _claim The `Claim` at the relative attack position.
    function attack(uint256 _parentIndex, Claim _claim) external payable;

    /// @notice Defend an agreed upon `Claim`.
    /// @param _parentIndex Index of the claim to defend in the `claimData` array.
    /// @param _claim The `Claim` at the relative defense position.
    function defend(uint256 _parentIndex, Claim _claim) external payable;

    /// @notice Perform an instruction step via an on-chain fault proof processor.
    /// @dev This function should point to a fault proof processor in order to execute
    ///      a step in the fault proof program on-chain. The interface of the fault proof
    ///      processor contract should adhere to the `IBigStepper` interface.
    /// @param _claimIndex The index of the challenged claim within `claimData`.
    /// @param _isAttack Whether or not the step is an attack or a defense.
    /// @param _stateData The stateData of the step is the preimage of the claim at the given
    ///        prestate, which is at `_stateIndex` if the move is an attack and `_claimIndex` if
    ///        the move is a defense. If the step is an attack on the first instruction, it is
    ///        the absolute prestate of the fault proof VM.
    /// @param _proof Proof to access memory nodes in the VM's merkle state tree.
    function step(uint256 _claimIndex, bool _isAttack, bytes calldata _stateData, bytes calldata _proof) external;

    /// @notice Posts the requested local data to the VM's `PreimageOralce`.
    /// @param _ident The local identifier of the data to post.
    /// @param _l2BlockNumber The L2 block number being disputed. This serves as the local context for the
    ///                       `PreimageOracle` key.
    /// @param _partOffset The offset of the data to post.
    function addLocalData(uint256 _ident, uint256 _l2BlockNumber, uint256 _partOffset) external;

    /// @notice Resolves the subgame rooted at the given claim index.
    /// @dev This function must be called bottom-up in the DAG
    ///      A subgame is a tree of claims that has a maximum depth of 1.
    ///      A subgame root claims is valid if, and only if, all of its child claims are invalid.
    ///      At the deepest level in the DAG, a claim is invalid if there's a successful step against it.
    /// @param _claimIndex The index of the subgame root claim to resolve.
    function resolveClaim(uint256 _claimIndex) external payable;

    /// @notice An L1 block hash that contains the disputed output root, fetched from the
    ///         `BlockOracle` and verified by referencing the timestamp associated with the
    ///         first L2 Output Proposal in the `L2OutputOracle` that contains the disputed
    ///         L2 block number.
    function l1Head() external view returns (Hash l1Head_);

    /// @notice The l2BlockNumber of the disputed output root in the `L2OutputOracle`.
    function l2BlockNumber() external view returns (uint256 l2BlockNumber_);

    /// @notice The l1BlockNumber that Cannon was ran from to generate the root claim.
    function l1BlockNumber() external view returns (uint256 l1BlockNumber_);
}
