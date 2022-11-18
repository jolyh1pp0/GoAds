package interfactor

import (
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"log"
	"time"
)

type passwordRecoveryInterfactor struct {
	PasswordRecoveryRepository repository.PasswordRecoveryRepository
}

type PasswordRecoveryInterfactor interface {
	GetEmailValidation(email string) (string, bool, error)
	GetTokenExists(email string) (time.Time, bool, error)
	CreateRecovery(passwdRecovery *model.PasswordRecovery) error
	DeleteRecovery(email string) error
	FindToken(token string) (bool, string, time.Time, error)
	UpdatePassword(userID string, password string) error
}

func NewPasswordRecoveryInterfactor(r repository.PasswordRecoveryRepository) PasswordRecoveryInterfactor {
	return &passwordRecoveryInterfactor{r}
}

func (pri passwordRecoveryInterfactor) GetEmailValidation(email string) (string, bool, error) {
	userID, valid, err := pri.PasswordRecoveryRepository.GetEmailValidation(email)
	if err != nil {
		log.Print(err)
		return "", false, err
	}

	return userID, valid, nil
}

func (pri passwordRecoveryInterfactor) GetTokenExists(email string) (time.Time, bool, error) {
	tokenExpiration, exists, err := pri.PasswordRecoveryRepository.GetTokenExists(email)
	if err != nil {
		log.Print(err)
		return time.Time{}, false, err
	}

	return tokenExpiration, exists, nil
}

func (pri passwordRecoveryInterfactor) CreateRecovery(passwdRecovery *model.PasswordRecovery) error {
	err := pri.PasswordRecoveryRepository.CreateRecovery(passwdRecovery)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (pri passwordRecoveryInterfactor) DeleteRecovery(userID string) error {
	err := pri.PasswordRecoveryRepository.DeleteRecovery(userID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (pri passwordRecoveryInterfactor) FindToken(token string) (bool, string, time.Time, error) {
	exists, userID, tokenExpiration, err := pri.PasswordRecoveryRepository.FindToken(token)
	if err != nil {
		log.Print(err)
		return false, "", time.Time{}, err
	}

	return exists, userID, tokenExpiration, nil
}

func (pri passwordRecoveryInterfactor) UpdatePassword(userID string, password string) error {
	err := pri.PasswordRecoveryRepository.UpdatePassword(userID, password)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
