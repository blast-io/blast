// SPDX-License-Identifier: BSL 1.1 - Copyright 2024 MetaLayer Labs Ltd.
pragma solidity 0.8.15;

/// @title  OptimistConstants
/// @notice Library for storing Optimist related constants that are shared in multiple contracts.
library OptimistConstants {
    /// @notice Attestation key issued by OptimistInviter allowing the attested account to mint.
    bytes32 internal constant OPTIMIST_CAN_MINT_FROM_INVITE_ATTESTATION_KEY = bytes32("optimist.can-mint-from-invite");
}
