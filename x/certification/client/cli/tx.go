package cli

import (
	"fmt"
	"strconv"
	"strings"

	"certifichain/x/certification/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdIssueCertification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-certification [id] [issuer-id] [learner-id] [title] [description] [issue-date] [expiry-date] [skills] [hash]",
		Short: "Issue a new certification",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			issueDate, err := strconv.ParseInt(args[5], 10, 64)
			if err != nil {
				return err
			}
			expiryDate, err := strconv.ParseInt(args[6], 10, 64)
			if err != nil {
				return err
			}

			skills := strings.Split(args[7], ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueCertification(
				args[0], // id
				args[1], // issuer-id
				args[2], // learner-id
				args[3], // title
				args[4], // description
				issueDate,
				expiryDate,
				skills,                              // skills (as a slice)
				args[8],                             // hash
				clientCtx.GetFromAddress().String(), // creator
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
