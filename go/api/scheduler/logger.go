package scheduler

import (
	"fmt"
	"time"
)

type Status string

const DT_FMT string = "02-01-2006 15:04:05 Monday"

const (
	SUCCESS   Status = "SUCCESS"
	PENDING   Status = "PENDING"
	FAILED    Status = "FAILED"
	TIMED_OUT Status = "TIMED_OUT"
)

type LogEntry struct {
	datetime time.Time
	status   Status
}

func NewLogEntry(status Status, datetime *time.Time) LogEntry {
	now := time.Now() // if no datetime was passed current time will be used
	if datetime == nil {
		datetime = &now
	}
	return LogEntry{
		status:   status,
		datetime: *datetime,
	}
}

func (entry LogEntry) String() string {
	return fmt.Sprintf("[%s] %s", entry.status, entry.datetime.Format(DT_FMT))
}
