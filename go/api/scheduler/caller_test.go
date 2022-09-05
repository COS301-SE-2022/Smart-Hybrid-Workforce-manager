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

func Test_mayCall(t *testing.T) {
	// testTime1 := time.Now() // Current time
	testTime3 := time.Date(2022, time.May, 10, 5, 3, 2, 1, time.UTC) // Same day as testTime2, different time

	type args struct {
		scheduledDay string
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
				now:          testTime3,
			},
			want: false,
		},
		{
			name: "Correct day, should be allowed",
			args: args{
				scheduledDay: testTime3.Weekday().String(),
				now:          testTime3,
			},
			want: true,
		},
	}
	assert := assert.New(t)
	for i, tt := range tests {
		assert.Equal(tt.want, mayCall(tt.args.scheduledDay, tt.args.now), tt.name+" Case: "+fmt.Sprint(i))
	}
}

func Test_timeOfNextWeekDay(t *testing.T) {
	type args struct {
		now     time.Time
		weekday string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Thursday 2022 21 July 01:02:03.004, next Thursday should be 2022 28 July 00:00:00.000...",
			args: args{
				now:     time.Date(2022, 7, 21, 1, 2, 3, 4, time.UTC),
				weekday: "Thursday",
			},
			want: time.Date(2022, 7, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Thursday 2022 21 July 01:02:03.004, next Friday should be 2022 22 July 00:00:00.000...",
			args: args{
				now:     time.Date(2022, 7, 21, 1, 2, 3, 4, time.UTC),
				weekday: "Friday",
			},
			want: time.Date(2022, 7, 22, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Thursday 2022 21 July 01:02:03.004, next Saturday should be 2022 23 July 00:00:00.000...",
			args: args{
				now:     time.Date(2022, 7, 21, 1, 2, 3, 4, time.UTC),
				weekday: "Saturday",
			},
			want: time.Date(2022, 7, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Thursday 2022 21 July 01:02:03.004, next Sunday should be 2022 24 July 00:00:00.000...",
			args: args{
				now:     time.Date(2022, 7, 21, 1, 2, 3, 4, time.UTC),
				weekday: "Sunday",
			},
			want: time.Date(2022, 7, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Thursday 2022 21 July 01:02:03.004, next Monday should be 2022 25 July 00:00:00.000...",
			args: args{
				now:     time.Date(2022, 7, 21, 1, 2, 3, 4, time.UTC),
				weekday: "Monday",
			},
			want: time.Date(2022, 7, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Thursday 2022 21 July 01:02:03.004, next Tuesday should be 2022 26 July 00:00:00.000...",
			args: args{
				now:     time.Date(2022, 7, 21, 1, 2, 3, 4, time.UTC),
				weekday: "Tuesday",
			},
			want: time.Date(2022, 7, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Thursday 2022 21 July 01:02:03.004, next Wednesday should be 2022 27 July 00:00:00.000...",
			args: args{
				now:     time.Date(2022, 7, 21, 1, 2, 3, 4, time.UTC),
				weekday: "Wednesday",
			},
			want: time.Date(2022, 7, 27, 0, 0, 0, 0, time.UTC),
		},
	}
	assert := assert.New(t)
	for i, tt := range tests {
		assert.Equal(tt.want, TimeOfNextWeekDay(tt.args.now, tt.args.weekday), tt.name+" Case: "+fmt.Sprint(i))
	}
}
