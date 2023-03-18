package main

import (
	"net/http"

	adapterHTTP "github.com/masato-MacbookPro/go-chat/pkg/adapter/http"
	"github.com/masato-MacbookPro/go-chat/pkg/config"
)

func main() {
	conf := config.LoadConfig()
	r := adapterHTTP.NewRouter()
	err := http.ListenAndServe(conf.HTTPInfo.Addr, r)
	if err != nil {
		panic(err)
	}
}
