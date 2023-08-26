package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/topup-storage/model/web"
)

type AccountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

// tidak perlu implementasi AccountRepository untuk update balance karena sudah otomatis diupdate di TransactionRepository

func (repository *AccountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request web.TopupRequest) error {
	selectQuery := `
		SELECT balance FROM accounts WHERE account_number = $1;
	`

	var currentBalance float64
	errQuery := tx.QueryRowContext(ctx, selectQuery, request.AccountNumber).Scan(&currentBalance)
	if errQuery != nil {
		err := fmt.Errorf("failed to retrieve current balance: %w", errQuery)
		log.Println(err.Error())
		return err
	}

	newBalance := currentBalance + request.Quantity

	fmt.Println("currentBalance:", currentBalance)
	fmt.Println("topup price:", request.Quantity)
	fmt.Println("newBalance:", newBalance)

	updateQuery := `
		UPDATE accounts SET balance = $1 WHERE account_number = $2;
	`

	_, errExec := tx.ExecContext(ctx, updateQuery, newBalance, request.AccountNumber)
	if errExec != nil {
		err := fmt.Errorf("failed to update balance: %w", errExec)
		log.Println(err.Error())
		return err
	}

	return nil
}
