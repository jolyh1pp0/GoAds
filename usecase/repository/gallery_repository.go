package repository

import "GoAds/domain/model"

type GalleryRepository interface {
	GetUserID(id string) (string, error)
	GetAdvertisementID(id string) (uint, error)
	Create(u *model.Gallery) error
	Delete(u *model.Gallery, id string) error
}
