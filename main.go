package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Aryan019/RSS-Scrapper/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// the apiConfig here is holding in the connection to the database
type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello from the Go server!")
	godotenv.Load()

	// Reading the port from the environment
	portString := os.Getenv("PORT")
	if portString == "" {
		fmt.Println("PORT not found in the environment variables")
		return
	}

	fmt.Println("Port ->", portString)

	// Create the main router
	router := chi.NewRouter()

	// Set up CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a versioned router
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)

	// Mount the versioned router onto the main router
	router.Mount("/v1", v1Router)

	// Create the server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Print("Server starting on port " + portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
