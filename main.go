package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie struct represents a movie entity
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct represents a director entity
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	// Further, we can add more information, like Image, image URL, DOB, Age, DOD, etc.
}

// movies is a slice to store movie data
var movies []Movie

// initialiseMovies initializes the movies slice with some sample data
func initialiseMovies() {
	movie_one := Movie{
		ID:    "1",
		Isbn:  "12345",
		Title: "Movie One",
		Director: &Director{
			FirstName: "Vishal",
			LastName:  "Singh",
		},
	}

	movie_two := Movie{
		ID:    "2",
		Isbn:  "823981",
		Title: "Movie Steve",
		Director: &Director{
			FirstName: "John",
			LastName:  "Tales",
		},
	}
	movies = append(movies, movie_one, movie_two)
}

// getMovies handles GET requests for all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	log.Println("INFO: GET request received for all movies")
}

// deleteMovie handles DELETE requests to delete a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			// Delete the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			log.Printf("INFO: DELETE request received for movie with ID: %s", params["id"])
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// getMovie handles GET requests to retrieve a movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			// Encode and send the movie details in the response
			json.NewEncoder(w).Encode(movie)
			log.Printf("INFO: GET request received for movie with ID: %s", params["id"])
			break
		}
	}
}

// createMovie handles POST requests to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	// Decode the request body into a Movie struct
	json.NewDecoder(r.Body).Decode(&movie)

	// Generate a random ID and append the new movie to the slice
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	// Encode and send the created movie details in the response
	json.NewEncoder(w).Encode(movie)
	log.Printf("INFO: POST request received for creating movie with ID: %s", movie.ID)
}

// updateMovie handles PUT requests to update a movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the existing movie and add the updated one
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			// Encode and send the updated movie details in the response
			json.NewEncoder(w).Encode(movie)
			log.Printf("INFO: PUT request received for updating movie with ID: %s", params["id"])
			return
		}
	}
}

// Main function
func main() {
	r := mux.NewRouter()
	initialiseMovies()

	// Define routes and their corresponding handler functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods
