package repositories

import (
	"cinema-booking/models"
	"database/sql"
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
	return nil, nil
}
