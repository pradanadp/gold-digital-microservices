package service

import (
	"context"

	"github.com/pradanadp/gold-digital-microservices/topup-service/model/web"
)

type TopupService interface {
	Create(ctx context.Context, request web.TopupRequest) ([]byte, error)
}
