package testutils

import (
	// "api/db"
	"fmt"
	"lib/collectionutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Colour string

// Used for test output formatting, if colours should be added
const (
	RESET  Colour = "\033[0m"
	RED    Colour = "\033[31m"
	GREEN  Colour = "\033[32m"
	YELLOW Colour = "\033[33m"
	BLUE   Colour = "\033[34m"
	PURPLE Colour = "\033[35m"
	CYAN   Colour = "\033[36m"
	GRAY   Colour = "\033[37m"
	WHITE  Colour = "\033[97m"
)

// Adds the colour to the string
func Scolour(colour Colour, msg string) string {
	return string(colour) + msg + string(RESET)
}

// Formats a string, and adds ansi escape sequences to colour the string
func Scolourf(colour Colour, format string, a ...interface{}) string {
	return fmt.Sprintf(Scolour(colour, format), a...)
}

// Returns a pointer to the value passed in, mainly used for passing in literals
func Ptr[T any](obj T) *T { return &obj }

// MapsWithcSlicesMatchLoosely function compares two maps that map two arrays, it checks that the keys of the
// maps are identical, and that the arrays mapped two match loosely, meaning the elements of the arrays match, but not necessarily the order
func MapsWithcSlicesMatchLoosely[K comparable, V any](t *testing.T, map1 map[K][]V, map2 map[K][]V, msgAndArgs ...interface{}) {
	if len(map1) != len(map2) {
		assert.Fail(t, "Len of maps should be same", msgAndArgs...)
	}

	for key, value := range map1 {
		if !collectionutils.MapHasKey(map2, key) {
			assert.Fail(t, "Maps should have matching keys", msgAndArgs...)
			return
		}
		assert.ElementsMatch(t, value, map2[key], msgAndArgs...)
		if t.Failed() {
			return
		}
	}
}
