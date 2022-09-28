package endpoints

import (
	"api/data"
	"encoding/json"
	"io/ioutil"
	ts "lib/test_setup"
	tu "lib/testutils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateRoleHandler(t *testing.T) {
	err := ts.ConnectDB(t)
	if err != nil {
		t.Fatal(err)
	}
	defer ts.DisconnectDB(t)

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
			CreateRoleHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("", "", "")}) // Make request
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

	//Good tests ================
	/*type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	goodTests := []struct {
		name    string
		request string
		args    args
		expect  goodExpect
	}{
		{
			name: "Valid role information",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/create`, strings.NewReader(`
					{
						"name" : "Test_Role"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    nil, // no response body for valid insecreatesrts
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateRoleHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("", "", "")}) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}*/

	// Further Bad Tests ================
	badTests := []struct {
		name    string
		request string
		args    args
		expect  badExpect
	}{
		{
			name: "Bad Data, null role name",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/create`, strings.NewReader(`
				{
					"role_name" : null

				}`)),
			},
			expect: badExpect{
				responseCode:    http.StatusInternalServerError,
				responseMessage: "",
			},
		},
		{
			name: "Bad Data, invalid role lead id",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/create`, strings.NewReader(`
				{
					"role_name" : "Test Role 2",
					"role_lead_id" : "inval"
				}`)),
			},
			expect: badExpect{
				responseCode:    http.StatusInternalServerError,
				responseMessage: "",
			},
		},
		{
			name: "Bad Data, duplicate role name",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/create`, strings.NewReader(`
				{
					"role_name" : "Test_Role"
				}`)),
			},
			expect: badExpect{
				responseCode:    http.StatusInternalServerError,
				responseMessage: "",
			},
		},
	}

	for _, tt := range badTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateRoleHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("", "", "")}) // Make request
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
			// if message, ok := _error.Error["message"]; ok {
			// 	if message != tt.expect.responseMessage {
			// 		t.Error(tu.Scolourf(tu.RED, "Incorrect message returned, expected '%s', got '%s'", tt.expect.responseMessage, message))
			// 	}
			// } else {
			// 	t.Error(tu.Scolourf(tu.RED, "Expected an error message, got none"))
			// }
		})
	}
}

func TestInformationRolesHandler(t *testing.T) {
	err := ts.ConnectDB(t)
	if err != nil {
		t.Fatal(err)
	}
	defer ts.DisconnectDB(t)
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/information`, strings.NewReader(`
				{
					role_name" : "Test Role"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/information`, strings.NewReader(`
				[
					{
    					"role_lead_id" : "11111111-1111-4a06-9983-8b374586e459"
					}
				]`)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
		{
			name: "Bad request, empty body",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/information`, strings.NewReader(``)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
	}

	for _, tt := range basicBadTests {
		t.Run(tt.name, func(t *testing.T) {
			InformationRolesHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("", "", "")}) // Make request
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

func TestDeleteRoleHandler(t *testing.T) {
	err := ts.ConnectDB(t)
	if err != nil {
		t.Fatal(err)
	}
	defer ts.DisconnectDB(t)

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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/remove`, strings.NewReader(`
				{
					role_name" : "Test Role"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/remove`, strings.NewReader(`
				[
					{
    					"role_lead_id" : "11111111-1111-4a06-9983-8b374586e459"
					}
				]`)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
		{
			name: "Bad request, empty body",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/role/remove`, strings.NewReader(``)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
	}

	for _, tt := range basicBadTests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteRoleHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("", "", "")}) // Make request
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

func TestCreateUserRoleHandler(t *testing.T) {

}

func TestInformationUserRolesHandler(t *testing.T) {

}

func TestDeleteUserRoleHandler(t *testing.T) {

}
