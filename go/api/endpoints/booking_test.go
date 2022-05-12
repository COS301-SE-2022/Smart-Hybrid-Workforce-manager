package endpoints

import (
	"api/data"
	"api/db"
	"api/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib/testutils"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCreateBookingHandler(t *testing.T) {
	dbDsnEnv := os.Getenv("DATABASE_DSN")
	dbMaxIdleEnv := os.Getenv("DATABASE_MAX_IDLE_CONNECTIONS")
	dbMaxOpenEnv := os.Getenv("DATABASE_MAX_OPENCONNECTIONS")

	// Check if all environment variables could be read
	if dbDsnEnv == "" {
		// DB must be running locally in order for this test to work
		// Set environment variables
		dbDsnEnv = "host=localhost port=5432 user=admin dbname=arche sslmode=disable"
		fmt.Println(testutils.Scolourf(testutils.CYAN, "DATABASE_DSN environmet var could not be read, using %s", dbDsnEnv))
		t.Setenv("DATABASE_DSN", dbDsnEnv)
	}

	if dbMaxIdleEnv == "" {
		// DB must be running locally in order for this test to work
		dbMaxIdleEnv = "5"
		fmt.Println(testutils.Scolourf(testutils.CYAN, "DATABASE_MAX_IDLE_CONNECTIONS environmet var could not be read, using %s", dbMaxIdleEnv))
		t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", "5")
	}

	if dbMaxOpenEnv == "" {
		dbMaxOpenEnv = "5"
		fmt.Println(testutils.Scolourf(testutils.CYAN, "DATABASE_MAX_OPEN_CONNECTIONS environmet var could not be read, using %s", dbMaxOpenEnv))
		t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", "5")
	}

	writer1 := httptest.NewRecorder()
	request1 := httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`{}`))
	err := db.RegisterAccess()
	if err != nil {
		t.Logf(testutils.Scolour(testutils.PURPLE, "Error while connecting to DB: %v, skipping test"), err)
		fmt.Printf(testutils.Scolour(testutils.PURPLE, "Error while connecting to DB: %v, skipping test"), err)
		t.SkipNow()
	}
	InformationBookingHandler(writer1, request1)
	response := writer1.Result()
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(testutils.Scolour(testutils.RED, "Error in response, could not read body, expected err to be nil, got %v"), err)
		t.FailNow()
	}
	var bookings []data.Booking
	utils.UnmarshalJSON(writer1, request1, bookings)
	err = json.Unmarshal(bodyBytes, &bookings)
	if err != nil {
		t.Errorf(testutils.Scolour(testutils.RED, "Error in response, invalid json, expected err to be nil, got %v"), err)
		t.FailNow()
	}

	for i, booking := range bookings {
		if booking.Id == nil {
			t.Errorf(testutils.Scolour(testutils.RED, "bookings[%d].Id is %v, expected a non-nil pointer string"), i, booking.Id)
			t.FailNow()
		}

		if *booking.Id == "" {
			t.Errorf(testutils.Scolour(testutils.RED, "*bookings[%d].Id is %s, expected a valid ID"), i, *booking.Id)
			t.FailNow()
		}

		if booking.UserId == nil {
			t.Errorf(testutils.Scolour(testutils.RED, "bookings[%d].UserId is %v, expected a non-nil pointer string"), i, booking.UserId)
			t.FailNow()
		}

		if *booking.UserId == "" {
			t.Errorf(testutils.Scolour(testutils.RED, "*bookings[%d].UserId is %s, expected a valid ID"), i, *booking.UserId)
			t.FailNow()
		}

		if booking.ResourceType == nil {
			t.Errorf(testutils.Scolour(testutils.RED, "bookings[%d].ResourceType is %v, expected a non-nil pointer string"), i, booking.ResourceType)
			t.FailNow()
		}

		if *booking.ResourceType == "" {
			t.Errorf(testutils.Scolour(testutils.RED, "*bookings[%d].ResourceType is %s, expected a valid ResourceType string"), i, *booking.ResourceType)
			t.FailNow()
		}

		if booking.Start == nil {
			t.Errorf(testutils.Scolour(testutils.RED, "bookings[%d].Start is %v, expected a non-nil pointer time"), i, booking.Start)
			t.FailNow()
		}

		if booking.End == nil {
			t.Errorf(testutils.Scolour(testutils.RED, "bookings[%d].End is %v, expected a non-nil pointer time"), i, booking.End)
			t.FailNow()
		}

		if booking.Booked == nil {
			t.Errorf(testutils.Scolour(testutils.RED, "bookings[%d].Booked is %v, expected a non-nil pointer boolean"), i, booking.Booked)
			t.FailNow()
		}

		if booking.DateCreated == nil {
			t.Errorf(testutils.Scolour(testutils.RED, "bookings[%d].DateCreated is %v, expected a non-nil pointer Date"), i, booking.DateCreated)
			t.FailNow()
		}
	}

}
