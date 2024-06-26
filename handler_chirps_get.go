package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// 5. Storage / 4. Get
func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	// 5. Storage / 4. Get
	// chirpIDString := r.PathValue("chirpID")
	chirpIDString := chi.URLParam(r, "chirpID")

	chirpID, err := strconv.Atoi(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	dbChirp, err := cfg.DB.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID: dbChirp.ID,
		// AuthorID: dbChirp.AuthorID,
		Body: dbChirp.Body,
	})
}

// 5. Storage / 1. Storage
// This endpoint should return an array of all chirps in the file, ordered by id in ascending order
func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {

	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	// authorID := -1
	// authorIDString := r.URL.Query().Get("author_id")
	// if authorIDString != "" {
	// 	authorID, err = strconv.Atoi(authorIDString)
	// 	if err != nil {
	// 		respondWithError(w, http.StatusBadRequest, "Invalid author ID")
	// 		return
	// 	}
	// }

	sortDirection := "asc"
	sortDirectionParam := r.URL.Query().Get("sort")
	if sortDirectionParam == "desc" {
		sortDirection = "desc"
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {

		// if authorID != -1 && dbChirp.AuthorID != authorID {
		// 	continue
		// }

		chirps = append(chirps, Chirp{
			ID: dbChirp.ID,
			// AuthorID: dbChirp.AuthorID,
			Body: dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		if sortDirection == "desc" {
			return chirps[i].ID > chirps[j].ID
		}
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, chirps)
}
