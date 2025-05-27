package storage

import "github.com/vishal130195/go-movies-crud/internal/models"

// MovieStore defines the interface for movie storage operations
type MovieStore interface {
	GetAll() ([]models.Movie, error)
	GetByID(id string) (*models.Movie, error)
	Create(movie *models.Movie) error
	Update(id string, movie *models.Movie) error
	Delete(id string) error
}
