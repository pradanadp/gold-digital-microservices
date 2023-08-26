package controller

import "net/http"

type AccountController interface {
	Get(w http.ResponseWriter, r *http.Request)
}
