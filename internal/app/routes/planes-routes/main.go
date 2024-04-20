package planes_routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kil0ba/flight-history-api/internal/app/controllers/planes"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
)

func AddAuthRoutes(ctx context.Context, app fiber.Router, server *flighthistoryserver.FlightHistoryServer) {
	server.Log.Info("Setting plane routes")
	app.Post("/getPlanes", planes.GetPlanesController(ctx, server))
}
