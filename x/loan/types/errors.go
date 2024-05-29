package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/loan module sentinel errors
var (
	ErrInvalidSigner    = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample           = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidLoanState = sdkerrors.Register(ModuleName, 1103, "invalid loan state")
	ErrDeadline         = sdkerrors.Register(ModuleName, 1104, "deadline")
)
