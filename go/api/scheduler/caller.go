package scheduler

import (
	"bytes"
	"context"
	"encoding/json"
	"lib/clock"
	"lib/logger"
	"lib/restclient"
	"lib/testutils"
	"net/http"
	"os"
	"time"
)

var (
	HTTPClient  restclient.HTTPClient
	Clock       clock.Clock   = &clock.RealClock{} // TODO: @JonathanEnslin make sure this is not a bad way of doing it
	timeout     time.Duration = 30 * time.Second
	endpointURL string
)

func init() {
	// get env vars here
	endpointURL = os.Getenv("SCHEDULER_ADDR")

	HTTPClient = &http.Client{
		Timeout: timeout,
	}
}

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

// TimeOfNextWeekDay returns the date/time of the next 'weekday'
func TimeOfNextWeekDay(now time.Time, weekday string) time.Time {
	day := int(DaysOfWeek[weekday])
	currentDay := int(now.Weekday())
	daysUntil := int((day - currentDay + 7) % 7) // +7 Is to ensure that the firs part of the expr is always >= 0
	if daysUntil == 0 {
		daysUntil = 7
	}
	y, m, d := now.AddDate(0, 0, daysUntil).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}

// CheckAndCall will access the logs, and then call the scheduler if the info log
// entry permits it
func checkAndCall(now time.Time, scheduledDay string) error {
	// scheduledDay := "Friday" // TODO: @JonathanEnslin Make env var
	lastEntry, err := ReadLastEntry()
	if err != nil {
		return err
	}
	if mayCall(scheduledDay, lastEntry, now) {
		// TODO: @JonathanEnslin move this into a seperate function that uses exponential backoff
		nextMonday := TimeOfNextWeekDay(now, "Monday")            // Start of next week
		nextSaturday := TimeOfNextWeekDay(nextMonday, "Saturday") // End of next work-week
		schedulerData, err := GetSchedulerData(nextMonday, nextSaturday)
		if err != nil {
			log_write_err := NewLogEntry(FAILED, &now).WriteLog()
			if log_write_err != nil {
				return log_write_err
			}
			return err
		}
		err = Call(schedulerData) // TODO: @JonathanEnslin handle the return data
		if err != nil {
			return err
		}
	}
	return nil
}

// Call Calls the scheduler endpoint and passes it the data
func Call(data *SchedulerData) error { // TODO: @JonathanEnslin improve this function
	body, _ := json.Marshal(data)
	bodyBytesBuff := bytes.NewBuffer(body)

	request, err := http.NewRequest(http.MethodPost, endpointURL, bodyBytesBuff)
	request.Header.Set("Content-Type", "application/json")
	// TODO: @JonathanEnslin determine request headers
	if err != nil {
		return err
	}
	response, err := HTTPClient.Do(request) // TODO: @JonathanEnslin URL env param
	logger.Info.Println(testutils.Scolour(testutils.GREEN, "HERE1"))
	now := Clock.Now()
	if err != nil {
		errType := FAILED
		if os.IsTimeout(err) {
			errType = TIMED_OUT // TODO: @JonathanEnslin look at implementing type of exp backoff for timeout
		}
		logErr := NewLogEntry(errType, &now).WriteLog()
		if logErr != nil {
			return logErr
		}
		return err
	}
	logger.Info.Println(testutils.Scolour(testutils.GREEN, "HERE2"))
	defer response.Body.Close()
	candidateBookings := &CandidateBookings{}
	err = json.NewDecoder(response.Body).Decode(candidateBookings)
	logger.Info.Println(testutils.Scolour(testutils.GREEN, "HERE3"))
	if err != nil {
		logErr := NewLogEntry(FAILED, &now).WriteLog()
		if logErr != nil {
			return logErr
		}
		return err
	}
	err = makeBookings(*candidateBookings)
	logger.Info.Println(testutils.Scolour(testutils.GREEN, "HERE4"))
	if err != nil {
		logErr := NewLogEntry(FAILED, &now).WriteLog()
		if logErr != nil {
			return logErr
		}
		return err
	}
	err = NewLogEntry(SUCCESS, &now).WriteLog()
	if err != nil { // TODO: IMPORTANT! @JonathanEnslin consider panicking when logger fails
		return err
	}
	return nil
}

// callOnDay will call checkAndCall() on each recurring certain day of the week,
// the method can be cancelled using the passed in context
func callOnDay(ctx context.Context, scheduledDay string) {
	// Initial call, for when the function initially gets called
	_ = checkAndCall(time.Now(), scheduledDay)
	// periodic calls
	stopLoop := false
	for !stopLoop {
		nextDay := TimeOfNextWeekDay(time.Now(), scheduledDay) // TODO: @JonathanEnslin, allow scheduled day to be changed
		timer := time.NewTimer(time.Until(nextDay))
		defer timer.Stop()
		select {
		case <-timer.C:
			_ = checkAndCall(time.Now(), scheduledDay)
		case <-ctx.Done():
			stopLoop = true
		}
	}
}
