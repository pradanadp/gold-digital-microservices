package controller

import (
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/input-harga-service/broker"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/helper"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/service"
)

type PriceControllerImpl struct {
	PriceService service.PriceService
}

func NewPriceController(priceService service.PriceService) PriceController {
	return &PriceControllerImpl{
		PriceService: priceService,
	}
}

func (controller *PriceControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	priceCreateRequest := web.PriceCreateRequest{}
	err := helper.ReadFromRequestBody(r, &priceCreateRequest)
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

	msg, err := controller.PriceService.Create(r.Context(), priceCreateRequest)
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
