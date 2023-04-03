package main

import (
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver"
)

func main() {
	server, err := httpserver.InitializeApp()
	if err != nil {
		panic(err)
	}
	server.Start()
}
