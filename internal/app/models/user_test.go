package model_test

import (
	"testing"

	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func getValidUser() *model.User {
	return &model.User{
		Uuid:     "valid-uuid",
		Email:    "test@example.com",
		Password: "securepassword123",
		Login:    "valid_user",
	}
}

func TestUser_Validate_ValidUser(t *testing.T) {
	user := getValidUser()

	err := user.Validate()
	assert.Nil(t, err, "Valid user should not produce validation error, err: ", err)
}

func TestUser_Validate_MissingEmail(t *testing.T) {
	user := &model.User{
		Uuid:     "valid-uuid",
		Password: "securepassword123",
		Login:    "valid_user",
	}

	err := user.Validate()
	assert.NotNil(t, err, "Missing email should produce validation error")
	assert.ErrorContains(t, err, "Email", "Error message should mention missing email")
}

func TestUser_Validate_InvalidEmail(t *testing.T) {
	user := &model.User{
		Uuid:     "valid-uuid",
		Email:    "invalidemail",
		Password: "securepassword123",
		Login:    "valid_user",
	}

	err := user.Validate()
	assert.NotNil(t, err, "Invalid email should produce validation error")
	assert.ErrorContains(t, err, "Email", "Error message should mention invalid email")
}

func TestUser_Validate_ShortPassword(t *testing.T) {
	user := &model.User{
		Uuid:     "valid-uuid",
		Email:    "test@example.com",
		Password: "short",
		Login:    "valid_user",
	}

	err := user.Validate()
	assert.NotNil(t, err, "Short password should produce validation error")
	assert.ErrorContains(t, err, "Password", "Error message should mention password length")
}

func TestUser_Validate_InvalidLogin(t *testing.T) {
	user := &model.User{
		Uuid:     "valid-uuid",
		Email:    "test@example.com",
		Password: "securepassword123",
		Login:    "invalid_login!",
	}

	err := user.Validate()
	assert.NotNil(t, err, "Invalid login should produce validation error")
	assert.ErrorContains(t, err, "Login", "Error message should mention invalid login format")

	user = &model.User{
		Uuid:     "valid-uuid",
		Email:    "test@example.com",
		Password: "securepassword123",
		Login:    "1invalid_login",
	}
	err = user.Validate()
	assert.NotNil(t, err, "Invalid login should produce validation error")
	assert.ErrorContains(t, err, "Login", "Error message should mention invalid login format")
}
