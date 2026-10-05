package services

import (
	"cinema-booking/models"
	"cinema-booking/repositories"
)

type BookingService struct {
	repository *repositories.BookingRepository
}

func NewBookingService(repository *repositories.BookingRepository) *BookingService {
	return &BookingService{
		repository: repository,
	}
}

func (s *BookingService) CreateBooking(booking *models.Booking) error {
	// Booking logic and concurrency handling will go here.
	return s.repository.CreateBooking(booking)
}
