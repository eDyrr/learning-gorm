package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"type:varchar(50);uniqueIndex"`
	LastName  string    `gorm:"type:varchar(50);uniqueIndex"`
	Email     string    `gorm:"type:varchar(50);not null"`
	Country   string    `gorm:"type:varchar(50);not null"`
	Role      string    `gorm:"type:varchar(50);not null"`
	Age       int       `gorm:"not null;size:3"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
