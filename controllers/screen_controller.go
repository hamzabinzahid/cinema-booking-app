package controllers

import (
	"cinema-booking/models"
	"cinema-booking/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ScreenController struct {
	service *services.ScreenService
}

func NewScreenController(service *services.ScreenService) *ScreenController {
	return &ScreenController{
		service: service,
	}
}

func (s *ScreenController) GetScreens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := s.service.GetScreens()
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Screens not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (s *ScreenController) GetScreenById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Unable to convert id to int", http.StatusInternalServerError)
	}

	data, err := s.service.GetScreenById(uint(idInt))
	if err != nil {
		http.Error(w, "Screen not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(data)
}

func (s *ScreenController) AddScreen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var screenData models.Screen
	err := json.NewDecoder(r.Body).Decode(&screenData)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	data, err := s.service.AddScreen(screenData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to add screen", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
