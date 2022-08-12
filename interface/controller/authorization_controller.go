package controller

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"github.com/kataras/jwt"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/mail"
	"time"
)

var key = []byte("example_key")

type authorizationController struct {
	authorizationInterfactor interfactor.AuthorizationInterfactor
}

type tokenClaims struct { // token lowercase TODO
	UserID string `json:"user_id"`
}

type AuthorizationController interface {
	CreateUser(c Context) error
	Login(c Context) error
}

func NewAuthorizationController(ai interfactor.AuthorizationInterfactor) AuthorizationController {
	return &authorizationController{ai}
}

func (ac *authorizationController) CreateUser(c Context) error {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		log.Print(err)
	}

	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidEmailAddress
	}

	const minEntropyBits = 60
	err = passwordvalidator.Validate(user.Password, minEntropyBits)
	if err != nil {
		log.Print(err)
		return domain.ErrInsecurePassword
	}

	bPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
	}
	user.Password = string(bPass)

	u, err := ac.authorizationInterfactor.Create(&user)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, u, "  ")
}


func GenerateJWT(userID string) (string, error) {
	claims := tokenClaims{
		UserID: userID,
	}

	token, err := jwt.Sign(jwt.HS256, key, claims, jwt.MaxAge(15*time.Minute))
	if err != nil {
		log.Println(err)
		return "", domain.ErrInvalidAccessToken
	}

	return string(token), nil
}

func ParseToken(accessToken string) (string, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, key, []byte(accessToken))
	if err != nil {
		log.Println(err)
		return "", domain.ErrInvalidAccessToken
	}

	claims := &tokenClaims{}
	err = verifiedToken.Claims(claims)
	if err != nil {
		log.Println(err)
		return "", domain.ErrInvalidAccessToken
	}

	return claims.UserID, nil
}

func (ac *authorizationController) Login(c Context) error {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		log.Print(err)
	}

	hPass, userID, err := ac.authorizationInterfactor.UserExists(user.Email)
	if err != nil {
		log.Print(err)
		return domain.ErrEmailIsNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(hPass), []byte(user.Password))
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidPassword
	}

	token, err := GenerateJWT(userID)
	if err != nil {
		log.Print(err)
	}

	return c.JSONPretty(http.StatusOK, "Successfully logged in. Access token: " + token , "")
}