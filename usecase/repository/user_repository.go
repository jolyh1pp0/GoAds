package repository

import "GoAds/domain/model"

type UserRepository interface {
	FindAll(u []*model.GetUsersResponseData) ([]*model.GetUsersResponseData, error)
	FindOne(u []*model.GetUsersResponseData, id string) ([]*model.GetUsersResponseData, error)
	GetUser(userID string) (string, error)
	Update(u *model.User, id string) (*model.User, error)
	Delete(u []*model.User, id string) ([]*model.User, error)
}
