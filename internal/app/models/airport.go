package model

type Airport struct {
	ID        int
	Name      string  `validate:"required"`
	Code      string  `validate:"required"`
	City      string  `validate:"required"`
	Country   string  `validate:"required"`
	Latitude  float32 `validate:"required"`
	Longitude float32 `validate:"required"`
	Timezone  string  `validate:"required"`
}

func (u *Airport) Validate() error {
	return validate.Struct(u)
}
