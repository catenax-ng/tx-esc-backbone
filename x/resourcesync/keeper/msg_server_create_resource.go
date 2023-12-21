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
	resourceKey, err := resource.ToResourceKey()
	if err != nil {
		return nil, err
	}
	if k.Keeper.HasResourceMapFor(ctx, resourceKey) {
		return nil, sdkerrors.Wrapf(types.ErrDuplicateResource, "resource %s/%s cannot be created: duplicate", resource.Originator, resource.OrigResId)
	}
	resourceMap := types.NewResourceMap(resource)
	k.Keeper.SetResourceMap(ctx, resourceMap)
	err2 := ctx.EventManager().EmitTypedEvent(&types.EventCreateResource{
		Creator:  msg.Creator,
		Resource: resource,
	})
	if err2 != nil {
		return nil, err2
	}
	return &types.MsgCreateResourceResponse{}, nil
}
