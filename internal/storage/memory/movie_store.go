package memory

import (
	"errors"
	"github.com/vishal130195/go-movies-crud/internal/models"
	"strconv"
	"sync"
)

var counter = 0

type MemoryMovieStore struct {
	mutex  sync.RWMutex
	movies []models.Movie
}

func NewMemoryMovieStore() *MemoryMovieStore {
	return &MemoryMovieStore{
		movies: make([]models.Movie, 0),
	}
}

func (s *MemoryMovieStore) GetAll() ([]models.Movie, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.movies, nil
}

func (s *MemoryMovieStore) GetByID(id string) (*models.Movie, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, movie := range s.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

// Create adds a new movie to the in-memory store, ensuring thread-safe access with a read lock and unlock mechanism.
func (s *MemoryMovieStore) Create(movie *models.Movie) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	counter++
	s.movies = append(s.movies, models.Movie{
		ID:    strconv.Itoa(counter),
		Isbn:  movie.Isbn,
		Title: movie.Title,
		Director: &models.Director{
			FirstName: movie.Director.FirstName,
			LastName:  movie.Director.LastName,
		},
	})
	return nil
}

func (s *MemoryMovieStore) Delete(ID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for index, movie := range s.movies {
		if movie.ID == ID {
			s.movies = append(s.movies[:index], s.movies[index+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}

func (s *MemoryMovieStore) Update(id string, movieIn *models.Movie) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, movie := range s.movies {
		if movie.ID == id {
			movie.Isbn = movieIn.Isbn
			movie.Title = movieIn.Title
			movie.Director.FirstName = movieIn.Director.FirstName
			movie.Director.LastName = movieIn.Director.LastName
			return nil
		}
	}
	return errors.New("movie not found")
}
