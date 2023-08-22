// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
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
		FS0        *FlatSegment
		S0         *BezierSegment
		S1         *BezierSegment
		S2         *FixedBezierSegment
		QS3        *FixedQuadraticSegment
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
				FS0: &FlatSegment{
					X0: sdk.MustNewDecFromStr("46479149.746915109636077002"),
					Y:  sdk.MustNewDecFromStr("0.017269768356181699"),
				},
				S0: &BezierSegment{
					P0Y:    sdk.MustNewDecFromStr("0.017269768356181699"),
					A:      sdk.MustNewDecFromStr("0.017269768356181699"),
					B:      sdk.MustNewDecFromStr("0.026010169307364557"),
					P1Y:    sdk.MustNewDecFromStr("0.034750570258547415"),
					P0X:    sdk.MustNewDecFromStr("46479149.746915109636077002"),
					P1X:    sdk.MustNewDecFromStr("1939426924.228026090155793888"),
					DeltaX: sdk.MustNewDecFromStr("1892947774.481110980519716886"),
				},
				S1: &BezierSegment{
					P0Y:    sdk.MustNewDecFromStr("0.034750570258547415"),
					A:      sdk.MustNewDecFromStr("0.039862217000934048"),
					B:      sdk.MustNewDecFromStr("0.069496742414167885"),
					P1Y:    sdk.MustNewDecFromStr("0.109347555736020564"),
					P0X:    sdk.MustNewDecFromStr("1939426924.228026090155793888"),
					P1X:    sdk.MustNewDecFromStr("3046479149.746915109636077002"),
					DeltaX: sdk.MustNewDecFromStr("1107052225.518889019480283114"),
				},
				S2: &FixedBezierSegment{
					BezierSegment: &BezierSegment{
						P0Y:    sdk.MustNewDecFromStr("0.100000000000000000"),
						A:      sdk.MustNewDecFromStr("0.200000000000000000"),
						B:      sdk.MustNewDecFromStr("0.333333333000000001"),
						P1Y:    sdk.MustNewDecFromStr("1.000000000000000000"),
						P0X:    sdk.MustNewDecFromStr("3000000000.000000000000000000"),
						P1X:    sdk.MustNewDecFromStr("6000000000.000000000000000000"),
						DeltaX: sdk.MustNewDecFromStr("3000000000.000000000000000000"),
					},
					IntervalP0X: sdk.MustNewDecFromStr("3046479149.746915109636077002"),
				},
				QS3: &FixedQuadraticSegment{
					A:             sdk.MustNewDecFromStr("0.533333334"),
					B:             sdk.MustNewDecFromStr("-5.733333341"),
					C:             sdk.MustNewDecFromStr("16.200000022"),
					ScalingFactor: sdk.MustNewDecFromStr("1000000000"),
					InitialX0:     sdk.MustNewDecFromStr("6000000000.000000000000000000"),
					CurrentX0:     sdk.MustNewDecFromStr("6000000000.000000000000000000"),
				},
				BPool:      sdk.MustNewDecFromStr("115125491.898143000000000000"),
				BPoolUnder: sdk.MustNewDecFromStr("114756324.035772563000000000"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := setupUbcAndByToken(t, tt.args.TokensToBuy)
			tt.wantErr(t, c.ShiftUp(tt.args.BPoolAdd, tt.args.DegirdingFactor),
				fmt.Sprintf("BuyToken(%v) then ShiftUp(%v, %v)", tt.args.TokensToBuy, tt.args.BPoolAdd, tt.args.DegirdingFactor))
			assert.EqualValues(t, tt.wants.FS0, c.FS0)
			assert.EqualValues(t, tt.wants.S0, c.S0)
			assert.EqualValues(t, tt.wants.S1, c.S1)
			assert.EqualValues(t, tt.wants.S2, c.S2)
			assert.EqualValues(t, tt.wants.QS3, c.QS3)
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
