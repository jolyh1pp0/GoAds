package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"log"
	"net/http"
)

type advertisementController struct {
	advertisementInterfactor interfactor.AdvertisementInterfactor
}

type AdvertisementController interface {
	GetAdvertisements(c Context) error
	GetOneAdvertisement(c Context) error
	UpdateAdvertisement(c Context) error
	CreateAdvertisement(c Context) error
	DeleteAdvertisement(c Context) error
}

func NewAdvertisementController(ad interfactor.AdvertisementInterfactor) AdvertisementController {
	return &advertisementController{ad}
}

func (ac *advertisementController) GetAdvertisements(c Context) error {
	var a []*model.Advertisement
	var orderQuery string
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}
	priceSort := c.QueryParam("priceSort")
	if priceSort == "cheap" {
		orderQuery = "price ASC"
	} else if priceSort == "expensive" {
		orderQuery = "price DESC"
	}

	dateSort := c.QueryParam("dateSort")
	if dateSort == "oldest" {
		orderQuery = "advertisements.created_at ASC"
	} else if dateSort == "newest" {
		orderQuery = "advertisements.created_at DESC"
	}

	a, err := ac.advertisementInterfactor.Get(a, limit, offset, orderQuery)
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
	var advertisement model.Advertisement

	err := c.Bind(&advertisement)
	if err != nil {
		log.Print(err)
	}

	err = ac.advertisementInterfactor.Create(&advertisement)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, "Advertisement created", "  ")
}

func (ac *advertisementController) UpdateAdvertisement(c Context) error {
	var advertisement model.Advertisement

	err := c.Bind(&advertisement)
	if err != nil {
		log.Print(err)
	}

	id := c.Param("id")

	a, err := ac.advertisementInterfactor.Update(&advertisement, id)
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

	return c.JSONPretty(http.StatusOK, "Advertisement №"+id+" deleted", "  ")
}
