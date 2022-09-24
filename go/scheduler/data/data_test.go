package data

import (
	tu "lib/testutils"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestExtractInverseUserIdMap(t *testing.T) {
	type args struct {
		userIdMap map[int]string
	}
	tests := []struct {
		name string
		args args
		want map[string][]int
	}{
		{
			name: "Test 1",
			args: args{
				userIdMap: map[int]string{
					1:  "Apple",
					2:  "Pear",
					30: "Coconut",
					40: "Apple",
				},
			},
			want: map[string][]int{
				"Apple":   {1, 40},
				"Pear":    {2},
				"Coconut": {30},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractInverseUserIdMap(tt.args.userIdMap)
			tu.MapsWithcSlicesMatchLoosely(t, got, tt.want, "ExtractInverseUserIdMap() = %v, want %v", got, tt.want)
		})
	}
}

// func TestSchedulerData_MapRooms(t *testing.T) {
// 	sptr := func(s string) *string { return &s }

// 	tests := []struct {
// 		name string
// 		data *SchedulerData
// 	}{
// 		{
// 			name: "Test 1",
// 			data: &SchedulerData{
// 				Rooms: []*RoomInfo{
// 					{Room: &Room{Id: sptr("Apple")}},
// 					{Room: &Room{Id: sptr("Pear")}}},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.data.MapRooms()
// 			for _, roomInfo := range tt.data.Rooms {
// 				assert.Same(t, roomInfo, tt.data.RoomsMap[*roomInfo.Id], "Expected pointers to point to the same object")
// 			}
// 		})
// 	}
// }

func TestSchedulerData_ApplyMapping(t *testing.T) {
	sptr := func(s string) *string { return &s }
	tests := []struct {
		name string
		data *SchedulerData
	}{
		{
			name: "Test 1",
			data: &SchedulerData{
				Rooms: []*RoomInfo{
					{Room: &Room{Id: sptr("Apple")}},
					{Room: &Room{Id: sptr("Pear")}}},
				Users: []*User{
					{Id: sptr("Coconut")},
					{Id: sptr("Pineapple")},
				},
				Teams: []*TeamInfo{
					{Team: &Team{Id: sptr("Orange")}},
					{Team: &Team{Id: sptr("Lime")}},
				},
				Resources: Resources{
					{Id: sptr("Strawberry")},
					{Id: sptr("Banana")},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.data.ApplyMapping()
			for _, roomInfo := range tt.data.Rooms {
				assert.Same(t, roomInfo, tt.data.RoomsMap[*roomInfo.Id],
					"Expected pointers for rooms to point to the same object")
			}
			for _, user := range tt.data.Users {
				assert.Same(t, user, tt.data.UserMap[*user.Id],
					"Expected pointers for users to point to the same object")
			}
			for _, resource := range tt.data.Resources {
				assert.Same(t, resource, tt.data.ResourcesMap[*resource.Id],
					"Expected pointers for resources to point to the same object")
			}
			for _, team := range tt.data.Teams {
				assert.Same(t, team, tt.data.TeamsMap[*team.Id],
					"Expected pointers for teams to point to the same object")
			}
		})
	}
}
