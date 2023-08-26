package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/helper"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/model/web"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/service"
)

type PriceControllerImpl struct {
	PriceService service.PriceService
}

func NewPriceController(categoryService service.PriceService) PriceController {
	return &PriceControllerImpl{
		PriceService: categoryService,
	}
}

func (controller *PriceControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	priceResponse, err := controller.PriceService.Get(context.Background())
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
		Data:  priceResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
