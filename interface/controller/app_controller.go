package controller

type AppController struct {
	Advertisement interface{ AdvertisementController }
	User          interface{ UserController }
	Comment       interface{ CommentController }
}
