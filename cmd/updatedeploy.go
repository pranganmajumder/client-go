package cmd

import (
	"github.com/pranganmajumder/client-go/api"
	"github.com/spf13/cobra"
)

//UpdateDeployCMD is the subcommand of command update . like: ./app update update-deploy
var updateDeployCMD = &cobra.Command{
	Use:   "update-deploy",
	Short: "It is a subcommand of command update",
	Long:  `It'll be used as a subcommand of command update, Like you can use it as ./app update update-deploy'`,

	Run: func(cmd *cobra.Command, args []string) {
		api.UpdateDeployment(updatedImage, updatedReplica)
	},
}

var (
	updatedImage   string
	updatedReplica int32
)

func init() {
	// add the sub-command deploy with main command update
	updateCMD.AddCommand(updateDeployCMD)

	// Now add flag with update-deploy
	updateDeployCMD.PersistentFlags().StringVarP(&updatedImage, "image", "i", "pranganmajumder/go-basic-restapi:1.0.0", "It'll update the current image set by the flag")
	updateDeployCMD.PersistentFlags().Int32VarP(&updatedReplica, "replica", "r", 1, "it'll update the number of replica")
}
