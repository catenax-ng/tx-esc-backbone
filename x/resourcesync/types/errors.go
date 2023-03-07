// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/resourcesync module sentinel errors
var (
	ErrSample              = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidOriginator   = sdkerrors.Register(ModuleName, 1101, "originator address is invalid %s")
	ErrInvalidResource     = sdkerrors.Register(ModuleName, 1102, "resource invalid")
	ErrDuplicateResource   = sdkerrors.Register(ModuleName, 1103, "resource duplicate")
	ErrNonexistentResource = sdkerrors.Register(ModuleName, 1104, "resource nonexistent")
)
