package handler

import (
	"event_scheduler/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var eventAdd model.EventAdd
	if err := c.ShouldBindJSON(&eventAdd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(eventAdd.Date)
}
