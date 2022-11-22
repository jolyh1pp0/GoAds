package interfactor

import (
	"GoAds/usecase/repository"
)

type chatInterfactor struct {
	ChatRepository repository.ChatRepository
}

type ChatInterfactor interface {
	GetUser(userID string) (string, error)
}

func NewChatInterfactor(r repository.ChatRepository) ChatInterfactor {
	return &chatInterfactor{r}
}

func (us *chatInterfactor) GetUser(userID string) (string, error) {
	username, err := us.ChatRepository.GetUser(userID)
	if err != nil {
		return "", err
	}

	return username, nil
}
