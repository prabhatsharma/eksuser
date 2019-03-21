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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/prabhatsharma/eksuser/pkg/add"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an IAM user to EKS",
	Long: `Add an IAM user to EKS. Specify compulsory flags --user and --group to which the user will be added. For example:
	$ eksuser add --user=prabhat --group=system:masters
	$ eksuser add --iamgroup=admin --group=system:masters
	$ eksuser add --user=prabhat --group=developer,ops`,
	Run: func(cmd *cobra.Command, args []string) {

		// workaround for https://github.com/spf13/viper/issues/233
		viper.BindPFlag("user1", cmd.Flags().Lookup("user"))
		viper.BindPFlag("group", cmd.Flags().Lookup("group"))
		viper.BindPFlag("iamgroup", cmd.Flags().Lookup("iamgroup"))

		user := viper.GetString("user1")
		iamgroup := viper.GetString("iamgroup")

		if user == "" && iamgroup == "" {
			fmt.Fprintf(os.Stderr, "Error: --user or --iamgroup value not specified. One of user or iamgroup must be provided\n")
			cmd.Usage()
			os.Exit(1)
		}

		group := viper.GetString("group")

		if group == "" {
			fmt.Fprintf(os.Stderr, "Error: kubernetes --group value not specified\n")
			cmd.Usage()
			os.Exit(1)
		} else {
			if user != "" {
				add.InsertUser(user, group)
			} else if iamgroup != "" {
				add.InsertIAMGroup(iamgroup, group)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("user", "u", "", "IAM user to be added to EKS. e.g. prabhat")
	addCmd.Flags().StringP("iamgroup", "i", "", "IAM group to be added to EKS. e.g. developers")
	addCmd.Flags().StringP("group", "g", "", "kubernetes group(s) to which user will be added. e.g. system:masters")

	viper.BindPFlag("user1", addCmd.Flags().Lookup("user"))
	viper.BindPFlag("group", addCmd.Flags().Lookup("group"))
	viper.BindPFlag("iamgroup", addCmd.Flags().Lookup("iamgroup"))
}
