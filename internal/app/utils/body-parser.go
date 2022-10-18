package utils

import "github.com/gofiber/fiber/v2"

func FillObjectWithInputParams(ctx *fiber.Ctx, obj interface{}) error {
	if err := ctx.BodyParser(obj); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			})
	}
	return nil
}
