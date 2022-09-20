package testutils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestMapsWithcSlicesMatchLoosely(t *testing.T) {
	type args struct {
		t          *testing.T
		map1       map[int][]int
		map2       map[int][]int
		msgAndArgs []interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantFailed bool
		msg        string
	}{
		{
			name: "Test 1",
			args: args{
				t: &testing.T{},
				map1: map[int][]int{
					1: {10},
				},
				map2: map[int][]int{
					1: {10},
				},
			},
			wantFailed: false,
			msg:        "Expected to not fail, maps identical",
		},
		{
			name: "Test 2",
			args: args{
				t: &testing.T{},
				map1: map[int][]int{
					1: {},
				},
				map2: map[int][]int{
					1: {},
				},
			},
			wantFailed: false,
			msg:        "Expected to not fail, maps identical",
		},
		{
			name: "Test 3",
			args: args{
				t:    &testing.T{},
				map1: map[int][]int{},
				map2: map[int][]int{},
			},
			wantFailed: false,
			msg:        "Expected to not fail, maps identical",
		},
		{
			name: "Test 4",
			args: args{
				t: &testing.T{},
				map1: map[int][]int{
					1: {10, 11, 12},
					2: {20, 21, 22},
				},
				map2: map[int][]int{
					2: {22, 20, 21},
					1: {12, 10, 11},
				},
			},
			wantFailed: false,
			msg:        "Expected to not fail, maps contain the same contents",
		},
		{
			name: "Test 5",
			args: args{
				t: &testing.T{},
				map1: map[int][]int{
					1: {10},
					2: {20},
				},
				map2: map[int][]int{
					2: {20},
				},
			},
			wantFailed: true,
			msg:        "Expected to fail, maps have different len",
		},
		{
			name: "Test 6",
			args: args{
				t: &testing.T{},
				map1: map[int][]int{
					1: {10},
					2: {20},
				},
				map2: map[int][]int{
					1:  {10},
					-2: {20},
				},
			},
			wantFailed: true,
			msg:        "Expected to fail, maps have different keys",
		},
		{
			name: "Test 7",
			args: args{
				t: &testing.T{},
				map1: map[int][]int{
					1: {10},
					2: {20},
				},
				map2: map[int][]int{
					1: {10},
					2: {20, 21},
				},
			},
			wantFailed: true,
			msg:        "Expected to fail, maps have different keys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapsWithcSlicesMatchLoosely(tt.args.t, tt.args.map1, tt.args.map2, tt.args.msgAndArgs...)
			assert.Equal(t, tt.wantFailed, tt.args.t.Failed(), tt.msg)
		})
	}
}
