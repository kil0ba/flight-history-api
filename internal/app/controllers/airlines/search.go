package airlines

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type SearchAirlinesRequest struct {
	Query string `json:"query"`
	Count int    `json:"count"`
	Page  int    `json:"page"`
}

type SearchAirlinesResponse struct {
	Airlines []model.Airline `json:"airlines"`
}

const searchAirline = "searchAirline: "

// SearchAirlinesController Search Airlines
//
//	@Summary   	  Search Airlines
//	@Tags         airlines
//	@Accept       json
//	@Produce      json
//	@Param        req body airlines.SearchAirlinesRequest true "airlines body"
//	@Success      200  {object}  airlines.SearchAirlinesResponse
//	@Failure      400  {object}  responses.DefaultResponse
//	@Failure      500  {string}  responses.DefaultResponse
//	@Router       /airlines/searchAirlines [post]
func SearchAirlinesController(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		searchAirlinesRequest := new(SearchAirlinesRequest)

		validateErrs := utils.FillObjectWithInputParams(fiberCtx, searchAirlinesRequest)
		if validateErrs != nil {
			server.Log.Debug(searchAirline, "validation Errors", validateErrs)
			return validateErrs
		}

		airlines, err := server.Store.AirlineRepository.Search(ctx, searchAirlinesRequest.Query, searchAirlinesRequest.Page, searchAirlinesRequest.Count)

		if err != nil {
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "error searching airports",
			})
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(SearchAirlinesResponse{
			Airlines: airlines,
		})
	}
}
