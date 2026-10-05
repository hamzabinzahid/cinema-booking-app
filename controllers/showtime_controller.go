package controllers

import (
	"net/http"

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
	// Implement later.
}

func (c *ShowtimeController) GetShowtimeByID(w http.ResponseWriter, r *http.Request) {
	// Implement later.
}
