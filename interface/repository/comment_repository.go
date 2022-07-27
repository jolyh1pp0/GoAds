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

func (cr *commentRepository) FindAll(c []*model.Comment) ([]*model.Comment, error) {
	err := cr.db.Select("id, user_id, advertisement_id, content, created_at").Find(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *commentRepository) FindOne(c []*model.Comment, id string) ([]*model.Comment, error) {
	err := cr.db.Select("id, user_id, advertisement_id, content, created_at").Where("id = ?", id).Find(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *commentRepository) Create(c *model.Comment) (*model.Comment, error) {
	err := cr.db.Create(c).Error

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
	err := cr.db.Where("id = ?", id).Delete(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}
