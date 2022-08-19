package repository

import "GoAds/domain/model"

type RoleRepository interface {
	FindAll(c []*model.Role) ([]*model.Role, error)
	FindOne(c []*model.Role, id string) ([]*model.Role, error)
	Create(c *model.Role) (*model.Role, error)
	Update(c *model.Role, id string) (*model.Role, error)
	Delete(c []*model.Role, id string) ([]*model.Role, error)
}
