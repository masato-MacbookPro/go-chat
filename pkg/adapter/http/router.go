package http

import "net/http"

func NewRouter() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/health_check", healthCheck)
	return r
}
