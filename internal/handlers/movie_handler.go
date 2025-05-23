package handlers

import (
	"encoding/json"
	"github.com/vishal130195/go-movies-crud/internal/models"
	"github.com/vishal130195/go-movies-crud/internal/storage"
	"log"
	// "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	store storage.MovieStore
	count int
}

func NewMovieHandler(store storage.MovieStore) *MovieHandler {
	return &MovieHandler{store: store,
		count: 0}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Implement CreateMovieHandler

// createMovie handles POST requests to create a new movie
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie

	// Decode the request body into a Movie struct
	json.NewDecoder(r.Body).Decode(&movie)
	h.count++
	movie.ID = strconv.Itoa(h.count)
	h.store.Create(&movie)
	json.NewEncoder(w).Encode(movie)
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
	writer.Write(data)
}

// Implement other handler methods (GetMovie, CreateMovie, UpdateMovie, DeleteMovie)
