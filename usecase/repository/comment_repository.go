package repository

import "GoAds/domain/model"

type CommentRepository interface {
	FindAll(c []*model.GetCommentsResponseData) ([]*model.GetCommentsResponseData, error)
	FindOne(c []*model.GetCommentsResponseData, id string) ([]*model.GetCommentsResponseData, error)
	Create(c *model.Comment) (*model.Comment, error)
	Update(c *model.Comment, id string) (*model.Comment, error)
	Delete(c []*model.Comment, id string) ([]*model.Comment, error)
}
