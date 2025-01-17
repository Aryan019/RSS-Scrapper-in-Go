package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Aryan019/RSS-Scrapper/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// apiConfig struct holds the database connection, which is used for interacting with the database.
type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Loading environment variables from the .env file
	fmt.Println("Hello from the Go server!")
	godotenv.Load()

	// Reading the port from the environment variables
	portString := os.Getenv("PORT")
	if portString == "" {
		// If PORT is not found, we log and stop further execution
		fmt.Println("PORT not found in the environment variables")
		return
	}

	// Reading the DB connection URL from the environment
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		// If DB_URL is not found, we log and stop further execution
		fmt.Println("DB Url not found in the environment variables")
		return
	}

	// Establishing a connection to the PostgreSQL database
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		// If connection fails, log the error and stop the program
		log.Fatal("Error connecting to the database", err)
	}

	// Confirmation that the port is correctly read
	fmt.Println("Port ->", portString)

	// Initializing database queries with the connection
	queries := database.New(conn)

	// Setting up the apiConfig struct to hold the queries and DB connection
	apiConf := apiConfig{
		DB: queries,
	}

	// Create the main router using Chi router
	router := chi.NewRouter()

	// Set up CORS middleware to handle cross-origin requests from specified origins
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},                   // Allow all origins for demo purposes (adjust in production)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow specific HTTP methods
		AllowedHeaders:   []string{"*"},                                       // Allow all headers
		ExposedHeaders:   []string{"Link"},                                    // Expose Link header to client
		AllowCredentials: false,                                               // Disable credentials sharing
		MaxAge:           300,                                                 // Cache preflight responses for 5 minutes
	}))

	// Create a versioned router for API v1
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)         // Health check endpoint
	v1Router.Get("/err", handleErr)                    // Sample error endpoint
	v1Router.Post("/users", apiConf.handlerCreateUser) // Create user endpoint

	// Mount the versioned router to the main router under the /v1 path
	router.Mount("/v1", v1Router)

	// Create the HTTP server and set it to listen on the specified port
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString, // Port from environment variable
	}

	// Log the server starting message
	log.Print("Server starting on port " + portString)

	// Start the server and handle any errors that occur
	err = server.ListenAndServe()
	if err != nil {
		// If server fails to start, log the error and terminate the program
		log.Fatal(err)
	}
}
