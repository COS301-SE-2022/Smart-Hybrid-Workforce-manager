/*
Base code from https://developers.google.com/calendar/api/quickstart/go
*/
/*
https://developers.google.com/calendar/api/guides/create-events#go
*/

package google_api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"strings"

	// "time"
	"api/data"
	"lib/logger"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	// "google.golang.org/genproto/googleapis/apps/script/type/calendar"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "/google_api/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	defer func() {
			err := f.Close()
			if err != nil {
				logger.Info.Printf("token error: %v\n",err);
			}
		}()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(filepath.Clean(path), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			logger.Info.Printf("token error: %v\n",err);
		}
	}()
	err = json.NewEncoder(f).Encode(token)
	_ = err
}

func createEvent(summary string, location *string, desc *string, starttime time.Time, endtime time.Time, attendee string) *calendar.Event {

	event := &calendar.Event{
		Summary: summary,
		Start: &calendar.EventDateTime{
			DateTime: starttime.Format(time.RFC3339),
			TimeZone: "Africa/Harare",
		},
		End: &calendar.EventDateTime{
			DateTime: endtime.Format(time.RFC3339),
			TimeZone: "Africa/Harare",
		},
		// Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		Attendees: []*calendar.EventAttendee{
		        &calendar.EventAttendee{Email:attendee},
		},
	}
	if location != nil {
		event.Location = *location;
	}
	if desc != nil {
		event.Description = *desc;
	}
    //Add Attendees

	return event
}

func CreateBooking(user *data.User ,booking *data.Booking) string{
	logger.Access.Printf("user: %v\nbooking: %v\n",*user,*booking)
	event := createEvent(*booking.ResourceType,nil,nil,*booking.Start,*booking.End,*user.Email)
	logger.Access.Printf("\nevent: %v\n",event)
	return "Calender Booking"
}

func TestingFunc() bool{
	ctx := context.Background()
	b, err := os.ReadFile("/google_api/credentials.json")
	if err != nil {
	    logger.Error.Printf("Unable to read client secret file: %v\n", err)
		return false
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		logger.Error.Printf("Unable to parse client secret file to config: %v\n", err)
		return false
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Error.Printf("Unable to retrieve Calendar client: %v\n", err)
		return false
	}
	event := createEvent("Testing",nil,nil,time.Now(),time.Now().Add(time.Hour * 5),"email@example.com")
	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		logger.Error.Printf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
	line := event.HtmlLink
	fmt.Println(line[strings.Index(line, "eid=")+4:])

	// fmt.Printf("\nEvent: %s\n", event)
	// fmt.Printf("\nEvent: %T\n", event)
	// // fmt.Printf("\nEvent: %s\n", event.id)
	return true
}


