package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/alice/checkers/x/checkers/types"
)

var _ = strconv.Itoa(0)

func CmdToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token [tokens]",
		Short: "Query token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqTokens := args[0]
			
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTokenRequest{
				
                Tokens: reqTokens, 
            }

            

			res, err := queryClient.Token(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}