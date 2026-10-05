package controllers

import (
	"net/http"

	"cinema-booking/services"
)

type BookingController struct {
	service *services.BookingService
}

func NewBookingController(service *services.BookingService) *BookingController {
	return &BookingController{
		service: service,
	}
}

func (c *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Implement later.
}
