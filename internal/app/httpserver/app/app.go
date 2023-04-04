package app

import (
	"log"
	"net/http"

	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/config"
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/presentation/handler"
)

type App struct {
	cfg     *config.AppConfig
	handler *handler.ChatHandler
}

func NewApp(
	cfg *config.AppConfig,
	handler *handler.ChatHandler,
) *App {
	return &App{
		cfg:     cfg,
		handler: handler,
	}
}

func (a *App) Start() {
	r := handler.NewRouter(a.handler)
	log.Print("server started!!!")

	err := http.ListenAndServe(a.cfg.HTTPInfo.Addr, r)
	if err != nil {
		panic(err)
	}
}
