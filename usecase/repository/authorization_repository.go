package repository

import "GoAds/domain/model"

type AuthorizationRepository interface {
	Create(u *model.User) (*model.User, error)
	CreateSession(s *model.Session) (*model.Session, error)
	UserExists(email string) (string, string, error)
	GetUserRoles(userID string) ([]int, error)
	GetRefreshTokenUUIDFromTable(token string) (string, error)
	Login(u []*model.User) ([]*model.User, error)
}
