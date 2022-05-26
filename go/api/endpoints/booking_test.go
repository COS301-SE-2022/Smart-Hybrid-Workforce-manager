package endpoints

import (
	"api/data"
	"api/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	migrate "lib/dbmigrate"
	dtdb "lib/dockertest_db"
	tu "lib/testutils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInformationBookingHandler(t *testing.T) {
	config := dtdb.TestDbConfig{
		Verbose:  true,
		HostPort: "5433",
		HostAdrr: "127.0.0.1",
	}
	testdb, err := dtdb.StartTestDb(config)
	defer dtdb.StopTestDbWithTest(testdb, t, false)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not start testDb, err: %v", err))
		t.FailNow()
	}

	migrator := migrate.AutoMigrate{
		MigratePath: "../../../db/sql",
		PathPatterns: []string{
			`*.schema.*sql`,
			`*.function.*sql`,
			`*mock.sql`,
		},
	}

	err = migrator.Migrate(testdb.Db)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not performed migration, err: %v", err))
	} else {
		fmt.Println(tu.Scolour(tu.GREEN, "Migration Performed"))
	}
	// Close after migration
	testdb.Db.Close()

	dbDsnEnv := testdb.Dsn
	// dbDsnEnv := "host=localhost port=5433 user=admin dbname=arche sslmode=disable"
	t.Setenv("DATABASE_DSN", dbDsnEnv)

	dbMaxIdleEnv := "5"
	t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", dbMaxIdleEnv)

	dbMaxOpenEnv := "5"
	t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", dbMaxOpenEnv)

	err = db.RegisterAccess()
	if err != nil {
		t.Logf(tu.Scolour(tu.PURPLE, "Error while connecting to DB: %v, skipping test"), err)
		t.FailNow()
	} else {
		fmt.Println(tu.Scolour(tu.GREEN, "DB connected"))
	}

	writer1 := httptest.NewRecorder()
	request1 := httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`{}`))
	InformationBookingHandler(writer1, request1, &data.Permissions{data.CreateGenericPermission("VIEW", "BOOKING", "USER")})
	response := writer1.Result()
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(tu.Scolour(tu.RED, "Error in response, could not read body, expected err to be nil, got %v"), err)
		t.FailNow()
	}
	var bookings []data.Booking
	// utils.UnmarshalJSON(writer1, request1, &bookings)
	// fmt.Println(bookings)
	err = json.Unmarshal(bodyBytes, &bookings)
	// fmt.Println(bookings)
	if err != nil {
		t.Errorf(tu.Scolour(tu.RED, "Error in response, invalid json, expected err to be nil, got %v"), err)
		t.FailNow()
	}

	for i, booking := range bookings {
		if booking.Id == nil {
			t.Errorf(tu.Scolour(tu.RED, "bookings[%d].Id is %v, expected a non-nil pointer string"), i, booking.Id)
			t.FailNow()
		}

		if *booking.Id == "" {
			t.Errorf(tu.Scolour(tu.RED, "*bookings[%d].Id is %s, expected a valid ID"), i, *booking.Id)
			t.FailNow()
		}

		if booking.UserId == nil {
			t.Errorf(tu.Scolour(tu.RED, "bookings[%d].UserId is %v, expected a non-nil pointer string"), i, booking.UserId)
			t.FailNow()
		}

		if *booking.UserId == "" {
			t.Errorf(tu.Scolour(tu.RED, "*bookings[%d].UserId is %s, expected a valid ID"), i, *booking.UserId)
			t.FailNow()
		}

		if booking.ResourceType == nil {
			t.Errorf(tu.Scolour(tu.RED, "bookings[%d].ResourceType is %v, expected a non-nil pointer string"), i, booking.ResourceType)
			t.FailNow()
		}

		if *booking.ResourceType == "" {
			t.Errorf(tu.Scolour(tu.RED, "*bookings[%d].ResourceType is %s, expected a valid ResourceType string"), i, *booking.ResourceType)
			t.FailNow()
		}

		if booking.Start == nil {
			t.Errorf(tu.Scolour(tu.RED, "bookings[%d].Start is %v, expected a non-nil pointer time"), i, booking.Start)
			t.FailNow()
		}

		if booking.End == nil {
			t.Errorf(tu.Scolour(tu.RED, "bookings[%d].End is %v, expected a non-nil pointer time"), i, booking.End)
			t.FailNow()
		}

		if booking.Booked == nil {
			t.Errorf(tu.Scolour(tu.RED, "bookings[%d].Booked is %v, expected a non-nil pointer boolean"), i, booking.Booked)
			t.FailNow()
		}

		if booking.DateCreated == nil {
			t.Errorf(tu.Scolour(tu.RED, "bookings[%d].DateCreated is %v, expected a non-nil pointer Date"), i, booking.DateCreated)
			t.FailNow()
		}
	}

}
