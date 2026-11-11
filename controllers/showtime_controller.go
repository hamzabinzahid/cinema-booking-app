package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cinema-booking/models"
	"cinema-booking/services"
)

type ShowtimeController struct {
	service *services.ShowtimeService
}

func NewShowtimeController(service *services.ShowtimeService) *ShowtimeController {
	return &ShowtimeController{
		service: service,
	}
}

func (c *ShowtimeController) GetShowtimes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := c.service.GetShowtimes()
	if err != nil {
		http.Error(w, "Showtimes not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (c *ShowtimeController) GetShowtimeByMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	strId := r.PathValue("movieid")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Unable to convert movie id", http.StatusInternalServerError)
		return
	}

	data, err := c.service.GetShowtimesByMovie(uint(intId))
	if err != nil {
		http.Error(w, "Showtimes not found", http.StatusNotFound)
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (c *ShowtimeController) GetShowtimeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	strId := r.PathValue("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Unable to convert movie id", http.StatusInternalServerError)
		return
	}

	data, err := c.service.GetShowtimesByMovie(uint(intId))
	if err != nil {
		http.Error(w, "Showtimes not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (c *ShowtimeController) AddShowtime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var showtimeDetails models.Showtime
	json.NewDecoder(r.Body).Decode(&showtimeDetails)

	data, err := c.service.AddShowtime(showtimeDetails)
	if err != nil {
		http.Error(w, "Unable to add showtime", http.StatusNotFound)
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(data)
}
