package migration

import (
	"govel/app/entity"

	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
