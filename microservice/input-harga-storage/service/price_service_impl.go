package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/input-harga-storage/model/domain"
	"github.com/pradanadp/gold-digital-microservices/input-harga-storage/model/web"
	"github.com/pradanadp/gold-digital-microservices/input-harga-storage/repository"
)

type PriceServiceImpl struct {
	PriceRepository repository.PriceRepository
	DB              *sql.DB
}

func NewPriceService(priceRepository repository.PriceRepository, DB *sql.DB) PriceService {
	return &PriceServiceImpl{
		PriceRepository: priceRepository,
		DB:              DB,
	}
}

func (service *PriceServiceImpl) Create(ctx context.Context, msg []byte) error {
	tx, errBegin := service.DB.Begin()
	if errBegin != nil {
		err := fmt.Errorf("failed to start transaction: %w", errBegin)
		log.Println(err.Error())
		return err
	}

	var priceRequest web.PriceCreateRequest

	errUnm := json.Unmarshal(msg, &priceRequest)
	if errUnm != nil {
		err := fmt.Errorf("error unmarshaling request: %w", errUnm)
		log.Println(err.Error())
		return err
	}

	var price = domain.Price{
		AdminId:      priceRequest.AdminId,
		TopupPrice:   priceRequest.TopupPrice,
		BuybackPrice: priceRequest.BuybackPrice,
	}

	_, errInsert := service.PriceRepository.Insert(ctx, tx, price)
	if errInsert != nil {
		tx.Rollback()
		err := fmt.Errorf("failed to insert transaction: %w", errInsert)
		log.Println(err.Error())
		return err
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		tx.Rollback()
		err := fmt.Errorf("failed to commit transaction: %w", errCommit)
		log.Println(err.Error())
		return err
	}

	return nil
}
