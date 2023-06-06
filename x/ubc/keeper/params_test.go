package keeper_test

import (
	"testing"

	testkeeper "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/x/ubc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.UbcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
