package registry

import (
	"GoAds/interface/controller"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Advertisement:    r.NewAdvertisementController(),
		Gallery:          r.NewGalleryController(),
		User:             r.NewUserController(),
		Comment:          r.NewCommentController(),
		Authorization:    r.NewAuthorizationController(),
		Role:             r.NewRoleController(),
		UserToRole:       r.NewUserToRoleController(),
		PasswordRecovery: r.NewPasswordRecoveryController(),
	}
}
