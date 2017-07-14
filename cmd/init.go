package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var forceInit bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long: `
Initialize a new project including installing any pre-commit hook
defined in the config (if any) and setting up a history table in the
DB for revision tracking.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().BoolVarP(&forceInit, "force", "f", false,
		"forcibly initialize the history table (wiping all old data)")
}
