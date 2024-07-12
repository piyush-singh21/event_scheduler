package database

import "event_scheduler/model"

func RegisterForEvent(eventId, key int) {
	var email string
	var name string
	DB.QueryRow("SELECT email,username FROM users WHERE id=?", key).Scan(&email, &name)
	DB.Exec("INSERT INTO register (email,eventId,name) VALUES (?,?,?)", email, eventId, name)

}

func FindEvent(key int) model.EventResp {
	var eventResp model.EventResp
	DB.QueryRow("SELECT id,title,description,StartDate,EndDate,Location FROM events WHERE id=?", key).Scan(&eventResp.ID, &eventResp.Title, &eventResp.Description, &eventResp.StartDate, &eventResp.EndDate, &eventResp.Location)
	return eventResp
}
