package middleware

import (
	"govel/app/exception"
	"govel/app/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(c *fiber.Ctx) error {
	token := helper.ParseECDSAToken(c.FormValue("token", ""), jwt.SigningMethodES256)
	if !token.Valid {
		exception.PanicResponse("Token invalid.")
	}
	return c.Next()
}
