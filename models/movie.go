package models

type Movie struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Duration uint   `json:"duration"`
}
