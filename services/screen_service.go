package services

import (
	"cinema-booking/models"
	"cinema-booking/repositories"
)

type ScreenService struct {
	repository *repositories.ScreenRepository
}

func NewScreenService(repository *repositories.ScreenRepository) *ScreenService {
	return &ScreenService{
		repository: repository,
	}
}

func (s *ScreenService) GetScreens() ([]models.Screen, error) {
	return s.repository.GetScreens()
}

func (s *ScreenService) GetScreenById(id uint) (*models.Screen, error) {
	return s.repository.GetScreenById(id)
}

func (s *ScreenService) AddScreen(screen models.Screen) (*models.Screen, error) {
	return s.repository.AddScreen(screen)
}
