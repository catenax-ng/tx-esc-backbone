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
	fmt "fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_Curve_ShiftUp tests if the values produced by shift up function
// is same as the values produced by the prototype.
func Test_Curve_ShiftUp(t *testing.T) {
	type wants struct {
		S0         *FlatSegment
		S1         *BezierSegment
		S2         *BezierSegment
		S3         *FixedBezierSegment
		S4         *FixedQuadraticSegment
		BPool      sdk.Dec
		BPoolUnder sdk.Dec
	}
	type args struct {
		BPoolAdd        sdk.Dec
		DegirdingFactor sdk.Dec
		TokensToBuy     sdk.Dec
	}
	tests := []struct {
		name    string
		args    args
		wants   wants
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Happy path",
			args: args{
				BPoolAdd:        sdk.NewDec(100e5),
				DegirdingFactor: sdk.NewDec(1),
				TokensToBuy:     sdk.NewDec(500e5),
			},
			wants: wants{
				S0: &FlatSegment{
					P1X: sdk.MustNewDecFromStr("46966373.186190389801025890861822048267737788"),
					Y:   sdk.MustNewDecFromStr("0.017254991440409323232454505835560365"),
				},
				S1: &BezierSegment{
					P0Y:    sdk.MustNewDecFromStr("0.017254991440409323232454505835560365"),
					A:      sdk.MustNewDecFromStr("0.017254991440409323232454505835560365"),
					B:      sdk.MustNewDecFromStr("0.026001798814231522750795441811313855"),
					P1Y:    sdk.MustNewDecFromStr("0.034748606188053722269136377787067345"),
					P0X:    sdk.MustNewDecFromStr("46966373.186190389801025890861822048267737788"),
					P1X:    sdk.MustNewDecFromStr("1938969127.118076932874869224554139114130514462"),
					DeltaX: sdk.MustNewDecFromStr("1892002753.931886543073843333692317065862776674"),
				},
				S2: &BezierSegment{
					P0Y:    sdk.MustNewDecFromStr("0.034748606188053722269136377787067345"),
					A:      sdk.MustNewDecFromStr("0.039870923511293814967035871232711658"),
					B:      sdk.MustNewDecFromStr("0.069534204713335437568568770228105743"),
					P1Y:    sdk.MustNewDecFromStr("0.109446130357262636831727731186190530"),
					P0X:    sdk.MustNewDecFromStr("1938969127.118076932874869224554139114130514462"),
					P1X:    sdk.MustNewDecFromStr("3046966373.186190389801025890861822048267737788"),
					DeltaX: sdk.MustNewDecFromStr("1107997246.068113456926156666307682934137223326"),
				},
				S3: &FixedBezierSegment{
					BezierSegment: &BezierSegment{
						P0Y:    sdk.MustNewDecFromStr("0.1"),
						A:      sdk.MustNewDecFromStr("0.2"),
						B:      sdk.MustNewDecFromStr("0.333333333333333333333333333000000001"),
						P1Y:    sdk.MustNewDecFromStr("1"),
						P0X:    sdk.MustNewDecFromStr("3000000000"),
						P1X:    sdk.MustNewDecFromStr("6000000000"),
						DeltaX: sdk.MustNewDecFromStr("3000000000"),
					},
					IntervalP0X: sdk.MustNewDecFromStr("3046966373.186190389801025890861822048267737788"),
				},
				S4: &FixedQuadraticSegment{
					A:             sdk.MustNewDecFromStr("0.533333333333333333333333334"),
					B:             sdk.MustNewDecFromStr("-5.733333333333333333333333341"),
					C:             sdk.MustNewDecFromStr("16.200000000000000000000000022"),
					ScalingFactor: sdk.MustNewDecFromStr("1000000000"),
					InitialX0:     sdk.MustNewDecFromStr("6000000000"),
					CurrentX0:     sdk.MustNewDecFromStr("6000000000"),
				},
				BPool:      sdk.MustNewDecFromStr("115125491.9"),
				BPoolUnder: sdk.MustNewDecFromStr("114807335.560533360132001356801641924"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := setupUbcAndByToken(t, tt.args.TokensToBuy)
			tt.wantErr(t, c.ShiftUp(tt.args.BPoolAdd, tt.args.DegirdingFactor),
				fmt.Sprintf("BuyToken(%v) then ShiftUp(%v, %v)", tt.args.TokensToBuy, tt.args.BPoolAdd, tt.args.DegirdingFactor))
			assert.EqualValues(t, tt.wants.S0, c.S0)
			assert.EqualValues(t, tt.wants.S1, c.S1)
			assert.EqualValues(t, tt.wants.S2, c.S2)
			assert.EqualValues(t, tt.wants.S3, c.S3)
			assert.EqualValues(t, tt.wants.S4, c.S4)
			assert.Equal(t, tt.wants.BPool, c.BPool)
			assert.Equal(t, tt.wants.BPoolUnder, c.BPoolUnder)
		})
	}
}

func setupUbcAndByToken(t *testing.T, tokensToBuy sdk.Dec) Curve {
	t.Helper()
	c := validCurve()
	require.NoError(t, c.Fit())
	IsEqualDecimal(t, "100000000", c.BPool)
	IsEqualDecimal(t, "100000000", c.BPoolUnder)
	IsEqualDecimal(t, "3000000000", c.CurrentSupply)
	c.Buy(tokensToBuy)
	return c
}
