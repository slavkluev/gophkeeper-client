package account

import (
	"github.com/spf13/cobra"

	"client/cmd"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "A brief description of your command",
}

func init() {
	cmd.RootCmd.AddCommand(accountCmd)
}
