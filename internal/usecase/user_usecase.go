package usecase

import (
	"context"

	"github.com/masato-MacbookPro/go-chat/internal/domain/model"
	"github.com/masato-MacbookPro/go-chat/internal/domain/repository"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepostitory
}

func NewUserUsecase(uu repository.UserRepostitory) UserUsecase {
	return &userUsecase{
		userRepository: uu,
	}
}

// GetUserByID implements UserUsecase
func (uu *userUsecase) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := uu.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
