package repositories

import (
	"database/sql"

	"cinema-booking/models"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetMovies() ([]models.Movie, error) {
	// Database query will be implemented here.
	return nil, nil
}
