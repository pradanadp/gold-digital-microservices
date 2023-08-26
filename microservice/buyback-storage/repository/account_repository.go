package repository

import (
	"context"
	"database/sql"

	"github.com/pradanadp/gold-digital-microservices/buyback-storage/model/web"
)

type AccountRepository interface {
	Update(ctx context.Context, tx *sql.Tx, request web.TopupRequest) error
}
