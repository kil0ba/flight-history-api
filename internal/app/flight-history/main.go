package flighthistory

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/routes"
)

func Start(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) *fiber.App {
	server.Log.Info("Starting server...")
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.AddRoutes(ctx, app, server)

	app.Listen(server.Config.BindAddr)

	return app
}
