package cli

import (
	"strconv"

	"encoding/json"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateResource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-resource [entry]",
		Short: "Broadcast message CreateResource",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEntry := new(types.Resource)
			err = json.Unmarshal([]byte(args[0]), argEntry)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateResource(
				clientCtx.GetFromAddress().String(),
				argEntry,
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
