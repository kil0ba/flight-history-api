package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	jwt_utils "github.com/kil0ba/flight-history-api/internal/app/utils/jwt-utils"
)

func AuthRequired(jwtManager *jwt_utils.JWTManager) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")

		if authHeader == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		// Check if header is in the format "Bearer <token>"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid authorization header format",
			})
		}

		// Extract token after "Bearer "
		tokenString := authHeader[7:]

		// Verify token
		claims, err := jwtManager.Verify(tokenString)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}
		ctx.Locals("auth", claims)

		// Token is valid, continue to next handler
		return ctx.Next()
	}
}
