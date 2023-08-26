package controller

import (
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/topup-service/broker"
	"github.com/pradanadp/gold-digital-microservices/topup-service/helper"
	"github.com/pradanadp/gold-digital-microservices/topup-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/topup-service/service"
)

type TopupControllerImpl struct {
	TopupService service.TopupService
}

func NewTopupController(topupService service.TopupService) TopupController {
	return &TopupControllerImpl{
		TopupService: topupService,
	}
}

func (controller *TopupControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	topupRequest := web.TopupRequest{}
	err := helper.ReadFromRequestBody(r, &topupRequest)
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

	msg, err := controller.TopupService.Create(r.Context(), topupRequest)
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
