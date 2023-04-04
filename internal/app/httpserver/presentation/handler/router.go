package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(h *ChatHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/health_check", HealthCheck).Methods("GET")
	r.HandleFunc("/users/{user_id}", h.GetUserByID).Methods("GET")
	return r
}
