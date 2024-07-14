package planes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type GetPlanesRequest struct {
	Count int `json:"count"`
	Page  int `json:"page"`
}

type GetPlanesResponse struct {
	Planes []model.Plane `json:"airports"`
}

// GetPlanesController Get Planes
//
//	@Summary   	  Get planes
//	@Tags         planes
//	@Accept       json
//	@Produce      json
//	@Param        req body planes.GetPlanesRequest true "airports body"
//	@Success      200  {object}  planes.GetPlanesResponse
//	@Failure      400  {object}  responses.DefaultResponse
//	@Failure      500  {string}  responses.DefaultResponse
//	@Router       /planes/getPlanes [post]
func GetPlanesController(ctx context.Context, server *flighthistoryserver.FlightHistoryServer) func(*fiber.Ctx) error {
	return func(fiberCtx *fiber.Ctx) error {
		req := new(GetPlanesRequest)

		validateErrs := utils.FillObjectWithInputParams(fiberCtx, req)

		if validateErrs != nil {
			return validateErrs
		}

		if req.Count <= 0 {
			req.Count = 10
		}

		if req.Page <= 0 {
			return fiber.ErrBadRequest
		}

		planes, err := server.Store.PlaneRepository.GetList(ctx, req.Count, req.Page)

		if err != nil {
			server.Log.Debug("[GetPlanesController] Cannot fetch planes", "err", err)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
			"planes": planes,
		})
	}
}
