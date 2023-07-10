// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"strconv"

	"github.com/catenax/esc-backbone/x/ubc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUndergird() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "undergird [voucherstoadd]",
		Short: "Broadcast message undergird",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVoucherstoadd := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUndergird(
				clientCtx.GetFromAddress().String(),
				argVoucherstoadd,
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
