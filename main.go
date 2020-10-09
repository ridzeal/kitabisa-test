package main

import (
	"strconv"
	"fmt"
	"net/http"
	"log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"

	"kitabisa-test/handler"
	"kitabisa-test/routes"
	"kitabisa-test/config"
)

func main() {
	// Prepare configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	port := strconv.Itoa(3000)
	r := chi.NewRouter()

	// Add useful middleware
	r.Use(middleware.Logger)

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