package main

import (
	"govel/app/exception"
	"govel/bootstrap"
	"govel/config"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	configuration.LoadEnv()

	app := bootstrap.Make(configuration)

	// Start App
	err := app.Listen(":" + configuration.Get("APP_PORT"))
	exception.PanicIfNeeded(err)
}
