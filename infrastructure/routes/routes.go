package routes

import (
	"GoAds/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.Advertisement.GetAdvertisements(context) })
	//e.POST("/users", func(context echo.Context) error { return c.Advertisement.CreateUser(context) })

	return e
}
