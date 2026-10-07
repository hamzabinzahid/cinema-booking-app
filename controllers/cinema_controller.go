package controllers

import (
	"cinema-booking/models"
	"cinema-booking/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type CinemaController struct {
	service *services.CinemaService
}

func NewCinemaController(service *services.CinemaService) *CinemaController {
	return &CinemaController{
		service: service,
	}
}

func (c *CinemaController) GetCinemas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cinemas, err := c.service.GetCinemas()
	if err != nil {
		http.Error(w, "Failed to get cinemas list", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(cinemas)
}

func (c *CinemaController) GetCinemaById(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.Request.PathValue("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Failed to convert Cinema ID", http.StatusInternalServerError)
	}

	cinema, err := c.service.GetCinemaById(idInt)
	if err != nil {
		http.Error(w, "Cinema not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(cinema)
}

func (c *CinemaController) addCinema(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cinema models.Cinema
	json.NewDecoder(r.Body).Decode(&cinema)

	data, err := c.service.AddCinema(cinema)
	if err != nil {
		http.Error(w, "Failed to add cinema", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(data)

}
