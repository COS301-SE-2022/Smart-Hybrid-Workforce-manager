/*
Base code from https://developers.google.com/calendar/api/quickstart/go
*/
/*
https://developers.google.com/calendar/api/guides/create-events#go
*/

package google_calendar

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	// "google.golang.org/genproto/googleapis/apps/script/type/calendar"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
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
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func createEvent(summary string, location *string, desc *string, starttime string, endtime string, attendess []string) *calendar.Event {

	event := &calendar.Event{
		Summary: summary,
		Start: &calendar.EventDateTime{
			DateTime: starttime,
			TimeZone: "Africa/Harare",
		},
		End: &calendar.EventDateTime{
			DateTime: endtime,
			TimeZone: "Africa/Harare",
		},
		// Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		Attendees: []*calendar.EventAttendee{
		        &calendar.EventAttendee{Email:"email@example.com"},
		        // &calendar.EventAttendee{Email:"sbrin@example.com"},
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

func TestingFunc() bool{
	// ctx := context.Background()
	// b, err := os.ReadFile("credentials.json")
	// if err != nil {
	//         log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// // If modifying these scopes, delete your previously saved token.json.
	// config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	// if err != nil {
	//         log.Fatalf("Unable to parse client secret file to config: %v", err)
	// }
	// client := getClient(config)

	// srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	// if err != nil {
	//         log.Fatalf("Unable to retrieve Calendar client: %v", err)
	// }
    // var arr []string
	// event := createEvent("Testing",nil,nil,time.Now().Format(time.RFC3339),time.Now().Add(time.Hour * 5).Format(time.RFC3339),arr)
	// calendarId := "primary"
	// event, err = srv.Events.Insert(calendarId, event).Do()
	// if err != nil {
	// log.Fatalf("Unable to create event. %v\n", err)
	// }
	// fmt.Printf("Event created: %s\n", event.HtmlLink)
	return true
}


