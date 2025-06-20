package memory

import (
	"errors"
	"github.com/vishal130195/go-movies-crud/internal/models"
	"github.com/vishal130195/go-movies-crud/internal/utils"
	"strconv"
	"sync"
)

type MemoryMovieStore struct {
	mutex   sync.RWMutex
	movies  []models.Movie
	counter uint64
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

// Create adds a new movie to the in-memory store, ensuring thread-safe access with a write lock and unlock mechanism.
func (s *MemoryMovieStore) Create(movie *models.Movie) (*models.Movie, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.counter++
	movie = &models.Movie{
		ID:    strconv.FormatUint(s.counter, 10),
		Isbn:  movie.Isbn,
		Title: movie.Title,
		Director: &models.Director{
			ID:        utils.GetUUID(),
			FirstName: movie.Director.FirstName,
			LastName:  movie.Director.LastName,
		},
	}
	s.movies = append(s.movies, *movie)
	return movie, nil
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
	for i := range s.movies {
		if s.movies[i].ID == id {
			s.movies[i].Isbn = movieIn.Isbn
			s.movies[i].Title = movieIn.Title
			// Keep the existing director ID, just update names
			s.movies[i].Director.FirstName = movieIn.Director.FirstName
			s.movies[i].Director.LastName = movieIn.Director.LastName
			return nil
		}
	}
	return errors.New("movie not found")
}
