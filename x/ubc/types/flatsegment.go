// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// integralX12 computes the integral of the curve segment with respect to "x",
// between limits x1 and x2.
func (fl *Flatsegment) integralX12(x1, x2 sdk.Dec) sdk.Dec {
	return fl.integralX1(x2).Sub(fl.integralX1(x1))
}

// integralX1 computes the integral of the curve segment with respect to "x",
// from the beginning of the curve until point x1.
func (fl *Flatsegment) integralX1(x1 sdk.Dec) sdk.Dec {
	return fl.Y.Mul(x1)
}
