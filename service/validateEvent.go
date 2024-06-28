package service

import (
	"errors"
	"event_scheduler/database"
	"event_scheduler/model"
	auth "event_scheduler/utils"
	"time"

	"github.com/gin-gonic/gin"
)

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
func ValidateDeleteEvent(deleteEvent model.DeleteEvent, c *gin.Context) error {
	id, _ := c.Get("id")
	if deleteEvent.ID != id {
		return errors.New("wrong id")
	}
	return nil
}
