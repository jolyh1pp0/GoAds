package controller

import (
	"GoAds/config"
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

var key = []byte(config.C.JWT.Key)

type authorizationController struct {
	authorizationInterfactor interfactor.AuthorizationInterfactor
}

type tokenClaims struct {
	UserID    string `json:"user_id"`
	UserRoles []int  `json:"user_roles"`
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

	return c.JSONPretty(http.StatusCreated, "Status 201. User " + u.ID + " created", "  ")
}


func GenerateJWT(userID string, userRoles []int) (string, error) {
	claims := tokenClaims{
		UserID: userID,
		UserRoles: userRoles,
	}

	token, err := jwt.Sign(jwt.HS256, key, claims, jwt.MaxAge(15*time.Minute))
	if err != nil {
		log.Println(err)
		return "", domain.ErrInvalidAccessToken
	}

	return string(token), nil
}

func ParseToken(accessToken string) (string, []int, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, key, []byte(accessToken))
	if err != nil {
		log.Println(err)
		return "", nil, domain.ErrInvalidAccessToken
	}

	claims := &tokenClaims{}
	err = verifiedToken.Claims(claims)
	if err != nil {
		log.Println(err)
		return "", nil, domain.ErrInvalidAccessToken
	}

	return claims.UserID, claims.UserRoles, nil
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

	userRoles, err := ac.authorizationInterfactor.GetUserRoles(userID)
	if err != nil {
		log.Print(err)
		return domain.ErrEmailIsNotFound // TODO:Error
	}

	token, err := GenerateJWT(userID, userRoles)
	if err != nil {
		log.Print(err)
	}

	return c.JSONPretty(http.StatusOK, "Status 200. Successfully logged in. Access token: " + token , "")
}

func getUserID(c Context) string {
	rawSessionData := c.Get(domain.SessionDataKey)
	sessionData, _ := rawSessionData.(domain.SessionData)

	return sessionData.UserID
}
