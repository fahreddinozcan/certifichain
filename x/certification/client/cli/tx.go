package cli

import (
	"fmt"
	"strconv"

	"certifichain/x/certification/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdIssueCertification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-certification [id] [issuer-id] [learner-id] [title] [description] [issue-date] [expiry-date] [skills]",
		Short: "Issue a new certification",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Convert string arguments to appropriate types
			issueDate, err := strconv.ParseInt(args[5], 10, 64)
			if err != nil {
				return err
			}
			expiryDate, err := strconv.ParseInt(args[6], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueCertification(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
				args[3],
				args[4],
				issueDate,
				expiryDate,
				args[7],
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

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdIssueCertification(),
		// Add other transaction commands here
	)

	return cmd
}
