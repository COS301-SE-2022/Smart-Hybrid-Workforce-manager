package scheduler

import (
	"bufio"
	"fmt"
	"os"
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

func (entry LogEntry) WriteLog(path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640) // rw-r-----
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(entry.String() + "\n")
	return err
}

func (entry LogEntry) ReadLastEntry(path string) (string, error) {
	// create the file if it does not yet exist
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0640) // rw-r-----
	if err != nil {
		return "", err
	}
	defer f.Close()

	var lastLine string // only last line should be returned
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return lastLine, nil
}
