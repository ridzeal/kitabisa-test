package main

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/spf13/viper"

	"kitabisa-test/entity"
	"kitabisa-test/models"
)

func main() {
	// Prepare configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var conf entity.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// Prepare DB Connection
	db, err := gorm.Open(sqlite.Open(conf.Database.Name + ".db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
	}
	
	db.AutoMigrate(&models.Team{}, &models.Player{})
}