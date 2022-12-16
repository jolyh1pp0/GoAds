package controller

import "C"
import (
	"GoAds/config"
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/interfactor"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/kataras/jwt"
	"github.com/twinj/uuid"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
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
	SessionUUID string `json:"session_uuid"`
	AccessUUID  string `json:"access_uuid"`
	RefreshUUID string `json:"refresh_uuid"`
}

type AuthorizationController interface {
	CreateUser(c Context) error
	Login(c Context) error
	Refresh(c Context) error
	Logout(c Context) error
	createSession(userID, accessTokenUUID, refreshTokenUUID string) (string, error)
	updateSession(currentRefreshTokenUUID, token, refreshToken string) error
}

func NewAuthorizationController(ai interfactor.AuthorizationInterfactor) AuthorizationController {
	return &authorizationController{ai}
}

func ValidatePassword(password string) ([]byte, error) {
	const minEntropyBits = 60
	err := passwordvalidator.Validate(password, minEntropyBits)
	if err != nil {
		log.Print(err)
		return nil, domain.ErrInsecurePassword
	}

	bPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return bPass, nil
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

	password, err := ValidatePassword(user.Password)
	if err != nil {
		log.Print(err)
		return err
	}
	user.Password = string(password)

	u, err := ac.authorizationInterfactor.Create(&user)
	if !errors.Is(err, nil) {
		return err
	}

	var userRole model.UserRole
	userRole.RoleID = 2
	userRole.UserID = u.ID

	err = ac.authorizationInterfactor.CreateUserToRole(userRole)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSONPretty(http.StatusCreated, u, "  ")
}

func GenerateJWT(userID string, userRoles []int, sessionUUID, accessUUID, refreshUUID string) (string, string, error) {
	claims := tokenClaims{
		UserID:      userID,
		UserRoles:   userRoles,
		SessionUUID: sessionUUID,
		AccessUUID:  accessUUID,
		RefreshUUID: refreshUUID,
	}

	token, err := jwt.Sign(jwt.HS256, key, claims, jwt.MaxAge(config.C.JWT.AccessTokenLifespan*time.Minute))
	if err != nil {
		log.Println(err)
		return "", "", domain.ErrInvalidAccessToken
	}

	refreshToken, err := jwt.Sign(jwt.HS256, refreshKey, claims, jwt.MaxAge(config.C.JWT.RefreshTokenLifespan*time.Minute))
	if err != nil {
		log.Println(err)
		return "", "", domain.ErrInvalidAccessToken
	}

	return string(token), string(refreshToken), nil
}

func ParseToken(accessToken string) (*tokenClaims, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, key, []byte(accessToken))
	if err != nil {
		log.Println(err)
		return nil, domain.ErrInvalidAccessToken
	}

	claims := &tokenClaims{}
	err = verifiedToken.Claims(claims)
	if err != nil {
		log.Println(err)
		return nil, domain.ErrInvalidAccessToken
	}

	return claims, nil
}

func (ac *authorizationController) Login(c Context) error {
	var user model.User

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := make(map[string]string)
	json.Unmarshal(bodyBytes, &bodyString)

	user.Email = bodyString["email"]
	user.Password = bodyString["password"]

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

	accessUUID := uuid.NewV4().String()
	refreshUUID := uuid.NewV4().String()

	sessionUUID, err := ac.createSession(userID, accessUUID, refreshUUID)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidSession
	}

	token, refreshToken, err := GenerateJWT(userID, userRoles, sessionUUID, accessUUID, refreshUUID)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidAccessToken
	}

	tokens := map[string]string{
		"access_token":  token,
		"refresh_token": refreshToken,
	}

	return c.JSONPretty(http.StatusOK, tokens, "")
}

func (ac *authorizationController) createSession(userID, accessTokenUUID, refreshTokenUUID string) (string, error) {
	var session model.Session
	session.UserID = userID
	session.AccessTokenUUID, session.RefreshTokenUUID = accessTokenUUID, refreshTokenUUID

	now := time.Now()
	session.RefreshTokenExpiresAt = now.Add(time.Hour * 24 * 2)
	session.ExpiresAt = now.Add(time.Hour * 24 * 30)

	sessionUU, err := ac.authorizationInterfactor.CreateSession(&session)
	if err != nil {
		log.Print(err)
	}

	return sessionUU.ID, nil
}

func (ac *authorizationController) updateSession(sessionUUID, accessUUID, refreshUUID string) error {
	var session model.Session

	session.AccessTokenUUID, session.RefreshTokenUUID = accessUUID, refreshUUID
	now := time.Now()
	session.RefreshTokenExpiresAt = now.Add(time.Hour * 24 * 2)

	err := ac.authorizationInterfactor.UpdateSession(sessionUUID, &session)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidSession
	}

	return nil
}

func (ac *authorizationController) Refresh(c Context) error {
	var session model.Session

	type Token struct {
		RefreshToken string `json:"refresh_token"`
	}

	var token Token
	err := c.Bind(&token)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidRefreshToken
	}

	claims, err := ParseToken(token.RefreshToken)
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidRefreshToken
	}

	session.ID, session.RefreshTokenUUID, session.AccessTokenUUID = claims.SessionUUID, claims.RefreshUUID, claims.AccessUUID
	userID, userRoles := claims.UserID, claims.UserRoles

	expiresAt, err := ac.authorizationInterfactor.GetSessionExpiration(claims.SessionUUID)
	if err != nil {
		log.Print(err)
	}
	if time.Now().After(expiresAt) {
		log.Println(domain.ErrSessionExpired)
		return domain.ErrSessionExpired
	}

	refreshTokenUUIDFromTable, err := ac.authorizationInterfactor.GetRefreshTokenUUIDFromTable(claims.SessionUUID)
	if err != nil {
		log.Print(err)
	}

	if refreshTokenUUIDFromTable != session.RefreshTokenUUID {
		log.Println(err)
		return domain.ErrInvalidRefreshToken
	}

	accessUUID := uuid.NewV4().String()
	refreshUUID := uuid.NewV4().String()

	accessToken, refreshToken, err := GenerateJWT(userID, userRoles, session.ID, accessUUID, refreshUUID)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidAccessToken
	}

	tokens := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	err = ac.updateSession(session.ID, accessUUID, refreshUUID)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidSession
	}

	return c.JSONPretty(http.StatusOK, tokens, "")
}

func (ac *authorizationController) Logout(c Context) error {
	type Token struct {
		AccessToken string `json:"access_token"`
	}

	var token Token
	err := c.Bind(&token)
	if err != nil {
		log.Print(err)
		return domain.ErrInvalidRefreshToken
	}

	claims, err := ParseToken(token.AccessToken)
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidRefreshToken
	}

	err = ac.authorizationInterfactor.Logout(claims.SessionUUID)
	if err != nil {
		log.Print(err)
	}

	return c.JSONPretty(http.StatusOK, "Logout. Session closed.", "")
}

func getUserID(c Context) string {
	rawSessionData := c.Get(domain.SessionDataKey)
	sessionData, _ := rawSessionData.(domain.SessionData)

	return sessionData.UserID
}

func getUserRole(c Context) []int {
	rawSessionData := c.Get(domain.SessionDataKey)
	sessionData, _ := rawSessionData.(domain.SessionData)

	return sessionData.UserRoles
}
