package services

import (
	"cinema-booking/models"
	"cinema-booking/repositories"
)

type CinemaService struct {
	service *repositories.CinemaRepository
}

func NewCinemaService(repository *repositories.CinemaRepository) *CinemaService {
	return &CinemaService{
		service: repository,
	}
}

func (c *CinemaService) GetCinemas() ([]models.Cinema, error) {
	return c.service.GetCinemas()
}

func (c *CinemaService) GetCinemaById(id int) (*models.Cinema, error) {
	return c.service.GetCinemaById(id)
}

func (c *CinemaService) AddCinema(cinemaDetails models.Cinema) (*models.Cinema, error) {
	return c.service.AddCinema(cinemaDetails)
}
