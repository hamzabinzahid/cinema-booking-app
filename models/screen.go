package models

type Screen struct {
	ScreenSingle
	CinemaDetails CinemaDetails `json:"cinema"`
}

type ScreenSingle struct {
	ID       uint   `json:"id"`
	CinemaID uint   `json:"cinema_id"`
	Name     string `json:"name"`
}
