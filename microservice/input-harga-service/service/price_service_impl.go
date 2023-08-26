package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/model/web"
)

type PriceServiceImpl struct {
	Validate *validator.Validate
}

func NewPriceService(validate *validator.Validate) PriceService {
	return &PriceServiceImpl{
		Validate: validate,
	}
}

func (service *PriceServiceImpl) Create(ctx context.Context, request web.PriceCreateRequest) (string, error) {
	if errVal := service.Validate.Struct(request); errVal != nil {
		err := errors.New("error validate struct")
		log.Println(err.Error())
		return "", err
	}

	data, err := json.Marshal(request)
	if err != nil {
		err := fmt.Errorf("error marshaling request, %w", err)
		log.Println(err.Error())
		return "", err
	}

	return string(data), nil
}
