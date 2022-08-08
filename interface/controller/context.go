package controller

import "net/http"

type Context interface {
	JSONPretty(code int, i interface{}, indent string) error
	Param(name string) string
	FormValue(name string) string
	QueryParam(name string) string
	Bind(i interface{}) error
	Request() *http.Request
	Set(key string, val interface{})
}
