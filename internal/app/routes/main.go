package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	authroutes "github.com/kil0ba/flight-history-api/internal/app/routes/auth-routes"
)

func AddRoutes(ctx context.Context, app *fiber.App, server *flighthistoryserver.FlightHistoryServer) {
	auth := app.Group("/auth")

	authroutes.AddAuthRoutes(ctx, auth, server)
}
