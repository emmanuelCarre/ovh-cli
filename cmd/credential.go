package cmd

import (
	"github.com/emmanuelCarre/ovh-cli/ovh"
	"github.com/emmanuelCarre/ovh-cli/utils"

	"github.com/spf13/cobra"
)

func getHeaderToDisplayForCreds() []string {
	return []string{"Id", "Expiration date", "Application ID", "Application name", "Right"}
}

var listCredentialCmd = &cobra.Command{
	Use:   "list",
	Short: "list credentials",
	Run: func(cmd *cobra.Command, args []string) {
		client := ovh.GetOvhClient()

		credentials := ovh.FetchCredentials(client)
		table := utils.GetTable()
		table.SetHeader(getHeaderToDisplay())
		for _, credential := range credentials {
			table.Append(credential.ToArrow())
		}
		table.Render()
	},
}

var deleteCredentialCmd = &cobra.Command{
	Use:   "delete [credential id]...",
	Short: "delete credentials",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := ovh.GetOvhClient()
		ovh.DeleteCredentials(client, args)
	},
}

var credentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "Manipulate credentials",
}

func init() {
	rootCmd.AddCommand(credentialCmd)
	credentialCmd.AddCommand(listCredentialCmd)
	credentialCmd.AddCommand(deleteCredentialCmd)
}
