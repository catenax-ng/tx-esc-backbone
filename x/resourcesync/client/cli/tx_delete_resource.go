package cli

import (
	"strconv"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDeleteResource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-resource [originator] [orig-res-id]",
		Short: "Broadcast message DeleteResource",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOriginator := args[0]
			argOrigResId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteResource(
				clientCtx.GetFromAddress().String(),
				argOriginator,
				argOrigResId,
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
