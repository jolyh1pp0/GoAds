package repository

import "GoAds/domain/model"

type AdvertisementRepository interface {
	FindAll(u []*model.GetAdvertisementsResponseData, limit string, offset string, orderQuery string) ([]*model.GetAdvertisementsResponseData, error)
	FindOne(u []*model.GetAdvertisementsResponseData, id string) ([]*model.GetAdvertisementsResponseData, error)
	Create(u *model.Advertisement) error
	Update(u *model.Advertisement, id string) (*model.Advertisement, error)
	Delete(u []*model.Advertisement, id string) ([]*model.Advertisement, error)
}
