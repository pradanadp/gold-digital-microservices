package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/domain"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/web"
)

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, request web.TransactionGetRequest) ([]domain.Transaction, error) {
	startDate, errParse := time.Parse("2006-01-02", request.StartDate)
	if errParse != nil {
		err := fmt.Errorf("failed to parse start date, %w", errParse)
		log.Println(err.Error())
		return nil, err
	}

	endDate, errParse := time.Parse("2006-01-02", request.EndDate)
	if errParse != nil {
		err := fmt.Errorf("failed to parse end date, %w", errParse)
		log.Println(err.Error())
		return nil, err
	}

	selectQuery := `
		SELECT * FROM transactions WHERE account_number = $1 AND created_at BETWEEN $2 AND $3;
	`

	rows, errQuery := tx.QueryContext(ctx, selectQuery, request.AccountNumber, startDate, endDate.Add(24*time.Hour)) // Add 24 hours to cover the whole day
	if errQuery != nil {
		err := fmt.Errorf("failed to query data, %w", errQuery)
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		if errScan := rows.Scan(
			&transaction.TransactionId,
			&transaction.AccountNumber,
			&transaction.Quantity,
			&transaction.TopupPrice,
			&transaction.BuybackPrice,
			&transaction.Type,
			&transaction.Balance,
			&transaction.CreatedAt,
		); errScan != nil {
			err := fmt.Errorf("failed to scan rows, %w", errScan)
			log.Println(err.Error())
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
