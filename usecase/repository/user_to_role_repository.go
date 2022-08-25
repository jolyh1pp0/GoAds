package repository

import "GoAds/domain/model"

type UserToRoleRepository interface {
	FindAll(c []*model.UserRoleResponseData) ([]*model.UserRoleResponseData, error)
	FindUserRoles(c []*model.UserRoleResponseData, userID string) ([]*model.UserRoleResponseData, error)
	Create(c *model.UserRole) (*model.UserRole, error)
	Update(c *model.UserRole, id string) error
	Delete(c []*model.UserRole, id string) ([]*model.UserRole, error)
}
