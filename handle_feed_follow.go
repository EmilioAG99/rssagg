package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/emilioag99/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}
	follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating feed_follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowtoFeedFollow(follow))
}

func (apiCfg *apiConfig) handleGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	follows, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feed_follow: %v", err))
		return
	}

	list := []FeedFollow{}
	for _, feed := range follows {
		list = append(list, databaseFeedFollowtoFeedFollow(feed))
	}
	respondWithJSON(w, 200, list)
}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId := chi.URLParam(r, "feedFollowID")
	if feedFollowId == "" {
		respondWithError(w, 400, "Error getting Feed id")
		return
	}
	feedFolowIdUUID, err := uuid.Parse(feedFollowId)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feed id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFolowIdUUID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feed_follow: %v", err))
		return
	}

	respondWithJSON(w, 200, fmt.Sprintf("Succesfuly unfollowed %v", feedFollowId))
}
