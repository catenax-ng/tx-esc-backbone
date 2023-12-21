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

func (m *Resource) GetOriginatorAddress() (originator sdk.AccAddress, err error) {
	originator, err = sdk.AccAddressFromBech32(m.Originator)
	return originator, sdkerrors.Wrapf(err, ErrInvalidOriginator.Error(), originator)
}
func (m *Resource) GetOrigResKey() (origResKey string, err error) {
	// FIXME enforce conditions on origResKey
	return m.OrigResId, nil
}
func (m *Resource) GetResourceKey() (resourceKey string, err error) {
	// FIXME enforce conditions on resourceKey
	return m.ResourceKey, nil
}
func (m *Resource) GetTargetSystem() (targetSystem string, err error) {
	// FIXME enforce conditions on targetSystem
	return m.TargetSystem, nil
}

func (m *Resource) GetDataHash() (dataHash []byte, err error) {
	// FIXME enforce conditions on dataHash
	return m.DataHash, nil
}

func (m *Resource) ToResourceKey() (ResourceKey, error) {
	return NewResourceKey(m.Originator, m.OrigResId)
}

func (m *Resource) Validate() (err error) {
	_, err = m.GetOriginatorAddress()
	if err != nil {
		return err
	}
	_, err = m.GetOrigResKey()
	if err != nil {
		return err
	}
	_, err = m.GetTargetSystem()
	if err != nil {
		return err
	}
	_, err = m.GetResourceKey()
	if err != nil {
		return err
	}

	_, err = m.GetDataHash()
	if err != nil {
		return err
	}
	return nil
}
