// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// integralX12 computes the integral of the curve segment with respect to "x",
// between limits x1 and x2, in the scaled domain.
func (qsg *Quadraticsegment) integralX12(x1, x2 sdk.Dec) sdk.Dec {
	integralX1 := func(qsg *Quadraticsegment, x1 sdk.Dec) sdk.Dec {
		// Results are computed in the scaled x-y domain and then converted back to
		// original domain.
		x1 = x1.Quo(qsg.ScalingFactor)

		part1 := x1.Power(3).Mul(qsg.A.Quo(sdk.NewDec(3)))
		part2 := x1.Power(2).Mul(qsg.B.Quo(sdk.NewDec(2)))
		part3 := x1.Mul(qsg.C)

		return qsg.ScalingFactor.Mul(part1.Add(part2).Add(part3))
	}

	return integralX1(qsg, x2).Sub(integralX1(qsg, x1))

}
