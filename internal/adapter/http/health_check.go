package http

import (
	"net/http"

	"github.com/masato-MacbookPro/go-chat/internal/adapter/response"
	"github.com/masato-MacbookPro/go-chat/internal/config"
	"github.com/masato-MacbookPro/go-chat/internal/infrastructure"
)

type healthCheckRespons struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	mysqlInfo := config.LoadConfig().MySQLInfo
	_, err := infrastructure.NewMySQLConnector(*mysqlInfo)
	if err != nil {
		rs := healthCheckRespons{
			Message: "failed mysql connect",
			Status:  http.StatusServiceUnavailable,
		}
		response.Json(w, rs, rs.Status)
		return
	}
	rs := healthCheckRespons{
		Message: "server started!!!",
		Status:  http.StatusOK,
	}
	response.Json(w, rs, rs.Status)
}
