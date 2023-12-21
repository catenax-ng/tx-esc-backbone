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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUndergird = "undergird"

var _ sdk.Msg = &MsgUndergird{}

func NewMsgUndergird(operator string, voucherstoadd string) *MsgUndergird {
	return &MsgUndergird{
		Operator:      operator,
		Voucherstoadd: voucherstoadd,
	}
}

func (msg *MsgUndergird) Route() string {
	return RouterKey
}

func (msg *MsgUndergird) Type() string {
	return TypeMsgUndergird
}

func (msg *MsgUndergird) GetSigners() []sdk.AccAddress {
	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operator}
}

func (msg *MsgUndergird) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUndergird) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}

	return validateVoucherCoin(msg.Voucherstoadd)
}
