package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	SocialId        string `gorm:"type:varchar(255);unique;default:null"`
	Email           string `gorm:"type:varchar(255);unique;not null"`
	Password        string `gorm:"type:varchar(255);default:null"`
	EmailVerifiedAt *time.Time
	Nick            string `gorm:"type:varchar(50);unique;not null"`
	Name            string `gorm:"type:varchar(255);index:,class:FULLTEXT;not null"`
	Pic             string `gorm:"type:varchar(255);not null;default:/assets/static/user.png"`
	Location        string `gorm:"type:varchar(255);default:Indonesia"`
	Desc            string `gorm:"type:varchar(255);default:null"`
	Role            int    `gorm:"type:tinyint(2);default:1"`
	Status          int    `gorm:"type:tinyint(2);default:0"`
	ApiToken        string `gorm:"type:varchar(80);default:null"`
	RememberToken   string `gorm:"type:varchar(100);default:null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
