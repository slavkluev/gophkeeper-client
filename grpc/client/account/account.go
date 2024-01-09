package account

import (
	"context"
	"time"

	accountsv1 "github.com/slavkluev/gophkeeper-contracts/gen/go/accounts"
	"google.golang.org/grpc"

	"client/domain/models"
)

type AccountClient struct {
	api     accountsv1.AccountsClient
	timeout time.Duration
}

func NewAccountClient(cc *grpc.ClientConn, timeout time.Duration) *AccountClient {
	return &AccountClient{
		api:     accountsv1.NewAccountsClient(cc),
		timeout: timeout,
	}
}

func (client *AccountClient) Save(
	ctx context.Context,
	login string,
	pass string,
	info string,
) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &accountsv1.SaveRequest{
		Login:    login,
		Password: pass,
		Info:     info,
	}

	res, err := client.api.Save(ctx, req)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (client *AccountClient) Update(
	ctx context.Context,
	id uint64,
	login string,
	pass string,
	info string,
) error {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &accountsv1.UpdateRequest{
		Id:       id,
		Login:    login,
		Password: pass,
		Info:     info,
	}

	_, err := client.api.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (client *AccountClient) List(ctx context.Context) ([]models.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, client.timeout)
	defer cancel()

	req := &accountsv1.GetAllRequest{}

	list, err := client.api.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var accounts []models.Account
	for _, account := range list.Accounts {
		accounts = append(accounts, models.Account{
			ID:    account.GetId(),
			Login: account.GetLogin(),
			Pass:  account.GetPassword(),
			Info:  account.GetInfo(),
		})
	}

	return accounts, nil
}
