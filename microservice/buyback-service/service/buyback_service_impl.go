package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/client"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/model/web"
)

type BuybackServiceImpl struct {
	Validate *validator.Validate
}

func NewBuybackService(validate *validator.Validate) BuybackService {
	return &BuybackServiceImpl{
		Validate: validate,
	}
}

func (service *BuybackServiceImpl) Create(ctx context.Context, request web.BuybackRequest) ([]byte, error) {
	if errVal := service.Validate.Struct(request); errVal != nil {
		err := fmt.Errorf("error validate struct, %w", errVal)
		log.Println(err.Error())
		return nil, err
	}

	price := client.CheckPrice()
	buybackPrice := price.Data.(map[string]any)["buyback_price"].(float64)

	acc := client.CheckBalance(web.AccountRequest{
		AccountNumber: request.AccountNumber,
	})
	balance := acc.Data.(map[string]any)["saldo"].(float64)

	if request.Quantity > balance {
		err := errors.New("insufficient balance")
		log.Println(err.Error())
		return nil, err
	}

	if request.Price != buybackPrice {
		err := errors.New("the price entered is different from the current buyback price")
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
