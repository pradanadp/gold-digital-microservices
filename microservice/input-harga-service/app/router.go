package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradanadp/gold-digital-microservices/input-harga-service/controller"
)

func InitRouter(priceController controller.PriceController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/input-harga", priceController.Create).Methods(http.MethodPost)

	return router
}
