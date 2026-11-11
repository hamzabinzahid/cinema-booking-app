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

func (s *ShowtimeService) GetShowtimes() ([]models.ShowTimeRel, error) {
	return s.repository.GetShowtimes()
}

func (s *ShowtimeService) GetShowtimesByMovie(movieId uint) ([]models.ShowTimeRel, error) {
	return s.repository.GetShowtimesByMovie(movieId)
}

func (s *ShowtimeService) GetShowtimesById(showTimeId uint) (*models.ShowTimeRel, error) {
	return s.repository.GetShowtimesById(showTimeId)
}

func (s *ShowtimeService) AddShowtime(showTime models.Showtime) (*models.ShowTimeRel, error) {
	return s.repository.AddShowtime(showTime)
}
