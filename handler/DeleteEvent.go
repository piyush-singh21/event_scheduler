package handler

import (
	"event_scheduler/database"
	"event_scheduler/model"
	"event_scheduler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Delete event
// @Schemes
// @Description Delete event mapped to user after logging in
// @Tags example
// @Accept json
// @Produce json
// @Param order body model.DeleteEvent true "Delete Event"
// @Success 200 {string} Event Deleted Successfully
// @Router /deleteEvent [delete]
func DeleteEvent(c *gin.Context) {
	var deleteEvent model.DeleteEvent
	if err := c.ShouldBindJSON(&deleteEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such event found"})
		return
	}
	err := service.ValidateDeleteEvent(deleteEvent, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DeleteEvent(deleteEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": "Event removed successfully"})
}
