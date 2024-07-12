package service

import (
	"event_scheduler/database"
	"event_scheduler/model"
)

func GetAllEvent() ([]model.EventResp, error) {
	var eventResp []model.EventResp
	result, err := database.DB.Query("SELECT id,title,description,StartDate,EndDate,Location FROM events")
	if err != nil {
		return []model.EventResp{}, nil
	}
	for result.Next() {
		var event model.EventResp
		result.Scan(&event.ID, &event.Title, &event.Description, &event.StartDate, &event.EndDate, &event.Location)
		eventResp = append(eventResp, event)
	}
	return eventResp, nil
}

func UpdateEvent() {

}
