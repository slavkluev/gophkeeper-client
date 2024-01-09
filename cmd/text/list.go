package text

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/domain/models"
	"client/grpc/app/text"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of texts",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("texts_target")
		timeout := viper.GetInt("timeout")

		textApp, err := text.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		texts, err := textApp.TextClient.List(context.Background())
		if err != nil {
			panic(err)
		}

		err = printTable(texts)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	textCmd.AddCommand(listCmd)
}

func printTable(texts []models.Text) error {
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

	_, err := fmt.Fprintln(w, "ID\tTEXT\tINFO")
	if err != nil {
		return err
	}

	for _, txt := range texts {
		row := fmt.Sprintf("%d\t%s\t%s", txt.ID, txt.Text, txt.Info)
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
