package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
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
	Login(u []*model.User) ([]*model.User, error)
	GetSession(userID string) (string, error)
	UpdateSession(userID string, s *model.Session) error
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

func (ai *authorizationInterfactor) GetRefreshTokenUUIDFromTable(token string) (string, error) {
	refreshTokenUUID, err := ai.AuthorizationRepository.GetRefreshTokenUUIDFromTable(token)
	if err != nil {
		return "", err
	}
	return refreshTokenUUID, nil
}

func (ai *authorizationInterfactor) Login(u []*model.User) ([]*model.User, error) {
	return nil, nil
}

func (ai *authorizationInterfactor) GetSession(userID string) (string, error) {
	session, err := ai.AuthorizationRepository.GetSession(userID)
	if err != nil {
		return "", err
	}

	return session, nil
}

func (ai *authorizationInterfactor) UpdateSession(userID string, s *model.Session) error {
	err := ai.AuthorizationRepository.UpdateSession(userID, s)
	if err != nil {
		return err
	}

	return nil
}
