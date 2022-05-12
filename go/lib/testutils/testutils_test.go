package testutils

import (
	"fmt"
	"strings"
	"testing"
)

func TestScolour(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		colour Colour
	}{
		{
			name:   "Red",
			str:    "This should be Red",
			colour: RED,
		},
		{
			name:   "Yellow, empty string",
			str:    "",
			colour: YELLOW,
		},
		{
			name:   "GREEN, string containing GREEN and RESET",
			str:    string(RESET) + "Reset then green" + string(GREEN),
			colour: GREEN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := Scolour(tt.colour, tt.str)
			if strings.Index(str, string(tt.colour)) != 0 {
				t.Errorf("Colour not inserted at correct index, expected at 0, found at %d", strings.Index(str, string(tt.colour)))
				t.FailNow()
			}
			if !strings.Contains(str, tt.str) {
				t.Errorf(`Expected to find "%s" in str, but not found`, tt.str)
			}
			if strings.LastIndex(str, string(RESET)) != len(str)-len(RESET) {
				t.Errorf("Content string not RESET correctly, expected RESET at index %d, but found at %d", len(str)-len(RESET), strings.LastIndex(str, string(RESET)))
				t.FailNow()
			}
		})
	}
}

func TestScolourf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
		colour Colour
	}{
		{
			name:   "Red, ints",
			format: "This should be Red, %d %f %s ",
			args:   []interface{}{1, 2.0, "three"},
			colour: RED,
		},
		{
			name:   "Yellow, empty string, no args",
			format: "",
			args:   []interface{}{},
			colour: YELLOW,
		},
		{
			name:   "GREEN, string containing GREEN and RESET, no args",
			format: string(RESET) + "Reset then green" + string(GREEN),
			args:   []interface{}{},
			colour: GREEN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := Scolourf(tt.colour, tt.format, tt.args...)
			if strings.Index(str, string(tt.colour)) != 0 {
				t.Errorf("Colour not inserted at correct index, expected at 0, found at %d", strings.Index(str, string(tt.colour)))
				t.FailNow()
			}
			if !strings.Contains(str, fmt.Sprintf(tt.format, tt.args...)) {
				t.Errorf(`Expected to find "%s" in str, but not found`+string(RESET), fmt.Sprintf(tt.format, tt.args...))
			}
			if strings.LastIndex(str, string(RESET)) != len(str)-len(RESET) {
				t.Errorf("Content string not RESET correctly, expected RESET at index %d, but found at %d", len(str)-len(RESET), strings.LastIndex(str, string(RESET)))
				t.FailNow()
			}
		})
	}
}
