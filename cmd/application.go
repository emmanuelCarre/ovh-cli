package cmd

import (
	"github.com/emmanuelCarre/ovh-cli/ovh"
	"github.com/emmanuelCarre/ovh-cli/utils"
	"github.com/spf13/cobra"
)

const applicationURI string = "/me/api/application/"

func getHeaderToDisplay() []string {
	return []string{"Id", "Key", "Name", "Description", "Status"}
}

var listApplicationCmd = &cobra.Command{
	Use:   "list",
	Short: "list applications",
	Run: func(cmd *cobra.Command, args []string) {
		client := ovh.GetOvhClient()

		applications := ovh.FetchApplications(client)
		table := utils.GetTable()
		table.SetHeader(getHeaderToDisplay())
		for _, application := range applications {
			table.Append(application.ToArrow())
		}
		table.Render()
	},
}

var deleteApplicationCmd = &cobra.Command{
	Use:   "delete [application id]...",
	Short: "delete applications",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := ovh.GetOvhClient()
		ovh.DeleteApplications(client, args)
	},
}

var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manipulate applications",
}

func init() {
	rootCmd.AddCommand(applicationCmd)
	applicationCmd.AddCommand(listApplicationCmd)
	applicationCmd.AddCommand(deleteApplicationCmd)
}
