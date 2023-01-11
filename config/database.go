package config

import (
	"govel/app/exception"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDatabase(appConfig Config) *gorm.DB {
	var database *gorm.DB
	var err error
	if appConfig.Get("DB_CONNECTION") == "mysql" {
		dsn := appConfig.Get("DB_USERNAME") + ":" + appConfig.Get("DB_PASSWORD") + "@tcp(" + appConfig.Get("DB_HOST") + ":" + appConfig.Get("DB_PORT") + ")/" + appConfig.Get("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=" + strings.ReplaceAll(appConfig.Get("DB_TIMEZONE"), "/", "%2F")
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if appConfig.Get("DB_CONNECTION") == "postgres" {
		dsn := "host=" + appConfig.Get("DB_HOST") + " user=" + appConfig.Get("DB_USERNAME") + " password=" + appConfig.Get("DB_PASSWORD") + " dbname=" + appConfig.Get("DB_DATABASE") + " port=" + appConfig.Get("DB_PORT") + " sslmode=disable TimeZone=" + appConfig.Get("DB_TIMEZONE")
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if appConfig.Get("DB_CONNECTION") == "sqlite" {
		database, err = gorm.Open(sqlite.Open(appConfig.Get("DB_DATABASE")), &gorm.Config{})
	} else if appConfig.Get("DB_CONNECTION") == "sqlserver" {
		dsn := "sqlserver://" + appConfig.Get("DB_USERNAME") + ":" + appConfig.Get("DB_PASSWORD") + "@" + appConfig.Get("DB_HOST") + ":" + appConfig.Get("DB_PORT") + "?database=" + appConfig.Get("DB_DATABASE")
		database, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	}
	exception.PanicIfNeeded(err)

	sqlDB, err := database.DB()
	exception.PanicIfNeeded(err)

	// Setup connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return database
}
