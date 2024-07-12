package handler

import (
	"event_scheduler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindEvent find all the events
// @Summary Get event
// @Schemes
// @Description Get list of all event after logging in
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} All Events
// @Router /events [get]
func FindEvent(c *gin.Context) {
	res, err := service.GetAllEvent()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}
