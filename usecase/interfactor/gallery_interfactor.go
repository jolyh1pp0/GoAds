package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
)

type galleryInterfactor struct {
	GalleryRepository repository.GalleryRepository
}

type GalleryInterfactor interface {
	GetAdvertisementUserId(id string) (string, error)
	GetAdvertisementId(id string) (uint, error)
	Create(g *model.Gallery) error
	Delete(g *model.Gallery, id string) error
}

func NewGalleryInterfactor(r repository.GalleryRepository) GalleryInterfactor {
	return &galleryInterfactor{r}
}

func (gi galleryInterfactor) GetAdvertisementUserId(id string) (string, error) {
	userID, err := gi.GalleryRepository.GetUserID(id)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (gi galleryInterfactor) GetAdvertisementId(id string) (uint, error) {
	adID, err := gi.GalleryRepository.GetAdvertisementID(id)
	if err != nil {
		return 0, err
	}

	return adID, nil
}

func (gi galleryInterfactor) Create(g *model.Gallery) error {
	err := gi.GalleryRepository.Create(g)
	if err != nil {
		return err
	}

	return err
}

func (gi galleryInterfactor) Delete(g *model.Gallery, id string) error {
	err := gi.GalleryRepository.Delete(g, id)
	if err != nil {
		return err
	}

	return err
}