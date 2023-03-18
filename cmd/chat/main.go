package main

import (
	"net/http"

	adapterHTTP "github.com/masato-MacbookPro/go-chat/pkg/adapter/http"
)

func main() {
	r := adapterHTTP.NewRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
