package cmd

import (
	"github.com/pranganmajumder/client-go/api"
	"github.com/spf13/cobra"
)

var deleteDeployCMD = &cobra.Command{
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
		api.DeleteDeployment()
	},
}

func init() {
	deleteCMD.AddCommand(deleteDeployCMD)
}
