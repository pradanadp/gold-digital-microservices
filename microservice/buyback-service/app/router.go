package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/controller"
)

func InitRouter(buybackController controller.BuybackController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/buyback", buybackController.Create).Methods(http.MethodPost)

	return router
}
