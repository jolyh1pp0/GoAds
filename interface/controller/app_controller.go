package controller

type AppController struct {
	Advertisement    interface{ AdvertisementController }
	Gallery          interface{ GalleryController }
	User             interface{ UserController }
	Comment          interface{ CommentController }
	Authorization    interface{ AuthorizationController }
	Role             interface{ RoleController }
	UserToRole       interface{ UserToRoleController }
	PasswordRecovery interface{ PasswordRecoveryController }
	Chat 			 interface{ ChatController }
}
