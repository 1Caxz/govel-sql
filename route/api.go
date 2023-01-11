package route

import (
	"govel/app/http/controller"
	"govel/app/repository"
	"govel/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Doc route rules https://docs.gofiber.io/
func APIRoute(route fiber.Router, database *gorm.DB) {
	// Setup Repository
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)
	userController.Route(route)
}
