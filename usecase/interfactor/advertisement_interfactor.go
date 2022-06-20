package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/presenter"
	"GoAds/usecase/repository"
)

type advertisementInterfactor struct {
	AdvertisementRepository repository.AdvertisementRepository
	AdvertisementPresenter  presenter.AdvertisementPresenter
}

type AdvertisementInterfactor interface {
	Get(u []*model.Advertisement) ([]*model.Advertisement, error)
	GetOne(u []*model.Advertisement, id string) ([]*model.Advertisement, error)
	Create(u *model.Advertisement) (*model.Advertisement, error)
	Delete(u []*model.Advertisement, id string) ([]*model.Advertisement, error)
}

func NewAdvertisementInterfactor(r repository.AdvertisementRepository, p presenter.AdvertisementPresenter) AdvertisementInterfactor {
	return &advertisementInterfactor{r, p}
}

func (us *advertisementInterfactor) Get(u []*model.Advertisement) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.AdvertisementPresenter.ResponseAdvertisements(u), nil
}

func (us *advertisementInterfactor) GetOne(u []*model.Advertisement, id string) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.FindOne(u, id)
	if err != nil {
		return nil, err
	}

	return us.AdvertisementPresenter.ResponseAdvertisements(u), nil
}

func (us *advertisementInterfactor) Create(u *model.Advertisement) (*model.Advertisement, error) {

	user := model.Advertisement{
		ID:          333,
		Title:       "333",
		Description: "333",
		Price:       333,
		Photo_1:     "333",
		Photo_2:     "",
		Photo_3:     "",
	}
	u, err := us.AdvertisementRepository.Create(&user)
	if err != nil {
		return nil, err
	}
	// TODO: Create
	return u, err
}

func (us *advertisementInterfactor) Delete(u []*model.Advertisement, id string) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.Delete(u, id)
	if err != nil {
		return nil, err
	}
	// TODO: Create
	return u, err
}
