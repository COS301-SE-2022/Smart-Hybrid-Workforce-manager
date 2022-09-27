package data

import (
	"testing"
	"time"
)

func TestTimeInIntervalInclusive(t *testing.T) {
	type args struct {
		check time.Time
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				check: time.Date(2022, 9, 27, 11, 30, 00, 00, time.UTC),
				start: time.Date(2022, 9, 27, 11, 20, 00, 00, time.UTC),
				end:   time.Date(2022, 9, 27, 11, 40, 00, 00, time.UTC),
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				check: time.Date(2022, 9, 27, 11, 20, 00, 00, time.UTC),
				start: time.Date(2022, 9, 27, 11, 20, 00, 00, time.UTC),
				end:   time.Date(2022, 9, 27, 11, 20, 00, 00, time.UTC),
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				check: time.Date(2022, 9, 27, 11, 30, 00, 00, time.UTC),
				start: time.Date(2022, 9, 27, 11, 30, 00, 00, time.UTC),
				end:   time.Date(2022, 9, 27, 11, 40, 00, 00, time.UTC),
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				check: time.Date(2022, 9, 27, 11, 30, 00, 00, time.UTC),
				start: time.Date(2022, 9, 27, 11, 20, 00, 00, time.UTC),
				end:   time.Date(2022, 9, 27, 11, 30, 00, 00, time.UTC),
			},
			want: true,
		},
		{
			name: "Test 5",
			args: args{
				check: time.Date(2022, 9, 27, 11, 50, 00, 00, time.UTC),
				start: time.Date(2022, 9, 27, 11, 20, 00, 00, time.UTC),
				end:   time.Date(2022, 9, 27, 11, 30, 00, 00, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeInIntervalInclusive(tt.args.check, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("TimeInIntervalInclusive() = %v, want %v", got, tt.want)
			}
		})
	}
}
