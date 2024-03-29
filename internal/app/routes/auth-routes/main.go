package authroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kil0ba/flight-history-api/internal/app/controllers/auth"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/middleware"
)

func AddAuthRoutes(app fiber.Router, server *flighthistoryserver.FlightHistoryServer) {
	server.Log.Info("Setting auth routes")
	app.Post("/signup", auth.SignupController(server))
	app.Post("/login", auth.LoginController(server))

	app.Get("/check-auth", middleware.AuthRequired(server.JwtManager), func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authenticated",
		})
	})
}
