package repository

import (
	"context"
	"database/sql"

	"github.com/pradanadp/gold-digital-microservices/input-harga-storage/model/domain"
)

type PriceRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, price domain.Price) (domain.Price, error)
}
