/*
Copyright © 2025 Honoka Toda, Shinya Ishitobi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goat-cli",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.AddCommand(newRenderCmd())

	return cmd
}

var rootCmd = newRootCmd()

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
