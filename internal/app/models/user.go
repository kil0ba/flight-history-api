package model

type User struct {
	Uuid              string
	Email             string `validate:"required,email"`
	Password          string `validate:"required,min=6"`
	EncryptedPassword string
	Login             string `validate:"required,min=6,username"`
}

func (u *User) Validate() error {
	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}
