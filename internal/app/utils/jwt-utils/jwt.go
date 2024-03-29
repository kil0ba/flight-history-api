package jwt_utils

import (
	"fmt"
	"time"

	model "github.com/kil0ba/flight-history-api/internal/app/models"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

type Claims struct {
	ExpiresAt int64
	Id        string
	Email     string
	Login     string
}

func (c Claims) Valid() error {
	return nil
}

func (manager *JWTManager) CreateToken(user *model.User) (string, error) {
	claims := Claims{
		ExpiresAt: time.Now().Add(manager.TokenDuration).Unix(),
		Id:        user.Uuid,
		Email:     user.Email,
		Login:     user.Login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.SecretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*Claims, error) {
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

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid Token")
	}

	return claims, nil
}
