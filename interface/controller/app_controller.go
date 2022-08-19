package controller

type AppController struct {
	Advertisement interface{ AdvertisementController }
	User          interface{ UserController }
	Comment       interface{ CommentController }
	Authorization interface{ AuthorizationController }
	Role          interface{ RoleController }
	UserToRole    interface{ UserToRoleController }
}
