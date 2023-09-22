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
					X0: sdk.MustNewDecFromStr("46479149.736614226371162451290185651416902218"),
					Y:  sdk.MustNewDecFromStr("0.017269768295289466057148867957072234"),
				},
				S0: &BezierSegment{
					P0Y:    sdk.MustNewDecFromStr("0.017269768295289466057148867957072234"),
					A:      sdk.MustNewDecFromStr("0.017269768295289466057148867957072234"),
					B:      sdk.MustNewDecFromStr("0.026010169252480982831302467523822519"),
					P1Y:    sdk.MustNewDecFromStr("0.034750570209672499605456067090572804"),
					P0X:    sdk.MustNewDecFromStr("46479149.736614226371162451290185651416902218"),
					P1X:    sdk.MustNewDecFromStr("1939426924.915664922710798117306552030770617419"),
					DeltaX: sdk.MustNewDecFromStr("1892947775.179050696339635666016366379353715201"),
				},
				S1: &BezierSegment{
					P0Y:    sdk.MustNewDecFromStr("0.034750570209672499605456067090572804"),
					A:      sdk.MustNewDecFromStr("0.039862216950465852771207048617955396"),
					B:      sdk.MustNewDecFromStr("0.069496742435242257408500890303316353"),
					P1Y:    sdk.MustNewDecFromStr("0.109347555734409267597369727249150170"),
					P0X:    sdk.MustNewDecFromStr("1939426924.915664922710798117306552030770617419"),
					P1X:    sdk.MustNewDecFromStr("3046479149.736614226371162451290185651416902218"),
					DeltaX: sdk.MustNewDecFromStr("1107052224.820949303660364333983633620646284799"),
				},
				S2: &FixedBezierSegment{
					BezierSegment: &BezierSegment{
						P0Y:    sdk.MustNewDecFromStr("0.1"),
						A:      sdk.MustNewDecFromStr("0.2"),
						B:      sdk.MustNewDecFromStr("0.333333333333333333333333333000000001"),
						P1Y:    sdk.MustNewDecFromStr("1"),
						P0X:    sdk.MustNewDecFromStr("3000000000"),
						P1X:    sdk.MustNewDecFromStr("6000000000"),
						DeltaX: sdk.MustNewDecFromStr("3000000000"),
					},
					IntervalP0X: sdk.MustNewDecFromStr("3046479149.736614226371162451290185651416902218"),
				},
				QS3: &FixedQuadraticSegment{
					A:             sdk.MustNewDecFromStr("0.533333333333333333333333334"),
					B:             sdk.MustNewDecFromStr("-5.733333333333333333333333341"),
					C:             sdk.MustNewDecFromStr("16.200000000000000000000000022"),
					ScalingFactor: sdk.MustNewDecFromStr("1000000000"),
					InitialX0:     sdk.MustNewDecFromStr("6000000000"),
					CurrentX0:     sdk.MustNewDecFromStr("6000000000"),
				},
				BPool:      sdk.MustNewDecFromStr("115125491.898148"),
				BPoolUnder: sdk.MustNewDecFromStr("114756324.034698007429998077480287124"),
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
