package handler

import (
	"event_scheduler/database"
	"event_scheduler/model"
	"event_scheduler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Update event
// @Schemes
// @Description Update event mapped to user after logging in
// @Tags example
// @Accept json
// @Produce json
// @Param order body model.UpdateEvent true "Update Event"
// @Success 200 {string} Event Updated Successfully
// @Router /updateEvent [put]

// UpdateEvent Update an already exisisting event
func UpdateEvent(c *gin.Context) {
	var updateEvent model.UpdateEvent
	if err := c.ShouldBindJSON(&updateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to update"})
		return
	}
	err := service.ValidateUpdateEvent(updateEvent, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.UpdateEvent(updateEvent)
	c.JSON(http.StatusOK, gin.H{"success": "event updated"})

}
