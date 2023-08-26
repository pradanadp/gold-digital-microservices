package service

import (
	"context"

	"github.com/pradanadp/gold-digital-microservices/buyback-service/model/web"
)

type BuybackService interface {
	Create(ctx context.Context, request web.BuybackRequest) ([]byte, error)
}
