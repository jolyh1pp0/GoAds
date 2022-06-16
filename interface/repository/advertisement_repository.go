package repository

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"errors"
	"github.com/jinzhu/gorm"
)

type advertisementRepository struct {
	db *gorm.DB
}

func NewAdvertisementRepository(db *gorm.DB) repository.AdvertisementRepository {
	return &advertisementRepository{db}
}

func (ar *advertisementRepository) FindAll(a []*model.Advertisement) ([]*model.Advertisement, error) {
	err := ar.db.Find(&a).Error

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (ar *advertisementRepository) Create(a *model.Advertisement) (*model.Advertisement, error) {
	if err := ar.db.Create(a).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return a, nil
}
