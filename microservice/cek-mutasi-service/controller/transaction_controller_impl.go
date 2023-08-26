package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/helper"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/service"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	var request web.TransactionGetRequest
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		log.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		helper.WriteToResponseBody(w, web.WebResponse{
			Error:   true,
			ReffID:  helper.GenerateId(),
			Message: "bad request. " + err.Error(),
		})
		return
	}

	transactions, err := controller.TransactionService.Get(context.Background(), request)
	if err != nil {
		log.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, web.WebResponse{
			Error:   true,
			ReffID:  helper.GenerateId(),
			Message: "Error while processing the request: " + err.Error(),
		})
		return
	}

	var resp []web.TransactionGetResponse
	for _, transaction := range transactions {
		resp = append(resp, web.TransactionGetResponse{
			CreatedAt:    transaction.CreatedAt,
			Type:         transaction.Type,
			Quantity:     transaction.Quantity,
			TopupPrice:   transaction.TopupPrice,
			BuybackPrice: transaction.BuybackPrice,
			Balance:      transaction.Balance,
		})
	}

	webResponse := web.WebResponse{
		Error: false,
		Data:  resp,
	}

	helper.WriteToResponseBody(w, webResponse)
}
