package data

import (
	"reflect"
	"testing"
)

func TestExtractAvailableDeskIds(t *testing.T) {
	sptr := func(s string) *string {
		return &s
	}
	type args struct {
		schedulerData *SchedulerData
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Normal extraction",
			args: args{
				schedulerData: &SchedulerData{
					Resources: Resources{
						&Resource{Id: sptr("Apple"), ResourceType: sptr("DESK")},
						&Resource{Id: sptr("Litchi"), ResourceType: sptr("DESK")},
						&Resource{Id: sptr("Pear"), ResourceType: sptr("DESK")},
						&Resource{Id: sptr("Banana"), ResourceType: sptr("DESK")},
						&Resource{Id: sptr("Coconut"), ResourceType: sptr("DESK")},
						&Resource{Id: sptr("Cabbage"), ResourceType: sptr("MEETINGROOM")},
					},
					CurrentBookings: &Bookings{
						&Booking{ResourceId: sptr("Apple")},
						&Booking{ResourceId: sptr("Pear")},
						&Booking{ResourceId: sptr("Coconut")},
					},
				},
			},
			want: []string{"Litchi", "Banana"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractAvailableDeskIds(tt.args.schedulerData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractAvailableDeskIds() = %v, want %v", got, tt.want)
			}
		})
	}
}
