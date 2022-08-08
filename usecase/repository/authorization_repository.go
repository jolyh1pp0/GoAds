package repository

import "GoAds/domain/model"

type AuthorizationRepository interface {
	Create(u *model.User) (*model.User, error)
	UserExists(email string) (string, string, error)
	Login(u []*model.User) ([]*model.User, error)
}
