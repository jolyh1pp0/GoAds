package repository

import (
	goAdsConf "GoAds/config"
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jinzhu/gorm"
)

type galleryRepository struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) repository.GalleryRepository {
	return &galleryRepository{db}
}

func (gr *galleryRepository) GetUserID(id string) (string, error) {
	a := model.Advertisement{}

	err := gr.db.Model(&a).Select("user_id").Where("advertisements.id = ?", id).Find(&a).Error

	if err != nil {
		return "", err
	}

	return a.UserID, nil
}

func (gr *galleryRepository) GetAdvertisementID(id string) (uint, error) {
	g := model.Gallery{}

	err := gr.db.Model(&g).Select("advertisement_id").Where("id = ?", id).Find(&g).Error

	if err != nil {
		return 0, err
	}

	return g.AdvertisementId, nil
}

func (gr *galleryRepository) Create(g *model.Gallery) error {
	var gallery model.Gallery
	var result *gorm.DB

	result = gr.db.Model(&g).Select("*").Where("advertisement_id = ?", g.AdvertisementId).Find(&gallery)

	err := gr.db.Transaction(func(tx *gorm.DB) error {
		if result.RowsAffected > 9 {
			return domain.PictureLimitReached
		}
		if err := gr.db.Model(&g).Create(g).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	if err != nil {
		return domain.PictureLimitReached
	}

	return nil
}

func deleteImageFromBucket(client *s3.Client, item string) error {
	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(goAdsConf.C.S3.BucketName),
		Key:    aws.String("images/" + item),
	})

	if err != nil {
		return err
	}

	return nil
}

func (gr *galleryRepository) Delete(g *model.Gallery, id string) error {
	err := gr.db.Transaction(func(tx *gorm.DB) error {
		err := gr.db.Model(&g).Select("file_name").Where("id = ?", id).Find(&g).Error
		if err != nil {
			return err
		}

		if err = gr.db.Model(&g).Where("id = ?", id).Delete(&g).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err = deleteImageFromBucket(goAdsConf.BucketClient, g.FileName); err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}