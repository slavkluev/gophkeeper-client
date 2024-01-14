package card

import (
	"context"
	"time"

	cardsv1 "github.com/slavkluev/gophkeeper-contracts/gen/go/cards"
	"google.golang.org/grpc"

	"client/domain/models"
)

type CardClient struct {
	api     cardsv1.CardsClient
	timeout time.Duration
}

func NewCardClient(cc *grpc.ClientConn, timeout time.Duration) *CardClient {
	return &CardClient{
		api:     cardsv1.NewCardsClient(cc),
		timeout: timeout,
	}
}

func (client *CardClient) Save(
	ctx context.Context,
	number string,
	cvv string,
	month string,
	year string,
	info string,
) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &cardsv1.SaveRequest{
		Number: number,
		Cvv:    cvv,
		Month:  month,
		Year:   year,
		Info:   info,
	}

	res, err := client.api.Save(ctx, req)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (client *CardClient) Update(
	ctx context.Context,
	id uint64,
	number string,
	cvv string,
	month string,
	year string,
	info string,
) error {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &cardsv1.UpdateRequest{
		Id:     id,
		Number: number,
		Cvv:    cvv,
		Month:  month,
		Year:   year,
		Info:   info,
	}

	_, err := client.api.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (client *CardClient) List(ctx context.Context) ([]models.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &cardsv1.GetAllRequest{}

	list, err := client.api.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var cards []models.Card
	for _, card := range list.Cards {
		cards = append(cards, models.Card{
			ID:     card.GetId(),
			Number: card.GetNumber(),
			CVV:    card.GetCvv(),
			Month:  card.GetMonth(),
			Year:   card.GetYear(),
			Info:   card.GetInfo(),
		})
	}

	return cards, nil
}
