package model

type Airline struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Alias    string `json:"alias"`
	IATA     string `json:"iata" validate:"len=2"`
	ICAO     string `json:"icao" validate:"len=3"`
	Callsign string `json:"callsign"`
	Country  string `json:"country" validate:"required"`
	Active   bool   `json:"active" validate:"required"`
}

func (u *Airline) Validate() error {
	return validate.Struct(u)
}
