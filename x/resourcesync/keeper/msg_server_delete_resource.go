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
	removed, found := k.Keeper.RemoveAndGetResourceMap(ctx, resourceKey)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNonexistentResource, "resource %s/%s cannot be deleted: nonexistent", resourceKey.GetOriginator(), resourceKey.GetOrigResKey())
	}
	err2 := ctx.EventManager().EmitTypedEvent(&types.EventDeleteResource{
		Creator:  msg.Creator,
		Resource: removed.Resource,
	})
	if err2 != nil {
		return nil, err2
	}
	return &types.MsgDeleteResourceResponse{}, nil
}
