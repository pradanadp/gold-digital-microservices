package service

import "context"

type AccountService interface {
	Update(ctx context.Context, msg []byte) error
}
