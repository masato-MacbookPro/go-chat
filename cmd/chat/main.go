package main

import (
	"net/http"

	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/config"
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/presentation/handler"
	"github.com/masato-MacbookPro/go-chat/internal/infrastructure"
)

func main() {
	conf := config.LoadConfig()
	mysqlInfo := config.LoadConfig().MySQLInfo
	db, err := infrastructure.NewMySQLConnector(*mysqlInfo)
	if err != nil {
		panic(err)
	}
	r := handler.NewRouter(db.Conn)
	err = http.ListenAndServe(conf.HTTPInfo.Addr, r)
	if err != nil {
		panic(err)
	}
}
