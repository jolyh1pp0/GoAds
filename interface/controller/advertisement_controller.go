package controller

import (
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
	var orderQuery string
	offset := c.QueryParam("offset")
	priceSort := c.QueryParam("priceSort")
	if priceSort == "cheap" {
		orderQuery = "price ASC"
	} else if priceSort == "expensive" {
		orderQuery = "price DESC"
	}

	dateSort := c.QueryParam("dateSort")
	if dateSort == "oldest" {
		orderQuery = "created_at ASC"
	} else if dateSort == "newest" {
		orderQuery = "created_at DESC"
	}

	a, err := ac.advertisementInterfactor.Get(a, "10", offset, orderQuery)
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
	title := c.FormValue("Title")
	description := c.FormValue("Description")
	price := c.FormValue("Price")
	price32, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Println(err)
	}
	photo_1 := c.FormValue("Photo_1")
	photo_2 := c.FormValue("Photo_2")
	photo_3 := c.FormValue("Photo_3")

	advertisement := model.Advertisement{
		Title:       title,
		Description: description,
		Price:       float32(price32),
		Photo_1:     photo_1,
		Photo_2:     photo_2,
		Photo_3:     photo_3,
	}

	a, err := ac.advertisementInterfactor.Create(&advertisement)
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
