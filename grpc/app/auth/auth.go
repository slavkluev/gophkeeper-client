package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"client/grpc/client/auth"
)

type AuthApp struct {
	AuthClient *auth.AuthClient
}

func New(
	target string,
	appId int64,
) (*AuthApp, error) {
	cc, err := grpc.DialContext(
		context.Background(),
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	accClient := auth.NewAuthClient(cc, appId)

	return &AuthApp{AuthClient: accClient}, nil
}
