package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check that your local alter-chain is well formed",
	Long: `
Determines if your alter-chain is well-formed. This includes such
things as determining if a root exists, each non-root alter has a
parent, each parent only has one child, etc.

These checks are run by default as part of many other commands. This
command is exposed for user scripts/manual-testing to more easily
identify issues with the alter-chain.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
	},
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
