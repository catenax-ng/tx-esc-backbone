// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (c *Curve) OfflineParamScan(factorFx2yAdjustment, allowedDiff sdk.Dec) error {
	c.FactorFxy = sdk.NewDec(1) // Override FactorFxy value to default initialization.

	if err := c.validateParameters(); err != nil {
		return err
	}
	c.initParameters()
	c.initSegmentsToZero()
	c.fitS3()
	c.fitS4()

	dx1 := sdk.NewDec(0)
	var midA1 sdk.Dec

	for dx1.LT(c.pX(3)) {
		c.fitS1S2()

		if c.validateCurvature() == nil {
			midA1 = c.lowerBoundA1().Add(c.upperBoundA1()).Quo(sdk.NewDec(2))

			if midA1.Mul(sdk.NewDec(1).Sub(allowedDiff)).LT(c.S2.A) &&
				midA1.Mul(sdk.NewDec(1).Add(allowedDiff)).GT(c.S2.A) {
				return nil
			}

		}
		c.FactorFxy = c.FactorFxy.Add(factorFx2yAdjustment)
		dx1 = c.dx1()
	}
	return nil
}

func (c *Curve) lowerBoundA1() sdk.Dec {
	return sdk.NewDec(-2).Quo(sdk.NewDec(3)).
		Mul(c.S2.DeltaX).
		Mul(c.S3.firstDerivativeT1(sdk.NewDec(0)).
			Quo(c.S3.DeltaX)).Add(c.pY(2))
}
func (c *Curve) upperBoundA1() sdk.Dec {
	return sdk.NewDecWithPrec(5, 1).
		Mul(c.pY(1)).
		Sub(
			sdk.NewDec(1).Quo(sdk.NewDec(6)).
				Mul(c.S2.DeltaX).
				Mul(c.S3.firstDerivativeT1(sdk.NewDec(0)).
					Quo(c.S3.DeltaX))).
		Add(sdk.NewDecWithPrec(5, 1).Mul(c.pY(2)))
}

func (c *Curve) dx1() sdk.Dec {
	return sdk.NewDec(1).Sub(c.FactorFxy).Mul(c.pY(2)).
		Sub(sdk.NewDec(1).Sub(c.FactorFxy).Mul(c.pY(0)))
}
