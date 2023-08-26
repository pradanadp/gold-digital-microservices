package service

import (
	"context"

	"github.com/pradanadp/gold-digital-microservices/input-harga-service/model/web"
)

type PriceService interface {
	Create(ctx context.Context, request web.PriceCreateRequest) (string, error)
}
