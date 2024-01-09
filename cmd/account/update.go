package account

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/grpc/app/account"
)

type UpdateRequest struct {
	id       uint64
	login    string
	password string
	info     string
}

var updateRequest UpdateRequest

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update account",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("accounts_target")
		timeout := viper.GetInt("timeout")

		accountApp, err := account.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		err = accountApp.AccClient.Update(
			context.Background(),
			updateRequest.id,
			updateRequest.login,
			updateRequest.password,
			updateRequest.info,
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("update was successful")
	},
}

func init() {
	accountCmd.AddCommand(updateCmd)

	updateCmd.Flags().Uint64VarP(&updateRequest.id, "id", "", 0, "Id")
	updateCmd.Flags().StringVarP(&updateRequest.login, "login", "l", "", "Login")
	updateCmd.Flags().StringVarP(&updateRequest.password, "password", "p", "", "Password")
	updateCmd.Flags().StringVarP(&updateRequest.info, "info", "i", "", "Info")
}
