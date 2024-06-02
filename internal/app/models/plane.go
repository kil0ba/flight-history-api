package model

type Plane struct {
	ID           int
	Name         string `validate:"required"`
	IataCode     string
	IcaoCode     string
	Manufacturer string
	Country      string `validate:"required"`
}

func (u *Plane) Validate() error {
	return validate.Struct(u)
}
