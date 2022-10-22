package controller

import (
	goAdsConf "GoAds/config"
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"net/http"
	"strconv"
)

type galleryController struct {
	galleryInterfactor interfactor.GalleryInterfactor
}

type GalleryController interface {
	AddPicture(c Context) error
	DeletePicture(c Context) error
}

func NewGalleryController(ad interfactor.GalleryInterfactor) GalleryController {
	return &galleryController{ad}
}

type Base64 struct {
	FileBase64 string   `json:"file_base_64"`
	FileName     string `json:"file_name"`
}

func checkIfAdmin(c Context) bool {
	roles := getUserRole(c)
	for _, role := range roles {
		if role == 5 {
			return true
		}
	}
	return false
}

func checkIfAccess(gc *galleryController, advertisementId string, userID string) bool {
	advertisementUserID, err := gc.galleryInterfactor.GetAdvertisementUserId(advertisementId)
	if err != nil {
		log.Print(err)
		return false
	}

	if userID != advertisementUserID {
		return false
	}

	return true
}

func decodeBase64(base64String string) []byte {
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Println(err)
	}

	return data
}

func createBucket() (*s3.Client , error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return client, nil
}

func uploadImageToBucket(client *s3.Client, base64String, fileName string) (string, error) {
	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{ // result
		Bucket: aws.String(goAdsConf.C.S3.BucketName),
		Key:    aws.String("images/" + fileName),
		Body:   bytes.NewReader(decodeBase64(base64String)),
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	return result.Location, nil
}

func (gc *galleryController) AddPicture(c Context) error {
	var gallery model.Gallery
	var base64Data Base64
	advertisementId := c.Param("id")

	if !checkIfAdmin(c) {
		userID := getUserID(c)
		if !checkIfAccess(gc, advertisementId, userID) {
			return domain.ErrForbidden
		}
	}

	err := c.Bind(&base64Data)
	if err != nil {
		log.Print(err)
		return err
	}

	fileExtension := ""

	fileExtensionByte := base64Data.FileBase64[0:1]
	if fileExtensionByte == "/" {
		fileExtension = ".jpg"
	} else if fileExtensionByte == "i" {
		fileExtension = ".png"
	} else if fileExtensionByte == "u" {
		fileExtension = ".webp"
	} else {
		return domain.ErrInvalidPictureExtension
	}

	gallery.FileName = base64Data.FileName + fileExtension

	adID, _ := strconv.ParseUint(advertisementId, 10, 32)
	gallery.AdvertisementId = uint(adID)

	bucketClient, err := createBucket()
	if err != nil {
		return err
	}

	filePath, err := uploadImageToBucket(bucketClient, base64Data.FileBase64, base64Data.FileName+fileExtension)
	if err != nil {
		log.Println(err)
		return err
	}

	gallery.FilePath = filePath

	err = gc.galleryInterfactor.Create(&gallery)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, "Picture for advertisement #" + advertisementId + " added.", " ")
}

func (gc *galleryController) DeletePicture(c Context) error {
	var gallery model.Gallery
	pictureId := c.Param("id")

	if !checkIfAdmin(c) {
		userID := getUserID(c)

		advertisementID, err := gc.galleryInterfactor.GetAdvertisementId(pictureId)
		if err != nil {
			log.Print(err)
			return err
		}

		if !checkIfAccess(gc, strconv.Itoa(int(advertisementID)), userID) {
			return domain.ErrForbidden
		}
	}

	err := gc.galleryInterfactor.Delete(&gallery, pictureId)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusOK, "Picture deleted", "  ")
}