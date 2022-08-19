package repository

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
)

type userToRoleRepository struct {
	db *gorm.DB
}

func NewUserToRoleRepository(db *gorm.DB) repository.UserToRoleRepository {
	return &userToRoleRepository{db}
}

func (urr *userToRoleRepository) FindAll(ur []*model.UserRoleResponseData) ([]*model.UserRoleResponseData, error) {
	err := urr.db.Model(&ur).Select("*").Preload("User").Preload("Role").Find(&ur).Error
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (urr *userToRoleRepository) FindOne(ur []*model.UserRoleResponseData, id string) ([]*model.UserRoleResponseData, error) {
	err := urr.db.Model(&ur).Select("*").Where("id = ?", id).Preload("User").Preload("Role").Find(&ur).Error
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (urr *userToRoleRepository) Create(ur *model.UserRole) (*model.UserRole, error) {
	err := urr.db.Model(&ur).Create(ur).Error

	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (urr *userToRoleRepository) Update(ur *model.UserRole, id string) error {
	result := urr.db.Model(&ur).Where("id = ?", id).Update(ur)
	if result.Error != nil {
		return domain.ErrCommentInternalServerError
	} else if result.RowsAffected == 0 {
		return domain.ErrForbidden
	}

	return nil
}

func (urr *userToRoleRepository) Delete(ur []*model.UserRole, id string) ([]*model.UserRole, error) {
	result := urr.db.Model(&ur).Where("id = ?", id).Delete(&ur)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, domain.ErrForbidden
	}

	return ur, nil
}
