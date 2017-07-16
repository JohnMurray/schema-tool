// <--
// Copyright © 2017 John Murray <me@johnmurray.io>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -->
package cmd

import (
	"fmt"
	"os"

	"github.com/johnmurray/schema-tool/chain"
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
		if _, err := chain.ScanDirectory(cwDirGlobal); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
