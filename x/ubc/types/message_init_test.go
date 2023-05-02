// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"errors"
	"testing"

	"github.com/catenax/esc-backbone/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgInit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		addr string
		err  error
	}{
		{
			name: "invalid address",
			addr: "invalid_address",
			err:  sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			addr: sample.AccAddress(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := validMsgInitUbc()
			msg.Creator = tt.addr
			err := msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgInit_ParseUbcobject_Happy(t *testing.T) {
	msg := validMsgInitUbc()
	ubc, err := msg.ParseUbcobject()
	require.NoError(t, err)
	IsEqualDecimal(t, "6000000000", ubc.RefTokenSupply)
	IsEqualDecimal(t, "1", ubc.RefTokenPrice)
	IsEqualDecimal(t, "10", ubc.RefProfitFactor)
	IsEqualDecimal(t, "100000000", ubc.BPool)
	IsEqualDecimal(t, "100000000", ubc.BPoolUnder)
	IsEqualDecimal(t, "0.0000000001", ubc.SlopeP2)
	IsEqualDecimal(t, "0.000000000666666667", ubc.SlopeP3)
	IsEqualDecimal(t, "0.2", ubc.FactorFy)
	IsEqualDecimal(t, "15832600001", ubc.FactorFxy)
}

func TestMsgInit_ParseUbcobject_Error(t *testing.T) {
	type test struct {
		name     string
		modifier func(msg *MsgInit)
	}

	tests := []test{
		{"RefTokenSupply",
			func(msg *MsgInit) { msg.RefTokenSupply = "abc" }},
		{"RefTokenPrice",
			func(msg *MsgInit) { msg.RefTokenPrice = "abc" }},
		{"RefProfitFactor",
			func(msg *MsgInit) { msg.RefProfitFactor = "abc" }},
		{"BPoolUnder",
			func(msg *MsgInit) { msg.BPoolUnder = "abc" }},
		{"SlopeP2",
			func(msg *MsgInit) { msg.SlopeP2 = "abc" }},
		{"SlopeP3",
			func(msg *MsgInit) { msg.SlopeP3 = "abc" }},
		{"FactorFy",
			func(msg *MsgInit) { msg.FactorFy = "abc" }},
		{"FactorFxy",
			func(msg *MsgInit) { msg.FactorFxy = "abc" }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := validMsgInitUbc()
			tc.modifier(msg)
			_, err := msg.ParseUbcobject()
			require.Error(t, err)
			require.True(t, errors.Is(err, ErrInvalidArg))
			t.Logf("%v", err)
		})
	}
}

func validMsgInitUbc() *MsgInit {
	return &MsgInit{
		RefTokenSupply:  "6000000000",
		RefTokenPrice:   "1",
		RefProfitFactor: "10",
		BPool:           "100000000",
		BPoolUnder:      "100000000",
		SlopeP2:         "0.0000000001",
		SlopeP3:         "0.000000000666666667",
		FactorFy:        "0.2",
		FactorFxy:       "15832600001",
	}
}
