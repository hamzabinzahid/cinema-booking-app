package models

import "time"

type Booking struct {
	ID         uint      `json:"id"`
	ShowtimeID uint      `json:"showtime_id"`
	UserID     uint      `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
