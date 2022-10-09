package flighthistory

import (
	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server"
)

func Start(server *flighthistoryserver.FlightHistoryServer) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(server.Config.BindAddr)

	return app
}
