package card

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/grpc/app/card"
)

type UpdateRequest struct {
	id     uint64
	number string
	cvv    string
	month  string
	year   string
	info   string
}

var updateRequest UpdateRequest

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update card",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("cards_target")
		timeout := viper.GetInt("timeout")

		cardApp, err := card.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		err = cardApp.CardClient.Update(
			context.Background(),
			updateRequest.id,
			updateRequest.number,
			updateRequest.cvv,
			updateRequest.month,
			updateRequest.year,
			updateRequest.info,
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("update was successful")
	},
}

func init() {
	cardCmd.AddCommand(updateCmd)

	updateCmd.Flags().Uint64VarP(&updateRequest.id, "id", "", 0, "Id")
	updateCmd.Flags().StringVarP(&updateRequest.number, "number", "n", "", "Number")
	updateCmd.Flags().StringVarP(&updateRequest.cvv, "cvv", "c", "", "CVV")
	updateCmd.Flags().StringVarP(&updateRequest.month, "month", "m", "", "Month")
	updateCmd.Flags().StringVarP(&updateRequest.year, "year", "y", "", "Year")
	updateCmd.Flags().StringVarP(&updateRequest.info, "info", "i", "", "Info")
}
