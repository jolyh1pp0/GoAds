package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"log"
	"net/http"
)

type roleController struct {
	roleInterfactor interfactor.RoleInterfactor
}

type RoleController interface {
	GetRoles(c Context) error
	GetOneRole(c Context) error
	CreateRole(c Context) error
	UpdateRole(c Context) error
	DeleteRole(c Context) error
}

func NewRoleController(ri interfactor.RoleInterfactor) RoleController {
	return &roleController{ri}
}

func (rc *roleController) GetRoles(c Context) error {
	var r []*model.Role

	r, err := rc.roleInterfactor.Get(r)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, r, "  ")
}

func (rc *roleController) GetOneRole(c Context) error {
	var r []*model.Role
	id := c.Param("id")

	r, err := rc.roleInterfactor.GetOne(r, id)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, r, "  ")
}

func (rc *roleController) CreateRole(c Context) error {
	var role model.Role

	err := c.Bind(&role)
	if err != nil {
		log.Print(err)
	}

	r, err := rc.roleInterfactor.Create(&role)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, "Status 201. Role " + r.Role + " created", "  ")
}

func (rc *roleController) UpdateRole(c Context) error {
	var role model.Role

	err := c.Bind(&role)
	if err != nil {
		log.Print(err)
	}

	id := c.Param("id")

	r, err := rc.roleInterfactor.Update(&role, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, "Status 201. Role " + r.Role + " updated", "  ")
}

func (rc *roleController) DeleteRole(c Context) error {
	var r []*model.Role

	id := c.Param("id")

	r, err := rc.roleInterfactor.Delete(r, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, "Status 200. Role №" + id + " deleted", "  ")
}