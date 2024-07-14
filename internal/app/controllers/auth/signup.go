package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type SignUpInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// SignupController Signup
//
//	@Summary   	  User Signup
//	@Description  create new account
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param        req body auth.SignUpInput true "login body"
//	@Success      200  {object}  responses.DefaultResponse
//	@Failure      400  {object}  responses.DefaultResponse
//	@Failure      500  {string}  responses.DefaultResponse
//	@Router       /auth/signup [post]
func SignupController(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		signUpInput := new(SignUpInput)

		valErr := utils.FillObjectWithInputParams(fiberCtx, signUpInput)

		if valErr != nil {
			return valErr
		}

		user := model.User{}
		user.Email = signUpInput.Email
		user.Login = signUpInput.Login
		user.Password = signUpInput.Password

		err := utils.ValidateStruct(&user)

		if err != nil {
			server.Log.Info("[SignupController]: Validation Error")
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(err)
		}

		existingUserByEmail, _ := server.Store.UserRepository.FindByEmail(ctx, user.Email)

		if existingUserByEmail != nil {
			server.Log.Info("[SignupController]: User exists")
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "user exists",
			})
		}

		existingUserByLogin, _ := server.Store.UserRepository.FindByLogin(ctx, user.Login)

		if existingUserByLogin != nil {
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "login exists",
			})
		}

		user.EncryptedPassword = signUpInput.Password

		if err := server.Store.UserRepository.Create(ctx, user); err != nil {
			server.Log.Error("Error with saving new user")
		}

		server.Log.Info("[SignupController]: User created")
		return fiberCtx.Status(fiber.StatusAlreadyReported).JSON(fiber.Map{
			"message": "user created",
		})
	}
}
