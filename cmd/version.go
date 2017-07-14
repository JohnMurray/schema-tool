package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of schema-tool",
	Long:  "All software has versions. This is schema-tool's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("schema-tool v1.0")
	},
}
