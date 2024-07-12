package model

type EventResp struct {
	ID          int64  `json:"event_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Location    string `json:"location"`
}
