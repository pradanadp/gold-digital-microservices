package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/repository"
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

func (service *PriceServiceImpl) Get(ctx context.Context) (web.PriceCheckResponse, error) {
	tx, errBegin := service.DB.Begin()
	if errBegin != nil {
		err := fmt.Errorf("failed to start transaction: %w", errBegin)
		log.Println(err.Error())
		return web.PriceCheckResponse{}, err
	}

	price, errService := service.PriceRepository.Get(ctx, tx)
	if errService != nil {
		err := fmt.Errorf("error implementing Get method repository: %w", errService)
		log.Println(err.Error())
		return web.PriceCheckResponse{}, err
	}

	if errService != nil {
		tx.Rollback()
		err := fmt.Errorf("failed to insert transaction: %w", errService)
		log.Println(err.Error())
		return web.PriceCheckResponse{}, err
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		tx.Rollback()
		err := fmt.Errorf("failed to commit transaction: %w", errCommit)
		log.Panicln(err.Error())
		return web.PriceCheckResponse{}, err
	}

	return web.PriceCheckResponse{
		TopupPrice:   price.TopupPrice,
		BuybackPrice: price.BuybackPrice,
	}, nil
}
