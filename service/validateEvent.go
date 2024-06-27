package service

import (
	"errors"
	"event_scheduler/model"

	"github.com/gin-gonic/gin"
)

func ValidateEvent(EventAdd model.EventAdd) {

}
func ValidateDeleteEvent(deleteEvent model.DeleteEvent, c *gin.Context) error {
	id, _ := c.Get("id")
	if deleteEvent.ID != id {
		return errors.New("wrong id")
	}
	return nil
}
