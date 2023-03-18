package http

import "net/http"

func NewRouter() http.Handler {
	r := http.NewServeMux()
	r.Handle("/health_check", http.HandlerFunc(healthCheck))
	return r
}
