package scheduler

import (
	"context"
	"time"
)

const LOG_PATH string = "////"

var DaysOfWeek = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
}

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

// timeOfNextWeekDay returns the date/time of the next 'weekday'
func timeOfNextWeekDay(now time.Time, weekday string) time.Time {
	day := int(DaysOfWeek[weekday])
	currentDay := int(now.Weekday())
	daysUntil := int((day - currentDay + 7) % 7) // +7 Is to ensure that the firs part of the expr is always >= 0
	y, m, d := now.AddDate(0, 0, daysUntil).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
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

// callOnDay will call checkAndCall() on each recurring certain day of the week,
// the method can be cancelled using the passed in context
func callOnDay(ctx context.Context, scheduledDay string) {
	// Initial call, for whenn the function initially gets called
	checkAndCall(scheduledDay)
	// periodic calls
	stopLoop := false
	for !stopLoop {
		nextDay := timeOfNextWeekDay(time.Now(), scheduledDay) // TODO: @JonathanEnslin, allow scheduled day to be changed
		timer := time.NewTimer(time.Until(nextDay))
		defer timer.Stop()
		select {
		case <-timer.C:
			checkAndCall(scheduledDay)
		case <-ctx.Done():
			stopLoop = true
		}
	}
}
