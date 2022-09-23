/*
Base code from https://developers.google.com/calendar/api/quickstart/go
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	// "time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	// "google.golang.org/genproto/googleapis/apps/script/type/calendar"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
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

func createEvent(summary *string, location *string, desc *string, starttime *string, endtime *string) (*calendar.Event){

        event := &calendar.Event{
                Summary: "Google I/O 2015",
                Location: "800 Howard St., San Francisco, CA 94103",
                Description: "A chance to hear more about Google's developer products.",
                Start: &calendar.EventDateTime{
                        DateTime: "2022-09-24T09:00:00-07:00",
                        TimeZone: "America/Los_Angeles",
                },
                End: &calendar.EventDateTime{
                        DateTime: "2022-09-24T17:00:00-07:00",
                        TimeZone: "America/Los_Angeles",
                },
                Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
                Attendees: []*calendar.EventAttendee{
                        &calendar.EventAttendee{Email:"lpage@example.com"},
                        &calendar.EventAttendee{Email:"sbrin@example.com"},
                },
                }
        return event
}

func main() {
        ctx := context.Background()
        b, err := os.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
        if err != nil {
                log.Fatalf("Unable to retrieve Calendar client: %v", err)
        }
        
        // // event := createEvent("Testing",nil,"")
        // calendarId := "primary"
        // event, err = srv.Events.Insert(calendarId, event).Do()
        // if err != nil {
        // log.Fatalf("Unable to create event. %v\n", err)
        // }
        // fmt.Printf("Event created: %s\n", event.HtmlLink)
              
}
