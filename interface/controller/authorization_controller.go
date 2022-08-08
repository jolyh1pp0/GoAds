package controller

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/mail"
	"time"
)

const key = "example_key"

type authorizationController struct {
	authorizationInterfactor interfactor.AuthorizationInterfactor
}

type tokenClaims struct {
	jwt.StandardClaims
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
		userID,
	})

	return token.SignedString([]byte(key))
}

func ParseToken(accessToken string) (string, error){
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims*")
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
	fmt.Println(token)

	return c.JSONPretty(http.StatusOK, "Successfully logged in", "")
}
