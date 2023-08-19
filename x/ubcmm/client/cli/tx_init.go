// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"strconv"

	"github.com/catenax/esc-backbone/x/ubcmm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [ref-token-supply] [ref-token-price] [ref-profit-factor] [b-pool] [b-pool-under] [slope-p-2] [slope-p-3] [factor-fy] [factor-fxy]",
		Short: "Broadcast message init",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRefTokenSupply := args[0]
			argRefTokenPrice := args[1]
			argRefProfitFactor := args[2]
			argBPool := args[3]
			argBPoolUnder := args[4]
			argSlopeP2 := args[5]
			argSlopeP3 := args[6]
			argFactorFy := args[7]
			argFactorFxy := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInit(
				clientCtx.GetFromAddress().String(),
				argRefTokenSupply,
				argRefTokenPrice,
				argRefProfitFactor,
				argBPool,
				argBPoolUnder,
				argSlopeP2,
				argSlopeP3,
				argFactorFy,
				argFactorFxy,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
