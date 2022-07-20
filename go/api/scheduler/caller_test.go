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
		assert.Equal(datesEqual(tt.args.t1, tt.args.t2), tt.want, tt.name+"Case: "+fmt.Sprint(i))
	}
}
