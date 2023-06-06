package resourcesync_test

import (
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/resourcesync"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ResourceMapList: []types.ResourceMap{
			{
				Originator: "0",
				OrigResId:  "0",
			},
			{
				Originator: "1",
				OrigResId:  "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ResourcesyncKeeper(t)
	resourcesync.InitGenesis(ctx, *k, genesisState)
	got := resourcesync.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ResourceMapList, got.ResourceMapList)
	// this line is used by starport scaffolding # genesis/test/assert
}
