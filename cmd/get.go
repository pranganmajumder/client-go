package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//CreateCMD is our main command
var getCMD = &cobra.Command{
	Use:   "get",
	Short: "It will get a resource you want",
	Long: `It show a resources like: ./app get list-deploy . Here list-deploy is our another child/
			sub-command of get command . You need to create the list-deploy sub-command manually and add it to getCMD`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get command called ")
	},
}

func init() {
	rootCmd.AddCommand(getCMD)
}
