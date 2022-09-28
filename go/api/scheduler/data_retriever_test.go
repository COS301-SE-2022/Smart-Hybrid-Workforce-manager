package scheduler

import (
	"api/data"
	"encoding/json"
	"fmt"
	ts "lib/test_setup"
	tu "lib/testutils"
	"testing"
	"time"
)

func TestGetSchedulerData(t *testing.T) {
	err := ts.ConnectDB(t)
	if err != nil {
		t.Fatal(err)
	}
	defer ts.DisconnectDB(t)
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *SchedulerData
		wantErr bool
	}{
		{
			name: "Get data successfully",
			args: args{
				from: time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC),
				to:   time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSchedulerData(tt.args.from, tt.args.to, tu.Ptr("DESK"))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchedulerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, err := json.MarshalIndent(got, "", "  ")
			if err != nil {
				t.Errorf("Could not unmarshall json")
			}
			fmt.Println(string(b))
		})
	}
}

func TestGroupByBuilding(t *testing.T) {

	__testSchedulerData := &SchedulerData{
		Buildings: []*BuildingInfo{
			{
				Building: &data.Building{
					Id: tu.Ptr("Building 1"),
				},
				RoomIds: []string{"Freezer", "Countertop"},
			},
			{
				Building: &data.Building{
					Id: tu.Ptr("Building 2"),
				},
				RoomIds: []string{"Fridge", "Pantry"},
			},
		},
		Teams: []*TeamInfo{
			{
				Team:    &data.Team{Id: tu.Ptr("Cabbage"), Priority: tu.Ptr(2)},
				UserIds: []string{"Lime", "Lemon", "Grapefruit", "Banana"},
			},
			{
				Team:    &data.Team{Id: tu.Ptr("Broccoli")},
				UserIds: []string{"Blueberry", "Gooseberry", "Lemon"},
			},
			{
				Team:    &data.Team{Id: tu.Ptr("Lettuce"), Priority: tu.Ptr(0)},
				UserIds: []string{"Strawberry", "Blueberry"},
			},
			{
				Team:    &data.Team{Id: tu.Ptr("Eggplant")},
				UserIds: []string{},
			},
		},
		Users: []*data.User{
			{Id: tu.Ptr("Lime"), PreferredDesk: tu.Ptr("Shelf3"), BuildingID: tu.Ptr("Building 1")},
			{Id: tu.Ptr("Lemon"), PreferredDesk: tu.Ptr("Shelf2"), BuildingID: tu.Ptr("Building 2")},
			{Id: tu.Ptr("Grapefruit"), PreferredDesk: tu.Ptr("Shelf1"), BuildingID: tu.Ptr("Building 1")},
			{Id: tu.Ptr("Banana"), PreferredDesk: tu.Ptr("Shelf10"), BuildingID: tu.Ptr("Building 1")},
			{Id: tu.Ptr("Blueberry"), PreferredDesk: tu.Ptr("Shelf100"), BuildingID: tu.Ptr("Building 2")},
			{Id: tu.Ptr("Gooseberry"), PreferredDesk: tu.Ptr("Shelf_not_exist"), BuildingID: tu.Ptr("Building 1")},
			{Id: tu.Ptr("Strawberry"), PreferredDesk: nil, BuildingID: tu.Ptr("Building 2")},
		},
		Rooms: []*RoomInfo{
			{Room: &data.Room{Id: tu.Ptr("Freezer"), BuildingId: tu.Ptr("Building 1")},
				ResourceIds: []string{"Shelf1", "Shelf2", "Shelf3", "Shelf4", "Shelf5"}},
			{Room: &data.Room{Id: tu.Ptr("Fridge"), BuildingId: tu.Ptr("Building 2")},
				ResourceIds: []string{"Shelf10", "Shelf20", "Shelf30", "Shelf40", "Shelf50"}},
			{Room: &data.Room{Id: tu.Ptr("Pantry"), BuildingId: tu.Ptr("Building 2")},
				ResourceIds: []string{"Shelf100", "Shelf200", "Shelf300"}},
			{Room: &data.Room{Id: tu.Ptr("Countertop"), BuildingId: tu.Ptr("Building 1")},
				ResourceIds: []string{"Shelf_1", "Shelf_2", "Shelf_3"}},
		},
		Resources: []*data.Resource{
			{Id: tu.Ptr("Shelf1"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(1.0), YCoord: tu.Ptr(1.0)},
			{Id: tu.Ptr("Shelf2"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(2.0), YCoord: tu.Ptr(2.0)},
			{Id: tu.Ptr("Shelf3"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(3.0), YCoord: tu.Ptr(3.0)},
			{Id: tu.Ptr("Shelf4"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(4.0), YCoord: tu.Ptr(4.0)},
			{Id: tu.Ptr("Shelf5"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(5.0), YCoord: tu.Ptr(5.0)},
			{Id: tu.Ptr("Shelf10"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(10.0), YCoord: tu.Ptr(10.0)},
			{Id: tu.Ptr("Shelf20"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(20.0), YCoord: tu.Ptr(20.0)},
			{Id: tu.Ptr("Shelf30"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(30.0), YCoord: tu.Ptr(30.0)},
			{Id: tu.Ptr("Shelf40"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(40.0), YCoord: tu.Ptr(40.0)},
			{Id: tu.Ptr("Shelf50"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(50.0), YCoord: tu.Ptr(50.0)},
			{Id: tu.Ptr("Shelf100"), RoomId: tu.Ptr("Pantry"), XCoord: tu.Ptr(100.0), YCoord: tu.Ptr(100.0)},
			{Id: tu.Ptr("Shelf200"), RoomId: tu.Ptr("Pantry"), XCoord: tu.Ptr(200.0), YCoord: tu.Ptr(200.0)},
			{Id: tu.Ptr("Shelf300"), RoomId: tu.Ptr("Pantry"), XCoord: tu.Ptr(300.0), YCoord: tu.Ptr(300.0)},
			{Id: tu.Ptr("Shelf_1"), RoomId: tu.Ptr("Countertop"), XCoord: tu.Ptr(0.1), YCoord: tu.Ptr(0.1)},
			{Id: tu.Ptr("Shelf_2"), RoomId: tu.Ptr("Countertop"), XCoord: tu.Ptr(0.2), YCoord: tu.Ptr(0.2)},
			{Id: tu.Ptr("Shelf_3"), RoomId: tu.Ptr("Countertop"), XCoord: tu.Ptr(0.3), YCoord: tu.Ptr(0.3)},
		},
		CurrentBookings: &data.Bookings{},
		PastBookings:    &data.Bookings{},
	}
	type args struct {
		schedulerData *SchedulerData
	}
	tests := []struct {
		name string
		args args
		want map[string]*SchedulerData
	}{
		{
			name: "Test 1",
			args: args{
				schedulerData: __testSchedulerData,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GroupByBuilding(tt.args.schedulerData)
		})
	}
}
