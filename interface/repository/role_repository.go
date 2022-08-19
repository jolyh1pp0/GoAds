package repository

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db}
}

func (rr *roleRepository) FindAll(r []*model.Role) ([]*model.Role, error) {
	err := rr.db.Model(&r).Select("*").Find(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (rr *roleRepository) FindOne(r []*model.Role, id string) ([]*model.Role, error) {
	err := rr.db.Model(&r).Select("*").Where("id = ?", id).Find(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (rr *roleRepository) Create(r *model.Role) (*model.Role, error) {
	err := rr.db.Model(&r).Create(r).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (rr *roleRepository) Update(r *model.Role, id string) (*model.Role, error) {
	result := rr.db.Model(&r).Where("id = ?", id).Update(r)
	if result.Error != nil {
		return nil, domain.ErrCommentInternalServerError
	} else if result.RowsAffected == 0 {
		return nil, domain.ErrForbidden
	}

	return r, nil
}

func (rr *roleRepository) Delete(r []*model.Role, id string) ([]*model.Role, error) {
	result := rr.db.Model(&r).Where("id = ?", id).Delete(&r)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, domain.ErrForbidden
	}

	return r, nil
}
