package card

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"client/domain/models"
	"client/grpc/app/card"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of cards",
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := viper.GetString("access_token")
		target := viper.GetString("cards_target")
		timeout := viper.GetInt("timeout")

		cardApp, err := card.New(accessToken, target, time.Duration(timeout)*time.Second)
		if err != nil {
			panic(err)
		}

		cards, err := cardApp.CardClient.List(context.Background())
		if err != nil {
			panic(err)
		}

		err = printTable(cards)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	cardCmd.AddCommand(listCmd)
}

func printTable(cards []models.Card) error {
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

	_, err := fmt.Fprintln(w, "ID\tNUMBER\tCVV\tMONTH\tYEAR\tINFO")
	if err != nil {
		return err
	}

	for _, crd := range cards {
		row := fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s", crd.ID, crd.Number, crd.CVV, crd.Month, crd.Year, crd.Info)
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
