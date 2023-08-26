package repository

import (
	"context"
	"database/sql"

	"github.com/pradanadp/gold-digital-microservices/topup-storage/model/web"
)

type AccountRepository interface {
	Update(ctx context.Context, tx *sql.Tx, request web.TopupRequest) error
}
