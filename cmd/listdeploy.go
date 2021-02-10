package cmd

import (
	"github.com/pranganmajumder/client-go/api"

	"github.com/spf13/cobra"
)

//listDeployCMD is the subcommand of command get . like: ./app get list-deploy
var listDeployCMD = &cobra.Command{
	Use:   "list-deploy",
	Short: "It will show the deployment you want",
	Long: `It show a resources like: ./app get list-deploy . Here list-deploy is our another child/
			sub-command of get command . You need to create the list-deploy sub-command manually and add it to getCMD`,

	Run: func(cmd *cobra.Command, args []string) {
		api.GetDeployment()
	},
}

func init() {
	getCMD.AddCommand(listDeployCMD)
}
