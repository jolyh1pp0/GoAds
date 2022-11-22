package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type Context interface {
	JSONPretty(code int, i interface{}, indent string) error
	Param(name string) string
	FormValue(name string) string
	QueryParam(name string) string
	Bind(i interface{}) error
	Request() *http.Request
	Response() *echo.Response
	Set(key string, val interface{})
	Get(key string) interface{}
}
