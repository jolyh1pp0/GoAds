package repository

import (
	"GoAds/domain/model"
	"time"
)

type AuthorizationRepository interface {
	Create(u *model.UserRegister) (*model.UserRegister, error)
	CreateSession(s *model.Session) (*model.Session, error)
	UserExists(email string) (string, string, error)
	GetUserRoles(userID string) ([]int, error)
	GetRefreshTokenUUIDFromTable(uuid string) (string, error)
	GetSessionUUID(userID string) (string, error)
	GetSessionExpiration(sessionUUID string) (time.Time, error)
	UpdateSession(sessionUUID string, s *model.Session) error
	Logout(sessionUUID string) error
	CreateUserToRole(userRole model.UserRole) error
}
