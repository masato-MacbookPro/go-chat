package repository

import (
	"context"

	"github.com/masato-MacbookPro/go-chat/pkg/domain/model"
)

type UserRepostitory interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}
