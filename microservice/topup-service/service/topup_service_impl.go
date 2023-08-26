package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/pradanadp/gold-digital-microservices/topup-service/client"
	"github.com/pradanadp/gold-digital-microservices/topup-service/model/web"
)

type TopupServiceImpl struct {
	Validate *validator.Validate
}

func NewTopupService(validate *validator.Validate) TopupService {
	return &TopupServiceImpl{
		Validate: validate,
	}
}

func (service *TopupServiceImpl) Create(ctx context.Context, request web.TopupRequest) ([]byte, error) {
	if errVal := service.Validate.Struct(request); errVal != nil {
		err := fmt.Errorf("error validate struct, %w", errVal)
		log.Println(err.Error())
		return nil, err
	}

	resp := client.CheckPrice()
	topupPrice := resp.Data.(map[string]any)["topup_price"].(float64)

	if request.Quantity < 0.001 {
		err := errors.New("minimum gold quantity is 0.001 gram")
		log.Println(err.Error())
		return nil, err
	}

	if request.Quantity*1000 != float64(int(request.Quantity*1000)) {
		err := errors.New("gold quantity must be a multiple of 0.001 gram")
		log.Println(err.Error())
		return nil, err
	}

	if request.Price != topupPrice {
		err := errors.New("the price entered is different from the current topup price")
		log.Println(err.Error())
		return nil, err
	}

	data, err := json.Marshal(request)
	if err != nil {
		err := fmt.Errorf("error marshaling request, %w", err)
		log.Println(err.Error())
		return nil, err
	}

	return data, nil
}
