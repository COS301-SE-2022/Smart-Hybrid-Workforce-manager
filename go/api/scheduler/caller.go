package scheduler

import "time"

const LOG_PATH string = "////"

// datesEqual checks the equality of dates but ignoring time
func datesEqual(t1 time.Time, t2 time.Time) bool {
	// TODO: @JonathanEnslin make sure time zones aren't an issue
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

// mayCall determines using the passed in params, and the current time, whether or not the scheduler
// may be called
func mayCall(scheduledDay string, lastEntry LogEntry, now time.Time) bool {
	// check if correct day of week
	if scheduledDay != now.Weekday().String() {
		return false
	}
	// compare the dates, if a different date, scheduling may occur
	if !datesEqual(now, lastEntry.datetime) {
		return true
	}
	// otherwise, inspect the log entry to determine
	if lastEntry.status == FAILED || lastEntry.status == TIMED_OUT {
		return true
	}
	// It was either succesful or it is pending
	return false
}

func run() error {
	scheduledDay := "Friday" // TODO: @JonathanEnslin Make env var
	now := time.Now()
	lastEntry, err := ReadLastEntry(LOG_PATH)
	if err != nil {
		return err
	}
	if mayCall(scheduledDay, *lastEntry, now) {
		call()
	} else {
		// TODO: @JonathanEnslin Report error if pending, or exp backoff
	}
	return nil
}

func call() {
	// TODO: @JonathanEnslin Implement
}

func callPeriodically() {
	// TODO @JonathanEnslin
	// Call the scheduler on set intervals
}
