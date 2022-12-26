package repository

import "GoAds/domain/model"

type AdvertisementRepository interface {
	FindAll(u []*model.GetAdvertisementsResponseData, limit string, offset string, orderQuery string) ([]*model.GetAdvertisementsResponseData, error)
	FindOne(u []*model.GetAdvertisementsResponseData, id string) ([]*model.GetAdvertisementsResponseData, error)
	Create(u *model.AdvertisementsRequestData) error
	Update(u *model.AdvertisementsUpdateRequestData, id string, userID string) (*model.AdvertisementsUpdateRequestData, error)
	Delete(u []*model.Advertisement, id string, userID string) ([]*model.Advertisement, error)
}
