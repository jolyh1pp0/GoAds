package repository

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) FindAll(c []*model.GetCommentsResponseData) ([]*model.GetCommentsResponseData, error) {
	err := cr.db.Model(&c).Select("*").Preload("User").Find(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *commentRepository) FindOne(c []*model.GetCommentsResponseData, id string) ([]*model.GetCommentsResponseData, error) {
	err := cr.db.Model(&c).Select("*").Preload("User").Where("comments.id = ?", id).Find(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *commentRepository) Create(c *model.GetCommentsCreateUpdateData) (*model.GetCommentsCreateUpdateData, error) {
	err := cr.db.Model(&c).Create(c).Error

	if err != nil {
		return nil, domain.ErrCommentInternalServerError
	}

	return c, nil
}

func (cr *commentRepository) Update(c *model.GetCommentsCreateUpdateData, id string, userID string) error {
	result := cr.db.Model(&c).Preload("User").Where("id = ? and user_id = ?", id, userID).Update(c)
	if result.Error != nil {
		return domain.ErrCommentInternalServerError
	} else if result.RowsAffected == 0 {
		return domain.ErrForbidden
	}

	return nil
}

func (cr *commentRepository) Delete(c []*model.Comment, id string, userID string) ([]*model.Comment, error) {
	result := cr.db.Model(&c).Where("id = ? and user_id = ?", id, userID).Delete(&c)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, domain.ErrForbidden
	}

	return c, nil
}
