package models

type Seat struct {
	ID       uint   `json:"id"`
	ScreenID uint   `json:"screen_id"`
	Number   string `json:"number"`
}
