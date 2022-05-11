package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testType struct {
	StrField   string  `json:"strfield"`
	FloatField float64 `json:"floatfield"`
	IntSlice   []int   `json:"intslice"`
}

func TestUnmarshalJSON(t *testing.T) {
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
