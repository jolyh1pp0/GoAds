package routes

import (
	"GoAds/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("example_key"),
		TokenLookup: "header:Authorization",
		Skipper: func(c echo.Context) bool {
			if c.Request().URL.Path == "/login" {
				return true
			} else if c.Request().URL.Path == "/register" {
				return true
			}
			return false
		},
	}))

	e.GET("/advertisements", func(context echo.Context) error { return c.Advertisement.GetAdvertisements(context) })
	e.GET("/advertisements/:id", func(context echo.Context) error { return c.Advertisement.GetOneAdvertisement(context) })
	e.PUT("/advertisements/:id", func(context echo.Context) error { return c.Advertisement.UpdateAdvertisement(context) })
	e.POST("/advertisements", func(context echo.Context) error { return c.Advertisement.CreateAdvertisement(context) })
	e.DELETE("advertisements/:id", func(context echo.Context) error { return c.Advertisement.DeleteAdvertisement(context) })

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	e.GET("/users/:id", func(context echo.Context) error { return c.User.GetOneUser(context) })
	e.PUT("/users/:id", func(context echo.Context) error { return c.User.UpdateUser(context) })
	e.DELETE("users/:id", func(context echo.Context) error { return c.User.DeleteUser(context) })

	e.GET("/comments", func(context echo.Context) error { return c.Comment.GetComments(context) })
	e.GET("/comments/:id", func(context echo.Context) error { return c.Comment.GetOneComment(context) })
	e.PUT("/comments/:id", func(context echo.Context) error { return c.Comment.UpdateComment(context) })
	e.POST("/comments", func(context echo.Context) error { return c.Comment.CreateComment(context) })
	e.DELETE("comments/:id", func(context echo.Context) error { return c.Comment.DeleteComment(context) })

	e.POST("/register", func(context echo.Context) error { return c.Authorization.CreateUser(context) })
	e.GET("/login", func(context echo.Context) error { return c.Authorization.Login(context) })

	return e
}
