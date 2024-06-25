package model

import "time"

type EventResp struct {
	ID          int64     `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"Date"`
}
