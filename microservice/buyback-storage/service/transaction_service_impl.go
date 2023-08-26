package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/buyback-storage/client"
	"github.com/pradanadp/gold-digital-microservices/buyback-storage/model/domain"
	"github.com/pradanadp/gold-digital-microservices/buyback-storage/model/web"
	"github.com/pradanadp/gold-digital-microservices/buyback-storage/repository"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
}

func NewTransactionService(transactionRepository repository.TransactionRepository, DB *sql.DB) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		DB:                    DB,
	}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, msg []byte) error {
	tx, errBegin := service.DB.Begin()
	if errBegin != nil {
		err := fmt.Errorf("failed to start transaction: %w", errBegin)
		log.Println(err.Error())
		return err
	}

	var topupRequest web.TopupRequest
	errUnmarshal := json.Unmarshal(msg, &topupRequest)
	if errUnmarshal != nil {
		err := fmt.Errorf("error unmarshaling request: %w", errUnmarshal)
		log.Println(err.Error())
		return err
	}

	resp := client.CheckPrice()
	topupPrice := resp.Data.(map[string]any)["topup_price"].(float64)
	buybackPrice := resp.Data.(map[string]any)["buyback_price"].(float64)

	var transaction = domain.Transaction{
		AccountNumber: topupRequest.AccountNumber,
		Quantity:      topupRequest.Quantity,
		TopupPrice:    topupPrice,
		BuybackPrice:  buybackPrice,
		Type:          "buyback",
	}

	_, errInsert := service.TransactionRepository.Insert(ctx, tx, transaction)
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
		log.Panicln(err.Error())
		return err
	}

	return nil
}
