package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
	"github.com/kil0ba/flight-history-api/internal/app/utils/responses"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// LoginController Login
//
//	@Summary   	  User login
//	@Description  get JWT token
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param        req body auth.LoginRequest true "login body"
//	@Success      200  {object}  auth.LoginResponse
//	@Failure      400  {object}  responses.DefaultResponse
//	@Failure      500  {string}  responses.DefaultResponse
//	@Router       /auth/login [post]
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
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(responses.DefaultResponse{
				Message: "Cannot find an user",
			})
		}

		if existingUser == nil {
			server.Log.Debug("[LoginController]: User not found")
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(responses.DefaultResponse{
				Message: "User not found",
			})
		}

		server.Log.Debug("[LoginController]: User founded")

		if loginInput.Password != existingUser.EncryptedPassword {
			server.Log.Debug("[LoginController]: Password is incorrect")
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(responses.DefaultResponse{
				Message: "Password is incorrect",
			})
		}

		token, err := server.JwtManager.CreateToken(existingUser)

		if err != nil {
			server.Log.Debug("[LoginController] cannot create token")
			return fiberCtx.SendStatus(fiber.StatusInternalServerError)
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(LoginResponse{
			Token: token,
		})
	}
}
