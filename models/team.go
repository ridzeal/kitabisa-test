package models

// Team table structure
type Team struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"unique"`
	Players []Player `gorm:"ForeignKey:Team;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}