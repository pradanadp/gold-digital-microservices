package service

import (
	"context"

	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/domain"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/web"
)

type TransactionService interface {
	Get(ctx context.Context, request web.TransactionGetRequest) ([]domain.Transaction, error)
}
