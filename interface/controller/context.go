package controller

type Context interface {
	JSON(code int, i interface{}) error
	JSONPretty(code int, i interface{}, indent string) error
	Param(name string) string
	Bind(i interface{}) error
}
