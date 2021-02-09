/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/pranganmajumder/client-go/api"
	"github.com/spf13/cobra"
)


// rootCmd represents the base command when called without any subcommands
var createCMD = &cobra.Command{
	Use:   "create-deploy",
	Short: "create a new deployment",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deployment called")
		api.CreateDeployment()
	},
}


var getCMD = &cobra.Command{
	Use:   "get-deploy",
	Short: "get the running deployment",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deployment called")
		api.GetDeployment()
	},
}

var updateCMD = &cobra.Command{
	Use:   "update-deploy",
	Short: "update the running deployment",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deployment called")
		api.UpdateDeployment()
	},
}


var deleteCMD = &cobra.Command{
	Use:   "delete-deploy",
	Short: "delete the running deployment",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deployment called")
		api.DeleteDeployment()
	},
}

// -------------------------------------- add the custom command to your rootCMD (root.go) --------------------------------------------------
func init() {
	rootCmd.AddCommand(createCMD)
	rootCmd.AddCommand(getCMD)
	rootCmd.AddCommand(updateCMD)
	rootCmd.AddCommand(deleteCMD)
}