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

	"github.com/catenax/esc-backbone/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUndergird_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUndergird
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUndergird{
				Operator:      "invalid_address",
				Voucherstoadd: "10" + VoucherDenom,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid vouchers to add (zero)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "0" + VoucherDenom,
			},
			err: ErrInvalidArg,
		}, {
			name: "invalid vouchers to add (negative)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "-10" + VoucherDenom,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid vouchers to add (decimal)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "10.0" + VoucherDenom,
			},
			err: nil,
		}, {
			name: "invalid vouchers to add (non number)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "abcd" + VoucherDenom,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid vouchers to add (empty)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "",
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "valid address",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "10" + VoucherDenom,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
