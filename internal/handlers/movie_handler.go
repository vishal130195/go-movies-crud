// Package handlers provides HTTP request handlers for the movies CRUD API.
// This package contains all the HTTP endpoint handlers that process incoming requests,
// interact with the storage layer, and return appropriate responses.
package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/vishal130195/go-movies-crud/internal/models"
	"github.com/vishal130195/go-movies-crud/internal/storage"
	"log"
	"net/http"
)

// MovieHandler handles all HTTP requests related to movie operations.
// It acts as a bridge between the HTTP layer and the storage layer,
// processing requests and managing responses for movie CRUD operations.
type MovieHandler struct {
	// store is the storage interface for persisting and retrieving movie data
	store storage.MovieStore
}

// NewMovieHandler creates and returns a new instance of MovieHandler.
// It initializes the handler with the provided storage implementation.
//
// Parameters:
//   - store: An implementation of the MovieStore interface for data persistence
//
// Returns:
//   - *MovieHandler: A pointer to the newly created MovieHandler instance
func NewMovieHandler(store storage.MovieStore) *MovieHandler {
	return &MovieHandler{store: store}
}

// GetMovies retrieves all movies from the storage.
// This endpoint returns a list of all movies currently stored in the system.
//
// @Summary Get all movies
// @Description Retrieve a list of all movies in the database
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {array} models.Movie "List of movies"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movies [get]
func (h *MovieHandler) GetMovies(w http.ResponseWriter, _ *http.Request) {
	movies, err := h.store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

// CreateMovie handles POST requests to create a new movie.
// This endpoint accepts a movie object in the request body and creates a new movie record.
//
// @Summary Create a new movie
// @Description Create a new movie with the provided information
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Movie object to be created"
// @Success 200 {object} models.Movie "Successfully created movie"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movies [post]
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie *models.Movie = new(models.Movie)

	// Decode the request body into a Movie struct
	err := json.NewDecoder(r.Body).Decode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie, err = h.store.Create(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("INFO: POST request received for creating movie with ID: %s", movie.ID)
}

// GetMovie retrieves a single movie by its ID.
// This endpoint returns the details of a specific movie identified by the provided ID.
//
// @Summary Get a movie by ID
// @Description Retrieve a specific movie using its unique identifier
// @Tags movies
// @Accept json
// @Produce json
// @Param id query string true "Movie ID"
// @Success 200 {object} models.Movie "Movie details"
// @Failure 400 {object} map[string]string "Missing or invalid ID parameter"
// @Failure 404 {object} map[string]string "Movie not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movie [get]
func (h *MovieHandler) GetMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// idValue := request.PathValue(request, "id")
	// Get the "id" query parameter: /movie?id=1
	id := request.URL.Query().Get("id")
	if id == "" {
		http.Error(writer, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	movie, err := h.store.GetByID(id)
	if err != nil {
		http.Error(writer, "Movie not found", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(movie)
	if err != nil {
		http.Error(writer, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(data)
	if err != nil {
		http.Error(writer, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// DeleteMovie removes a movie from the storage by its ID.
// This endpoint permanently deletes a movie record from the system.
//
// @Summary Delete a movie
// @Description Remove a movie from the database using its ID
// @Tags movies
// @Accept json
// @Produce json
// @Param id query string true "Movie ID to delete"
// @Success 200 {object} map[string]string "Movie successfully deleted"
// @Failure 400 {object} map[string]string "Missing or invalid ID parameter"
// @Failure 404 {object} map[string]string "Movie not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movie/delete [delete]
func (h *MovieHandler) DeleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := request.URL.Query().Get("id")
	if id == "" {
		http.Error(writer, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}
	err := h.store.Delete(id)
	if err != nil {
		http.Error(writer, "Movie not found", http.StatusNotFound)
		return
	}
	log.Printf("INFO: DELETE request for movie with ID: %s", id)
}

// UpdateMovie modifies an existing movie record.
// This endpoint accepts a complete movie object and updates the corresponding record.
//
// @Summary Update an existing movie
// @Description Update a movie with the provided information
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Updated movie object"
// @Success 200 {object} models.Movie "Successfully updated movie"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 404 {object} map[string]string "Movie not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movie/update [put]
func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	// Decode the request body into a Movie struct
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.store.Update(movie.ID, &movie)
	if err != nil {

		http.Error(w, fmt.Errorf("Errored out while updating movie: %w", err).Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(movie)
	if err != nil {

	}
	log.Printf("INFO: POST request received for updateing movie with ID: %s", movie.ID)
}
