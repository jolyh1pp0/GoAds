package repository

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) repository.ChatRepository {
	return &chatRepository{db}
}

func (c chatRepository) GetUser(userID string) (string, error) {
	var u model.GetUsersResponseData

	err := c.db.Model(&u).Select("*").Where("id = ?", userID).Find(&u).Error

	if err != nil {
		return "", err
	}

	return u.FirstName + " " + u.LastName, nil
}
