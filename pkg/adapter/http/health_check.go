package http

import (
	"net/http"

	"github.com/masato-MacbookPro/go-chat/pkg/response"
)

type healthCheckRespons struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	rs := healthCheckRespons{
		Message: "server started!!!",
		Status:  http.StatusOK,
	}
	response.Write(w, rs, http.StatusOK)
}
