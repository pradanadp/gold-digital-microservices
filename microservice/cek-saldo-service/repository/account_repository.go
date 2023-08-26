package repository

import (
	"context"
	"database/sql"

	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/model/domain"
)

type AccountRepository interface {
	Get(ctx context.Context, tx *sql.Tx, accountNumber string) (domain.Account, error)
}
