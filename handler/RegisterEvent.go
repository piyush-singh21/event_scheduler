package handler

import (
	"context"
	"event_scheduler/database"
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

func RegisterEvent(c *gin.Context) {
	var registerEvent model.RegisterEvent
	if err := c.ShouldBindJSON(&registerEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	keyInterface, _ := c.Get("id")
	key := auth.Convert(keyInterface)
	eventResp, err := service.ValidateRegistrationId(key, registerEvent.EventId, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body := fmt.Sprintf("Dear user,\n\n You registered to event successfully :\n\nTitle: %s\n Description: %s\nStartTime: %s\nEndTime: %s\nLocation: %s\n\n\nBest regards", eventResp.Title, eventResp.Description, eventResp.StartDate, eventResp.EndDate, eventResp.Location)
	service.GetDataToSendMail(key, body, "register")
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
	id := database.GetLastEntryId()
	event := service.GetCalendarEvent(id, svr)
	userMail := database.GetUserMail(key)
	event.Attendees = append(event.Attendees, &calendar.EventAttendee{Email: userMail})
	service.UpdateCalendarEvent(event.Id, event, svr)
	c.JSON(http.StatusOK, gin.H{"success": "Registered to event"})
}
