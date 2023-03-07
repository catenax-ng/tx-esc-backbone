// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
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
