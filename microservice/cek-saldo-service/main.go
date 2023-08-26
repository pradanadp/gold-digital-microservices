package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/app"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/controller"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/repository"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/service"
)

func main() {
	config := app.InitConfig("local.env")
	db, err := app.InitDB(config)
	if err != nil {
		panic(err.Error())
	}

	accountRepository := repository.NewAccountRepository()
	accountService := service.NewAccountService(accountRepository, db)
	accountController := controller.NewAccountController(accountService)

	router := app.InitRouter(accountController)

	server := http.Server{
		Addr:    "localhost:3003",
		Handler: router,
	}

	fmt.Println("server started at localhost:3003")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
