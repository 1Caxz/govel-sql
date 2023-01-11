package main

import (
	"govel/config"
	"govel/database/migration"
	"govel/database/seeder"
	"os"
)

func main() {
	// Setup Configuration
	appConfig := config.New()
	appConfig.LoadEnv()

	// Setup Database
	database := config.NewDatabase(appConfig)

	if len(os.Args) > 1 {
		if os.Args[1] == "start" {
			migration.Migrator(database)
		} else if os.Args[1] == "seed" {
			seeder.Seeder(database)
		}
	}
}
