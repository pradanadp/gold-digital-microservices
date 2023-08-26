package controller

import (
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/buyback-service/broker"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/helper"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/service"
)

type BuybackControllerImpl struct {
	BuybackService service.BuybackService
}

func NewBuybackController(buybackService service.BuybackService) BuybackController {
	return &BuybackControllerImpl{
		BuybackService: buybackService,
	}
}

func (controller *BuybackControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	buybackRequest := web.BuybackRequest{}
	err := helper.ReadFromRequestBody(r, &buybackRequest)
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

	msg, err := controller.BuybackService.Create(r.Context(), buybackRequest)
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

	if err := broker.PublishMessage([]byte(msg)); err != nil {
		log.Println("error: ", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, web.WebResponse{
			Error:   true,
			ReffID:  helper.GenerateId(),
			Message: "kafka is not ready",
		})
		return
	}

	webResponse := web.WebResponse{
		Error:  false,
		ReffID: helper.GenerateId(),
	}

	helper.WriteToResponseBody(w, webResponse)
}
