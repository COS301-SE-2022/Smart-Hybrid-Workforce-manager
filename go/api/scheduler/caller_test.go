package scheduler

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_datesEqual(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Dates are equal, expected true",
			args: args{
				t1: time.Date(2022, time.May, 10, 2, 3, 4, 5, time.UTC),
				t2: time.Date(2022, time.May, 10, 2, 3, 4, 5, time.UTC),
			},
			want: true,
		},
		{
			name: "Dates are equal, expected true",
			args: args{
				t1: time.Date(2022, time.May, 10, 2, 3, 4, 5, time.UTC),
				t2: time.Date(2022, time.May, 10, 4, 3, 2, 1, time.UTC),
			},
			want: true,
		},
		{
			name: "Months are not equal, expected false",
			args: args{
				t1: time.Date(2022, time.June, 10, 2, 3, 4, 5, time.UTC),
				t2: time.Date(2022, time.May, 10, 4, 3, 2, 1, time.UTC),
			},
			want: false,
		},
		{
			name: "Years are not equal, expected false",
			args: args{
				t1: time.Date(2021, time.May, 10, 2, 3, 4, 5, time.UTC),
				t2: time.Date(2022, time.May, 10, 4, 3, 2, 1, time.UTC),
			},
			want: false,
		},
		{
			name: "Days are not equal, expected false",
			args: args{
				t1: time.Date(2022, time.May, 20, 2, 3, 4, 5, time.UTC),
				t2: time.Date(2022, time.May, 10, 4, 3, 2, 1, time.UTC),
			},
			want: false,
		},
	}
	assert := assert.New(t)
	for i, tt := range tests {
		assert.Equal(tt.want, datesEqual(tt.args.t1, tt.args.t2), tt.name+" Case: "+fmt.Sprint(i))
	}
}

func makeLogEntryPtr(entry LogEntry) *LogEntry {
	return &entry
}

func Test_mayCall(t *testing.T) {
	// testTime1 := time.Now() // Current time
	testTime2 := time.Date(2022, time.May, 10, 4, 3, 2, 1, time.UTC)
	testTime3 := time.Date(2022, time.May, 10, 5, 3, 2, 1, time.UTC) // Same day as testTime2, different time

	type args struct {
		scheduledDay string
		lastEntry    *LogEntry
		now          time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Called on wrong day, not scheduled day",
			args: args{
				scheduledDay: testTime3.AddDate(0, 0, 1).Weekday().String(),
				lastEntry:    nil,
				now:          testTime3,
			},
			want: false,
		},
		{
			name: "Last entry is success on same day, should not be allowed",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				lastEntry:    makeLogEntryPtr(NewLogEntry(SUCCESS, &testTime2)),
				now:          testTime3,
			},
			want: false,
		},
		{
			name: "Correct day, no log entry, should be allowed",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				lastEntry:    nil,
				now:          testTime3,
			},
			want: true,
		},
		{
			name: "Correct day, FAILURE log entry on same day",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				lastEntry:    makeLogEntryPtr(NewLogEntry(FAILED, &testTime3)),
				now:          testTime3,
			},
			want: true,
		},
		{
			name: "Correct day, FAILED log entry on different day",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				lastEntry:    makeLogEntryPtr(NewLogEntry(FAILED, &testTime3)),
				now:          testTime3.AddDate(0, 0, 7),
			},
			want: true,
		},
		{
			name: "Correct day, TIMED_OUT log entry on same day",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				lastEntry:    makeLogEntryPtr(NewLogEntry(TIMED_OUT, &testTime3)),
				now:          testTime3,
			},
			want: true,
		},
		{
			name: "Correct day, SUCCESS log entry on different day",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				lastEntry:    makeLogEntryPtr(NewLogEntry(SUCCESS, &testTime3)),
				now:          testTime3.AddDate(0, 0, 7),
			},
			want: true,
		}, {
			name: "Correct day, SUCCESS log entry on different day, no passed in time",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				lastEntry:    makeLogEntryPtr(NewLogEntry(SUCCESS, nil)),
				now:          testTime3.AddDate(0, 0, 7),
			},
			want: true,
		},
	}
	assert := assert.New(t)
	for i, tt := range tests {
		assert.Equal(tt.want, mayCall(tt.args.scheduledDay, tt.args.lastEntry, tt.args.now), tt.name+" Case: "+fmt.Sprint(i))
	}
}
