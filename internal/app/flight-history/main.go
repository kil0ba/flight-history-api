package flighthistory

import (
	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server"
	"github.com/kil0ba/flight-history-api/internal/app/routes"
)

func Start(server *flighthistoryserver.FlightHistoryServer) *fiber.App {
	server.Log.Info("Starting server...")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.AddRoutes(app, server)

	app.Listen(server.Config.BindAddr)

	return app
}
