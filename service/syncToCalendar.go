package service

import (
	"event_scheduler/database"
	"event_scheduler/model"
	"fmt"

	"google.golang.org/api/calendar/v3"
)

var calendarSvc *calendar.Service

func SyncToCalendar(userId int, eventAdd model.EventAdd) error {
	email := database.GetEmail(userId)
	// startTime, err := time.Parse(time.RFC3339, eventAdd.Date)
	// if err != nil {
	// 	fmt.Println("Invalid startime format")
	// }
	// endtime := startTime.Add(30 * time.Minute)
	calEvent := &calendar.Event{
		Summary:     eventAdd.Title,
		Description: eventAdd.Description,
		Start: &calendar.EventDateTime{
			DateTime: eventAdd.Date,
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: eventAdd.Date,
			TimeZone: "UTC",
		},
	}
	_, err := calendarSvc.Events.Insert(email, calEvent).Do()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
