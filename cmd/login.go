package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/grpc/app/auth"
)

var (
	login    string
	password string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		target := viper.GetString("sso_target")
		appID := viper.GetInt64("app_id")

		authApp, err := auth.New(target, appID)
		if err != nil {
			panic(err)
		}

		accessToken, err := authApp.AuthClient.Login(context.Background(), login, password)
		if err != nil {
			panic(err)
		}

		viper.Set("access_token", accessToken)

		err = viper.WriteConfig()
		if err != nil {
			panic(err)
		}

		fmt.Println("login was successful")
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&login, "login", "l", "", "Login")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
}
