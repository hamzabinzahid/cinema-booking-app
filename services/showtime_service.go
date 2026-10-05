package services

import (
	"cinema-booking/models"
	"cinema-booking/repositories"
)

type ShowtimeService struct {
	repository *repositories.ShowtimeRepository
}

func NewShowtimeService(repository *repositories.ShowtimeRepository) *ShowtimeService {
	return &ShowtimeService{
		repository: repository,
	}
}

func (s *ShowtimeService) GetShowtimes() ([]models.Showtime, error) {
	return s.repository.GetShowtimes()
}
