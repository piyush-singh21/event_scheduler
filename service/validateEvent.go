package service

import (
	"errors"
	"event_scheduler/database"
	"event_scheduler/model"
	auth "event_scheduler/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Used for validating the existing event
func ValidateEvent(EventAdd model.EventAdd, key int) error {
	parsedStartTime, err := auth.ParseDate(EventAdd.StartDate)
	if err != nil {
		return err
	}
	parsedEndTime, err := auth.ParseDate(EventAdd.EndDate)
	if err != nil {
		return err
	}
	if parsedStartTime.Before(time.Now()) {
		return errors.New("time should be in future")
	}
	if parsedEndTime.Compare(parsedStartTime) <= 0 {
		return errors.New("end time should be ahead of start time")
	}

	database.AddEvent(EventAdd, parsedStartTime, parsedEndTime, key)
	return nil
}

// Check if user is deleting his events only
func ValidateDeleteEvent(deleteEvent model.DeleteEvent, c *gin.Context) error {
	id, _ := c.Get("id")
	if deleteEvent.ID != id {
		return errors.New("wrong id")
	}
	return nil
}

// Check if user is updating his events only
func ValidateUpdateEvent(updateEvent model.UpdateEvent, c *gin.Context) error {
	id, _ := c.Get("id")
	authId := auth.Convert(id)
	var userId int
	database.DB.QueryRow("SELECT userId FROM events WHERE id=?", updateEvent.ID).Scan(&userId)
	if userId != authId {
		return errors.New("you are not allowed to update this event")
	}
	parsedStartTime, err := auth.ParseDate(updateEvent.StartDate)
	if err != nil {
		return err
	}
	parsedEndTime, err := auth.ParseDate(updateEvent.EndDate)
	if err != nil {
		return err
	}
	if parsedStartTime.Before(time.Now()) {
		return errors.New("time should be in future")
	}
	if parsedEndTime.Compare(parsedStartTime) <= 0 {
		return errors.New("end time should be ahead of start time")
	}
	return nil

}
