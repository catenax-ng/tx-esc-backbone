package keeper

import (
	"context"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateResource(goCtx context.Context, msg *types.MsgUpdateResource) (*types.MsgUpdateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateResourceResponse{}, nil
}
