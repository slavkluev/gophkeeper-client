package text

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/grpc/app/text"
)

type AddRequest struct {
	text string
	info string
}

var addRequest AddRequest

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add text",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("texts_target")
		timeout := viper.GetInt("timeout")

		textApp, err := text.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		_, err = textApp.TextClient.Save(
			context.Background(),
			addRequest.text,
			addRequest.info,
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("add was successful")
	},
}

func init() {
	textCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addRequest.text, "text", "t", "", "Text")
	addCmd.Flags().StringVarP(&addRequest.info, "info", "i", "", "Info")
}
