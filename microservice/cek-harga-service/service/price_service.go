package service

import (
	"context"

	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/model/web"
)

type PriceService interface {
	Get(ctx context.Context) (web.PriceCheckResponse, error)
}
