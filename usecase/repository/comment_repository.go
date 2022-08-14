package repository

import "GoAds/domain/model"

type CommentRepository interface {
	FindAll(c []*model.GetCommentsResponseData) ([]*model.GetCommentsResponseData, error)
	FindOne(c []*model.GetCommentsResponseData, id string) ([]*model.GetCommentsResponseData, error)
	Create(c *model.GetCommentsCreateUpdateData) (*model.GetCommentsCreateUpdateData, error)
	Update(c *model.GetCommentsCreateUpdateData, id string, userID string) (*model.GetCommentsCreateUpdateData, error)
	Delete(c []*model.Comment, id string, userID string) ([]*model.Comment, error)
}
