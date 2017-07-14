package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var filenameNew string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new alter",
	Long: `
Create a new alter in the chain. Two files will be created for an
'up' and 'down' alter. The current alter chain must be valid
(verified by the 'check' command) as the newly created alter will
be added to the end of the chain.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

func init() {
	RootCmd.AddCommand(newCmd)

	newCmd.PersistentFlags().StringVarP(&filenameNew, "file", "f", "",
		"The name of the new file (without file-extension)")
}
