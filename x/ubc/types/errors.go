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

// x/ubc module sentinel errors
var (
	ErrSample       = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidArg   = sdkerrors.Register(ModuleName, 1101, "invalid argument")
	ErrCurveFitting = sdkerrors.Register(ModuleName, 1102, "fitting the curve")
	ErrFundHandling = sdkerrors.Register(ModuleName, 1103, "fund handling")
	ErrComputation  = sdkerrors.Register(ModuleName, 1104, "computation error")
)
