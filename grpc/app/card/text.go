package card

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"client/grpc/client/card"
	"client/grpc/interceptor"
)

type CardApp struct {
	CardClient *card.CardClient
}

func New(
	accessToken string,
	target string,
	timeout time.Duration,
) (*CardApp, error) {
	authInterceptor := interceptor.NewAuthInterceptor(accessToken)

	cc, err := grpc.DialContext(
		context.Background(),
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		return nil, err
	}

	cardClient := card.NewCardClient(cc, timeout)

	return &CardApp{CardClient: cardClient}, nil
}
