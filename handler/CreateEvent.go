package handler

import (
	"context"
	"event_scheduler/model"
	"event_scheduler/service"
	auth "event_scheduler/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// CreateEvent Any registered user can create an event
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
	service.AddEvent(srv, eventAdd)

	c.JSON(http.StatusCreated, gin.H{"success": "Event added successfully"})
}
