package text

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"client/grpc/client/text"
	"client/grpc/interceptor"
)

type TextApp struct {
	TextClient *text.TextClient
}

func New(
	accessToken string,
	target string,
	timeout time.Duration,
) (*TextApp, error) {
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

	textClient := text.NewTextClient(cc, timeout)

	return &TextApp{TextClient: textClient}, nil
}
