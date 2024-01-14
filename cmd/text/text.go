package text

import (
	"github.com/spf13/cobra"

	"client/cmd"
)

var textCmd = &cobra.Command{
	Use: "text",
}

func init() {
	cmd.RootCmd.AddCommand(textCmd)
}
