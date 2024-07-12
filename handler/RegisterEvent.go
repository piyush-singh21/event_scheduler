package handler

import (
	"event_scheduler/model"
	"event_scheduler/service"
	auth "event_scheduler/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEvent(c *gin.Context) {
	var registerEvent model.RegisterEvent
	if err := c.ShouldBindJSON(&registerEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	eventResp, err := service.ValidateRegistrationId(registerEvent.EventId, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	keyInterface, _ := c.Get("id")
	key := auth.Convert(keyInterface)
	body := fmt.Sprintf("Dear user,\n\n You registered to event successfully :\n\nTitle: %s\n Description: %s\nStartTime: %s\nEndTime: %s\nLocation: %s\n\n\nBest regards", eventResp.Title, eventResp.Description, eventResp.StartDate, eventResp.EndDate, eventResp.Location)
	service.GetDataToSendMail(key, body, "register")
	c.JSON(http.StatusOK, gin.H{"success": "Registered to event"})
}
