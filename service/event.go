package service

import (
	"event_scheduler/database"
	"event_scheduler/model"
	"fmt"
	"time"
)

func GetAllEvent() ([]model.EventResp, error) {
	var eventResp []model.EventResp
	result, err := database.DB.Query("SELECT id,title,description,date FROM events")
	if err != nil {
		return []model.EventResp{}, nil
	}
	// const layout = "2006-Jan-02"
	for result.Next() {
		var event model.EventResp
		var dateString time.Time
		result.Scan(&event.ID, &event.Title, &event.Description, &event.Date)
		fmt.Println(event.Date)
		dateString = time.Time(event.Date)
		// fmt.Println(dateString)
		event.Date = dateString
		eventResp = append(eventResp, event)
	}
	return eventResp, nil
}
