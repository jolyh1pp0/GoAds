package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type userToRoleInterfactor struct {
	UserToRoleRepository repository.UserToRoleRepository
}

type UserToRoleInterfactor interface {
	Get(r []*model.UserRoleResponseData) ([]*model.UserRoleResponseData, error)
	GetUserRoles(r []*model.UserRoleResponseData, userID string) ([]*model.UserRoleResponseData, error)
	Create(r *model.UserRole) (*model.UserRole, error)
	Update(r *model.UserRole, id string) error
	Delete(r []*model.UserRole, id string) ([]*model.UserRole, error)
}

func NewUserToRoleInterfactor(r repository.UserToRoleRepository) UserToRoleInterfactor {
	return &userToRoleInterfactor{r}
}

func (uri *userToRoleInterfactor) Get(ur []*model.UserRoleResponseData) ([]*model.UserRoleResponseData, error) {
	ur, err := uri.UserToRoleRepository.FindAll(ur)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (uri *userToRoleInterfactor) GetUserRoles(ur []*model.UserRoleResponseData, userID string) ([]*model.UserRoleResponseData, error) {
	ur, err := uri.UserToRoleRepository.FindUserRoles(ur, userID)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (uri *userToRoleInterfactor) Create(ur *model.UserRole) (*model.UserRole, error) {
	ur, err := uri.UserToRoleRepository.Create(ur)
	if err != nil {
		return nil, err
	}

	return ur, err
}

func (uri *userToRoleInterfactor) Update(ur *model.UserRole, id string) error {
	err := uri.UserToRoleRepository.Update(ur, id)
	if err != nil {
		return err
	}

	return err
}

func (uri *userToRoleInterfactor) Delete(ur []*model.UserRole, id string) ([]*model.UserRole, error) {
	ur, err := uri.UserToRoleRepository.Delete(ur, id)
	if err != nil {
		return nil, err
	}

	return ur, err
}
