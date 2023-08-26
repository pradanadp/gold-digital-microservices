package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/input-harga-storage/helper"
	"github.com/pradanadp/gold-digital-microservices/input-harga-storage/model/domain"
)

type PriceRepositoryImpl struct{}

func NewPriceRepository() PriceRepository {
	return &PriceRepositoryImpl{}
}

func (repository *PriceRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, price domain.Price) (domain.Price, error) {
	insertQuery := `
		INSERT INTO prices(price_id, admin_id, topup_price, buyback_price) values ($1, $2, $3, $4)
	`

	priceId := helper.GenerateId()
	_, errExec := tx.ExecContext(ctx, insertQuery, priceId, price.AdminId, price.TopupPrice, price.BuybackPrice)
	if errExec != nil {
		err := fmt.Errorf("failed to insert price: %w", errExec)
		log.Println(err.Error())
		return domain.Price{}, err
	}

	price.PriceId = priceId
	return price, nil
}
