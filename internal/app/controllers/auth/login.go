package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginController(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		loginInput := new(LoginRequest)

		validateErrs := utils.FillObjectWithInputParams(fiberCtx, loginInput)

		if validateErrs != nil {
			return validateErrs
		}

		existingUser, err := server.Store.UserRepository.FindByLogin(ctx, loginInput.Login)

		if err != nil {
			server.Log.Debug("[LoginController] Cannot find an user")
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "can't find user",
			})
		}

		if existingUser == nil {
			server.Log.Debug("[LoginController]: User not found")
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user doesn't exist",
			})
		}

		server.Log.Debug("[LoginController]: User founded")

		if loginInput.Password != existingUser.EncryptedPassword {
			server.Log.Debug("[LoginController]: Password is incorrect")
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user doesn't exist",
			})
		}

		token, err := server.JwtManager.CreateToken(existingUser)

		if err != nil {
			server.Log.Debug("[LoginController] cannot create token")
			return fiberCtx.SendStatus(fiber.StatusInternalServerError)
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	}
}
