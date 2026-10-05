package controllers

import (
	"encoding/json"
	"net/http"

	"cinema-booking/services"
)

type MovieController struct {
	service *services.MovieService
}

func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{
		service: service,
	}
}

func (c *MovieController) GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movies, err := c.service.GetMovies()
	if err != nil {
		http.Error(w, "Failed to get movies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

func (c *MovieController) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	// Implement later.
}

func (c *MovieController) AddMovie(w http.ResponseWriter, r *http.Request) {
	// Implement later.
}
