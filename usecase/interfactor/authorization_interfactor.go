package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"time"
)

type authorizationInterfactor struct {
	AuthorizationRepository repository.AuthorizationRepository
}

type AuthorizationInterfactor interface {
	Create(u *model.User) (*model.User, error)
	CreateSession(s *model.Session) (*model.Session, error)
	UserExists(email string) (string, string, error)
	GetUserRoles(userID string) ([]int, error)
	GetRefreshTokenUUIDFromTable(token string) (string, error)
	Logout(sessionUUID string) error
	GetSessionExpiration(sessionUUID string) (time.Time, error)
	GetSessionUUID(userID string) (string, error)
	UpdateSession(sessionUUID string, s *model.Session) error
}

func NewAuthorizationInterfactor(r repository.AuthorizationRepository) AuthorizationInterfactor {
	return &authorizationInterfactor{r}
}

func (ai *authorizationInterfactor) Create(u *model.User) (*model.User, error) {
	u, err := ai.AuthorizationRepository.Create(u)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (ai *authorizationInterfactor) CreateSession(s *model.Session) (*model.Session, error) {
	s, err := ai.AuthorizationRepository.CreateSession(s)
	if err != nil {
		return nil, err
	}

	return s, err
}

func (ai *authorizationInterfactor) UserExists(email string) (string, string, error) {
	password, userID, err := ai.AuthorizationRepository.UserExists(email)
	if err != nil {
		return "", "", err
	}
	return password, userID, nil
}

func (ai *authorizationInterfactor) GetUserRoles(userID string) ([]int, error) {
	userRoles, err := ai.AuthorizationRepository.GetUserRoles(userID)
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (ai *authorizationInterfactor) GetRefreshTokenUUIDFromTable(uuid string) (string, error) {
	refreshTokenUUID, err := ai.AuthorizationRepository.GetRefreshTokenUUIDFromTable(uuid)
	if err != nil {
		return "", err
	}
	return refreshTokenUUID, nil
}

func (ai *authorizationInterfactor) GetSessionExpiration(sessionUUID string) (time.Time, error) {
	uuid, err := ai.AuthorizationRepository.GetSessionExpiration(sessionUUID)
	if err != nil {
		return time.Time{}, err
	}

	return uuid, nil
}

func (ai *authorizationInterfactor) GetSessionUUID(userID string) (string, error) {
	uuid, err := ai.AuthorizationRepository.GetSessionUUID(userID)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (ai *authorizationInterfactor) UpdateSession(sessionUUID string, s *model.Session) error {
	err := ai.AuthorizationRepository.UpdateSession(sessionUUID, s)
	if err != nil {
		return err
	}

	return nil
}

func (ai *authorizationInterfactor) Logout(sessionUUID string) error {
	err := ai.AuthorizationRepository.Logout(sessionUUID)
	if err != nil {
		return err
	}

	return nil
}
