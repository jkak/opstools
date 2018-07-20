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
	"github.com/jkak/opstools/os/cmd/dusub2"
	"github.com/spf13/cobra"
)

// dusub2Cmd represents the dusub2 command
var dusub2Cmd = &cobra.Command{
	Use:   "dusub2",
	Short: "summary the disk usage of each sub dir for given path like linux du",
	Long:  "summary the disk usage of each sub dir for given path like linux du",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		for _, p := range args {
			dusub2.Dusub2(p, ignores)
		}
	},
}

var ignores string

func init() {
	rootCmd.AddCommand(dusub2Cmd)

	dusub2Cmd.Flags().StringVarP(&ignores, "ignore", "i", "", "ignore dir name seperate by ','")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dusub2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dusub2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
