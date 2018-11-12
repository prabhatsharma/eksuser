// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	Short: "Delete an EKS user. It does not delete the corresponding IAM user.",
	Long:  `Delete an EKS user. It does not delete the corresponding IAM user.`,
	Run: func(cmd *cobra.Command, args []string) {

		// workaround for https://github.com/spf13/viper/issues/233
		viper.BindPFlag("user", cmd.Flags().Lookup("user"))

		user := viper.GetString("user")

		if user == "" {
			fmt.Fprintf(os.Stderr, "Error: user not specified\n")
			cmd.Usage()
			os.Exit(1)
		}

		del.DeleteUser(user)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("user", "u", "", "IAM user to be deleted from EKS. e.g. prabhat")

	viper.BindPFlag("user", deleteCmd.Flags().Lookup("user"))
}
