package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"lib/testutils"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	type testType struct {
		StrField   string  `json:"strfield"`
		FloatField float64 `json:"floatfield"`
		IntSlice   []int   `json:"intslice"`
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	goodTests := []struct {
		name   string
		args   args
		expect testType
	}{
		{
			name: "testType all fields filled",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/irrelevant/", strings.NewReader(`{"strfield":"This is a String Field.","floatfield":12.5,"intslice":[1,2,3]}`)),
			},
			expect: testType{
				StrField:   "This is a String Field.",
				FloatField: 12.5,
				IntSlice:   []int{1, 2, 3},
			},
		},
		{
			name: "testType all fields empty",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/irrelevant/", strings.NewReader(`{}`)),
			},
			expect: testType{
				StrField:   "",
				FloatField: 0.0,
				IntSlice:   []int{},
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			var testObj testType
			err := UnmarshalJSON(tt.args.w, tt.args.r, &testObj)
			// @todo - look into, close Request body
			if err != nil {
				t.Errorf("\033[31mExpected err to be nil, got %v instead\033[0m", err)
				t.FailNow()
			}
			if testObj.StrField != tt.expect.StrField {
				t.Errorf("\033[31mstring field incorrectly decoded, expected '%s', got '%s' instead\033[0m", tt.expect.StrField, testObj.StrField)
				t.FailNow()
			}
			if testObj.FloatField != tt.expect.FloatField {
				t.Errorf("\033[31mfloat field incorrectly decoded, expected %f, got %f instead\033[0m", tt.expect.FloatField, testObj.FloatField)
				t.FailNow()
			}
			if len(tt.expect.IntSlice) != len(testObj.IntSlice) {
				t.Errorf("\033[31mintSlice decoded incorrectly, expected slice of len %d, got slice of len %d\033[0m", len(tt.expect.IntSlice), len(testObj.IntSlice))
				t.FailNow()
			}
			for i, v := range tt.expect.IntSlice {
				if v != testObj.IntSlice[i] {
					t.Errorf("\033[31mintSlice decoded incorrectly, expected value %d at index %d, got value %d instead\033[0m", v, i, testObj.IntSlice[i])
					t.FailNow()
				}
			}
		})
	}

	badTests := []struct {
		name   string
		args   args
		expect testType
	}{
		{
			name: "Bad JSON",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/irrelevant/", strings.NewReader(`{{"strfield":"This is a String Field.","floatfield":12.5,"intslice":[1,2,3]]}`)),
			},
			expect: testType{},
		},
	}
	for _, tt := range badTests {
		t.Run(tt.name, func(t *testing.T) {
			var testObj testType
			err := UnmarshalJSON(tt.args.w, tt.args.r, testObj)
			// @todo - look into, close Request body
			if err == nil {
				t.Errorf("\033[31mExpected err to be not be nil, got %v instead\033[0m", err)
				t.FailNow()
			}
		})
	}
}

func Test_jsonResponse(t *testing.T) {
	type testType struct {
		StrField   string  `json:"strfield"`
		FloatField float64 `json:"floatfield"`
		IntSlice   []int   `json:"intslice"`
	}
	type args struct {
		w       *httptest.ResponseRecorder
		r       *http.Request
		payload interface{}
	}

	type expect struct {
		statusCode int
	}

	tests := []struct {
		name   string
		args   args
		expect expect
	}{
		{
			name: "Valid Payload",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/irrelevant/", nil),
				payload: testType{
					StrField:   "This is a String Field.",
					FloatField: 12.5,
					IntSlice:   []int{1, 2, 3},
				},
			},
			expect: expect{
				statusCode: 200,
			},
		},
		{
			name: "Empty Payload",
			args: args{
				w:       httptest.NewRecorder(),
				r:       httptest.NewRequest(http.MethodGet, "/irrelevant/", nil),
				payload: testType{},
			},
			expect: expect{
				statusCode: 200,
			},
		},
		{
			name: "Status 400",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/irrelevant/", nil),
				payload: testType{
					StrField:   "This is a String Field.",
					FloatField: 12.5,
					IntSlice:   []int{1, 2, 3},
				},
			},
			expect: expect{
				statusCode: 400,
			},
		},
		{
			name: "Single value JSON",
			args: args{
				w:       httptest.NewRecorder(),
				r:       httptest.NewRequest(http.MethodGet, "/irrelevant/", nil),
				payload: 123,
			},
			expect: expect{
				statusCode: 400,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := tt.args.w
			jsonResponse(writer, tt.args.r, tt.expect.statusCode, tt.args.payload)
			resp := writer.Result()
			defer resp.Body.Close()
			if resp.StatusCode != tt.expect.statusCode {
				t.Errorf("\033[31mExpected HTTP status code %d, got %d instead\033[0m", resp.StatusCode, tt.expect.statusCode)
				t.FailNow()
			}
			// 'Content-Type' hardcoded since it will always be JSON as per the jsonResponse function
			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("\033[31mExpected Content-type of 'application/json', got %s instead\033[0m", resp.Header.Get("Content-type"))
				t.FailNow()
			}
			// Test contents, use UnmarshallJSON to test correctness
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("\033[31mError reading response body, expected err to be %v, got %v instead\033[0m", nil, err)
				t.FailNow()
			}
			body := string(bodyBytes)
			expectedBodyBytes, _ := json.Marshal(tt.args.payload)
			expectedBody := string(expectedBodyBytes) + "\n" // added since json.Encoder delimits with '\n'
			if body != expectedBody {
				t.Errorf("\033[31mPayload not marshalled correctly, expected to payload to be %s, got %s instead\033[0m", expectedBody, body)
				t.FailNow()
			}
		})
	}

	badTests := []struct {
		name   string
		args   args
		expect expect
	}{
		{
			name: "Invalid value Payload",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/irrelevant/", nil),
				// payload: math.Inf,
				payload: math.Inf,
			},
			expect: expect{
				statusCode: 200,
			},
		},
	}

	for _, tt := range badTests {
		t.Run(tt.name, func(t *testing.T) {
			writer := tt.args.w
			jsonResponse(writer, tt.args.r, tt.expect.statusCode, tt.args.payload)
			resp := writer.Result()
			defer resp.Body.Close()
			if resp.StatusCode != tt.expect.statusCode {
				t.Errorf("\033[31mExpected HTTP status code %d, got %d instead\033[0m", resp.StatusCode, tt.expect.statusCode)
				t.FailNow()
			}
			// 'Content-Type' hardcoded since it will always be JSON as per the jsonResponse function
			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("\033[31mExpected Content-type of 'application/json', got %s instead\033[0m", resp.Header.Get("Content-type"))
				t.FailNow()
			}
			respBytes, _ := ioutil.ReadAll(resp.Body)
			respString := string(respBytes)
			if respString != "" {
				t.Errorf("Expected response body to be \"\", got %s instead", respString)
				t.FailNow()
			}
			// Update at some point to catch log output
		})
	}
}

func TestBadRequest(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/somepath/", strings.NewReader("stringbody"))
	message := "bad request made"
	BadRequest(writer, request, message)

	response := writer.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf(testutils.Scolour(testutils.RED, "Expected status code %d, got %d instead"), http.StatusBadRequest, response.StatusCode)
		t.FailNow()
	}
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(testutils.Scolour(testutils.RED, "Expected err to be nil but got %v instead"), err)
	}
	var errResponse errorResponse
	err = json.Unmarshal(dataBytes, &errResponse)
	if err != nil {
		t.Errorf(testutils.Scolour(testutils.RED, "Error unmarshalling response body, expected err to nil, got %v instead"), err)
		t.FailNow()
	}
	if status, ok := errResponse.Error["status"]; ok {
		statusInt := int(status.(float64))
		if statusInt != http.StatusBadRequest {
			t.Errorf(testutils.Scolour(testutils.RED, "Expected status to be %v (%T), got %v (%T) instead"), http.StatusBadRequest, http.StatusBadRequest, statusInt, statusInt)
			t.FailNow()
		}
	} else {
		t.Errorf(testutils.Scolour(testutils.RED, "Status not set, expected status to be %v"), http.StatusBadRequest)
		t.FailNow()
	}

	if msg, ok := errResponse.Error["message"]; ok {
		if msg != message {
			t.Errorf(testutils.Scolour(testutils.RED, "Expected message to be '%s' (%T), got '%s' (%T) instead"), message, message, msg, msg)
			t.FailNow()
		}
	} else {
		t.Errorf(testutils.Scolour(testutils.RED, "Message not set, expected message to be '%s'"), message)
		t.FailNow()
	}
}

func TestInternalServerError(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/somepath/", strings.NewReader("stringbody"))
	message := errors.New("Mock Error")
	InternalServerError(writer, request, message)
	response := writer.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf(testutils.Scolour(testutils.RED, "Expected status code %d, got %d instead"), http.StatusInternalServerError, response.StatusCode)
		t.FailNow()
	}
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(testutils.Scolour(testutils.RED, "Expected err to be nil but got %v instead"), err)
	}
	var errResponse errorResponse
	err = json.Unmarshal(dataBytes, &errResponse)
	if err != nil {
		t.Errorf(testutils.Scolour(testutils.RED, "Error unmarshalling response body, expected err to nil, got %v instead"), err)
		t.FailNow()
	}
	if status, ok := errResponse.Error["status"]; ok {
		statusInt := int(status.(float64))
		if statusInt != http.StatusInternalServerError {
			t.Errorf(testutils.Scolour(testutils.RED, "Expected status to be %v (%T), got %v (%T) instead"), http.StatusInternalServerError, http.StatusInternalServerError, statusInt, statusInt)
			t.FailNow()
		}
	} else {
		t.Errorf(testutils.Scolour(testutils.RED, "Status not set, expected status to be %v"), http.StatusInternalServerError)
		t.FailNow()
	}
	if msg, ok := errResponse.Error["message"]; ok {
		if msg != message {
			t.Errorf(testutils.Scolour(testutils.RED, "Expected message to be '%v' (%T), got '%v' (%T) instead"), message, message, msg, msg)
			t.FailNow()
		}
	} else {
		t.Errorf(testutils.Scolour(testutils.RED, "Message not set, expected message to be '%v'"), message)
		t.FailNow()
	}
}

func TestOk(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/somepath/", nil)
	Ok(writer, request)
	resp := writer.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf(testutils.Scolour(testutils.RED, "Expected response status code to be %d but got %d instead"), http.StatusOK, resp.StatusCode)
	}
}
