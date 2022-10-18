package jwt_utils

import (
	"fmt"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

func (manager *JWTManager) CreateToken(user *model.User) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(manager.TokenDuration).Unix(),
		Id:        user.Uuid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.SecretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return "", fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.SecretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("invalid Token")
	}

	return claims, nil
}
