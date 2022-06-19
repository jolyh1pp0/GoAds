package presenter

import (
	"GoAds/domain/model"
	"GoAds/usecase/presenter"
)

type advertisementPresenter struct{}

func NewAdvertisementPresenter() presenter.AdvertisementPresenter {
	return &advertisementPresenter{}
}

func (ap *advertisementPresenter) ResponseAdvertisements(us []*model.Advertisement) []*model.Advertisement {
	//for _, u := range us {
	//	u.Title = "Mr." + u.Title
	//}

	return us
}
