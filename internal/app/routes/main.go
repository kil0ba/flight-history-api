package routes

import (
	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server"
	authroutes "github.com/kil0ba/flight-history-api/internal/app/routes/auth-routes"
)

func AddRoutes(app *fiber.App, server *flighthistoryserver.FlightHistoryServer) {
	auth := app.Group("/auth")

	authroutes.AddAuthRoutes(auth, server)
}
