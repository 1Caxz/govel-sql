package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Middleware for all request
func AppMiddleware(c *fiber.Ctx) error {
	//
	return c.Next()
}
