package models

type Screen struct {
	ID       uint   `json:"id"`
	CinemaID uint   `json:"cinema_id"`
	Name     string `json:"name"`
}
