package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Middleware under API route path only
func APIMiddleware(c *fiber.Ctx) error {
	//
	return c.Next()
}
