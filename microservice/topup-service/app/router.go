package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradanadp/gold-digital-microservices/topup-service/controller"
)

func InitRouter(topupController controller.TopupController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/topup", topupController.Create).Methods(http.MethodPost)

	return router
}
