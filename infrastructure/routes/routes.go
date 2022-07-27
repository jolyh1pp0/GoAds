package routes

import (
	"GoAds/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/advertisements", func(context echo.Context) error { return c.Advertisement.GetAdvertisements(context) })
	e.GET("/advertisements/:id", func(context echo.Context) error { return c.Advertisement.GetOneAdvertisement(context) })
	e.PUT("/advertisements/:id", func(context echo.Context) error { return c.Advertisement.UpdateAdvertisement(context) })
	e.POST("/advertisements", func(context echo.Context) error { return c.Advertisement.CreateAdvertisement(context) })
	e.DELETE("advertisements/:id", func(context echo.Context) error { return c.Advertisement.DeleteAdvertisement(context) })

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	e.GET("/users/:id", func(context echo.Context) error { return c.User.GetOneUser(context) })
	e.PUT("/users/:id", func(context echo.Context) error { return c.User.UpdateUser(context) })
	e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })
	e.DELETE("users/:id", func(context echo.Context) error { return c.User.DeleteUser(context) })

	e.GET("/comments", func(context echo.Context) error { return c.Comment.GetComments(context) })
	e.GET("/comments/:id", func(context echo.Context) error { return c.Comment.GetOneComment(context) })
	e.PUT("/comments/:id", func(context echo.Context) error { return c.Comment.UpdateComment(context) })
	e.POST("/comments", func(context echo.Context) error { return c.Comment.CreateComment(context) })
	e.DELETE("comments/:id", func(context echo.Context) error { return c.Comment.DeleteComment(context) })

	return e
}
