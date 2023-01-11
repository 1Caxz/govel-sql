package seeder

import (
	"govel/app/entity"
	"govel/app/exception"

	"github.com/bxcodec/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	for i := 0; i < 30; i++ {
		hashed, err := bcrypt.GenerateFromPassword([]byte(faker.Word()), bcrypt.DefaultCost)
		exception.PanicIfNeeded(err)
		db.Create(&entity.User{
			Email:    faker.Word() + "@gmail.com",
			Password: string(hashed),
			Name:     faker.Word(),
			Nick:     faker.Word(),
			Role:     1,
			Status:   1,
		})
	}
}
