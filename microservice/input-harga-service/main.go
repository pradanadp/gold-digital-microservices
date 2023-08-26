package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/app"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/controller"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/service"
)

func main() {
	validate := validator.New()
	priceService := service.NewPriceService(validate)
	priceController := controller.NewPriceController(priceService)

	router := app.InitRouter(priceController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("server started at localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
