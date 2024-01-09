package auth

import (
	"context"

	ssov1 "github.com/slavkluev/gophkeeper-contracts/gen/go/sso"
	"google.golang.org/grpc"
)

type AuthClient struct {
	api   ssov1.AuthClient
	appId int64
}

func NewAuthClient(cc *grpc.ClientConn, appId int64) *AuthClient {
	return &AuthClient{
		api:   ssov1.NewAuthClient(cc),
		appId: appId,
	}
}

func (client *AuthClient) Login(ctx context.Context, email string, password string) (string, error) {
	req := &ssov1.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    client.appId,
	}

	res, err := client.api.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetToken(), nil
}
