package bootstrap

import (
	"context"
	"govel/app/http/middleware"
	"govel/config"
	"govel/route"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Make(configuration config.Config) *fiber.App {
	// Setup database
	database := config.NewDatabase(configuration)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())

	// Setup App Middleware
	app.Use(middleware.AppMiddleware)

	// Save databse object to fiber context
	app.Use(func(c *fiber.Ctx) error {
		timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		c.Locals("DB", database.WithContext(timeoutContext))
		return c.Next()
	})

	// Setup Routing
	apiRoute := app.Group("/api", middleware.APIMiddleware)
	route.APIRoute(apiRoute, database)
	webRoute := app.Group("/", middleware.WebMiddleware)
	route.WebRoute(webRoute, database)

	// 404 respond status if route path not found
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return app
}
