package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Middleware under web route path only
func WebMiddleware(c *fiber.Ctx) error {
	//
	return c.Next()
}
