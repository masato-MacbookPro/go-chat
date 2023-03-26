package http

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masato-MacbookPro/go-chat/pkg/domain/usecase"
	"github.com/masato-MacbookPro/go-chat/pkg/infrastructure/repositoryimpl"
)

func NewRouter(db *sql.DB) http.Handler {
	userRepositoryImpl := repositoryimpl.NewUserRepositoryImpl(db)
	userUsecase := usecase.NewUserUsecase(userRepositoryImpl)
	userHandler := NewUserHandler(userUsecase)

	r := mux.NewRouter()
	r.HandleFunc("/health_check", healthCheck).Methods("GET")
	r.HandleFunc("/users/{user_id}", userHandler.GetUserByID).Methods("GET")
	return r
}
