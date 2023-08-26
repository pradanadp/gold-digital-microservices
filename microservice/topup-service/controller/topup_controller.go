package controller

import (
	"net/http"
)

type TopupController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
