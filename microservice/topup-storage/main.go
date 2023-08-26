package main

import (
	_ "github.com/lib/pq"
	"github.com/pradanadp/gold-digital-microservices/topup-storage/app"
	"github.com/pradanadp/gold-digital-microservices/topup-storage/broker"
)

func main() {
	config := app.InitConfig("local.env")
	db, err := app.InitDB(config)
	if err != nil {
		panic(err.Error())
	}

	broker.SubscribeMessage(db)
}
