package test

import (
	"govel/bootstrap"
	"govel/config"

	"github.com/gofiber/fiber/v2"
)

func CreateApplication() (app *fiber.App) {
	// Setup Configuration
	configuration := config.New()
	configuration.LoadEnv("../.env.test")

	return bootstrap.Make(configuration)
}

var app = CreateApplication()
