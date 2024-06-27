package model

import "time"

type EventAdd struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"Date"`
}
