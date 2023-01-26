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
