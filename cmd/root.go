package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ovh-cli",
	Short: "CLI to interact with OVH me API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute root command
func Execute() {
	rootCmd.Execute()
}
