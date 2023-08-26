package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradanadp/gold-digital-microservices/cek-saldo-service/controller"
)

func InitRouter(accountController controller.AccountController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/cek-saldo", accountController.Get).Methods(http.MethodGet)

	return router
}
