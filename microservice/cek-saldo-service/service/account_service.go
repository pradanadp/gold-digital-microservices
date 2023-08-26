package service

import (
	"context"

	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/model/web"
)

type AccountService interface {
	Get(ctx context.Context, accountNumber string) (web.BalanceCheckResponse, error)
}
