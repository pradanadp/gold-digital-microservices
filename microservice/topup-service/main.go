package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pradanadp/gold-digital-microservices/topup-service/app"
	"github.com/pradanadp/gold-digital-microservices/topup-service/controller"
	"github.com/pradanadp/gold-digital-microservices/topup-service/service"
)

func main() {
	validate := validator.New()
	topupService := service.NewTopupService(validate)
	topupController := controller.NewTopupController(topupService)

	router := app.InitRouter(topupController)

	server := http.Server{
		Addr:    "localhost:3030",
		Handler: router,
	}

	fmt.Println("server started at localhost:3030")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
