package handler

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masato-MacbookPro/go-chat/internal/infrastructure/repositoryimpl"
	"github.com/masato-MacbookPro/go-chat/internal/usecase"
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
