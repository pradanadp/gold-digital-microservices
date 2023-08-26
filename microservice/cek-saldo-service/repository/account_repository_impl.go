package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/model/domain"
)

type AccountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

func (repository *AccountRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, accountNumber string) (domain.Account, error) {
	selectQuery := `
		SELECT * FROM accounts WHERE account_number = $1;
	`

	var account domain.Account
	errQuery := tx.QueryRowContext(ctx, selectQuery, accountNumber).Scan(
		&account.AccountNumber, &account.Balance, &account.CreatedAt, &account.UpdatedAt,
	)
	if errQuery != nil {
		err := fmt.Errorf("failed to retrieve account: %w", errQuery)
		log.Println(err.Error())
		return domain.Account{}, err
	}

	return account, nil
}
