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
var ErrInvalidPassword = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Invalid password")

var ErrInvalidAccessToken = echo.NewHTTPError(http.StatusInternalServerError, "Invalid or expired access token")
var ErrInvalidRefreshToken = echo.NewHTTPError(http.StatusInternalServerError, "Invalid or expired refresh token")
var ErrEmptyJWTKey = errors.New("Empty key is not allowed.")

var ErrForbidden = echo.NewHTTPError(http.StatusForbidden, "Status 403 Forbidden.")

var ErrInvalidSession = echo.NewHTTPError(http.StatusInternalServerError, "Invalid Session.")
var ErrSessionExpired = echo.NewHTTPError(http.StatusInternalServerError, "Session Expired.")

var ErrInvalidPictureExtension = echo.NewHTTPError(http.StatusInternalServerError, "Invalid file extension.")
var PictureLimitReached = echo.NewHTTPError(http.StatusInternalServerError, "Only 10 images allowed to add.")

var ErrInvalidEmail = echo.NewHTTPError(http.StatusInternalServerError, "Invalid email.")
var ErrTokenActive = echo.NewHTTPError(http.StatusInternalServerError, "This email already has an active token.")
var ErrTokenInvalid = echo.NewHTTPError(http.StatusInternalServerError, "Invalid token.")
