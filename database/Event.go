package database

import (
	"database/sql"
	"errors"
	"event_scheduler/model"
	"time"
)

func DeleteEvent(deleteEvent model.DeleteEvent) error {
	var title string
	err := DB.QueryRow("SELECT title FROM events WHERE title=?", deleteEvent.Title).Scan(&title)
	// fmt.Println(title)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no such title found")
		}
	}
	DB.Exec("DELETE FROM events WHERE title=? AND userId=?", deleteEvent.Title, deleteEvent.ID)
	return nil
}
func AddEvent(EventAdd model.EventAdd, parsedTime time.Time, key int) {
	DB.Exec("INSERT INTO events (title,description,userId,date) VALUES (?,?,?,?)", EventAdd.Title, EventAdd.Description, key, parsedTime)
}
