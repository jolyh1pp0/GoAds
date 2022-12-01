package repository

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.GetUsersResponseData) ([]*model.GetUsersResponseData, error) {
	err := ur.db.Model(&u).Select("*").Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) FindOne(u []*model.GetUsersResponseData, id string) ([]*model.GetUsersResponseData, error) {
	err := ur.db.Model(&u).Select("*").Where("id = ?", id).Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) GetUser(userID string) (string, error) {
	var u model.GetUsersResponseData

	err := ur.db.Model(&u).Select("*").Where("id = ?", userID).Find(&u).Error

	if err != nil {
		return "", err
	}

	return u.FirstName + " " + u.LastName, nil
}

func (ur *userRepository) Update(u *model.User, id string) (*model.User, error) {
	err := ur.db.Model(&u).Where("id = ?", id).Update(u).Error

	if err != nil {
		if err.Error() == domain.ErrUserAlreadyWithEmail {
			return nil, domain.ErrUserEmailAlreadyExists
		} else if err.Error() == domain.ErrUserAlreadyWithPhone {
			return nil, domain.ErrUserPhoneAlreadyExists
		}
		return nil, domain.ErrUserInternalServerError
	}

	return u, nil
}

func (ur *userRepository) Delete(u []*model.User, id string) ([]*model.User, error) {
	err := ur.db.Model(&u).Where("id = ?", id).Delete(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
