package models

import (
	"time"
)

type Person struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	FirstName string     `gorm:"not null" json:"first_name"`
	LastName  string     `gorm:"not null" json:"last_name"`
	CreatedAt *time.Time `gorm:"default:null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"default:null" json:"deleted_at"`
}
