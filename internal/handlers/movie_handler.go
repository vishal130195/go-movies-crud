package handlers

import (
	"encoding/json"
	"github.com/vishal130195/go-movies-crud/internal/models"
	"github.com/vishal130195/go-movies-crud/internal/storage"
	"log"
	"net/http"
)

type MovieHandler struct {
	store storage.MovieStore
}

func NewMovieHandler(store storage.MovieStore) *MovieHandler {
	return &MovieHandler{store: store}
}

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

// Implement CreateMovieHandler

// createMovie handles POST requests to create a new movie
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie

	// Decode the request body into a Movie struct
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.store.Create(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("INFO: POST request received for creating movie with ID: %s", movie.ID)
}

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
		http.Error(w, "Errored out while updating movie", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("INFO: POST request received for updateing movie with ID: %s", movie.ID)
}
