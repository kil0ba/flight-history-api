package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/kil0ba/flight-history-api/docs"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	airport_routes "github.com/kil0ba/flight-history-api/internal/app/routes/airport-routes"
	authroutes "github.com/kil0ba/flight-history-api/internal/app/routes/auth-routes"
	planesroutes "github.com/kil0ba/flight-history-api/internal/app/routes/planes-routes"
)

func AddRoutes(ctx context.Context, app *fiber.App, server *flighthistoryserver.FlightHistoryServer) {
	auth := app.Group("/auth")
	planes := app.Group("/planes")
	airports := app.Group("/airports")
	app.Get("/swagger/*", swagger.HandlerDefault)

	authroutes.AddAuthRoutes(ctx, auth, server)
	planesroutes.AddPlanesRoutes(ctx, planes, server)
	airport_routes.AddAirportRoutes(ctx, airports, server)
}
