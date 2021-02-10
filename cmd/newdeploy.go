package cmd

import (
	"fmt"

	"github.com/pranganmajumder/client-go/api"

	"github.com/spf13/cobra"
)

//CreateDeployCMD is the subcommand of command create . like: ./app create new-deploy
var createDeployCMD = &cobra.Command{
	Use:   "new-deploy",
	Short: "It is a subcommand of command create",
	Long:  `It'll be used as a subcommand of command create, Like you can use it as ./app create new-deploy'`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create deployment called")
		api.CreateDeployment(defaultImage, defaultReplica)
	},
}

var (
	defaultImage   string
	defaultReplica int32
)

func init() {
	// add the sub-command deploy with main command create
	createCMD.AddCommand(createDeployCMD)

	// Now add flag with new-deploy
	createDeployCMD.PersistentFlags().StringVarP(&defaultImage, "image", "i", "pranganmajumder/go-basic-restapi:1.0.0", "It sets the custom image you want")
	createDeployCMD.PersistentFlags().Int32VarP(&defaultReplica, "replica", "r", 2, "It sets the number of replica user want")

}
