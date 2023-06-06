package keeper

import (
	"context"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ResourceMapAll(goCtx context.Context, req *types.QueryAllResourceMapRequest) (*types.QueryAllResourceMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var resourceMaps []types.ResourceMap
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	resourceMapStore := prefix.NewStore(store, types.KeyPrefix(types.ResourceMapKeyPrefix))

	pageRes, err := query.Paginate(resourceMapStore, req.Pagination, func(key []byte, value []byte) error {
		var resourceMap types.ResourceMap
		if err := k.cdc.Unmarshal(value, &resourceMap); err != nil {
			return err
		}

		resourceMaps = append(resourceMaps, resourceMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllResourceMapResponse{ResourceMap: resourceMaps, Pagination: pageRes}, nil
}

func (k Keeper) ResourceMap(goCtx context.Context, req *types.QueryGetResourceMapRequest) (*types.QueryGetResourceMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	resourceKey, err := types.NewResourceKey(req.Originator, req.OrigResId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetResourceMap(
		ctx,
		resourceKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetResourceMapResponse{ResourceMap: val}, nil
}
