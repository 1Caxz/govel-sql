package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func WebRoute(route fiber.Router, database *gorm.DB) {
	route.Static("/", "./public").Name("root")
}
