package repository

import "GoAds/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	FindOne(u []*model.User, id string) ([]*model.User, error)
	Update(u *model.User, id string) (*model.User, error)
	Delete(u []*model.User, id string) ([]*model.User, error)
}
