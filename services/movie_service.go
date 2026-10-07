package services

import (
	"cinema-booking/models"
	"cinema-booking/repositories"
)

type MovieService struct {
	repository *repositories.MovieRepository
}

func NewMovieService(repository *repositories.MovieRepository) *MovieService {
	return &MovieService{
		repository: repository,
	}
}

func (s *MovieService) GetMovies() ([]models.Movie, error) {
	return s.repository.GetMovies()
}

func (s *MovieService) AddMovie(movieData *models.Movie) error {
	return s.repository.AddMovie(movieData)
}

func (s *MovieService) GetMovieByID(id uint) (*models.Movie, error) {
	return s.repository.GetMovieByID(id)
}
