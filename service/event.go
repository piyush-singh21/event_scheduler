package service

import (
	"event_scheduler/database"
	"event_scheduler/model"
	"fmt"
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
		result.Scan(&event.ID, &event.Title, &event.Description, &event.Date)
		fmt.Println(event.Date.Format("2006-04-13"))
		eventResp = append(eventResp, event)
	}
	return eventResp, nil
}

func UpdateEvent() {

}
