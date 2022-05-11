package testutils

import "fmt"

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
