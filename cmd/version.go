package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"client/config"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"Build version: %s\nBuild date: %s\nBuild commit: %s\n",
			config.BuildVersion,
			config.BuildDate,
			config.BuildCommit,
		)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
