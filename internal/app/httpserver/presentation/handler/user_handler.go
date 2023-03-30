package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/presentation/response"
	"github.com/masato-MacbookPro/go-chat/internal/usecase"
)

// TODO: errorハンドリングを共通化する
type errorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	const keyUserID = "user_id"

	ctx := r.Context()

	pathVars := mux.Vars(r)
	userIDString, exist := pathVars[keyUserID]
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		rs := errorResponse{
			Message: "not found params user id",
			Status:  http.StatusBadRequest,
		}
		response.Json(w, rs, rs.Status)
		return
	}
	if !exist {
		rs := errorResponse{
			Message: "not found params user id",
			Status:  http.StatusBadRequest,
		}
		response.Json(w, rs, rs.Status)
		return
	}

	user, err := uh.usecase.GetUserByID(ctx, userID)
	if err != nil {
		rs := errorResponse{
			Message: "not found error",
			Status:  http.StatusNotFound,
		}
		response.Json(w, rs, rs.Status)
		return
	}
	response.Json(w, user, http.StatusOK)
}
