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
	parsedTime, err := auth.ParseDate(EventAdd.Date)
	if err != nil {
		return err
	}
	if parsedTime.Before(time.Now()) {
		return errors.New("time should be in future")
	}

	database.AddEvent(EventAdd, parsedTime, key)
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
	var userId int
	database.DB.QueryRow("SELECT userId FROM events WHERE id=?", updateEvent.ID).Scan(&userId)
	if userId != id {
		return errors.New("you are not allowed to update this event")
	}
	return nil

}
