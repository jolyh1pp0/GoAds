package routes

import (
	"GoAds/domain"
	"GoAds/domain/model"
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

		claims, err := controller.ParseToken(headerParts[1])
		userID, userRoles := claims.UserID, claims.UserRoles

		if err != nil {
			log.Println(err)
			return domain.ErrInvalidAccessToken
		}

		c.Set(domain.SessionDataKey, domain.SessionData{
			UserID:    userID,
			UserRoles: userRoles,
		})

		return next(c)
	}
}

func Role(roleID ...int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRoles := getUserRolesID(c)

			for _, v := range roleID {
				for _, v2 := range userRoles {
					if v == v2 {
						return next(c)
					}
				}
			}

			return domain.ErrForbidden
		}
	}
}

func getUserRolesID(c echo.Context) []int {
	rawSessionData := c.Get(domain.SessionDataKey)
	sessionData, _ := rawSessionData.(domain.SessionData)

	return sessionData.UserRoles
}

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Recover())

	advertisementGroup := e.Group("/advertisements", Auth, Role(model.RoleAdvertisementID, model.RoleAdminID))
	{
		advertisementGroup.GET("", func(context echo.Context) error { return c.Advertisement.GetAdvertisements(context) })
		advertisementGroup.GET("/:id", func(context echo.Context) error { return c.Advertisement.GetOneAdvertisement(context) })
		advertisementGroup.PUT("/:id", func(context echo.Context) error { return c.Advertisement.UpdateAdvertisement(context) })
		advertisementGroup.POST("", func(context echo.Context) error { return c.Advertisement.CreateAdvertisement(context) })
		advertisementGroup.DELETE("/:id", func(context echo.Context) error { return c.Advertisement.DeleteAdvertisement(context) })
	}

	galleryGroup := e.Group("/gallery", Auth, Role(model.RoleAdvertisementID, model.RoleAdminID))
	{
		galleryGroup.POST("/:id", func(context echo.Context) error { return c.Gallery.AddPicture(context) })
		galleryGroup.DELETE("/:id", func(context echo.Context) error { return c.Gallery.DeletePicture(context) })
	}

	userGroup := e.Group("/users", Auth, Role(model.RoleUserID, model.RoleAdminID))
	{
		userGroup.GET("", func(context echo.Context) error { return c.User.GetUsers(context) })
		userGroup.GET("/:id", func(context echo.Context) error { return c.User.GetOneUser(context) })
		userGroup.PUT("/:id", func(context echo.Context) error { return c.User.UpdateUser(context) })
		userGroup.DELETE("/:id", func(context echo.Context) error { return c.User.DeleteUser(context) })
	}

	passwordRecoveryGroup := e.Group("/password-recovery")
	{
		passwordRecoveryGroup.POST("/reset", func(context echo.Context) error { return c.PasswordRecovery.ResetPassword(context) })
		passwordRecoveryGroup.POST("/set/:token", func(context echo.Context) error { return c.PasswordRecovery.SetPassword(context) })
	}

	commentGroup := e.Group("/comments", Auth, Role(model.RoleCommentID, model.RoleAdminID))
	{
		commentGroup.GET("", func(context echo.Context) error { return c.Comment.GetComments(context) })
		commentGroup.GET("/:id", func(context echo.Context) error { return c.Comment.GetOneComment(context) })
		commentGroup.PUT("/:id", func(context echo.Context) error { return c.Comment.UpdateComment(context) })
		commentGroup.POST("", func(context echo.Context) error { return c.Comment.CreateComment(context) })
		commentGroup.DELETE("/:id", func(context echo.Context) error { return c.Comment.DeleteComment(context) })
	}

	//roleGroup := e.Group("/roles", Auth)
	//{
	//	roleGroup.GET("", func(context echo.Context) error { return c.Role.GetRoles(context) })
	//	roleGroup.GET("/:id", func(context echo.Context) error { return c.Role.GetOneRole(context) })
	//	roleGroup.PUT("/:id", func(context echo.Context) error { return c.Role.UpdateRole(context) })
	//	roleGroup.POST("", func(context echo.Context) error { return c.Role.CreateRole(context) })
	//	roleGroup.DELETE("/:id", func(context echo.Context) error { return c.Role.DeleteRole(context) })
	//}

	userToRoleGroup := e.Group("/user-to-role", Auth, Role(model.RoleUserToRoleID, model.RoleAdminID))
	{
		userToRoleGroup.GET("", func(context echo.Context) error { return c.UserToRole.GetUserToRoles(context) })
		userToRoleGroup.GET("/", func(context echo.Context) error { return c.UserToRole.GetUserRoles(context) })
		userToRoleGroup.PUT("/:id", func(context echo.Context) error { return c.UserToRole.UpdateUserToRole(context) })
		userToRoleGroup.POST("", func(context echo.Context) error { return c.UserToRole.CreateUserToRole(context) })
		userToRoleGroup.DELETE("/:id", func(context echo.Context) error { return c.UserToRole.DeleteUserToRole(context) })
	}

	chatGroup := e.Group("/chat", Auth, Role(model.RoleUserID, model.RoleAdminID))
	{
		hub := controller.NewHub()
		go hub.Run()
		chatGroup.GET("", func(context echo.Context) error { return c.Chat.ServeWs(context, hub) })
	}

	e.POST("/register", func(context echo.Context) error { return c.Authorization.CreateUser(context) })
	e.GET("/login", func(context echo.Context) error { return c.Authorization.Login(context) })
	e.GET("/refresh", func(context echo.Context) error { return c.Authorization.Refresh(context) })
	e.GET("/logout", func(context echo.Context) error { return c.Authorization.Logout(context) })

	return e
}
