package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello from the go server ")
	godotenv.Load()

	// Reading in the port from the file we have
	portString := os.Getenv("PORT")

	if portString == "" {
		fmt.Println("PORT not found in the environment variables")
		return
	}

	fmt.Println("Port -> ", portString)

	// Creating in a new router
	router := chi.NewRouter()

	// Setting up the router file including cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)

	// Mounting in tge router
	v1Router.Mount("/v1", router)

	// Creating in the server
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
