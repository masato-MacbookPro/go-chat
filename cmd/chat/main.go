package main

import (
	"net/http"

	adapterHTTP "github.com/masato-MacbookPro/go-chat/pkg/adapter/http"
	"github.com/masato-MacbookPro/go-chat/pkg/config"
	"github.com/masato-MacbookPro/go-chat/pkg/infrastructure"
)

func main() {
	conf := config.LoadConfig()
	mysqlInfo := config.LoadConfig().MySQLInfo
	db, err := infrastructure.NewMySQLConnector(*mysqlInfo)
	if err != nil {
		panic(err)
	}
	r := adapterHTTP.NewRouter(db.Conn)
	err = http.ListenAndServe(conf.HTTPInfo.Addr, r)
	if err != nil {
		panic(err)
	}
}
