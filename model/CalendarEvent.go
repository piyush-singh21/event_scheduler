package model

import "time"

type CalendarEvent struct {
	Summary     string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}
