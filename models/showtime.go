package models

import "time"

type Showtime struct {
	ID        uint      `json:"id"`
	MovieID   uint      `json:"movie_id"`
	ScreenID  uint      `json:"screen_id"`
	StartTime time.Time `json:"start_time"`
}
