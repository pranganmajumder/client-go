package cmd

import (
	"github.com/spf13/cobra"
)

//UpdateCMD is the root cammand . like : ./app update
var updateCMD = &cobra.Command{
	Use:   "update",
	Short: "It is a a root command , which will update the resource you want",
	Long:  `It'll be used as a root of command, Like you can use it as ./app update'`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(updateCMD)
}
