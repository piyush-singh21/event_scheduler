package service

import (
	"context"
	"encoding/json"
	"event_scheduler/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

// type CalendarService struct {
// 	srv *calendar.Service
// }

// var (
// 	config *oauth2.Config
// )

// func NewCalendarService() *CalendarService {
// 	return &CalendarService{}
// }
// func LoadAuthConfig() (*oauth2.Config, error) {
// 	b, err := ioutil.ReadFile("credentials.json")
// 	// fmt.Println("token\n", b)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to read client secret file :%v", err)
// 	}
// 	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
// 	}
// 	return config, nil
// }

// func HandleGoogleLogin(c *gin.Context, config *oauth2.Config) {
// 	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	c.Redirect(http.StatusTemporaryRedirect, url)
// }
// func HandleGoogleCallback(config *oauth2.Config, code string) (*oauth2.Token, error) {
// 	tok, err := config.Exchange(context.Background(), code)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to retrive web token %v", err)
// 	}
// 	return tok, nil
// }

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
func AddEvent(svr *calendar.Service, event model.EventAdd) {
	layout := "2006-01-02T15:04:05"
	timeString := event.Date
	startTime, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println(err)
	}
	endTime := startTime.Add(30 * time.Minute)
	endTimeString := endTime.Format(time.RFC3339)
	// timeString += "Z"
	// endTimeString += "Z"
	e := &calendar.Event{
		Summary:     event.Title,
		Description: event.Description,
		Start: &calendar.EventDateTime{
			DateTime: timeString,
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: endTimeString,
			TimeZone: "UTC",
		},
	}
	_, err = svr.Events.Insert("primary", e).Do()
	if err != nil {
		log.Fatalf("Unable to add event %v", err)
	}

}
