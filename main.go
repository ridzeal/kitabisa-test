package main

import (
	"strconv"
	"fmt"
	"net/http"
	"log"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
  "gorm.io/driver/sqlite"

	"kitabisa-test/handler"
	"kitabisa-test/routes"
	"kitabisa-test/config"
	"kitabisa-test/entity"
)

func main() {
	// Prepare configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var conf config.Configuration

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
	app := &application{Config: conf, DB: db}

	// Init Router
	port := strconv.Itoa(conf.Server.Port)
	r := chi.NewRouter()

	// Add useful middleware
	r.Use(middleware.Logger)
	r.Use(app.appCtx)

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

type application struct {
	Config config.Configuration
	DB *gorm.DB
}

func (app *application) appCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), entity.AppCtx, app)
		next.ServeHTTP(w, r.WithContext(ctx))
  })
}