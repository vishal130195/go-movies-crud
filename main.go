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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	// Further we can add mmore information, Like Image, image url, DOB, Age, DOD etc.
}

var movies []Movie

func initialiseMovies() /*[]Movie*/ {
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
	//return movies
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// We need to write json data to writer and what data in json that is movies that need to be encoded.
	json.NewEncoder(w).Encode(movies)
}

// delete function

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// we need ID to delete the movie
	w.Header().Set("Content-Type", "application/json")
	// We need id here from request

	//Vars returns the route variables for the current request, if any.
	params := mux.Vars(r)

	// As here we have set of data set, we can not identify the data in map access here we need to check in loop for id one by one
	for index, movie := range movies {
		if movie.ID == params["id"] {
			// delete this..
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// we need ID to delete the movie
	w.Header().Set("Content-Type", "application/json")
	// We need id here from request

	//Vars returns the route variables for the current request, if any.
	params := mux.Vars(r)

	// As here we have set of data set, we can not identify the data in map access here we need to check in loop for id one by one
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	// create a random ID
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get id from params

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// delete movie and add new
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	// find and delete that movie with id
}

// Main function
func main() {
	r := mux.NewRouter()
	initialiseMovies()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

/*
 5 Routes
 1. Get All = getMovies()
 2. Get by id  = getMovie(id) // ID required
 3. Create = createMovie()
 4. Update = updateMovie() // ID required and values required
 5. Delete = deleteMovie() // ID required
*/
