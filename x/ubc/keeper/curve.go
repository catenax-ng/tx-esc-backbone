// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package keeper

import (
	"github.com/catenax/esc-backbone/x/ubc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetUbcobject set ubcobject in the store
func (k Keeper) SetUbcobject(ctx sdk.Context, ubcobject types.Ubcobject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UbcobjectKey))
	b := k.cdc.MustMarshal(&ubcobject)
	store.Set([]byte{0}, b)
}

// GetUbcobject returns ubcobject
func (k Keeper) GetUbcobject(ctx sdk.Context) (val types.Ubcobject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UbcobjectKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUbcobject removes ubcobject from the store
func (k Keeper) RemoveUbcobject(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UbcobjectKey))
	store.Delete([]byte{0})
}
