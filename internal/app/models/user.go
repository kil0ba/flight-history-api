package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/kil0ba/flight-history-api/internal/app/models/validators"
)

type User struct {
	Uuid              string
	Email             string `validate:"required,email"`
	Password          string `validate:"required,min=6"`
	EncryptedPassword string
	Login             string `validate:"required,min=6,username"`
}

func (u *User) Validate() error {
	validate := validator.New()
	// Register a custom validation function for the "usernameformat" tag
	err := validate.RegisterValidation("username", validators.Username)
	if err != nil {
		return err
	}

	err = validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}
