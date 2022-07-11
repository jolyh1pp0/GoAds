package domain

import (
	"github.com/labstack/echo"
	"net/http"
)

const ErrDBAlreadyWithTitle = "pq: duplicate key value violates unique constraint \"advertisements_title_key\""

var ErrAdvertisementInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "Internal server error.")
var ErrAdvertisementTitleAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "Status 400 Bad Request. Title already exists.")
