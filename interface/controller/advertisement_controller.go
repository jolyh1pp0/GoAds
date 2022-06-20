package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"net/http"
)

type advertisementController struct {
	advertisementInterfactor interfactor.AdvertisementInterfactor
}

type AdvertisementController interface {
	GetAdvertisements(c Context) error
	GetOneAdvertisement(c Context) error
	CreateAdvertisement(c Context) error
	DeleteAdvertisement(c Context) error
}

func NewAdvertisementController(us interfactor.AdvertisementInterfactor) AdvertisementController {
	return &advertisementController{us}
}

func (ac *advertisementController) GetAdvertisements(c Context) error {
	var a []*model.Advertisement

	a, err := ac.advertisementInterfactor.Get(a)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, a, "  ")
}

func (ac *advertisementController) GetOneAdvertisement(c Context) error {
	var a []*model.Advertisement
	id := c.Param("id")

	a, err := ac.advertisementInterfactor.GetOne(a, id)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, a, "  ")
}

func (ac *advertisementController) CreateAdvertisement(c Context) error {
	var params model.Advertisement
	// TODO: PARAMS HERE
	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	a, err := ac.advertisementInterfactor.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, a, "  ")
}

func (ac *advertisementController) DeleteAdvertisement(c Context) error {
	var a []*model.Advertisement

	id := c.Param("id")

	a, err := ac.advertisementInterfactor.Delete(a, id)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, a, "  ")
}
