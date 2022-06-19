package routes

import (
	"GoAds/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/advertisements", func(context echo.Context) error { return c.Advertisement.GetAdvertisements(context) })
	e.GET("/advertisements/:id", func(context echo.Context) error { return c.Advertisement.GetOneAdvertisement(context) })

	e.POST("/create", func(context echo.Context) error { return c.Advertisement.CreateAdvertisement(context) })
	e.GET("/create", func(context echo.Context) error { return c.Advertisement.CreateAdvertisement(context) })
	return e
}
