package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cinema-booking/models"
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
	w.Header().Set("Content-Type", "application/json")
	movieIdStr := r.PathValue("id")
	movieIdInt, err := strconv.Atoi(movieIdStr)
	if err != nil {
		http.Error(w, "Failed to convert movie ID", http.StatusInternalServerError)
	}

	movieRecord, err := c.service.GetMovieByID(uint(movieIdInt))
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(movieRecord)

}

func (c *MovieController) AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = c.service.AddMovie(&movie)
	if err != nil {
		http.Error(w, "Failed to create movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}
