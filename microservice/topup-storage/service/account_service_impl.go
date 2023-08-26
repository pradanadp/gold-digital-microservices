package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/topup-storage/model/web"
	"github.com/pradanadp/gold-digital-microservices/topup-storage/repository"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
	DB                *sql.DB
}

func NewAccountService(accountRepository repository.AccountRepository, DB *sql.DB) AccountService {
	return &AccountServiceImpl{
		AccountRepository: accountRepository,
		DB:                DB,
	}
}

func (service *AccountServiceImpl) Update(ctx context.Context, msg []byte) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	var topupRequest web.TopupRequest
	err = json.Unmarshal(msg, &topupRequest)
	if err != nil {
		log.Println("error unmarshaling request")
		return fmt.Errorf("error unmarshaling request: %w", err)
	}

	err = service.AccountRepository.Update(ctx, tx, topupRequest)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
