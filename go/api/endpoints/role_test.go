package endpoints

import (
	"api/data"
	"encoding/json"
	"io/ioutil"
	dtdb "lib/dockertest_db"
	tu "lib/testutils"
	ts "lib/test_setup"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateRoleHandler(t *testing.T) {
	testdb := ts.SetupTest(t)
	defer dtdb.StopTestDbWithTest(testdb, t, false)

	// ==================
	// Perform tests ====
	// ==================
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	type badExpect struct {
		responseCode    int
		responseMessage string
	}
	
	// Basic Bad tests ================
	basicBadTests := []struct {
		name    string
		request string
		args    args
		expect  badExpect
	}{
		{
			name: "Bad JSON, syntax error",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/create`, strings.NewReader(`
				{
					role_name : "Test Role"
    				role_lead_id : "11111111-1111-4a06-9983-8b374586e459"
				}`)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
		{
			name: "Bad JSON, array",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/create`, strings.NewReader(`
				[
					{
						"role_name" : "Test Role",
    					"role_lead_id" : "11111111-1111-4a06-9983-8b374586e459"
					}
				]`)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
	}

	for _, tt := range basicBadTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateRoleHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "BUILDING")}) // Make request
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
			var _error errorResponse
			err = json.Unmarshal(bodyBytes, &_error) // decode body
			if err != nil {
				t.Errorf(tu.Scolourf(tu.RED, "Invalid JSON, could not decode, err: %v", err))
				t.FailNow()
			}
			if message, ok := _error.Error["message"]; ok {
				if message != tt.expect.responseMessage {
					t.Error(tu.Scolourf(tu.RED, "Incorrect message returned, expected '%s', got '%s'", tt.expect.responseMessage, message))
				}
			} else {
				t.Error(tu.Scolourf(tu.RED, "Expected an error message, got none"))
			}
		})
	}
}
