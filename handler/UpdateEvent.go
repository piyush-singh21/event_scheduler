package handler

import (
	"context"
	"event_scheduler/database"
	"event_scheduler/model"
	"event_scheduler/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// UpdateEvent Update an already exisisting event
// @Summary Update event
// @Schemes
// @Description Update event mapped to user after logging in
// @Tags example
// @Accept json
// @Produce json
// @Param order body model.UpdateEvent true "Update Event"
// @Success 200 {string} Event Updated Successfully
// @Router /updateEvent [put]
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
	svr, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	// id := strconv.Itoa(updateEvent.ID) + "event"
	// event := service.GetCalendarEvent(updateEvent.ID, svr)
	service.UpdateCalendarData(updateEvent, svr)
	c.JSON(http.StatusOK, gin.H{"success": "event updated"})

}
