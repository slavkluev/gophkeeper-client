package card

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/grpc/app/card"
)

type AddRequest struct {
	number string
	cvv    string
	month  string
	year   string
	info   string
}

var addRequest AddRequest

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add card",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("cards_target")
		timeout := viper.GetInt("timeout")

		cardApp, err := card.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		_, err = cardApp.CardClient.Save(
			context.Background(),
			addRequest.number,
			addRequest.cvv,
			addRequest.month,
			addRequest.year,
			addRequest.info,
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("add was successful")
	},
}

func init() {
	cardCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addRequest.number, "number", "n", "", "Number")
	addCmd.Flags().StringVarP(&addRequest.cvv, "cvv", "c", "", "CVV")
	addCmd.Flags().StringVarP(&addRequest.month, "month", "m", "", "Month")
	addCmd.Flags().StringVarP(&addRequest.year, "year", "y", "", "Year")
	addCmd.Flags().StringVarP(&addRequest.info, "info", "i", "", "Info")
}
