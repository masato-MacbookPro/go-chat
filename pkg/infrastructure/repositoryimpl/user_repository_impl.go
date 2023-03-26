package repositoryimpl

import (
	"context"
	"database/sql"

	"github.com/masato-MacbookPro/go-chat/pkg/domain/entity"
	"github.com/masato-MacbookPro/go-chat/pkg/domain/model"
	"github.com/masato-MacbookPro/go-chat/pkg/domain/repository"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) repository.UserRepostitory {
	return &userRepositoryImpl{
		DB: db,
	}
}

func (ur *userRepositoryImpl) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := entity.FindUser(ctx, ur.DB, int64(id))
	if err != nil {
		return nil, err
	}
	return model.NewUserFromRepository(model.UserID(user.ID), user.Name, user.Email), nil
}
