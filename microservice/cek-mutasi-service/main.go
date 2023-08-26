package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/app"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/controller"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/repository"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/service"
)

func main() {
	config := app.InitConfig("local.env")
	db, err := app.InitDB(config)
	if err != nil {
		panic(err.Error())
	}

	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, db)
	transactionController := controller.NewTransactionController(transactionService)

	router := app.InitRouter(transactionController)

	server := http.Server{
		Addr:    "localhost:3303",
		Handler: router,
	}

	fmt.Println("server started at localhost:3303")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
