package repository

import "GoAds/domain/model"

type CommentRepository interface {
	FindAll(c []*model.Comment) ([]*model.Comment, error)
	FindOne(c []*model.Comment, id string) ([]*model.Comment, error)
	Create(c *model.Comment) (*model.Comment, error)
	Update(c *model.Comment, id string) (*model.Comment, error)
	Delete(c []*model.Comment, id string) ([]*model.Comment, error)
}
