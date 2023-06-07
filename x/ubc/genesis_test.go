// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package ubc_test

import (
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/ubc"
	"github.com/catenax/esc-backbone/x/ubc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Ubcobject: &types.Ubcobject{
			FS0:             "97",
			S0:              "8",
			S1:              "5",
			S2:              "19",
			QS3:             "44",
			RefProfitFactor: "42",
			RefTokenSupply:  "8",
			RefTokenPrice:   "91",
			BPool:           "99",
			BPoolUnder:      "35",
			FactorFy:        "57",
			FactorFxy:       "8",
			TradingPoint:    "96",
			CurrentSupply:   "41",
			Slopep2:         "69",
			Slopep3:         "62",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UbcKeeper(t)
	ubc.InitGenesis(ctx, *k, genesisState)
	got := ubc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Ubcobject, got.Ubcobject)
	// this line is used by starport scaffolding # genesis/test/assert
}
