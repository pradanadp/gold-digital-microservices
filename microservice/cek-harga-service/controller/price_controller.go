package controller

import "net/http"

type PriceController interface {
	Get(w http.ResponseWriter, r *http.Request)
}
