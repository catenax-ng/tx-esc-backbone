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

func TestMsgBuy_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgBuy
		err  error
	}{
		{
			name: "valid",
			msg: MsgBuy{
				Buyer: sample.AccAddress(),
				Value: "10" + SystemTokenDenom,
			},
			err: nil,
		}, {
			name: "zero value",
			msg: MsgBuy{
				Buyer: sample.AccAddress(),
				Value: "0" + SystemTokenDenom,
			},
			err: ErrInvalidArg,
		}, {
			name: "negative value",
			msg: MsgBuy{
				Buyer: sample.AccAddress(),
				Value: "-5" + SystemTokenDenom,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid address",
			msg: MsgBuy{
				Buyer: "invalid_address",
				Value: "10" + SystemTokenDenom,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid value",
			msg: MsgBuy{
				Buyer: sample.AccAddress(),
				Value: "abcd",
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid denom",
			msg: MsgBuy{
				Buyer: sample.AccAddress(),
				Value: "2" + SystemTokenDenom + "x",
			},
			err: ErrInvalidArg,
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
