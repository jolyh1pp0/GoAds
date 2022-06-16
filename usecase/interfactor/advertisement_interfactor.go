package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/presenter"
	"GoAds/usecase/repository"
)

type advertisementInterfactor struct {
	AdvertisementRepository repository.AdvertisementRepository
	AdvertisementPresenter  presenter.AdvertisementPresenter
	//DBRepository   repository.DBRepository
}

type AdvertisementInterfactor interface {
	Get(u []*model.Advertisement) ([]*model.Advertisement, error)
	Create(u *model.Advertisement) (*model.Advertisement, error)
}

func NewAdvertisementInteractor(r repository.AdvertisementRepository, p presenter.AdvertisementPresenter) AdvertisementInterfactor {
	return &advertisementInterfactor{r, p}
}

func (us *advertisementInterfactor) Get(u []*model.Advertisement) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.AdvertisementPresenter.ResponseAdvertisements(u), nil
}

func (us *advertisementInterfactor) Create(u *model.Advertisement) (*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.Create(u)

	// do mailing
	// do logging
	// do another process
	return u, err
}
