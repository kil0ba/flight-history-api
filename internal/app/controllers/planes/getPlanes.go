package planes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/utils"
)

type GetPlanesRequest struct {
	Count int `json:"count"`
	Page  int `json:"page"`
}

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
