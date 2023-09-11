package main

import (
	"fmt"
	"net/http"

	"github.com/julianinsua/RSSAgregator/internal/database"
)

func (cfg *apiConfig) handleGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	// type Parameters struct {
	// 	Limit int `json:"limit"`
	// }
	//
	// params := Parameters{}
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&params)
	// if err != nil {
	// 	respondWithError(w, 422, fmt.Sprintf("unable to process json payload: %v", err))
	// }

	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(10),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("unable to get posts: %v", err))
		return
	}
	ps := Posts{}
	ps.FromDB(posts)

	respondWithJson(w, 200, ps)

}
