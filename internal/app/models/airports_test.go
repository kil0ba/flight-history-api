package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirport_Validate(t *testing.T) {
	tests := []struct {
		name    string
		airport Airport
		wantErr bool
	}{
		{
			name: "Valid Airport",
			airport: Airport{
				ID:        1,
				Name:      "John F. Kennedy International Airport",
				Code:      "JFK",
				City:      "New York",
				Country:   "USA",
				Latitude:  40.6413,
				Longitude: -73.7781,
			},
			wantErr: false,
		},
		{
			name: "Missing Name",
			airport: Airport{
				ID:        2,
				Code:      "LAX",
				City:      "Los Angeles",
				Country:   "USA",
				Latitude:  33.9416,
				Longitude: -118.4085,
			},
			wantErr: true,
		},
		{
			name: "Missing Code",
			airport: Airport{
				ID:        3,
				Name:      "Heathrow Airport",
				City:      "London",
				Country:   "UK",
				Latitude:  51.4700,
				Longitude: -0.4543,
			},
			wantErr: true,
		},
		{
			name: "Missing City",
			airport: Airport{
				ID:        4,
				Name:      "Charles de Gaulle Airport",
				Code:      "CDG",
				Country:   "France",
				Latitude:  49.0097,
				Longitude: 2.5479,
			},
			wantErr: true,
		},
		{
			name: "Missing Country",
			airport: Airport{
				ID:        5,
				Name:      "Tokyo Haneda Airport",
				Code:      "HND",
				City:      "Tokyo",
				Latitude:  35.5533,
				Longitude: 139.7811,
			},
			wantErr: true,
		},
		{
			name: "Missing Latitude",
			airport: Airport{
				ID:        6,
				Name:      "Dubai International Airport",
				Code:      "DXB",
				City:      "Dubai",
				Country:   "UAE",
				Longitude: 55.3644,
			},
			wantErr: true,
		},
		{
			name: "Missing Longitude",
			airport: Airport{
				ID:       7,
				Name:     "Singapore Changi Airport",
				Code:     "SIN",
				City:     "Singapore",
				Country:  "Singapore",
				Latitude: 1.3644,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.airport.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
