package scheduler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const SEP string = " ~ "

type Status string

const DT_FMT string = "02-01-2006 15:04:05 Monday"

var (
	logPath string = "scheduler.log"
)

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

// NewLogEntry Generates a LogEntry struct, if datetime is nil, Now is used
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

// Generates string representation
func (entry LogEntry) String() string {
	return fmt.Sprintf("%s%s%s", entry.status, SEP, entry.datetime.Format(DT_FMT))
}

// Parse a string into a struct, representation returned by String is used
func Parse(str string) (*LogEntry, error) {
	parts := strings.Split(str, SEP)
	status := parts[0]
	dateStr := parts[1]
	date, err := time.Parse(DT_FMT, dateStr)
	if err != nil {
		return nil, err
	}
	entry := NewLogEntry(Status(status), &date)
	return &entry, nil
}

// WriteLog Writes the entry to the passed file
func (entry LogEntry) WriteLog() error {
	// func (entry LogEntry) WriteLog(path string) error {
	// TODO: @JonathanEnslin investigate why 0644 perms aren't allowed
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600) // rw-------
	if err != nil {
		return err
	}
	_, err = f.WriteString(entry.String() + "\n")
	_ = f.Close() // TODO: @JonathanEnslin handle error
	return err
}

// ReadLastEntry reads the last entry from a log file and returns the entry, or nil
// if file is empty
func ReadLastEntry() (*LogEntry, error) {
	// create the file if it does not yet exist
	// TODO: @JonathanEnslin investigate why 0644 perms aren't allowed
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDONLY, 0600) // rw-------
	if err != nil {
		return nil, err
	}

	var lastLine string // only last line should be returned
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if lastLine == "" {
		return nil, nil
	}
	_ = f.Close() // TODO: @JonathanEnslin handle error
	return Parse(lastLine)
}
