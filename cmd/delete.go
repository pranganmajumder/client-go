package cmd

import (
	"github.com/spf13/cobra"
)

//deleteCMD is our main command to delete a resource like deployment, pod etc
var deleteCMD = &cobra.Command{
	Use:   "delete",
	Short: "It will delete a resource you want",
	Long: `It delete a resources like: ./app delete delete-deploy . Here delete-deploy is our another child/
			sub-command of delete command . You need to create the delete-deploy sub-command manually and add it to deleteCMD`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(deleteCMD)
}
