package scheduler

import (
	"context"
	"time"
)

const LOG_PATH string = "////"

// datesEqual checks the equality of dates but ignoring time
func datesEqual(t1 time.Time, t2 time.Time) bool {
	// TODO: @JonathanEnslin make sure time zones aren't an issue
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

// mayCall determines using the passed in params, and the current time, whether or not the scheduler
// may be called
func mayCall(scheduledDay string, lastEntry *LogEntry, now time.Time) bool {
	// check if correct day of week
	if scheduledDay != now.Weekday().String() {
		return false
	}
	// compare the dates, if a different date, scheduling may occur
	if lastEntry == nil || !datesEqual(now, lastEntry.datetime) {
		return true
	}
	// otherwise, inspect the log entry to determine
	if lastEntry.status == FAILED || lastEntry.status == TIMED_OUT {
		return true
	}
	// It was either succesful or it is pending
	return false
}

// CheckAndCall will access the logs, and then call the scheduler if the info log
// entry permits it
func checkAndCall(scheduledDay string) error {
	// scheduledDay := "Friday" // TODO: @JonathanEnslin Make env var
	now := time.Now()
	lastEntry, err := ReadLastEntry(LOG_PATH)
	if err != nil {
		return err
	}
	if mayCall(scheduledDay, lastEntry, now) {
		call()
	} else {
		// TODO: @JonathanEnslin Report error if pending, or exp backoff
	}
	return nil
}

func call() {
	// TODO: @JonathanEnslin Implement
}

// callOnDay will call checkAndCall() on each recurring certain day of the week
func callOnDay(ctx context.Context, scheduledDay string) {
	// Initial call, for whenn the function initially gets called
	checkAndCall(scheduledDay)
	// periodic calls
	stopLoop := false
	for !stopLoop {
		timer := time.NewTimer(time.Until(time.Now())) // TODO: @JonathanEnslin Change stubbed until time
		select {
		case <-timer.C:
			checkAndCall(scheduledDay)
		case <-ctx.Done():
			stopLoop = true
		}
	}
}
