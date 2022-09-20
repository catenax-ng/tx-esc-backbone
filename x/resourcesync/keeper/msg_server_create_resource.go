package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateResource(goCtx context.Context, msg *types.MsgCreateResource) (*types.MsgCreateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	resource := *msg.Entry
	err := resource.Validate()
	if err != nil {
		return nil, err
	}
	if k.Keeper.HasResourceMapFor(ctx, resource) {
		return nil, sdkerrors.Wrapf(types.ErrDuplicateResource, "resource %s/%s cannot be created: duplicate", resource.Originator, resource.OrigResId)
	}
	resourceMap := types.NewResourceMap(resource)
	k.Keeper.SetResourceMap(ctx, resourceMap)
	return &types.MsgCreateResourceResponse{}, nil
}
