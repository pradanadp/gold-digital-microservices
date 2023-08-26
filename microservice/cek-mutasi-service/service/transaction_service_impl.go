package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/domain"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/repository"
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

func (service *TransactionServiceImpl) Get(ctx context.Context, request web.TransactionGetRequest) ([]domain.Transaction, error) {
	tx, errBegin := service.DB.Begin()
	if errBegin != nil {
		err := fmt.Errorf("failed to start transaction: %w", errBegin)
		log.Println(err.Error())
		return nil, err
	}

	transactions, err := service.TransactionRepository.Get(ctx, tx, request)
	if err != nil {
		tx.Rollback()
		err := fmt.Errorf("transaction service impl error: error get transaction")
		log.Println(err.Error())
		return nil, err
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		tx.Rollback()
		err := fmt.Errorf("failed to commit transaction: %w", errCommit)
		log.Panicln(err.Error())
		return nil, err
	}

	return transactions, nil
}
