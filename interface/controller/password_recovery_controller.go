package controller

import (
	"GoAds/config"
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"context"
	"github.com/twinj/uuid"
	"log"
	"net/http"
	"time"
)

type passwordRecoveryController struct {
	passwordRecoveryInterfactor interfactor.PasswordRecoveryInterfactor
}

type PasswordRecoveryController interface {
	ResetPassword(c Context) error
	SetPassword(c Context) error
}

func NewPasswordRecoveryController(ad interfactor.PasswordRecoveryInterfactor) PasswordRecoveryController {
	return &passwordRecoveryController{ad}
}

func sendEmail(recipient string, token string) error {
	// any recipient cannot be used as it is a test domain
	// for testing we use an authorized recipient, watch readme how to add it
	authorizedRecipient := "" // your authorized recipient
	sender := "test@GoAds"
	subject := "Password recovery"
	body := "Your link: http://localhost:8080/password_recovery/set/" + token + "\nPaste this into Postman."
	message := config.MG.NewMessage(sender, subject, body, authorizedRecipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := config.MG.Send(ctx, message)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (prc *passwordRecoveryController) ResetPassword(c Context) error {
	var passwdRecovery model.PasswordRecovery

	err := c.Bind(&passwdRecovery)
	if err != nil {
		log.Print(err)
		return err
	}

	userID, valid, err := prc.passwordRecoveryInterfactor.GetEmailValidation(passwdRecovery.UserEmail)
	if err != nil {
		log.Print(err)
		return err
	}
	if !valid {
		return domain.ErrInvalidEmail
	}

	tokenUUID := uuid.NewV4().String()
	passwdRecovery.UserID = userID
	passwdRecovery.Token = tokenUUID
	passwdRecovery.ExpiresAt = time.Now().Add(time.Hour*10)

	existsTokenExpiration, exists, err := prc.passwordRecoveryInterfactor.GetTokenExists(passwdRecovery.UserEmail)
	if err != nil {
		log.Print(err)
		return err
	}

	if exists {
		if time.Now().After(existsTokenExpiration) {
			err = prc.passwordRecoveryInterfactor.DeleteRecovery(passwdRecovery.UserID)
			if err != nil {
				log.Print(err)
				return err
			}
			err = prc.passwordRecoveryInterfactor.CreateRecovery(&passwdRecovery)
			if err != nil {
				log.Print(err)
				return err
			}
			err = sendEmail(passwdRecovery.UserEmail, tokenUUID)
			if err != nil {
				log.Print(err)
				return err
			}

		} else {
			return domain.ErrTokenActive
		}
	} else {
		err = prc.passwordRecoveryInterfactor.CreateRecovery(&passwdRecovery)
		if err != nil {
			log.Print(err)
			return err
		}

		err = sendEmail(passwdRecovery.UserEmail, tokenUUID)
		if err != nil {
			log.Print(err)
			return err
		}
	}

	return c.JSONPretty(http.StatusOK, "A password reset link has been sent to your email.", " ")
}

type PasswordForm struct {
	Password string
}

func (prc *passwordRecoveryController) SetPassword(c Context) error {
	var passwdForm PasswordForm
	token := c.Param("token")

	err := c.Bind(&passwdForm)
	if err != nil {
		log.Print(err)
		return err
	}

	exists, userID, existsTokenExpiration, err := prc.passwordRecoveryInterfactor.FindToken(token)
	if err != nil {
		log.Print(err)
		return domain.ErrTokenInvalid
	}

	if time.Now().After(existsTokenExpiration) {
		err = prc.passwordRecoveryInterfactor.DeleteRecovery(userID)
		if err != nil {
			log.Print(err)
			return err
		}
		return domain.ErrTokenInvalid
	}

	if !exists {
		return domain.ErrTokenInvalid
	}

	validatedPassword, err := ValidatePassword(passwdForm.Password)
	if err != nil {
		log.Print(err)
		return err
	}

	err = prc.passwordRecoveryInterfactor.UpdatePassword(userID, string(validatedPassword))
	if err != nil {
		log.Print(err)
		return err
	}

	err = prc.passwordRecoveryInterfactor.DeleteRecovery(userID)
	if err != nil {
		log.Print(err)
		return err
	}

	return c.JSONPretty(http.StatusOK, "Password updated successfully.", " ")
}