package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/repository"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
	DB                *sql.DB
}

func NewAccountService(priceRepository repository.AccountRepository, DB *sql.DB) AccountService {
	return &AccountServiceImpl{
		AccountRepository: priceRepository,
		DB:                DB,
	}
}

func (service *AccountServiceImpl) Get(ctx context.Context, accountNumber string) (web.BalanceCheckResponse, error) {
	tx, errBegin := service.DB.Begin()
	if errBegin != nil {
		err := fmt.Errorf("failed to start transaction: %w", errBegin)
		log.Println(err.Error())
		return web.BalanceCheckResponse{}, err
	}

	account, errService := service.AccountRepository.Get(ctx, tx, accountNumber)
	if errService != nil {
		log.Println(errService.Error())
		return web.BalanceCheckResponse{}, errService
	}

	if errService != nil {
		tx.Rollback()
		err := fmt.Errorf("failed to insert transaction: %w", errService)
		log.Println(err.Error())
		return web.BalanceCheckResponse{}, err
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		tx.Rollback()
		err := fmt.Errorf("failed to commit transaction: %w", errCommit)
		log.Panicln(err.Error())
		return web.BalanceCheckResponse{}, err
	}

	resp := web.BalanceCheckResponse{
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
	}

	return resp, nil
}
