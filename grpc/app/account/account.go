package account

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"client/grpc/client/account"
	"client/grpc/interceptor"
)

type AccountApp struct {
	AccClient *account.AccountClient
}

func New(
	accessToken string,
	target string,
	timeout time.Duration,
) (*AccountApp, error) {
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

	accClient := account.NewAccountClient(cc, timeout)

	return &AccountApp{AccClient: accClient}, nil
}
