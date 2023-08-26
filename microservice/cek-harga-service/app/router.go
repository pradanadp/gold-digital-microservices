package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradanadp/gold-digital-microservices/cek-harga-service/controller"
)

func InitRouter(priceController controller.PriceController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/cek-harga", priceController.Get).Methods(http.MethodGet)

	return router
}
