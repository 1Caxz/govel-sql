package exception

import (
	"govel/app/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(400).JSON(model.WebResponse{
			Code:    400,
			Message: "BAD_REQUEST",
			Data:    err.Error(),
		})
	}

	return ctx.Status(500).JSON(model.WebResponse{
		Code:    500,
		Message: "INTERNAL_SERVER_ERROR",
		Data:    err.Error(),
	})
}
