package controller

type Context interface {
	JSONPretty(code int, i interface{}, indent string) error
	Param(name string) string
	Bind(i interface{}) error
}
