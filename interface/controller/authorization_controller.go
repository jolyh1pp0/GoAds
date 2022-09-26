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
	createSession(userID string, token string, refreshToken string) error
	checkSession(userID string) (bool, error)
	updateSession(userID, token string, refreshToken string) error
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

	token, err := jwt.Sign(jwt.HS256, key, claims, jwt.MaxAge(30*time.Minute))
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
		return domain.ErrInvalidAccessToken
	}

	tokens := map[string]string{
		"access_token":  token,
		"refresh_token": refreshToken,
	}

	exists, err := ac.checkSession(userID)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidSession
	}

	if exists == false {
		err = ac.createSession(userID, token, refreshToken)
		if err != nil {
			log.Print(err)
			return domain.ErrInvalidSession
		}
	} else if exists == true {
		err = ac.updateSession(userID, token, refreshToken)
		if err != nil {
			log.Print(err)
			return domain.ErrInvalidSession
		}
	}

	return c.JSONPretty(http.StatusOK, tokens, "")
}

func(ac *authorizationController) createSession(userID, accessToken, refreshToken string) error {
	var session model.Session
	session.UserID = userID
	session.AccessToken = accessToken
	session.RefreshToken = refreshToken
	session.AccessTokenUUID, session.RefreshTokenUUID, _ = getUUIDs(accessToken)

	now := time.Now()
	session.RefreshTokenExpiresAt = now.Add(time.Hour*24*2)
	session.ExpiresAt = now.Add(time.Hour*24*30)

	_, err := ac.authorizationInterfactor.CreateSession(&session)
	if err != nil {
		log.Print(err)
	}

	return nil
}

func(ac *authorizationController) checkSession(userID string) (bool, error) {
	session, err := ac.authorizationInterfactor.GetSession(userID)
	if err != nil {
		log.Print(err)
		return false, domain.ErrInvalidSession
	}
	if session == "" {
		return false, nil
	} else {
		return true, nil
	}
}

func(ac *authorizationController) updateSession(userID, accessToken, refreshToken string) error {
	var session model.Session
	session.AccessToken = accessToken
	session.RefreshToken = refreshToken
	session.AccessTokenUUID, session.RefreshTokenUUID, _ = getUUIDs(accessToken)
	now := time.Now()
	session.RefreshTokenExpiresAt = now.Add(time.Hour*24*2)
	session.ExpiresAt = now.Add(time.Hour*24*30)

	err := ac.authorizationInterfactor.UpdateSession(userID, &session)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidSession
	}

	return nil
}

func (ac *authorizationController) Refresh(c Context) error {
	var session model.Session

	err := c.Bind(&session)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidRefreshToken
	}

	_, session.RefreshTokenUUID, err = getUUIDs(session.RefreshToken)
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidRefreshToken
	}

	userID, userRoles, err := ParseToken(session.RefreshToken)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidRefreshToken
	}

	// TODO: check session expiration

	refreshTokenUUID, err := ac.authorizationInterfactor.GetRefreshTokenUUIDFromTable(session.RefreshTokenUUID)
	if err != nil {
		log.Print(err)
	}

	if refreshTokenUUID != session.RefreshTokenUUID {
		log.Println(err)
		return domain.ErrInvalidRefreshToken
	}

	accessToken, refreshToken, err := GenerateJWT(userID, userRoles)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidAccessToken
	}

	tokens := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	err = ac.updateSession(userID, accessToken, refreshToken)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidAccessToken
	}

	return c.JSONPretty(http.StatusOK, tokens, "")
}

func getUserID(c Context) string {
	rawSessionData := c.Get(domain.SessionDataKey)
	sessionData, _ := rawSessionData.(domain.SessionData)

	return sessionData.UserID
}
