package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reverseList bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the current alter chain",
	Long: `
List the current alters in order for how they are applied.
Note that a '*' indicates whether or not the alter has been
applied to the database in the current configs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolVarP(&reverseList, "reverse", "r", false,
		"List the contents of current alter chain in reverse order")
}
