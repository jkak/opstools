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
	"github.com/jkak/opstools/domain/cmd/tls"
	"github.com/spf13/cobra"
)

// httpsCmd represents the https command
var httpsCmd = &cobra.Command{
	Use:   "https",
	Short: "check the expire time for https",
	Long:  "check the expire time for https of given domains",
	Run: func(cmd *cobra.Command, args []string) {
		tls.Check(gapDays, domains)
	},
}

var (
	gapDays int32
	domains string
)

func init() {
	rootCmd.AddCommand(httpsCmd)
	httpsCmd.Flags().StringVarP(&domains, "domain", "d", "baidu.com", "domains like: abc.com,z.cn")
	httpsCmd.Flags().Int32VarP(&gapDays, "gaps", "g", 60, "default gap of expire days for domain")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
