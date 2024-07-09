package airports

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type SearchAirportsRequest struct {
	Query string `json:"query"`
	Count int    `json:"count"`
	Page  int    `json:"page"`
}

const searchAirport = "searchAirport: "

func SearchAirportsController(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		searchAirportsRequest := new(SearchAirportsRequest)

		validateErrs := utils.FillObjectWithInputParams(fiberCtx, searchAirportsRequest)
		if validateErrs != nil {
			server.Log.Debug(searchAirport, "validation Errors", validateErrs)
			return validateErrs
		}

		airport, err := server.Store.AirportRepository.Search(ctx, searchAirportsRequest.Query, searchAirportsRequest.Count)

		if err != nil {
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "error searching airports",
			})
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
			"airport": airport,
		})
	}
}
