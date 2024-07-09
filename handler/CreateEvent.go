package handler

import (
	"event_scheduler/model"
	"event_scheduler/service"
	auth "event_scheduler/utils"
	"fmt"
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
	// fmt.Println(val)
	err := service.ValidateEvent(eventAdd, val)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.GetDataToSendMail(val, eventAdd)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusCreated, gin.H{"success": "Event added successfully"})
}
