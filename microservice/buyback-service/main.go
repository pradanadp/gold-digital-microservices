package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/app"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/controller"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/service"
)

func main() {
	validate := validator.New()
	buybackService := service.NewBuybackService(validate)
	buybackController := controller.NewBuybackController(buybackService)

	router := app.InitRouter(buybackController)

	server := http.Server{
		Addr:    "localhost:3330",
		Handler: router,
	}

	fmt.Println("server started at localhost:3330")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
