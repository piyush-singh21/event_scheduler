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
func GetId(deleteEvent model.DeleteEvent) int {
	var id int
	DB.QueryRow("SELECT id FROM events WHERE title=? AND userId=?", deleteEvent.Title, deleteEvent.ID).Scan(&id)
	return id
}
func AddEvent(EventAdd model.EventAdd, parsedStartTime time.Time, parsedEndTime time.Time, key int) {
	DB.Exec("INSERT INTO events (title,description,userId,StartDate,EndDate,Location) VALUES (?,?,?,?,?,?)", EventAdd.Title, EventAdd.Description, key, parsedStartTime, parsedEndTime, EventAdd.Location)
}

func UpdateEvent(updateEvent model.UpdateEvent) {
	DB.Exec("UPDATE events SET title=?,description=?,StartDate=?,EndDate=?,Location=? WHERE id=?", updateEvent.Title, updateEvent.Description, updateEvent.StartDate, updateEvent.EndDate, updateEvent.Location, updateEvent.ID)
}
func GetEmail(userId int) string {
	var email string
	DB.QueryRow("SELECT email FROM users WHERE id=?", userId).Scan(&email)
	return email
}
func GetLastEntryId() int {
	var id int
	DB.QueryRow("SELECT id FROM events order by id desc limit 1").Scan(&id)
	return id
}
