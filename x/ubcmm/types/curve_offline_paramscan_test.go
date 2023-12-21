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
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Curve_OfflineParamScan(t *testing.T) {
	newValidCurve := func() Curve {
		return Curve{
			RefTokenSupply:  sdk.NewDec(6e9),
			RefTokenPrice:   sdk.NewDec(1),
			RefProfitFactor: sdk.NewDec(10),
			BPool:           sdk.NewDec(100e6),
			BPoolUnder:      sdk.NewDec(100e6),
			SlopeP2:         sdk.NewDecWithPrec(3, 1).Quo(sdk.NewDec(3e9)),
			SlopeP3:         sdk.NewDec(2).Quo(sdk.NewDec(3e9)),
			FactorFy:        sdk.NewDecWithPrec(2, 1),
			FactorFxy:       sdk.NewDec(15832600001),
		}
	}

	tests := []struct {
		factorFy          sdk.Dec
		expectedFactorFxy string
	}{
		{sdk.NewDecWithPrec(1, 2), "17889500001.000000000000000000000000000000000000"}, // python produces: 17446200001.0
		{sdk.NewDecWithPrec(2, 1), "16179000001.000000000000000000000000000000000000"}, // python produces: 15832600001.0
	}
	// Note: The difference in results produced by python could be due to the floating point arithmetics and difference in precision.

	for i := range tests {

		c := newValidCurve()
		c.FactorFy = tests[i].factorFy

		err := c.OfflineParamScan(sdk.NewDec(100000), sdk.NewDecWithPrec(1, 4))
		require.NoError(t, err)
		assert.Equal(t, c.FactorFxy.String(), tests[i].expectedFactorFxy)
		assert.True(t, c.IsIntegralEqualToBPool())
	}
}
