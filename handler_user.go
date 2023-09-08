package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julianinsua/RSSAgregator/internal/database"
)

func (ap *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 422, fmt.Sprintf("Can't process request: %v", err))
	}

	user, err := ap.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Internal Error %v", err))
	}

	resUsr := User{}
	resUsr.FromDB(user)

	respondWithJson(w, 201, resUsr)
}

func (ap *apiConfig) handleGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	resUser := User{}
	resUser.FromDB(user)
	respondWithJson(w, 200, resUser)
}
