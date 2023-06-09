package repositoryimpl

import (
	"context"
	"database/sql"

	"github.com/masato-MacbookPro/go-chat/internal/domain/entity"
	"github.com/masato-MacbookPro/go-chat/internal/domain/model"
	"github.com/masato-MacbookPro/go-chat/internal/domain/repository"
)

type ChatRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) repository.ChatRepostitory {
	return &ChatRepositoryImpl{
		DB: db,
	}
}

func (ur *ChatRepositoryImpl) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := entity.FindUser(ctx, ur.DB, int64(id))
	if err != nil {
		return nil, err
	}
	return model.NewUserFromRepository(model.UserID(user.ID), user.Name, user.Email), nil
}
