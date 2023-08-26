package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/helper"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/service"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &AccountControllerImpl{
		AccountService: accountService,
	}
}

func (controller *AccountControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	var accountReq web.AccountRequest
	err := helper.ReadFromRequestBody(r, &accountReq)
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

	resp, err := controller.AccountService.Get(context.Background(), accountReq.AccountNumber)
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

	webResponse := web.WebResponse{
		Error: false,
		Data:  resp,
	}

	helper.WriteToResponseBody(w, webResponse)
}
