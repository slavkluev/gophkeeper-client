package account

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/domain/models"
	"client/grpc/app/account"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of accounts",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("accounts_target")
		timeout := viper.GetInt("timeout")

		accountApp, err := account.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		accounts, err := accountApp.AccClient.List(context.Background())
		if err != nil {
			panic(err)
		}

		err = printTable(accounts)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	accountCmd.AddCommand(listCmd)
}

func printTable(accounts []models.Account) error {
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

	_, err := fmt.Fprintln(w, "ID\tLOGIN\tPASSWORD\tINFO")
	if err != nil {
		return err
	}

	for _, acc := range accounts {
		row := fmt.Sprintf("%d\t%s\t%s\t%s", acc.ID, acc.Login, acc.Pass, acc.Info)
		_, err := fmt.Fprintln(w, row)
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}
