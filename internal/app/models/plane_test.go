package model_test

import (
	"testing"

	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestPlane_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		plane    model.Plane
		expected error
	}{
		{
			name: "Valid Plane",
			plane: model.Plane{
				ID:           1,
				Name:         "Boeing 747",
				IataCode:     "BA",
				IcaoCode:     "BOE",
				Manufacturer: "Boeing",
				Country:      "USA",
			},
			expected: nil,
		},
		{
			name: "Invalid Plane - Missing Name",
			plane: model.Plane{
				ID:           2,
				IataCode:     "AA",
				IcaoCode:     "AIR",
				Manufacturer: "Airbus",
				Country:      "France",
			},
			expected: assert.AnError,
		},
		{
			name: "Invalid Plane - Missing Country",
			plane: model.Plane{
				ID:           3,
				Name:         "Cessna 172",
				IataCode:     "CS",
				IcaoCode:     "CES",
				Manufacturer: "Cessna",
			},
			expected: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.plane.Validate()
			if tc.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
