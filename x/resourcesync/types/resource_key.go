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

type resourceKey struct {
	originator string
	origResId  string
}

type ResourceKey interface {
	GetOriginator() (originator string)
	GetOriginatorAddress() (originator sdk.AccAddress)
	GetOrigResKey() (origResKey string)
}

func NewResourceKeyForDelete(deleteMsg *MsgDeleteResource) (ResourceKey, error) {
	return NewResourceKey(deleteMsg.Originator, deleteMsg.OrigResId)
}
func NewResourceKey(originator string, origResId string) (ResourceKey, error) {
	err := checkOriginator(originator)
	if err != nil {
		return nil, err
	}
	return &resourceKey{
		originator: originator,
		origResId:  origResId,
	}, nil
}

func checkOriginator(originator string) error {
	_, err := sdk.AccAddressFromBech32(originator)
	return err
}

func (m resourceKey) GetOriginator() (originator string) {
	return m.originator
}

func (m resourceKey) GetOriginatorAddress() (originator sdk.AccAddress) {
	originator, _ = sdk.AccAddressFromBech32(m.originator)
	return originator
}

func (m resourceKey) GetOrigResKey() (origResKey string) {
	return m.origResId
}
