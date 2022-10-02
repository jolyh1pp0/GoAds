package repository

import "GoAds/domain/model"

type AuthorizationRepository interface {
	Create(u *model.User) (*model.User, error)
	CreateSession(s *model.Session) (*model.Session, error)
	UserExists(email string) (string, string, error)
	GetUserRoles(userID string) ([]int, error)
	GetRefreshTokenUUIDFromTable(uuid string) (string, error)
	Login(u []*model.User) ([]*model.User, error)
	GetSession(userID string) (int, error)
	GetSessionUUID(userID string) (string, error)
	UpdateSession(sessionUUID string, s *model.Session) error
}
