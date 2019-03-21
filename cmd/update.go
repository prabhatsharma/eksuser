// Copyright © 2018 Prabhat Sharma <hi.prabhat@gmail.com>
//
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

package cmd

import (
	"fmt"
	"os"

	"github.com/prabhatsharma/eksuser/pkg/add"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an IAM user in EKS",
	Long: `Update an IAM user to EKS. Specify compulsory flags --user and --group to which the user will be added. For example:
	$ eksuser update --user=prabhat --group=system:masters
	$ eksuser update --user=prabhat --group=developer,ops
	$ eksuser update --iamgroup=admin --group=system:masters`,
	Run: func(cmd *cobra.Command, args []string) {

		// workaround for https://github.com/spf13/viper/issues/233
		// viper.BindPFlag("user1", cmd.Flags().Lookup("user"))
		// viper.BindPFlag("group", cmd.Flags().Lookup("group"))
		// viper.BindPFlag("iamgroup", updateCmd.Flags().Lookup("iamgroup"))

		user := viper.GetString("user1")

		if user == "" {
			fmt.Fprintf(os.Stderr, "Error: user not specified\n")
			cmd.Usage()
			os.Exit(1)
		}

		group := viper.GetString("group")

		if group == "" {
			fmt.Fprintf(os.Stderr, "Error: group not specified\n")
			cmd.Usage()
			os.Exit(1)
		} else {
			// Add and Update offer same functionality. So just call add in update
			add.InsertUser(user, group)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("user", "u", "", "IAM user to be updated in EKS. e.g. prabhat")
	updateCmd.Flags().StringP("group", "g", "", "kubernetes group(s) to which user will be added. e.g. system:masters")
	updateCmd.Flags().StringP("iamgroup", "i", "", "IAM group to be added to EKS. e.g. developers")

	viper.BindPFlag("user", updateCmd.Flags().Lookup("user"))
	viper.BindPFlag("group", updateCmd.Flags().Lookup("group"))
	viper.BindPFlag("iamgroup", updateCmd.Flags().Lookup("iamgroup"))
}
