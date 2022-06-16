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
	CreateAdvertisement(c Context) error
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

	return c.JSON(http.StatusOK, a)
}

func (ac *advertisementController) CreateAdvertisement(c Context) error {
	var params model.Advertisement

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	a, err := ac.advertisementInterfactor.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusCreated, a)
}
