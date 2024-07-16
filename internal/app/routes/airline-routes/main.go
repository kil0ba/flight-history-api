package airline_routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kil0ba/flight-history-api/internal/app/controllers/airlines"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
)

func AddAirlinesRoutes(ctx context.Context, app fiber.Router, server *flighthistoryserver.FlightHistoryServer) {
	server.Log.Info("Setting airlines routes")
	app.Post("/searchAirlines", airlines.SearchAirlinesController(ctx, server))
}
