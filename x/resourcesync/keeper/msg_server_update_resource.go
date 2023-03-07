// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateResource(goCtx context.Context, msg *types.MsgUpdateResource) (*types.MsgUpdateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	resource := *msg.Entry
	err := resource.Validate()
	if err != nil {
		return nil, err
	}
	resourceKey, err := resource.ToResourceKey()
	if err != nil {
		return nil, err
	}
	resourceMap, found := k.Keeper.GetResourceMap(ctx, resourceKey)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNonexistentResource, "resource %s/%s cannot be updated: does not exist", resource.Originator, resource.OrigResId)
	}
	resourceMap.Resource = resource
	k.Keeper.SetResourceMap(ctx, resourceMap)
	err2 := ctx.EventManager().EmitTypedEvent(&types.EventUpdateResource{
		Creator:  msg.Creator,
		Resource: resourceMap.Resource,
	})
	if err2 != nil {
		return nil, err2
	}
	return &types.MsgUpdateResourceResponse{}, nil
}
