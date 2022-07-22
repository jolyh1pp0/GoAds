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
		Advertisement: r.NewAdvertisementController(),
		User:          r.NewUserController(),
		Comment:       r.NewCommentController(),
	}
}
