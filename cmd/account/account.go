package account

import (
	"github.com/spf13/cobra"

	"client/cmd"
)

var accountCmd = &cobra.Command{
	Use: "account",
}

func init() {
	cmd.RootCmd.AddCommand(accountCmd)
}
