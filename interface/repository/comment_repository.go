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

func (cr *commentRepository) Create(c *model.Comment) (*model.Comment, error) {
	err := cr.db.Model(&c).Create(c).Error

	if err != nil {
		return nil, domain.ErrCommentInternalServerError
	}

	return c, nil
}

func (cr *commentRepository) Update(c *model.Comment, id string) (*model.Comment, error) {
	err := cr.db.Model(&c).Where("id = ?", id).Update(c).Error
	if err != nil {
		return nil, domain.ErrCommentInternalServerError
	}

	return c, nil
}

func (cr *commentRepository) Delete(c []*model.Comment, id string) ([]*model.Comment, error) {
	err := cr.db.Model(&c).Where("id = ?", id).Delete(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}
