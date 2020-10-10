package main

import (
	"strconv"
	"fmt"
	"net/http"
	"log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
  "gorm.io/driver/sqlite"

	"kitabisa-test/handler"
	"kitabisa-test/routes"
	"kitabisa-test/entity"
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

	// Create app context
	app := &entity.Application{Config: conf, DB: db}

	// Init
	port := strconv.Itoa(conf.Server.Port)
	r := chi.NewRouter()

	// Add useful middleware
	r.Use(middleware.Logger)
	r.Use(app.AppCtx)

	// Router
	r.Get("/", welcome)
	r.Post("/bundle", handler.Bundle)
	r.Route("/soccer", routes.SoccerRouter)

	// Start server
	fmt.Println("Server starting on port " + port)
	http.ListenAndServe(":" + port, r)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}