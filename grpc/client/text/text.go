package text

import (
	"context"
	"time"

	textsv1 "github.com/slavkluev/gophkeeper-contracts/gen/go/texts"
	"google.golang.org/grpc"

	"client/domain/models"
)

type TextClient struct {
	api     textsv1.TextsClient
	timeout time.Duration
}

func NewTextClient(cc *grpc.ClientConn, timeout time.Duration) *TextClient {
	return &TextClient{
		api:     textsv1.NewTextsClient(cc),
		timeout: timeout,
	}
}

func (client *TextClient) Save(
	ctx context.Context,
	text string,
	info string,
) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &textsv1.SaveRequest{
		Text: text,
		Info: info,
	}

	res, err := client.api.Save(ctx, req)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (client *TextClient) Update(
	ctx context.Context,
	id uint64,
	text string,
	info string,
) error {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &textsv1.UpdateRequest{
		Id:   id,
		Text: text,
		Info: info,
	}

	_, err := client.api.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (client *TextClient) List(ctx context.Context) ([]models.Text, error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &textsv1.GetAllRequest{}

	list, err := client.api.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var texts []models.Text
	for _, text := range list.Texts {
		texts = append(texts, models.Text{
			ID:   text.GetId(),
			Text: text.GetText(),
			Info: text.GetInfo(),
		})
	}

	return texts, nil
}
