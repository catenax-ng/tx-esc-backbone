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
