package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"log"
	"net/http"
)

type userToRoleController struct {
	userToRoleInterfactor interfactor.UserToRoleInterfactor
}

type UserToRoleController interface {
	GetUserToRoles(c Context) error
	GetUserRoles(c Context) error
	CreateUserToRole(c Context) error
	UpdateUserToRole(c Context) error
	DeleteUserToRole(c Context) error
}

func NewUserToRoleController(ri interfactor.UserToRoleInterfactor) UserToRoleController {
	return &userToRoleController{ri}
}

func (uc *userToRoleController) GetUserToRoles(c Context) error {
	var ur []*model.UserRoleResponseData

	ur, err := uc.userToRoleInterfactor.Get(ur)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, ur, "  ")
}

func (uc *userToRoleController) GetUserRoles(c Context) error {
	var ur []*model.UserRoleResponseData
	userID := c.QueryParam("UserID")
	ur, err := uc.userToRoleInterfactor.GetUserRoles(ur, userID)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, ur, "  ")
}

func (uc *userToRoleController) CreateUserToRole(c Context) error {
	var userRole model.UserRole

	err := c.Bind(&userRole)
	if err != nil {
		log.Print(err)
	}

	ur, err := uc.userToRoleInterfactor.Create(&userRole)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, "Status 201. User " + ur.UserID + " assigned to new role", "  ")
}

func (uc *userToRoleController) UpdateUserToRole(c Context) error {
	var userRole model.UserRole

	err := c.Bind(&userRole)
	if err != nil {
		log.Print(err)
	}

	id := c.Param("id")

	err = uc.userToRoleInterfactor.Update(&userRole, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, "Status 201. User " + id + " reassigned to new role", "  ")
}

func (uc *userToRoleController) DeleteUserToRole(c Context) error {
	var ur []*model.UserRole

	id := c.Param("id")

	ur, err := uc.userToRoleInterfactor.Delete(ur, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, "Status 200. User to role " + id + " deleted", "  ")
}