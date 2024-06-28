package handler

import (
	"event_scheduler/model"
	"event_scheduler/service"
	auth "event_scheduler/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var eventAdd model.EventAdd
	if err := c.ShouldBindJSON(&eventAdd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	key, _ := c.Get("id")
	val := auth.Convert(key)
	err := service.ValidateEvent(eventAdd, val)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Event added successfully"})
}
