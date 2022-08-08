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
	UserExists(email string) (string, string, error)
	Login(u []*model.User) ([]*model.User, error)
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

func (ai *authorizationInterfactor) UserExists(email string) (string, string, error) {
	password, userID, err := ai.AuthorizationRepository.UserExists(email)
	if err != nil {
		return "", "", err
	}
	return password, userID, nil
}

func (ai *authorizationInterfactor) Login(u []*model.User) ([]*model.User, error) {
	return nil, nil
}
