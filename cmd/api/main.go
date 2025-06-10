package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vishal130195/go-movies-crud/internal/handlers"
	"github.com/vishal130195/go-movies-crud/internal/storage/memory"
	"log"
	"net/http"
)

func main() {
	// Initialize storage
	movieStore := memory.NewMemoryMovieStore()

	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(movieStore)

	// Setup router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/movie", movieHandler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
	r.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/update", movieHandler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/delete", movieHandler.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
