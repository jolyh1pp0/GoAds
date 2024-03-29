package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"log"
	"net/http"
)

type userController struct {
	userInterfactor interfactor.UserInterfactor
}

type UserController interface {
	GetUsers(c Context) error
	GetOneUser(c Context) error
	UpdateUser(c Context) error
	DeleteUser(c Context) error
}

func NewUserController(us interfactor.UserInterfactor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.GetUsersResponseData

	u, err := uc.userInterfactor.Get(u)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, u, "  ")
}

func (uc *userController) GetOneUser(c Context) error {
	var u []*model.GetUsersResponseData
	id := c.Param("id")

	u, err := uc.userInterfactor.GetOne(u, id)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, u, "  ")
}

func (uc *userController) UpdateUser(c Context) error {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		log.Print(err)
	}

	id := c.Param("id")

	u, err := uc.userInterfactor.Update(&user, id)
	if !errors.Is(err, nil) {
		return err
	}

	u.UpdatedAt = nil
	return c.JSONPretty(http.StatusCreated, u, "  ")
}

func (uc *userController) DeleteUser(c Context) error {
	var u []*model.User

	id := c.Param("id")

	u, err := uc.userInterfactor.Delete(u, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, "User "+id+" deleted", "  ")
}
