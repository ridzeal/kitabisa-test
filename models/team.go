package models

import (
	"gorm.io/gorm"
)

// Team table structure
type Team struct {
	gorm.Model
	ID int `gorm:"primary_key"`
	Name string `gorm:"unique"`
	Players []Player `gorm:"ForeignKey:Team;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}