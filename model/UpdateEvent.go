package model

type UpdateEvent struct {
	ID int `json:"id"`
	// UserId      int    `json:"userId"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	EndDate     string `json:"endDate,omitempty"`
	Location    string `json:"location,omitempty"`
}
