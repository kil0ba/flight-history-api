package model

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Uuid              string
	Email             string `validate:"required,email"`
	Password          string `validate:"required,min=6"`
	EncryptedPassword string
	Login             string `validate:"required,min=6"`
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}
