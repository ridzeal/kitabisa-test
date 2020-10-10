package models

// Player table structure
type Player struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"unique"`
	Team int
}