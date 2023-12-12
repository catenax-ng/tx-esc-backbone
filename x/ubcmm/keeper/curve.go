// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package keeper

import (
	"github.com/catenax/esc-backbone/x/ubcmm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCurve set curve in the store
func (k Keeper) SetCurve(ctx sdk.Context, curve types.Curve) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurveKey))
	b := k.cdc.MustMarshal(&curve)
	store.Set([]byte{0}, b)
}

// GetCurve returns curve
func (k Keeper) GetCurve(ctx sdk.Context) (val types.Curve, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurveKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	val.PopulateSegments()
	return val, true
}

// RemoveCurve removes curve from the store
func (k Keeper) RemoveCurve(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurveKey))
	store.Delete([]byte{0})
}
