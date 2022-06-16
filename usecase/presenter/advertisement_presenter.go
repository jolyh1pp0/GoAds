package presenter

import "GoAds/domain/model"

type AdvertisementPresenter interface {
	ResponseAdvertisements(u []*model.Advertisement) []*model.Advertisement
}
