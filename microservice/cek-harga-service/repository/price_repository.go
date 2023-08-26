package repository

import (
	"context"
	"database/sql"

	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/model/domain"
)

type PriceRepository interface {
	Get(ctx context.Context, tx *sql.Tx) (domain.Price, error)
}
