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

func createBuilding(id string, name string, location string, dimension string) data.Building {
	var building data.Building
	building.Id = &id
	building.Name = &name
	building.Location = &location
	building.Dimension = &dimension
	return building
}

/*func createUserTeam(teamId string, userId string) data.UserTeam {
	var userTeam data.UserTeam
	userTeam.TeamId = &teamId
	userTeam.UserId = &userId
	return userTeam
}

func createTeamAssociation(teamId string, teamAssociationId string) data.TeamAssociation {
	var userTeam data.TeamAssociation
	userTeam.TeamId = &teamId
	userTeam.TeamIdAssociation = &teamAssociationId
	return userTeam
}*/

//Buildings
func TestInformationBuildingsHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/information`, strings.NewReader(`
				{
					id: null,
					name: "Building A",
					location: "Building A Location..."
					dimensions: "5x5"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/information`, strings.NewReader(`
				[
					{
						"id": null,
						"name": "Building A",
						"location": "Building A Location...",
						"dimensions": "5x5"
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
			InformationBuildingsHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "BUILDING")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		expect  goodExpect
	}{
		{
			name: "OK building information",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/information`, strings.NewReader(`
					{
						"id": null,
						"name": "Building A",
						"location": "Building A Location...",
						"dimensions": "5x5"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			InformationBuildingsHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "BUILDING")}) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestCreateBuildingHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
				{
					"id": null
					"name": "Building A"
					"location": "Building A Location..."
					"dimensions": "5x5"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
				[
					{
						"id": null,
						"name": "Building A",
						"location": "Building A Location...",
						"dimensions": "5x5"
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
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		expect  goodExpect
	}{
		{
			name: "OK building creation",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
					{
						"id": null,
						"name": "Building A",
						"location": "Building A Location",
						"dimension": "5x5"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestDeleteBuildingHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/remove`, strings.NewReader(`
				{
					"id": null
					"name": "Building A"
					"location": "Building A Location..."
					"dimensions": "5x5"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/remove`, strings.NewReader(`
				[
					{
						"id": null,
						"name": "Building A",
						"location": "Building A Location...",
						"dimensions": "5x5"
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
			DeleteBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "BUILDING")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		args2   args
		expect  goodExpect
	}{
		{
			name: "OK building deletion",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/remove`, strings.NewReader(`
					{
						"id": null
					}`)),
			},
			args2: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
				{
					"id": null,
					"name": "Building A",
					"location": "Building A Location",
					"dimension": "5x5"
				}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args2.w, tt.args2.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")})
			DeleteBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "BUILDING")}) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

//Rooms
func TestInformationRoomsHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/information`, strings.NewReader(`
				{
					"id": null
					"buildingId": null
					"name": "Room A"
					"location": "Room A Location..."
					"dimensions": "5x5"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/information`, strings.NewReader(`
				[
					{
						"id": null,
						"buildingId": null,
						"name": "Room A",
						"location": "Room A Location...",
						"dimensions": "5x5"
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
			InformationRoomsHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOM")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		expect  goodExpect
	}{
		{
			name: "OK room information",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/information`, strings.NewReader(`
					{
						"id": null
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			InformationRoomsHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOM")}) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestCreateRoomHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
				{
					"id": null
					"buildingId": null
					"name": "Room A"
					"location": "Room A Location..."
					"dimensions": "5x5"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
				[
					{
						"id": null,
						"buildingId": null,
						"name": "Room A",
						"location": "Room A Location...",
						"dimensions": "5x5"
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
			CreateRoomHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		args2   args
		expect  goodExpect
	}{
		{
			name: "OK room creation",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
					{
						"id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Building A",
						"location": "Building A Location",
						"dimension": "5x5"
					}`)),
			},
			args2: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": null,
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
			CreateRoomHandler(tt.args2.w, tt.args2.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestDeleteRoomHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/remove`, strings.NewReader(`
				{
					"id": null
					"buildingId": null
					"name": "Room A"
					"location": "Room A Location..."
					"dimensions": "5x5"
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/remove`, strings.NewReader(`
				[
					{
						"id": null,
						"buildingId": null,
						"name": "Room A",
						"location": "Room A Location...",
						"dimensions": "5x5"
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
			DeleteRoomHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "ROOM")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		args2   args
		args3   args
		expect  goodExpect
	}{
		{
			name: "OK room creation",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
					{
						"id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Building A",
						"location": "Building A Location",
						"dimension": "5x5"
					}`)),
			},
			args2: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": null,
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			args3: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/remove`, strings.NewReader(`
					{
						"id": null
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
			CreateRoomHandler(tt.args2.w, tt.args2.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			DeleteRoomHandler(tt.args3.w, tt.args3.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "ROOM")})
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestInformationIdentifiersHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/information`, strings.NewReader(`
				{
					"id": null
					"roomId": null
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/information`, strings.NewReader(`
				[
					{
						"id": null
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
			InformationIdentifiersHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "IDENTIFIER")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		expect  goodExpect
	}{
		{
			name: "OK resource information",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/information`, strings.NewReader(`
					{
						"id": null
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			InformationIdentifiersHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "IDENTIFIER")}) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestCreateIdentifierHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/create`, strings.NewReader(`
				{
					"id": null
					"roomId": null
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/create`, strings.NewReader(`
				[
					{
						"id": null
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
			CreateIdentifierHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "IDENTIFIER")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		args2   args
		args3   args
		expect  goodExpect
	}{
		{
			name: "OK resource creation",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
					{
						"id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Building A",
						"location": "Building A Location",
						"dimension": "5x5"
					}`)),
			},
			args2: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": "14141414-dc08-4a06-9983-8b374586e459",
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			args3: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/create`, strings.NewReader(`
					{
						"id": null,
						"room_id": "14141414-dc08-4a06-9983-8b374586e459",
						"name": "Desk 1",
						"xcoord": "10",
						"ycoord": "10",
						"width": "200",
						"height": "200",
						"rotation": "0",
						"role_id": null,
						"resource_type": "DESK",
						"decorations": "{\"computer\": true}"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
			CreateRoomHandler(tt.args2.w, tt.args2.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			CreateIdentifierHandler(tt.args3.w, tt.args3.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "IDENTIFIER")})
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestDeleteIdentifierHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/delete`, strings.NewReader(`
				{
					"id": null
					"roomId": null
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/delete`, strings.NewReader(`
				[
					{
						"id": null
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
			DeleteIdentifierHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "IDENTIFIER")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		args2   args
		args3   args
		args4   args
		expect  goodExpect
	}{
		{
			name: "OK resource creation",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
					{
						"id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Building A",
						"location": "Building A Location",
						"dimension": "5x5"
					}`)),
			},
			args2: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": "14141414-dc08-4a06-9983-8b374586e459",
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			args3: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/create`, strings.NewReader(`
					{
						"id": "22222222-dc08-4a06-9983-8b374586e459",
						"room_id": "14141414-dc08-4a06-9983-8b374586e459",
						"name": "Desk 1",
						"xcoord": "10",
						"ycoord": "10",
						"width": "200",
						"height": "200",
						"rotation": "0",
						"role_id": null,
						"resource_type": "DESK",
						"decorations": "{\"computer\": true}"
					}`)),
			},
			args4: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/delete`, strings.NewReader(`
					{
						"id": "22222222-dc08-4a06-9983-8b374586e459"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
			CreateRoomHandler(tt.args2.w, tt.args2.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			CreateIdentifierHandler(tt.args3.w, tt.args3.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "IDENTIFIER")})
			DeleteIdentifierHandler(tt.args4.w, tt.args4.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "IDENTIFIER")})
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

//Room Association
func TestInformationRoomAssociationsHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/information`, strings.NewReader(`
				{
					"id": null
					"roomId": null
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/information`, strings.NewReader(`
				[
					{
						"id": null,
						"roomId": null
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
			InformationRoomAssociationsHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOMASSOCIATION")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		expect  goodExpect
	}{
		{
			name: "OK resource information",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/information`, strings.NewReader(`
					{
						"room_id": null
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			InformationRoomAssociationsHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("VIEW", "RESOURCE", "ROOMASSOCIATION")}) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestCreateRoomAssociationHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/create`, strings.NewReader(`
				{
					"id": null
					"roomId": null
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/create`, strings.NewReader(`
				[
					{
						"id": null,
						"roomId": null
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
			CreateRoomAssociationHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOMASSOCIATION")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		args2   args
		args3   args
		args4   args
		expect  goodExpect
	}{
		{
			name: "OK resource creation",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
					{
						"id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Building A",
						"location": "Building A Location",
						"dimension": "5x5"
					}`)),
			},
			args2: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": "14141414-dc08-4a06-9983-8b374586e459",
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			args3: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": "15151515-dc08-4a06-9983-8b374586e459",
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			args4: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/create`, strings.NewReader(`
					{
						"room_id": "14141414-dc08-4a06-9983-8b374586e459",
						"room_id_association": "15151515-dc08-4a06-9983-8b374586e459"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
			CreateRoomHandler(tt.args2.w, tt.args2.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			CreateRoomHandler(tt.args3.w, tt.args3.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			CreateRoomAssociationHandler(tt.args4.w, tt.args4.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOMASSOCIATION")})
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}

func TestDeleteRoomAssociationHandler(t *testing.T) {
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resources/room/association/remove`, strings.NewReader(`
				{
					"id": null
					"roomId": null
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
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resources/room/association/remove`, strings.NewReader(`
				[
					{
						"id": null,
						"roomId": null
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
			DeleteRoomAssociationHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "ROOMASSOCIATION")}) // Make request
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

	// Good tests ================
	type goodExpect struct {
		responseCode    int
		responseBody    *string
		responseMessage string
	}

	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		args2   args
		args3   args
		args4   args
		args5   args
		expect  goodExpect
	}{
		{
			name: "OK resource creation",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/building/create`, strings.NewReader(`
					{
						"id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Building A",
						"location": "Building A Location",
						"dimension": "5x5"
					}`)),
			},
			args2: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": "14141414-dc08-4a06-9983-8b374586e459",
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			args3: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/create`, strings.NewReader(`
					{
						"id": "15151515-dc08-4a06-9983-8b374586e459",
						"building_id": "98989898-dc08-4a06-9983-8b374586e459",
						"name": "Room A",
						"location": "Room A Location",
						"dimension": "5x5"
					}`)),
			},
			args4: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/create`, strings.NewReader(`
					{
						"room_id": "14141414-dc08-4a06-9983-8b374586e459",
						"room_id_association": "15151515-dc08-4a06-9983-8b374586e459"
					}`)),
			},
			args5: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/resource/room/association/remove`, strings.NewReader(`
					{
						"room_id": "14141414-dc08-4a06-9983-8b374586e459",
						"room_id_association": "15151515-dc08-4a06-9983-8b374586e459"
					}`)),
			},
			expect: goodExpect{
				responseCode:    http.StatusOK,
				responseBody:    requestBodies[0], // do not use yet, API returns null
				responseMessage: "request_ok",
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBuildingHandler(tt.args.w, tt.args.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "BUILDING")}) // Make request
			CreateRoomHandler(tt.args2.w, tt.args2.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			CreateRoomHandler(tt.args3.w, tt.args3.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOM")})
			CreateRoomAssociationHandler(tt.args4.w, tt.args4.r, &data.Permissions{data.CreateGenericPermission("CREATE", "RESOURCE", "ROOMASSOCIATION")})
			DeleteRoomAssociationHandler(tt.args5.w, tt.args5.r, &data.Permissions{data.CreateGenericPermission("DELETE", "RESOURCE", "ROOMASSOCIATION")})
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
		})
	}
}
