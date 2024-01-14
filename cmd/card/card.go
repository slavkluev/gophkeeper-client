package card

import (
	"github.com/spf13/cobra"

	"client/cmd"
)

var cardCmd = &cobra.Command{
	Use: "card",
}

func init() {
	cmd.RootCmd.AddCommand(cardCmd)
}
