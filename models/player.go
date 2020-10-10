package models

import (
	"gorm.io/gorm"
)

// Player table structure
type Player struct {
	gorm.Model
	ID int `gorm:"primary_key"`
	Name string `gorm:"unique"`
	Team int
}