package account

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/grpc/app/account"
)

type AddRequest struct {
	login    string
	password string
	info     string
}

var addRequest AddRequest

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add account",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("accounts_target")
		timeout := viper.GetInt("timeout")

		accountApp, err := account.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		_, err = accountApp.AccClient.Save(
			context.Background(),
			addRequest.login,
			addRequest.password,
			addRequest.info,
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("add was successful")
	},
}

func init() {
	accountCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addRequest.login, "login", "l", "", "Login")
	addCmd.Flags().StringVarP(&addRequest.password, "password", "p", "", "Password")
	addCmd.Flags().StringVarP(&addRequest.info, "info", "i", "", "Info")
}
