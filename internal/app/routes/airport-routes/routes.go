package airport_routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kil0ba/flight-history-api/internal/app/controllers/airports"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
)

func AddAirportRoutes(ctx context.Context, app fiber.Router, server *flighthistoryserver.FlightHistoryServer) {
	server.Log.Info("Setting airport routes")
	app.Post("/searchAirports", airports.SearchAirportsController(ctx, server))
}
