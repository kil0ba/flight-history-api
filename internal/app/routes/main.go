package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	authroutes "github.com/kil0ba/flight-history-api/internal/app/routes/auth-routes"
	planesroutes "github.com/kil0ba/flight-history-api/internal/app/routes/planes-routes"
)

func AddRoutes(ctx context.Context, app *fiber.App, server *flighthistoryserver.FlightHistoryServer) {
	auth := app.Group("/auth")
	planes := app.Group("/planes")

	authroutes.AddAuthRoutes(ctx, auth, server)
	planesroutes.AddAuthRoutes(ctx, planes, server)
}
