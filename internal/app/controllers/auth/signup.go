package auth

import (
	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type SignUpInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func SignupController(server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		signUpInput := new(SignUpInput)

		if err := ctx.BodyParser(signUpInput); err != nil {
			server.Log.Error("[SignupController]: Error with parsing body")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		user := model.User{}
		user.Email = signUpInput.Email
		user.Login = signUpInput.Login
		user.Password = signUpInput.Password

		err := utils.ValidateStruct(user)

		if err != nil {
			server.Log.Info("[SignupController]: Validation Error")
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		existingUser, _ := server.Store.UserRepository.FindByEmail(user.Email)

		if existingUser != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user exists",
			})
		}

		user.EncryptedPassword = signUpInput.Password

		if err := server.Store.UserRepository.Create(user); err != nil {
			server.Log.Error("Error with saving new user")
		}

		return ctx.Status(fiber.StatusAlreadyReported).JSON(fiber.Map{
			"message": "user created",
		})
	}
}
