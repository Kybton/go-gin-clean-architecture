package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}
