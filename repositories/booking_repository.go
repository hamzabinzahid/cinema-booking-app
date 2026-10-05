package repositories

import (
	"database/sql"

	"cinema-booking/models"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) CreateBooking(booking *models.Booking) error {
	// Database insert will be implemented here.
	return nil
}
