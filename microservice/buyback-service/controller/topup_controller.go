package controller

import (
	"net/http"
)

type BuybackController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
