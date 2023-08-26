package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/topup-storage/helper"
	"github.com/pradanadp/gold-digital-microservices/topup-storage/model/domain"
)

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) (domain.Transaction, error) {
	selectQuery := `
		SELECT balance FROM accounts WHERE account_number = $1;
	`

	var currentBalance float64
	errQuery := tx.QueryRowContext(ctx, selectQuery, transaction.AccountNumber).Scan(&currentBalance)
	if errQuery != nil {
		err := fmt.Errorf("failed to retrieve current balance: %w", errQuery)
		log.Println(err.Error())
		return domain.Transaction{}, err
	}

	newBalance := currentBalance + transaction.Quantity

	fmt.Println("currentBalance:", currentBalance)
	fmt.Println("topup price:", transaction.Quantity)
	fmt.Println("newBalance:", newBalance)

	updateQuery := `
		UPDATE accounts SET balance = $1 WHERE account_number = $2;
	`

	_, errExec := tx.ExecContext(ctx, updateQuery, newBalance, transaction.AccountNumber)
	if errExec != nil {
		err := fmt.Errorf("failed to update balance: %w", errExec)
		log.Println(err.Error())
		return domain.Transaction{}, err
	}

	insertQuery := `
		INSERT INTO transactions(
			transaction_id,
			account_number,
			quantity, topup_price,
			buyback_price, type,
			balance
		) VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	transactionId := helper.GenerateId()

	_, errExec = tx.ExecContext(
		ctx, insertQuery,
		transactionId,
		transaction.AccountNumber,
		transaction.Quantity,
		transaction.TopupPrice,
		transaction.BuybackPrice,
		transaction.Type,
		newBalance,
	)
	if errExec != nil {
		err := fmt.Errorf("failed to insert transaction data: %w", errExec)
		log.Println(err.Error())
		return domain.Transaction{}, err
	}

	transaction.TransactionId = transactionId
	return transaction, nil
}
