package repository

import (
	"context"

	"github.com/masato-MacbookPro/go-chat/internal/domain/model"
)

type UserRepostitory interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}
