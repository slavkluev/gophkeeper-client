package text

import (
	"github.com/spf13/cobra"

	"client/cmd"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "A brief description of your command",
}

func init() {
	cmd.RootCmd.AddCommand(textCmd)
}
