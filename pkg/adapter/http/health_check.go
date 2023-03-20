package http

import (
	"net/http"

	"github.com/masato-MacbookPro/go-chat/pkg/adapter/response"
	"github.com/masato-MacbookPro/go-chat/pkg/config"
	"github.com/masato-MacbookPro/go-chat/pkg/infra"
)

type healthCheckRespons struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	mysqlInfo := config.LoadConfig().MySQLInfo
	_, err := infra.NewMySQLConnector(*mysqlInfo)
	if err != nil {
		rs := healthCheckRespons{
			Message: "failed mysql connect",
			Status:  http.StatusServiceUnavailable,
		}
		response.Write(w, rs, rs.Status)
		return
	}
	rs := healthCheckRespons{
		Message: "server started!!!",
		Status:  http.StatusOK,
	}
	response.Write(w, rs, rs.Status)
}
