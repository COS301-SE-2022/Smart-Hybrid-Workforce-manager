package clock

import "time"

// The clock interface is used to make code more testable by allowing the Now function to be mocked

type Clock interface {
	Now() time.Time
}

type RealClock struct {
}

func (c *RealClock) Now() time.Time {
	return time.Now()
}
