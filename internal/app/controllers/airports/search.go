package airports

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type SearchAirportsRequest struct {
	Query string `json:"query"`
	Count int    `json:"count"`
	Page  int    `json:"page"`
}

type SearchAirportsResponse struct {
	Airports []model.Airport `json:"airports"`
}

const searchAirport = "searchAirport: "

// SearchAirportsController Search Airports
//
//	@Summary   	  Search Airports
//	@Tags         airports
//	@Accept       json
//	@Produce      json
//	@Param        req body airports.SearchAirportsRequest true "airports body"
//	@Success      200  {object}  airports.SearchAirportsResponse
//	@Failure      400  {object}  responses.DefaultResponse
//	@Failure      500  {string}  responses.DefaultResponse
//	@Router       /airports/searchAirports [post]
func SearchAirportsController(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		searchAirportsRequest := new(SearchAirportsRequest)

		validateErrs := utils.FillObjectWithInputParams(fiberCtx, searchAirportsRequest)
		if validateErrs != nil {
			server.Log.Debug(searchAirport, "validation Errors", validateErrs)
			return validateErrs
		}

		airport, err := server.Store.AirportRepository.Search(ctx, searchAirportsRequest.Query, searchAirportsRequest.Page, searchAirportsRequest.Count)

		if err != nil {
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "error searching airports",
			})
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(SearchAirportsResponse{
			Airports: airport,
		})
	}
}
