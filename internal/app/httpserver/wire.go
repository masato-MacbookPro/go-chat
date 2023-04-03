//go:build wireinject
// +build wireinject

package httpserver

import (
	"github.com/google/wire"
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/app"
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/config"
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/presentation/handler"
	"github.com/masato-MacbookPro/go-chat/internal/infrastructure"
	"github.com/masato-MacbookPro/go-chat/internal/infrastructure/repositoryimpl"
	"github.com/masato-MacbookPro/go-chat/internal/usecase"
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		app.NewApp,
		config.LoadConfig,
		wire.FieldsOf(new(*config.AppConfig), "MySQLInfo"),
		infrastructure.NewMySQLConnector,
		handler.NewUserHandler,
		usecase.NewUserUsecase,
		repositoryimpl.NewUserRepositoryImpl,
	)
	return nil, nil
}
