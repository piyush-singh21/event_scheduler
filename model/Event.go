package model

import "time"

type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"Date"`
	UserID      int64     `json:"UserId"`
}
