package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/kil0ba/flight-history-api/internal/app/models/validators"
)

var validate = validator.New()

func init() {
	err := validate.RegisterValidation("username", validators.Username)
	if err != nil {
		panic("failed to register 'username' validation function")
	}
}
