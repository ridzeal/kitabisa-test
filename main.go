package main

import (
	"strconv"
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := strconv.Itoa(3000)
	r := chi.NewRouter()

	// Add useful middleware
	r.Use(middleware.Logger)

	// Router
	r.Get("/", welcome)

	// Start server
	fmt.Println("Server starting on port " + port)
	http.ListenAndServe(":" + port, r)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}