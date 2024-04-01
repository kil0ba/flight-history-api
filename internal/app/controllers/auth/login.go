package auth

import (
	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginController(server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		loginInput := new(LoginInput)

		validateErrs := utils.FillObjectWithInputParams(ctx, loginInput)

		if validateErrs != nil {
			return validateErrs
		}

		existingUser, err := server.Store.UserRepository.FindByLogin(loginInput.Login)

		if err != nil {
			server.Log.Debug("[LoginController] Cannot find an user")
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		if existingUser == nil {
			server.Log.Debug("[LoginController]: User not found")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user doesn't exist",
			})
		}

		server.Log.Debug("[LoginController]: User founded")

		if loginInput.Password != existingUser.EncryptedPassword {
			server.Log.Debug("[LoginController]: Password is incorrect")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user doesn't exist",
			})
		}

		token, err := server.JwtManager.CreateToken(existingUser)

		if err != nil {
			server.Log.Debug("[LoginController] cannot create token")
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	}
}
