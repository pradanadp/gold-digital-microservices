package controller

import "net/http"

type TransactionController interface {
	Get(w http.ResponseWriter, r *http.Request)
}
