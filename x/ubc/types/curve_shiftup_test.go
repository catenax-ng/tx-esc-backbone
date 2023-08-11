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

// Test_Ubcobject_ShiftUp tests if the values produced by shift up function
// is same as the values produced by the prototype.
func Test_Ubcobject_ShiftUp(t *testing.T) {
	type wants struct {
		FS0        *Flatsegment
		S0         *Segment
		S1         *Segment
		S2         *Fixedsegment
		QS3        *Quadraticsegment
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
				FS0: &Flatsegment{
					X0: sdk.MustNewDecFromStr("46479149.746915109636077002"),
					Y:  sdk.MustNewDecFromStr("0.020408732850910393"),
				},
				S0: &Segment{
					P0:     sdk.MustNewDecFromStr("0.020408732850910393"),
					A:      sdk.MustNewDecFromStr("0.020408732850910393"),
					B:      sdk.MustNewDecFromStr("0.028367859565819354"),
					P1:     sdk.MustNewDecFromStr("0.036326986280728314"),
					P0X:    sdk.MustNewDecFromStr("46479149.746915109636077002"),
					P1X:    sdk.MustNewDecFromStr("2038258170.761775649954248590"),
					DeltaX: sdk.MustNewDecFromStr("1991779021.014860540318171588"),
				},
				S1: &Segment{
					P0:     sdk.MustNewDecFromStr("0.036326986280728314"),
					A:      sdk.MustNewDecFromStr("0.040355826048446913"),
					B:      sdk.MustNewDecFromStr("0.066392634033828685"),
					P1:     sdk.MustNewDecFromStr("0.104673777868010282"),
					P0X:    sdk.MustNewDecFromStr("2038258170.761775649954248590"),
					P1X:    sdk.MustNewDecFromStr("3046479149.746915109636077002"),
					DeltaX: sdk.MustNewDecFromStr("1008220978.985139459681828412"),
				},
				S2: &Fixedsegment{
					Segment: &Segment{
						P0:     sdk.MustNewDecFromStr("0.100000000000000000"),
						A:      sdk.MustNewDecFromStr("0.200000000000000000"),
						B:      sdk.MustNewDecFromStr("0.333333333000000001"),
						P1:     sdk.MustNewDecFromStr("1.000000000000000000"),
						P0X:    sdk.MustNewDecFromStr("3000000000.000000000000000000"),
						P1X:    sdk.MustNewDecFromStr("6000000000.000000000000000000"),
						DeltaX: sdk.MustNewDecFromStr("3000000000.000000000000000000"),
					},
					IntervalP0X: sdk.MustNewDecFromStr("3046479149.746915109636077002"),
				},
				QS3: &Quadraticsegment{
					A:             sdk.MustNewDecFromStr("0.533333334"),
					B:             sdk.MustNewDecFromStr("-5.733333341"),
					C:             sdk.MustNewDecFromStr("16.200000022"),
					ScalingFactor: sdk.MustNewDecFromStr("1000000000"),
				},
				BPool:      sdk.MustNewDecFromStr("115125491.898143000000000000"),
				BPoolUnder: sdk.MustNewDecFromStr("114756324.035772563000000000"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ubc := setupUbcAndByToken(t, tt.args.TokensToBuy)
			tt.wantErr(t, ubc.ShiftUp(tt.args.BPoolAdd, tt.args.DegirdingFactor),
				fmt.Sprintf("BuyToken(%v) then ShiftUp(%v, %v)", tt.args.TokensToBuy, tt.args.BPoolAdd, tt.args.DegirdingFactor))
			assert.EqualValues(t, tt.wants.FS0, ubc.FS0)
			assert.EqualValues(t, tt.wants.S0, ubc.S0)
			assert.EqualValues(t, tt.wants.S1, ubc.S1)
			assert.EqualValues(t, tt.wants.S2, ubc.S2)
			assert.EqualValues(t, tt.wants.QS3, ubc.QS3)
			assert.Equal(t, tt.wants.BPool, ubc.BPool)
			assert.Equal(t, tt.wants.BPoolUnder, ubc.BPoolUnder)
		})
	}
}

func setupUbcAndByToken(t *testing.T, tokensToBuy sdk.Dec) Curve {
	t.Helper()
	ubc := validUbcParams()
	require.NoError(t, ubc.Fit())
	IsEqualDecimal(t, "100000000", ubc.BPool)
	IsEqualDecimal(t, "100000000", ubc.BPoolUnder)
	IsEqualDecimal(t, "3000000000", ubc.CurrentSupply)
	ubc.Buy(tokensToBuy)
	return ubc
}
