package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type advertisementInterfactor struct {
	AdvertisementRepository repository.AdvertisementRepository
}

type AdvertisementInterfactor interface {
	Get(u []*model.Advertisement, limit string, offset string, orderQuery string) ([]*model.Advertisement, error)
	GetOne(u []*model.Advertisement, id string) ([]*model.Advertisement, error)
	Create(u *model.Advertisement) (*model.Advertisement, error)
	Update(u *model.Advertisement, id string) (*model.Advertisement, error)
	Delete(u []*model.Advertisement, id string) ([]*model.Advertisement, error)
}

func NewAdvertisementInterfactor(r repository.AdvertisementRepository) AdvertisementInterfactor {
	return &advertisementInterfactor{r}
}

func (us *advertisementInterfactor) Get(u []*model.Advertisement, limit string, offset string, orderQuery string) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.FindAll(u, limit, offset, orderQuery)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *advertisementInterfactor) GetOne(u []*model.Advertisement, id string) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.FindOne(u, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *advertisementInterfactor) Create(u *model.Advertisement) (*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.Create(u)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (us *advertisementInterfactor) Update(u *model.Advertisement, id string) (*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.Update(u, id)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (us *advertisementInterfactor) Delete(u []*model.Advertisement, id string) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.Delete(u, id)
	if err != nil {
		return nil, err
	}

	return u, err
}
