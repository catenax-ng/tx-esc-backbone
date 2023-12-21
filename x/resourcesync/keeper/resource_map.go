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
package keeper

import (
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetResourceMap set a specific resourceMap in the store from its index
func (k Keeper) SetResourceMap(ctx sdk.Context, resourceMap types.ResourceMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceMapKeyPrefix))
	b := k.cdc.MustMarshal(&resourceMap)
	store.Set(types.ResourceMapKey(
		resourceMap.Originator,
		resourceMap.OrigResId,
	), b)
}

// GetResourceMap returns a resourceMap from its index
func (k Keeper) GetResourceMap(
	ctx sdk.Context,
	resourceKey types.ResourceKey,
) (val types.ResourceMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceMapKeyPrefix))

	b := store.Get(types.ResourceMapKeyOf(resourceKey))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveResourceMap removes a resourceMap from the store
func (k Keeper) RemoveResourceMap(
	ctx sdk.Context,
	resourceKey types.ResourceKey,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceMapKeyPrefix))
	store.Delete(types.ResourceMapKeyOf(resourceKey))
}

func (k Keeper) RemoveAndGetResourceMap(ctx sdk.Context, resourceKey types.ResourceKey) (*types.ResourceMap, bool) {
	result, found := k.GetResourceMap(ctx, resourceKey)
	removed := &result
	if found {
		k.RemoveResourceMap(ctx, resourceKey)
	} else {
		removed = nil
	}
	return removed, found
}

// GetAllResourceMap returns all resourceMap
func (k Keeper) GetAllResourceMap(ctx sdk.Context) (list []types.ResourceMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ResourceMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) HasResourceMapFor(ctx sdk.Context, resourceKey types.ResourceKey) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceMapKeyPrefix))
	key := types.ResourceMapKeyOf(resourceKey)
	return store.Has(key)
}
