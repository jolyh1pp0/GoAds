package routes

import (
	"GoAds/domain"
	"GoAds/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"strings"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return domain.ErrInvalidAccessToken
		}

		if headerParts[0] != "Bearer" {
			return domain.ErrInvalidAccessToken
		}

		userID, err := controller.ParseToken(headerParts[1])
		if err != nil {
			log.Println(err)
			return domain.ErrInvalidAccessToken
		}

		c.Set("UserID", userID)

	return next(c)
	}
}

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Recover())

	advertisementGroup := e.Group("/advertisements", Auth)
	{
		advertisementGroup.GET("", func(context echo.Context) error { return c.Advertisement.GetAdvertisements(context) })
		advertisementGroup.GET("/:id", func(context echo.Context) error { return c.Advertisement.GetOneAdvertisement(context) })
		advertisementGroup.PUT("/:id", func(context echo.Context) error { return c.Advertisement.UpdateAdvertisement(context) })
		advertisementGroup.POST("", func(context echo.Context) error { return c.Advertisement.CreateAdvertisement(context) })
		advertisementGroup.DELETE("/:id", func(context echo.Context) error { return c.Advertisement.DeleteAdvertisement(context) })
	}

	userGroup := e.Group("/users", Auth)
	{
		userGroup.GET("", func(context echo.Context) error { return c.User.GetUsers(context) })
		userGroup.GET("/:id", func(context echo.Context) error { return c.User.GetOneUser(context) })
		userGroup.PUT("/:id", func(context echo.Context) error { return c.User.UpdateUser(context) })
		userGroup.DELETE("/:id", func(context echo.Context) error { return c.User.DeleteUser(context) })
	}

	commentGroup := e.Group("/comments", Auth)
	{
		commentGroup.GET("", func(context echo.Context) error { return c.Comment.GetComments(context) })
		commentGroup.GET("/:id", func(context echo.Context) error { return c.Comment.GetOneComment(context) })
		commentGroup.PUT("/:id", func(context echo.Context) error { return c.Comment.UpdateComment(context) })
		commentGroup.POST("", func(context echo.Context) error { return c.Comment.CreateComment(context) })
		commentGroup.DELETE("/:id", func(context echo.Context) error { return c.Comment.DeleteComment(context) })
	}

	e.POST("/register", func(context echo.Context) error { return c.Authorization.CreateUser(context) })
	e.GET("/login", func(context echo.Context) error { return c.Authorization.Login(context) })

	return e
}
