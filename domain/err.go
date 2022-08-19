package domain

import (
	"errors"
	"github.com/labstack/echo"
	"net/http"
)

const (
	ErrAdvertisementAlreadyWithTitle = "pq: duplicate key value violates unique constraint \"advertisements_title_key\""
	ErrUserAlreadyWithEmail          = "pq: duplicate key value violates unique constraint \"users_email_key\""
	ErrUserAlreadyWithPhone          = "pq: duplicate key value violates unique constraint \"users_phone_key\""
)

var ErrAdvertisementInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error.")
var ErrAdvertisementTitleAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Title already exists.")

var ErrUserInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error.")
var ErrUserEmailAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Email already exists.")
var ErrUserPhoneAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Phone already exists.")

var ErrCommentInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error.")

var ErrInvalidEmailAddress = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error. Invalid email address")
var ErrEmailIsNotFound = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error. User with this email doesn't exist")

var ErrInsecurePassword = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error. Insecure password, try including more special characters, using lowercase letters, using uppercase letters or using a longer password")
var ErrInvalidPassword = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error. Invalid password")

var ErrInvalidAccessToken = echo.NewHTTPError(http.StatusInternalServerError, "Status 500 Internal server error. Invalid or expired access token")
var ErrEmptyJWTKey = errors.New("Empty JWT key is not allowed.")

var ErrForbidden = echo.NewHTTPError(http.StatusForbidden, "Status 403 Forbidden.")
