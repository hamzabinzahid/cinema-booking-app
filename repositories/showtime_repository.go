package repositories

import (
	"database/sql"

	"cinema-booking/models"
)

type ShowtimeRepository struct {
	db *sql.DB
}

func NewShowtimeRepository(db *sql.DB) *ShowtimeRepository {
	return &ShowtimeRepository{
		db: db,
	}
}

func (r *ShowtimeRepository) GetShowtimes() ([]models.Showtime, error) {
	// Database query will be implemented here.
	return nil, nil
}
