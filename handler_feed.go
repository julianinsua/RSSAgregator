package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julianinsua/RSSAgregator/internal/database"
)

func (cfg *apiConfig) handleAddFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 422, fmt.Sprintf("can't process request: %v", err))
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Unable to save feed: %v", err))
		return
	}

	fd := Feed{}
	fd.FromDB(feed)
	respondWithJson(w, 200, fd)
}

func (cfg *apiConfig) handleGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	dbFeeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Unable to get feeds: %v", err))
	}

	feeds := Feeds{}
	feeds.FromDB(dbFeeds)

	respondWithJson(w, 200, feeds)
}
