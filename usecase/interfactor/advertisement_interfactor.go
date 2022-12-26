package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type advertisementInterfactor struct {
	AdvertisementRepository repository.AdvertisementRepository
}

type AdvertisementInterfactor interface {
	Get(u []*model.GetAdvertisementsResponseData, limit string, offset string, orderQuery string) ([]*model.GetAdvertisementsResponseData, error)
	GetOne(u []*model.GetAdvertisementsResponseData, id string) ([]*model.GetAdvertisementsResponseData, error)
	Create(u *model.AdvertisementsRequestData) error
	Update(u *model.AdvertisementsUpdateRequestData, id string, userID string) (*model.AdvertisementsUpdateRequestData, error)
	Delete(u []*model.Advertisement, id string, userID string) ([]*model.Advertisement, error)
}

func NewAdvertisementInterfactor(r repository.AdvertisementRepository) AdvertisementInterfactor {
	return &advertisementInterfactor{r}
}

func (us *advertisementInterfactor) Get(u []*model.GetAdvertisementsResponseData, limit string, offset string, orderQuery string) ([]*model.GetAdvertisementsResponseData, error) {
	u, err := us.AdvertisementRepository.FindAll(u, limit, offset, orderQuery)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *advertisementInterfactor) GetOne(u []*model.GetAdvertisementsResponseData, id string) ([]*model.GetAdvertisementsResponseData, error) {
	u, err := us.AdvertisementRepository.FindOne(u, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *advertisementInterfactor) Create(u *model.AdvertisementsRequestData) error {
	err := us.AdvertisementRepository.Create(u)
	if err != nil {
		return err
	}

	return err
}

func (us *advertisementInterfactor) Update(u *model.AdvertisementsUpdateRequestData, id string, userID string) (*model.AdvertisementsUpdateRequestData, error) {
	u, err := us.AdvertisementRepository.Update(u, id, userID)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (us *advertisementInterfactor) Delete(u []*model.Advertisement, id string, userID string) ([]*model.Advertisement, error) {
	u, err := us.AdvertisementRepository.Delete(u, id, userID)
	if err != nil {
		return nil, err
	}

	return u, err
}
