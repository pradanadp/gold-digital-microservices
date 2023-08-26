package service

import "context"

type TransactionService interface {
	Create(ctx context.Context, msg []byte) error
}
