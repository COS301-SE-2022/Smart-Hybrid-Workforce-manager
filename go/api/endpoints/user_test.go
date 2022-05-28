package endpoints

import (
	"api/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	dbm "lib/dbmigrate"
	dtdb "lib/dockertest_db"
	tu "lib/testutils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	// Taken from lib/utils, copied since it is not exported
	type errorResponse struct {
		Error map[string]interface{} `json:"error"`
	}

	// SETUP =============
	{
		config := dtdb.TestDbConfig{
			Verbose:  true,
			HostPort: "5433",
			HostAdrr: "127.0.0.1",
		}
		testdb, err := dtdb.StartTestDb(config) // Start DB
		defer dtdb.StopTestDbWithTest(testdb, t, false)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not start testDb, err: %v", err))
			t.FailNow()
		}
		// Perform migration
		migrator := dbm.AutoMigrate{
			MigratePath: "../../../db/sql",
			PathPatterns: []string{
				`*.schema.*sql`,   // schema files first
				`*.function.*sql`, // function files second
				// `*mock.sql`,    // will use own mock data
			},
		}
		migrator.Migrate(testdb.Db)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not performed migration, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.RED, "GOT HERE"))
		// update env vars for db connection
		t.Setenv("DATABASE_DSN", testdb.Dsn)

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
	}
	// ==================

	// Perform tests ====
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	type expect struct {
		responseCode    int
		responseMessage string
	}

	basicBadTests := []struct {
		name    string
		request string
		args    args
		expect  expect
	}{
		{
			name: "Bad JSON, syntax error",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`
				{
					first_name: "John",
					last_name: "Smith",
					email: "johnsmith@test.com",
					password : "password"
				}`)),
			},
			expect: expect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
		{
			name: "Bad JSON, array",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`
				[
					{
						"first_name": "anonymous",
						"last_name": "anonymous",
						"email": "asd@test.com",
						"password" : "password"
					}
				]`)),
			},
			expect: expect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
		{
			name: "Bad Email", // Regex might have to be tested more extensively
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`
				{
					"first_name": "anonymous",
					"last_name": "anonymous",
					"email": "asd@@test.com",
					"password" : "password"
				}`)),
			},
			expect: expect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_email",
			},
		},
	}

	for _, tt := range basicBadTests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterUserHandler(tt.args.w, tt.args.r) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Error(tu.Scolourf(tu.RED, "Could not read response body, err: %v", err))
				t.FailNow()
			}
			var error errorResponse
			err = json.Unmarshal(bodyBytes, &error) // decode body
			if err != nil {
				t.Errorf(tu.Scolourf(tu.RED, "Invalid JSON, could not decode, err: %v", err))
				t.FailNow()
			}
			if message, ok := error.Error["message"]; ok {
				if message != tt.expect.responseMessage {
					t.Error(tu.Scolourf(tu.RED, "Incorrect message returned, expected '%s', got '%s'", tt.expect.responseMessage, message))
				}
			} else {
				t.Error(tu.Scolourf(tu.RED, "Expected an error message, got none"))
			}
		})
	}

	// ==================
}
