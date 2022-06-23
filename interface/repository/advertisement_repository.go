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

func (ar *advertisementRepository) FindAll(a []*model.Advertisement, limit string, offset string, orderQuery string) ([]*model.Advertisement, error) {
	err := ar.db.Limit(limit).Offset(offset).Select("title, photo_1, price").Order(orderQuery).Find(&a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (ar *advertisementRepository) FindOne(a []*model.Advertisement, id string) ([]*model.Advertisement, error) {
	err := ar.db.Select("title, description, photo_1, photo_2, photo_3, price").Where("id = ?", id).Find(&a).Error
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

func (ar *advertisementRepository) Delete(a []*model.Advertisement, id string) ([]*model.Advertisement, error) {
	err := ar.db.Where("id = ?", id).Delete(&a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}
