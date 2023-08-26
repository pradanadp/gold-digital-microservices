package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/model/domain"
)

type PriceRepositoryImpl struct{}

func NewPriceRepository() PriceRepository {
	return &PriceRepositoryImpl{}
}

func (repository *PriceRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) (domain.Price, error) {
	selectQuery := `
		SELECT * FROM prices ORDER BY created_at DESC LIMIT 1
	`

	rows, errQuery := tx.QueryContext(ctx, selectQuery)
	if errQuery != nil {
		err := fmt.Errorf("failed to query data, %w", errQuery)
		log.Println(err)
		return domain.Price{}, err
	}
	defer rows.Close()

	price := domain.Price{}
	if rows.Next() {
		errScan := rows.Scan(
			&price.PriceId, &price.AdminId, &price.TopupPrice, &price.BuybackPrice, &price.CreatedAt,
		)
		if errScan != nil {
			err := fmt.Errorf("failed to scan rows, %w", errQuery)
			log.Println(err)
			return domain.Price{}, err
		}
		return price, nil
	} else {
		return price, errors.New("data is not found")
	}
}
