package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	//
	return c.Next()
}
