package main

import (
	"fmt"
	"net/http"

	"github.com/julianinsua/RSSAgregator/internal/auth"
	"github.com/julianinsua/RSSAgregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("unauthorized: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("unauthorized: %v", err))
			return
		}

		handler(w, r, user)
	}
}
