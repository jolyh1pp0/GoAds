package repository

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type passwordRecoveryRepository struct {
	db *gorm.DB
}

func NewPasswordRecoveryRepository(db *gorm.DB) repository.PasswordRecoveryRepository {
	return &passwordRecoveryRepository{db}
}

func (gr *passwordRecoveryRepository) GetEmailValidation(email string) (string, bool, error) {
	u := model.User{}

	result := gr.db.Model(&u).Select("id, email").Where("email = ?", email).Find(&u)

	if result.Error == gorm.ErrRecordNotFound {
		return "", false, nil
	} else if result.RowsAffected >= 1 {
		return u.ID, true, nil
	} else {
		log.Print(result.Error)
		return "", false, result.Error
	}
}

func (gr *passwordRecoveryRepository) GetTokenExists(email string) (time.Time, bool, error) {
	pr := model.PasswordRecovery{}

	result := gr.db.Model(&pr).Select("user_email, expires_at").Where("user_email = ?", email).Find(&pr)

	if result.Error == gorm.ErrRecordNotFound {
		return time.Time{}, false, nil
	} else if result.RowsAffected >= 1 {
		return pr.ExpiresAt, true, nil
	} else {
		log.Print(result.Error)
		return time.Time{}, false, result.Error
	}
}

func (gr *passwordRecoveryRepository) CreateRecovery(passwdRecovery *model.PasswordRecovery) error {
	err := gr.db.Model(&passwdRecovery).Create(passwdRecovery).Error
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (gr *passwordRecoveryRepository) DeleteRecovery(userID string) error {
	var passwdRecovery model.PasswordRecovery

	err := gr.db.Model(&passwdRecovery).Where("user_id = ?", userID).Delete(&passwdRecovery).Error
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (gr *passwordRecoveryRepository) FindToken(token string) (bool, string, time.Time, error) {
	pr := model.PasswordRecovery{}

	result := gr.db.Model(&pr).Select("token, user_id, expires_at").Where("token = ?", token).Find(&pr)

	if result.Error == gorm.ErrRecordNotFound {
		return false, "", time.Time{}, nil
	} else if result.RowsAffected >= 1 {
		return true, pr.UserID, pr.ExpiresAt, nil
	} else {
		log.Print(result.Error)
		return false, "", time.Time{}, result.Error
	}
}

func (gr *passwordRecoveryRepository) UpdatePassword(userID string, password string) error {
	u := model.User{}
	u.Password = password

	err := gr.db.Model(&u).Where("id = ?", userID).Update(u).Error
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
