package domain

import (
	"github.com/labstack/echo"
	"net/http"
)

const (
	ErrAdvertisementAlreadyWithTitle = "pq: duplicate key value violates unique constraint \"advertisements_title_key\""
	ErrUserAlreadyWithEmail          = "pq: duplicate key value violates unique constraint \"users_email_key\""
	ErrUserAlreadyWithPhone          = "pq: duplicate key value violates unique constraint \"users_phone_key\""
)

var ErrAdvertisementInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "Internal server error.")
var ErrAdvertisementTitleAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Title already exists.")

var ErrUserInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "Internal server error.")
var ErrUserEmailAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Email already exists.")
var ErrUserPhoneAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Phone already exists.")
