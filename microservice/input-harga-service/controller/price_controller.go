package controller

import (
	"net/http"
)

type PriceController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
