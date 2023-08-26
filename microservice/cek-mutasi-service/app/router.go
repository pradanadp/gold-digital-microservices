package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradanadp/gold-digital-microservices/cek-mutasi-service/controller"
)

func InitRouter(transactionController controller.TransactionController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/mutasi", transactionController.Get).Methods(http.MethodGet)

	return router
}
