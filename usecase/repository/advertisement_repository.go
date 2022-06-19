package repository

import "GoAds/domain/model"

type AdvertisementRepository interface {
	FindAll(u []*model.Advertisement) ([]*model.Advertisement, error)
	FindOne(u []*model.Advertisement, id string) ([]*model.Advertisement, error)
	Create(u *model.Advertisement) (*model.Advertisement, error)
}
