package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ovh-cli",
	Short: "cli to interact with OVH me API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute root command
func Execute() {
	rootCmd.Execute()
}
