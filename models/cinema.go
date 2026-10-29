package models

type Cinema struct {
	CinemaDetails
	Screens []ScreenSingle `json:"screens"`
}

type CinemaDetails struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
