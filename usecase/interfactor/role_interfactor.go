package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type roleInterfactor struct {
	RoleRepository repository.RoleRepository
}

type RoleInterfactor interface {
	Get(r []*model.Role) ([]*model.Role, error)
	GetOne(r []*model.Role, id string) ([]*model.Role, error)
	Create(r *model.Role) (*model.Role, error)
	Update(r *model.Role, id string) (*model.Role, error)
	Delete(r []*model.Role, id string) ([]*model.Role, error)
}

func NewRoleInterfactor(r repository.RoleRepository) RoleInterfactor {
	return &roleInterfactor{r}
}

func (ri *roleInterfactor) Get(r []*model.Role) ([]*model.Role, error) {
	r, err := ri.RoleRepository.FindAll(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ri *roleInterfactor) GetOne(r []*model.Role, id string) ([]*model.Role, error) {
	r, err := ri.RoleRepository.FindOne(r, id)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ri *roleInterfactor) Create(r *model.Role) (*model.Role, error) {
	r, err := ri.RoleRepository.Create(r)
	if err != nil {
		return nil, err
	}

	return r, err
}

func (ri *roleInterfactor) Update(r *model.Role, id string) (*model.Role, error) {
	r, err := ri.RoleRepository.Update(r, id)
	if err != nil {
		return nil, err
	}

	return r, err
}

func (ri *roleInterfactor) Delete(r []*model.Role, id string) ([]*model.Role, error) {
	r, err := ri.RoleRepository.Delete(r, id)
	if err != nil {
		return nil, err
	}

	return r, err
}
