package handler

import (
	"event_scheduler/model"
	"event_scheduler/service"
	auth "event_scheduler/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create event
// @Schemes
// @Description Create event after logging in
// @Tags example
// @Accept json
// @Produce json
// @Param order body model.EventAdd true "Create Event"
// @Success 200 {string} Event Created Successfully
// @Router /createEvent [post]
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
	err = service.GetDataToSendMail(val, eventAdd)
	if err != nil {
		fmt.Println(err)
	}
	// err = service.SyncToCalendar(val, eventAdd)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	c.JSON(http.StatusCreated, gin.H{"success": "Event added successfully"})
}
