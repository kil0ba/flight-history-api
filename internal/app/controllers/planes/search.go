package planes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type SearchPlanesRequest struct {
	Query string `json:"query"`
	Count int    `json:"count"`
	Page  int    `json:"page"`
}

type SearchPlanesResponse struct {
	Planes *[]model.Plane `json:"planes"`
}

const searchPlanes = "searchPlanes: "

// SearchPlanesController Search Planes
//
//	@Summary   	  Search Planes
//	@Tags         planes
//	@Accept       json
//	@Produce      json
//	@Param        req body planes.SearchPlanesRequest true "airports body"
//	@Success      200  {object}  planes.SearchPlanesResponse
//	@Failure      400  {object}  responses.DefaultResponse
//	@Failure      500  {string}  responses.DefaultResponse
//	@Router       /planes/searchPlanes [post]
func SearchPlanesController(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		searchPlanesRequest := new(SearchPlanesRequest)

		validateErrs := utils.FillObjectWithInputParams(fiberCtx, searchPlanesRequest)
		if validateErrs != nil {
			server.Log.Debug(searchPlanes, "validation Errors", validateErrs)
			return validateErrs
		}

		planes, err := server.Store.PlaneRepository.Search(ctx, searchPlanesRequest.Query, searchPlanesRequest.Count, searchPlanesRequest.Page)

		if err != nil {
			return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "error searching planes",
			})
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(SearchPlanesResponse{
			Planes: planes,
		})
	}
}
