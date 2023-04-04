package usecase

import (
	"context"

	"github.com/masato-MacbookPro/go-chat/internal/domain/model"
	"github.com/masato-MacbookPro/go-chat/internal/domain/repository"
)

type ChatUsecase interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type chatUsecase struct {
	chatRepository repository.ChatRepostitory
}

func NewChatUsecase(cr repository.ChatRepostitory) ChatUsecase {
	return &chatUsecase{
		chatRepository: cr,
	}
}

// GetUserByID implements ChatUsecase
func (cu *chatUsecase) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := cu.chatRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
