package endpoints

import (
	"api/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	ts "lib/test_setup"
	"lib/testutils"
	"lib/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateBookingHandler(t *testing.T) {
	t.SkipNow()
	err := ts.ConnectDB(t)
	if err != nil {
		t.Fatal(err)
	}
	defer ts.DisconnectDB(t)

	writer1 := httptest.NewRecorder()
	request1 := httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`{}`))
	if err != nil {
		t.Logf(testutils.Scolour(testutils.PURPLE, "Error while connecting to DB: %v, skipping test"), err)
		fmt.Printf(testutils.Scolour(testutils.PURPLE, "Error while connecting to DB: %v, skipping test"), err)
		t.SkipNow()
	}
	InformationBookingHandler(writer1, request1, nil)
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
