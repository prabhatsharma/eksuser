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

	"github.com/prabhatsharma/eksuser/pkg/del"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an EKS user or users in an iamgroup. It does not delete the corresponding IAM user.",
	Long:  `Delete an EKS user or users in an iamgroup. It does not delete the corresponding IAM user.`,
	Run: func(cmd *cobra.Command, args []string) {

		// workaround for https://github.com/spf13/viper/issues/233
		viper.BindPFlag("iamgroup", cmd.Flags().Lookup("iamgroup"))
		viper.BindPFlag("user1", cmd.Flags().Lookup("user"))

		user := viper.GetString("user1")
		iamgroup := viper.GetString("iamgroup")

		if user != "" {
			del.DeleteUser(user)
		} else if iamgroup != "" {
			del.DeleteIAMGroup(iamgroup)
		} else {
			fmt.Fprintf(os.Stderr, "Error: --user or --iamgroup value not specified\n")
			cmd.Usage()
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("user", "u", "", "IAM user to be deleted from EKS. e.g. prabhat")
	deleteCmd.Flags().StringP("iamgroup", "i", "", "IAM group to be deleted from EKS. e.g. developers")

	viper.BindPFlag("user1", deleteCmd.Flags().Lookup("user"))
	viper.BindPFlag("iamgroup", deleteCmd.Flags().Lookup("iamgroup"))
}
