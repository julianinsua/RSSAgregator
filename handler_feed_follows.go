package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/julianinsua/RSSAgregator/internal/database"
)

func (cfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feedId"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 422, fmt.Sprintf("Unable to process json body: %v", err))
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	ff := FeedFollow{}
	ff.FromDB(feedFollow)

	respondWithJson(w, 200, ff)
}

func (cfg *apiConfig) handleListFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't get feed follows: %v", err))
		return
	}

	ffs := FeedFollows{}
	ffs.FromDB(feedFollows)
	respondWithJson(w, 200, ffs)
}

func (cfg *apiConfig) handleUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	ffIDStr := chi.URLParam(r, "ffID")
	ffID, err := uuid.Parse(ffIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("unable to parse uuid: %v", err))
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     ffID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("couldn't delete feed follow: %v", err))
		return
	}

	respondWithJson(w, 200, fmt.Sprintf("feed follow %v deleted successfully", ffID))
}
