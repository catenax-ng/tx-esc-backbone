package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
)

func CmdListResourceMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-resource-map",
		Short: "list all resourceMap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllResourceMapRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ResourceMapAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowResourceMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-resource-map [originator] [orig-res-id]",
		Short: "shows a resourceMap",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argOriginator := args[0]
			argOrigResId := args[1]

			params := &types.QueryGetResourceMapRequest{
				Originator: argOriginator,
				OrigResId:  argOrigResId,
			}

			res, err := queryClient.ResourceMap(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
