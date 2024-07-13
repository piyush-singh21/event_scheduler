package handler

import (
	"context"
	"event_scheduler/database"
	"event_scheduler/model"
	"event_scheduler/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// DeleteEvent User can only delete events created by him
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
	id := database.GetId(deleteEvent)
	err = database.DeleteEvent(deleteEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read secret file %v\n", err)
		return
	}
	scopes := []string{
		calendar.CalendarEventsScope,
		calendar.CalendarScope,
		// Add other scopes as needed
	}
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		log.Fatalf("unable to parse client secret file to config: %v\n", err)
		return
	}
	client := service.GetClient(config)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	fmt.Println(id)
	service.DeleteCalenderEvent(id, srv)
	c.JSON(http.StatusAccepted, gin.H{"success": "Event removed successfully"})
}
