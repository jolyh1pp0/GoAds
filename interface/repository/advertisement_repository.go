package repository

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
)

type advertisementRepository struct {
	db *gorm.DB
}

func NewAdvertisementRepository(db *gorm.DB) repository.AdvertisementRepository {
	return &advertisementRepository{db}
}

func (ar *advertisementRepository) FindAll(a []*model.GetAdvertisementsResponseData, limit string, offset string, orderQuery string) ([]*model.GetAdvertisementsResponseData, error) {
	err := ar.db.Limit(limit).Offset(offset).Model(&a).Select("title, description, photo_1, price, created_at, user_id").Order(orderQuery).Preload("User").Find(&a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (ar *advertisementRepository) FindOne(a []*model.GetAdvertisementsResponseData, id string) ([]*model.GetAdvertisementsResponseData, error) {
	err := ar.db.Model(&a).Select("*").Where("advertisements.id = ?", id).Preload("User").Preload("Comments.User").Find(&a).Error

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (ar *advertisementRepository) Create(a *model.Advertisement) error {
	err := ar.db.Model(&a).Create(a).Error

	if err != nil {
		if err.Error() == domain.ErrAdvertisementAlreadyWithTitle {
			return domain.ErrAdvertisementTitleAlreadyExists
		}
		return domain.ErrAdvertisementInternalServerError
	}
	return nil
}

func (ar *advertisementRepository) Update(a *model.Advertisement, id string) (*model.Advertisement, error) {
	err := ar.db.Model(&a).Where("id = ?", id).Update(a).Error

	if err != nil {
		if err.Error() == domain.ErrAdvertisementAlreadyWithTitle {
			return nil, domain.ErrAdvertisementTitleAlreadyExists
		}
		return nil, domain.ErrAdvertisementInternalServerError
	}

	return a, nil
}

func (ar *advertisementRepository) Delete(a []*model.Advertisement, id string) ([]*model.Advertisement, error) {
	err := ar.db.Model(&a).Where("id = ?", id).Delete(&a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}
