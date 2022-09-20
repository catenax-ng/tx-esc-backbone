package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/resourcesync module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidOriginator = sdkerrors.Register(ModuleName, 1101, "originator address is invalid %s")
	ErrInvalidResource   = sdkerrors.Register(ModuleName, 1102, "resource invalid")
)
