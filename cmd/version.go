package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// GitCommit is git commit hash and define during build.
	GitCommit string
	// Version is application version and define during build.
	Version string
	// GoVersion is Golang version and define during build.
	GoVersion string
	// OsArchi is Golang architectuer and define during build.
	OsArchi string
	// BuildDate is build date and define during build.
	BuildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:    %s\n", Version)
		fmt.Printf("Commit:     %s\n", GitCommit)
		fmt.Printf("Build date: %s\n", BuildDate)
		fmt.Printf("Go version: %s\n", GoVersion)
		fmt.Printf("OS/Arch::   %s\n", OsArchi)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
