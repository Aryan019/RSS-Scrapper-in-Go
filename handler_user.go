package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Aryan019/RSS-Scrapper/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Reached in here")
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	log.Print("After the decoder line")

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	log.Print("the name value is ", params.Name)

	user, err := apiCfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		log.Printf("Error creating user: %v", err) // Log the error to help debug
		respondWithError(w, 400, "Couldn't create user")
		return
	}

	log.Print("He is reaching in the last line ")
	respondWithJSON(w, 200, user)
}
