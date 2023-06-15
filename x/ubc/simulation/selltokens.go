// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package simulation

import (
	"math/rand"

	"github.com/catenax/esc-backbone/x/ubc/keeper"
	"github.com/catenax/esc-backbone/x/ubc/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSelltokens(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSelltokens{
			Seller: simAccount.Address.String(),
		}

		// TODO: Handling the Selltokens simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Selltokens simulation not implemented"), nil, nil
	}
}
