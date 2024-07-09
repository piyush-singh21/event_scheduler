package model

type UpdateEvent struct {
	ID          int    `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Date        string `json:"Date,omitempty"`
	// UserID      int64  `json:"UserId,omitempty"`
}
