package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vishal130195/go-movies-crud/internal/handlers"
	"github.com/vishal130195/go-movies-crud/internal/storage/memory"
	"log"
	"net/http"
)

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize storage
	movieStore := memory.NewMemoryMovieStore()

	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(movieStore)
	// Write a simple log message to indicate the server is starting
	log.Println("Initializing movie handler with in-memory storage")
	// Setup router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/movie", movieHandler.GetMovie).Methods("GET", "OPTIONS")
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET", "OPTIONS")
	r.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST", "OPTIONS")
	r.HandleFunc("/movie/update", movieHandler.UpdateMovie).Methods("PUT", "OPTIONS")
	r.HandleFunc("/movie/delete", movieHandler.DeleteMovie).Methods("DELETE", "OPTIONS")

	// Enable CORS
	r.Use(enableCORS)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
