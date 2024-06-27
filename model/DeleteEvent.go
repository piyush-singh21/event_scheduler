package model

type DeleteEvent struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}
