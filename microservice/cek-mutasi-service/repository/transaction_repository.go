package repository

import (
	"context"
	"database/sql"

	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/domain"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/web"
)

type TransactionRepository interface {
	Get(ctx context.Context, tx *sql.Tx, request web.TransactionGetRequest) ([]domain.Transaction, error)
}
