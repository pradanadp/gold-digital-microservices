package repository

import (
	"context"
	"database/sql"

	"github.com/pradanadp/gold-digital-microservices/topup-storage/model/domain"
)

type TransactionRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) (domain.Transaction, error)
}
