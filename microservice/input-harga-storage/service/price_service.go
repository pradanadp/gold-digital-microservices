package service

import (
	"context"
)

type PriceService interface {
	Create(ctx context.Context, msg []byte) error
}
