package controller

import (
	"GoAds/config"
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"errors"
	"github.com/kataras/jwt"
	"github.com/twinj/uuid"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/mail"
	"time"
)

var key = []byte(config.C.JWT.Key)
var refreshKey = []byte(config.C.JWT.RefreshKey)

type authorizationController struct {
	authorizationInterfactor interfactor.AuthorizationInterfactor
}

type tokenClaims struct {
	UserID      string `json:"user_id"`
	UserRoles   []int  `json:"user_roles"`
	AccessUUID  string `json:"access_uuid"`
	RefreshUUID string `json:"refresh_uuid"`
}

type AuthorizationController interface {
	CreateUser(c Context) error
	Login(c Context) error
	Refresh(c Context) error
	createSession(token string, refreshToken string) error
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

func GenerateJWT(userID string, userRoles []int) (string, string, error) {
	accessUUID := uuid.NewV4().String()
	refreshUUID := uuid.NewV4().String()

	claims := tokenClaims{
		UserID:      userID,
		UserRoles:   userRoles,
		AccessUUID:  accessUUID,
		RefreshUUID: refreshUUID,
	}

	token, err := jwt.Sign(jwt.HS256, key, claims, jwt.MaxAge(15*time.Minute))
	if err != nil {
		log.Println(err)
		return "", "", domain.ErrInvalidAccessToken
	}

	refreshToken, err := jwt.Sign(jwt.HS256, refreshKey, claims, jwt.MaxAge(time.Hour*24*2))
	if err != nil {
		log.Println(err)
		return "", "", domain.ErrInvalidAccessToken
	}

	return string(token), string(refreshToken), nil
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

func getUUIDs(accessToken string) (string, string, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, key, []byte(accessToken))
	if err != nil {
		log.Println(err)
		return "", "", domain.ErrInvalidAccessToken
	}

	claims := &tokenClaims{}
	err = verifiedToken.Claims(claims)
	if err != nil {
		log.Println(err)
		return "", "", domain.ErrInvalidAccessToken
	}

	return claims.AccessUUID, claims.RefreshUUID, nil
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
		return domain.ErrEmailIsNotFound
	}

	token, refreshToken, err := GenerateJWT(userID, userRoles)
	if err != nil {
		log.Print(err)
	}

	tokens := map[string]string{
		"access_token":  token,
		"refresh_token": refreshToken,
	}

	err = ac.createSession(token, refreshToken)
	if err != nil {
		log.Print(err)
	}

	return c.JSONPretty(http.StatusOK, tokens, "")
}

func(ac *authorizationController) createSession(token string, refreshToken string) error {
	var session model.Session
	session.AccessToken = token
	session.RefreshToken = refreshToken
	session.AccessTokenUUID, session.RefreshTokenUUID, _ = getUUIDs(token)

	_, err := ac.authorizationInterfactor.CreateSession(&session)
	if err != nil {
		log.Print(err)
	}

	return nil
}

func (ac *authorizationController) Refresh(c Context) error {
	var session model.Session

	err := c.Bind(&session)
	if err != nil {
		log.Print(err)
	}

	refreshToken := session.RefreshToken

	_, err = jwt.Verify(jwt.HS256, key, []byte(refreshToken))
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidRefreshToken
	}

	_, session.RefreshTokenUUID, _ = getUUIDs(refreshToken)

	userID, userRoles, err := ParseToken(session.AccessToken)
	if err != nil {
		log.Print(err)
	}

	token, refreshToken, err := GenerateJWT(userID, userRoles)
	if err != nil {
		log.Print(err)
	}

	tokens := map[string]string{
		"access_token":  token,
		"refresh_token": refreshToken,
	}

	err = ac.createSession(token, refreshToken)
	if err != nil {
		log.Print(err)
	}

	//refreshTokenUUID, err := ac.authorizationInterfactor.GetRefreshTokenUUIDFromTable(session.RefreshTokenUUID)
	//if err != nil {
	//	log.Print(err)
	//}
	//
	//if refreshTokenUUID != session.RefreshTokenUUID {
	//	log.Println(err)
	//	return domain.ErrInvalidRefreshToken
	//}

	return c.JSONPretty(http.StatusOK, tokens, "")
}

func getUserID(c Context) string {
	rawSessionData := c.Get(domain.SessionDataKey)
	sessionData, _ := rawSessionData.(domain.SessionData)

	return sessionData.UserID
}
