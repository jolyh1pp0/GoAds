package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type userInterfactor struct {
	UserRepository repository.UserRepository
}

type UserInterfactor interface {
	Get(u []*model.GetUsersResponseData) ([]*model.GetUsersResponseData, error)
	GetOne(u []*model.GetUsersResponseData, id string) ([]*model.GetUsersResponseData, error)
	GetUser(userID string) (string, error)
	Update(u *model.User, id string) (*model.User, error)
	Delete(u []*model.User, id string) ([]*model.User, error)
}

func NewUserInterfactor(r repository.UserRepository) UserInterfactor {
	return &userInterfactor{r}
}

func (us *userInterfactor) Get(u []*model.GetUsersResponseData) ([]*model.GetUsersResponseData, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *userInterfactor) GetOne(u []*model.GetUsersResponseData, id string) ([]*model.GetUsersResponseData, error) {
	u, err := us.UserRepository.FindOne(u, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *userInterfactor) GetUser(userID string) (string, error) {
	username, err := us.UserRepository.GetUser(userID)
	if err != nil {
		return "", err
	}

	return username, nil
}

func (us *userInterfactor) Update(u *model.User, id string) (*model.User, error) {
	u, err := us.UserRepository.Update(u, id)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (us *userInterfactor) Delete(u []*model.User, id string) ([]*model.User, error) {
	u, err := us.UserRepository.Delete(u, id)
	if err != nil {
		return nil, err
	}

	return u, err
}
