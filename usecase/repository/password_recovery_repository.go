package repository

import (
	"GoAds/domain/model"
	"time"
)

type PasswordRecoveryRepository interface {
	GetEmailValidation(email string) (string, bool, error)
	GetTokenExists(email string) (time.Time, bool, error)
	CreateRecovery(passwdRecovery *model.PasswordRecovery) error
	DeleteRecovery(userID string) error
	FindToken(token string) (bool, string, time.Time, error)
	UpdatePassword(userID string, password string) error
}
