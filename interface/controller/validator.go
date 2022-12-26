package controller

import (
	"context"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
	"log"
)

var validate *validator.Validate

func BindModAndValidate(model interface{}, c Context) error {
	err := c.Bind(&model)
	if err != nil {
		log.Print(err)
		return err
	}

	conform := modifiers.New()
	err = conform.Struct(context.Background(), model)
	if err != nil {
		log.Panic(err)
	}

	validate = validator.New()
	err = validate.Struct(model)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
