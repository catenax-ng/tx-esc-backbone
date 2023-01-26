package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
)

func (k msgServer) DeleteResource(goCtx context.Context, msg *types.MsgDeleteResource) (*types.MsgDeleteResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	resourceKey, err := types.NewResourceKeyForDelete(msg)
	if err != nil {
		return nil, err
	}
	if !k.Keeper.HasResourceMapFor(ctx, resourceKey) {
		return nil, sdkerrors.Wrapf(types.ErrNonexistentResource, "resource %s/%s cannot be deleted: nonexistent", resourceKey.GetOriginator(), resourceKey.GetOrigResKey())
	}
	k.Keeper.RemoveResourceMap(ctx, resourceKey)
	return &types.MsgDeleteResourceResponse{}, nil
}
