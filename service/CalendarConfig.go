package service

import (
	"context"
	"encoding/json"
	"event_scheduler/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	fmt.Println(authCode)

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
func AddEvent(id int, svr *calendar.Service, event model.EventAdd) {
	layout := "2006-01-02T15:04:05"
	// timeString := event.Date
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Unable to load location: %v", err)
	}
	startTime, err := time.ParseInLocation(layout, event.StartDate, location)
	if err != nil {
		fmt.Println(err)
		return
	}
	endTime, err := time.ParseInLocation(layout, event.EndDate, location)
	if err != nil {
		fmt.Println(err)
		return
	}
	endTimeString := endTime.Format(time.RFC3339)
	startTimeString := startTime.Format(time.RFC3339)
	eventId := strconv.Itoa(id) + "event"
	// fmt.Println(eventId)
	e := &calendar.Event{
		Id:          eventId,
		Summary:     event.Title,
		Description: event.Description,
		Start: &calendar.EventDateTime{
			DateTime: startTimeString,
			TimeZone: "Asia/Kolkata",
		},
		End: &calendar.EventDateTime{
			DateTime: endTimeString,
			TimeZone: "Asia/Kolkata",
		},
		Location: event.Location,
	}
	_, err = svr.Events.Insert("primary", e).Do()
	if err != nil {
		log.Fatalf("Unable to add event %v", err)
	}

}
func GetCalendarEvent(eventId int, svr *calendar.Service) *calendar.Event {
	id := convertToString(eventId)
	event, _ := svr.Events.Get("primary", id).Do()
	return event
}
func RegisterCalendarEvent(id string, event *calendar.Event, svr *calendar.Service) {
	svr.Events.Update("primary", id, event).Do()
}
func UpdateCalendarData(updateEvent model.UpdateEvent, attendees []*calendar.EventAttendee, svr *calendar.Service) {
	eventId := convertToString(updateEvent.ID)
	layout := "2006-01-02T15:04:05"
	// timeString := event.Date
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Unable to load location: %v", err)
	}
	startTime, err := time.ParseInLocation(layout, updateEvent.StartDate, location)
	if err != nil {
		fmt.Println(err)
		return
	}
	endTime, err := time.ParseInLocation(layout, updateEvent.EndDate, location)
	if err != nil {
		fmt.Println(err)
		return
	}
	endTimeString := endTime.Format(time.RFC3339)
	startTimeString := startTime.Format(time.RFC3339)
	e := &calendar.Event{
		Id:          eventId,
		Summary:     updateEvent.Title,
		Description: updateEvent.Description,
		Start: &calendar.EventDateTime{
			DateTime: startTimeString,
			TimeZone: "Asia/Kolkata",
		},
		End: &calendar.EventDateTime{
			DateTime: endTimeString,
			TimeZone: "Asia/Kolkata",
		},
		Attendees: attendees,
		Location:  updateEvent.Location,
	}
	svr.Events.Update("primary", eventId, e).Do()

}
func DeleteCalenderEvent(id int, svr *calendar.Service) {
	eventId := convertToString(id)
	svr.Events.Delete("primary", eventId).Do()
}
func convertToString(key int) string {
	return strconv.Itoa(key) + "event"
}
