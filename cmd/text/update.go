package text

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/grpc/app/text"
)

type UpdateRequest struct {
	id   uint64
	text string
	info string
}

var updateRequest UpdateRequest

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update text",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("texts_target")
		timeout := viper.GetInt("timeout")

		textApp, err := text.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		err = textApp.TextClient.Update(
			context.Background(),
			updateRequest.id,
			updateRequest.text,
			updateRequest.info,
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("update was successful")
	},
}

func init() {
	textCmd.AddCommand(updateCmd)

	updateCmd.Flags().Uint64VarP(&updateRequest.id, "id", "", 0, "Id")
	updateCmd.Flags().StringVarP(&updateRequest.text, "text", "t", "", "Text")
	updateCmd.Flags().StringVarP(&updateRequest.info, "info", "i", "", "Info")
}
