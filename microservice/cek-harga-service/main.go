package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/app"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/controller"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/repository"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/service"
)

func main() {
	config := app.InitConfig("local.env")
	db, err := app.InitDB(config)
	if err != nil {
		panic(err.Error())
	}

	priceRepository := repository.NewPriceRepository()
	priceService := service.NewPriceService(priceRepository, db)
	priceController := controller.NewPriceController(priceService)

	router := app.InitRouter(priceController)

	server := http.Server{
		Addr:    "localhost:3300",
		Handler: router,
	}

	fmt.Println("server started at localhost:3300")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
