package ga

import (
	tu "lib/testutils"
	"reflect"
	"scheduler/data"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var __test_domain *Domain = &Domain{
	SchedulerData: &data.SchedulerData{
		Teams: []*data.TeamInfo{
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
			{Id: tu.Ptr("Lime"), PreferredDesk: tu.Ptr("Shelf3")},
			{Id: tu.Ptr("Lemon"), PreferredDesk: tu.Ptr("Shelf2")},
			{Id: tu.Ptr("Grapefruit"), PreferredDesk: tu.Ptr("Shelf1")},
			{Id: tu.Ptr("Banana"), PreferredDesk: tu.Ptr("Shelf10")},
			{Id: tu.Ptr("Blueberry"), PreferredDesk: tu.Ptr("Shelf100")},
			{Id: tu.Ptr("Gooseberry"), PreferredDesk: tu.Ptr("Shelf_not_exist")},
			{Id: tu.Ptr("Strawberry"), PreferredDesk: nil},
		},
		Rooms: []*data.RoomInfo{
			{Room: &data.Room{Id: tu.Ptr("Freezer")}},
			{Room: &data.Room{Id: tu.Ptr("Fridge")}},
			{Room: &data.Room{Id: tu.Ptr("Pantry")}},
			{Room: &data.Room{Id: tu.Ptr("Countertop")}},
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
	},
	Map: map[int]string{
		0: "Lime",
		1: "Lemon",
		2: "Grapefruit",
		3: "Lime",
		4: "Grapefruit",
		5: "Blueberry",
		6: "Gooseberry",
		7: "Strawberry",
		8: "Blueberry",
	},
	InverseMap: map[string][]int{
		"Lime":       {0, 3},
		"Lemon":      {1},
		"Grapefruit": {4, 2},
		"Blueberry":  {5, 8},
		"Gooseberry": {6},
		"Strawberry": {7},
	},
}

func Test_distanceRadicand(t *testing.T) {
	type args struct {
		origin []float64
		coord  []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				origin: []float64{0.0, 0.0},
				coord:  []float64{4.0, 5.0},
			},
			want: 41.0,
		},
		{
			name: "Test 2",
			args: args{
				origin: []float64{1.0, 2.0},
				coord:  []float64{4.0, 6.0},
			},
			want: 25.0,
		},
		{
			name: "Test 3",
			args: args{
				origin: []float64{1.0, 2.5},
				coord:  []float64{-1.0, 6.0},
			},
			want: 16.25,
		},
		{
			name: "Test 4",
			args: args{
				origin: []float64{1.0},
				coord:  []float64{-1.0},
			},
			want: 4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceRadicand(tt.args.origin, tt.args.coord); got != tt.want {
				t.Errorf("distanceRadicand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCentroid(t *testing.T) {
	type args struct {
		coords [][]float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test 1",
			args: args{
				coords: [][]float64{{-1, -1}, {1, 3}, {0, 0}, {2, -1}, {2, 3}},
			},
			want: []float64{0.8, 0.8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCentroid(tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCentroid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avgDistanceFromCentroid(t *testing.T) {
	type args struct {
		coords [][]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				coords: [][]float64{{-1, -1}, {1, 3}, {0, 0}, {2, -1}, {2, 3}},
			},
			want: 2.1110702096228455,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avgDistanceFromCentroid(tt.args.coords); got != tt.want {
				t.Errorf("avgDistanceFromCentroid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndividual_getTeamsGroupedByRooms(t *testing.T) {
	type args struct {
		domain *Domain
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
		want       []teamRoomGroups
	}{
		{
			name: "Test 1",
			args: args{
				domain: __test_domain,
			},
			individual: &Individual{
				Gene: [][]string{
					//    0        1         2           3          4          5         6         7         8
					// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
					{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
				},
			},
			want: []teamRoomGroups{
				{
					teamId: "Broccoli",
					roomGroups: map[string][]int{
						"Freezer": {5, 6, 1},
						"Fridge":  {8},
					},
				},
				{
					teamId: "Cabbage",
					roomGroups: map[string][]int{
						"Freezer": {0, 1},
						"Fridge":  {2, 4},
						"Pantry":  {3},
					},
				},
				{
					teamId:     "Eggplant",
					roomGroups: map[string][]int{},
				},
				{
					teamId: "Lettuce",
					roomGroups: map[string][]int{
						"Freezer": {7, 5},
						"Fridge":  {8},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.domain.SchedulerData.ApplyMapping()
			got := tt.individual.getTeamsGroupedByRooms(tt.args.domain)
			assert.Len(t, got, len(tt.want), "Expected got to have a length of %v but got len %v", len(tt.want), len(got))
			sort.Slice(got, func(i, j int) bool {
				return got[i].teamId < got[j].teamId
			})
			for i := range tt.want {
				tu.MapsWithcSlicesMatchLoosely(t, got[i].roomGroups, tt.want[i].roomGroups)
			}
		})
	}
}

func TestIndividual_getUserCoordinate(t *testing.T) {
	mockSchedulerData := data.SchedulerData{
		Resources: data.Resources{
			&data.Resource{Id: tu.Ptr("Apple"), XCoord: tu.Ptr(1.1), YCoord: tu.Ptr(11.1)},
			&data.Resource{Id: tu.Ptr("Pear"), XCoord: tu.Ptr(2.2), YCoord: tu.Ptr(22.2)},
			&data.Resource{Id: tu.Ptr("Banana"), XCoord: tu.Ptr(3.3), YCoord: tu.Ptr(33.3)},
		},
	}
	mockDomain := &Domain{
		SchedulerData: &mockSchedulerData,
	}
	mockIndividual := &Individual{
		Gene: [][]string{
			{"Apple", "Pear", "Banana"},
		},
	}
	type args struct {
		domain *Domain
		index  int
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
		want       []float64
	}{
		{
			name:       "Test 1",
			individual: mockIndividual,
			args: args{
				domain: mockDomain,
				index:  0,
			},
			want: []float64{1.1, 11.1},
		},
		{
			name:       "Test 2",
			individual: mockIndividual,
			args: args{
				domain: mockDomain,
				index:  1,
			},
			want: []float64{2.2, 22.2},
		},
		{
			name:       "Test 3",
			individual: mockIndividual,
			args: args{
				domain: mockDomain,
				index:  2,
			},
			want: []float64{3.3, 33.3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.domain.SchedulerData.ApplyMapping()
			if got := tt.individual.getUserCoordinate(tt.args.domain, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Individual.getUserCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndividual_getTeamRoomProximities(t *testing.T) {
	type args struct {
		domain *Domain
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
		want       []teamRoomProximity
	}{
		{
			name: "Test 1",
			args: args{
				domain: __test_domain,
			},
			individual: &Individual{
				Gene: [][]string{
					//    0        1         2           3          4          5         6         7         8
					// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
					{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
				},
			},
			want: []teamRoomProximity{
				{
					teamRoomGroups: teamRoomGroups{
						teamId: "Broccoli",
						roomGroups: map[string][]int{
							"Freezer": {5, 6, 1},
							"Fridge":  {8},
						},
					},
					roomProximities: map[string]float64{
						"Freezer": 1.5713484026367723,
						"Fridge":  0.0,
					},
				},
				{
					teamRoomGroups: teamRoomGroups{
						teamId: "Cabbage",
						roomGroups: map[string][]int{
							"Freezer": {0, 1},
							"Fridge":  {2, 4},
							"Pantry":  {3},
						},
					},
					roomProximities: map[string]float64{
						"Freezer": 0.7071067811865476,
						"Fridge":  14.142135623730951,
						"Pantry":  0.0,
					},
				},
				{
					teamRoomGroups: teamRoomGroups{
						teamId:     "Eggplant",
						roomGroups: map[string][]int{},
					},
					roomProximities: map[string]float64{},
				},
				{
					teamRoomGroups: teamRoomGroups{
						teamId: "Lettuce",
						roomGroups: map[string][]int{
							"Freezer": {7, 5},
							"Fridge":  {8},
						},
					},
					roomProximities: map[string]float64{
						"Freezer": 0.7071067811865476,
						"Fridge":  0.0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.domain.SchedulerData.ApplyMapping()
			got := tt.individual.getTeamRoomProximities(tt.args.domain)
			assert.Len(t, got, len(tt.want), "Expected got to have a length of %v but got len %v", len(tt.want), len(got))
			sort.Slice(got, func(i, j int) bool {
				return got[i].teamId < got[j].teamId
			})
			for i := range tt.want {
				assert.InDeltaMapValues(t, got[i].roomProximities, tt.want[i].roomProximities, 0.0001, "Individual.getTeamRoomProximities() = %v, want %v")
			}
		})
	}
}

func Test_individualTeamProximityScore(t *testing.T) {
	type args struct {
		teamRoomProx teamRoomProximity
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				teamRoomProximity{
					teamRoomGroups: teamRoomGroups{
						teamId: "Broccoli",
						roomGroups: map[string][]int{
							"Freezer": {5, 6, 1},
							"Fridge":  {8},
						},
					},
					roomProximities: map[string]float64{
						"Freezer": 1.5713484026367723,
						"Fridge":  0.0,
					},
				},
			},
			want: 1.571348,
		},
		{
			name: "Test 2",
			args: args{
				teamRoomProximity{
					teamRoomGroups: teamRoomGroups{
						teamId: "Cabbage",
						roomGroups: map[string][]int{
							"Freezer": {0, 1},
							"Fridge":  {2, 4},
							"Pantry":  {3},
						},
					},
					roomProximities: map[string]float64{
						"Freezer": 0.7071067811865476,
						"Fridge":  14.142135623730951,
						"Pantry":  0.0,
					},
				},
			},
			want: 0.7071067 + 14.1421356,
		},
		{
			name: "Test 3",
			args: args{
				teamRoomProximity{
					teamRoomGroups: teamRoomGroups{
						teamId:     "Eggplant",
						roomGroups: map[string][]int{},
					},
					roomProximities: map[string]float64{},
				},
			},
			want: 0.0,
		},
		{
			name: "Test 4",
			args: args{
				teamRoomProximity{
					teamRoomGroups: teamRoomGroups{
						teamId: "Lettuce",
						roomGroups: map[string][]int{
							"Freezer": {7, 5},
							"Fridge":  {8},
						},
					},
					roomProximities: map[string]float64{
						"Freezer": 0.7071067811865476,
						"Fridge":  0.0,
					},
				},
			},
			want: 0.7071067 + 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := individualTeamProximityScore(tt.args.teamRoomProx)
			assert.InDelta(t, tt.want, got, 0.0001, "individualTeamProximityScore() = %v, want %v", got, tt.want)
		})
	}
}

func TestIndividual_teamProximityScore(t *testing.T) {
	type args struct {
		domain *Domain
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
		want       float64
	}{
		{
			name: "Test 1",
			args: args{
				domain: __test_domain,
			},
			individual: &Individual{
				Gene: [][]string{
					//    0        1         2           3          4          5         6         7         8
					// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
					{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
				},
			},
			// want: 1.0/(0.7071067+0.0+1.0) + 2.0/(0.7071067+14.1421356+1.0) + 1.0/(1.571348+1.0) + 1.0/(0.0+1.0),
			want: 6.575731493546652,
		},
		{
			name: "Test 2",
			args: args{
				domain: &Domain{
					SchedulerData: &data.SchedulerData{
						Teams: []*data.TeamInfo{
							{
								Team:    &data.Team{Id: tu.Ptr("Cabbage")},
								UserIds: []string{"Lime", "Lemon", "Grapefruit", "Banana"},
							},
							{
								Team:    &data.Team{Id: tu.Ptr("Broccoli")},
								UserIds: []string{"Blueberry", "Gooseberry", "Lemon"},
							},
							{
								Team:    &data.Team{Id: tu.Ptr("Lettuce")},
								UserIds: []string{"Strawberry", "Blueberry"},
							},
							{
								Team:    &data.Team{Id: tu.Ptr("Eggplant")},
								UserIds: []string{},
							},
						},
						Rooms: []*data.RoomInfo{
							{Room: &data.Room{Id: tu.Ptr("Freezer")}},
							{Room: &data.Room{Id: tu.Ptr("Fridge")}},
							{Room: &data.Room{Id: tu.Ptr("Pantry")}},
							{Room: &data.Room{Id: tu.Ptr("Countertop")}},
						},
						Resources: []*data.Resource{
							{Id: tu.Ptr("Shelf1"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(11.26), YCoord: tu.Ptr(-52.67)},
							{Id: tu.Ptr("Shelf2"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(88.190), YCoord: tu.Ptr(45.005)},
							{Id: tu.Ptr("Shelf3"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(118.19), YCoord: tu.Ptr(198.85)},
							{Id: tu.Ptr("Shelf4"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(254.34), YCoord: tu.Ptr(17.31)},
							{Id: tu.Ptr("Shelf5"), RoomId: tu.Ptr("Freezer"), XCoord: tu.Ptr(21.26), YCoord: tu.Ptr(200.39)},
							{Id: tu.Ptr("Shelf10"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(373.57), YCoord: tu.Ptr(44.23)},
							{Id: tu.Ptr("Shelf20"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(88.19), YCoord: tu.Ptr(88.19)},
							{Id: tu.Ptr("Shelf30"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(87.42), YCoord: tu.Ptr(-23.47)},
							{Id: tu.Ptr("Shelf40"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(298.0), YCoord: tu.Ptr(43.46)},
							{Id: tu.Ptr("Shelf50"), RoomId: tu.Ptr("Fridge"), XCoord: tu.Ptr(165.0), YCoord: tu.Ptr(-24.23)},
							{Id: tu.Ptr("Shelf100"), RoomId: tu.Ptr("Pantry"), XCoord: tu.Ptr(178.96), YCoord: tu.Ptr(12.69)},
							{Id: tu.Ptr("Shelf200"), RoomId: tu.Ptr("Pantry"), XCoord: tu.Ptr(87.42), YCoord: tu.Ptr(-23.47)},
							{Id: tu.Ptr("Shelf300"), RoomId: tu.Ptr("Pantry"), XCoord: tu.Ptr(250.0), YCoord: tu.Ptr(50.0)},
							{Id: tu.Ptr("Shelf_1"), RoomId: tu.Ptr("Countertop"), XCoord: tu.Ptr(0.1), YCoord: tu.Ptr(0.1)},
							{Id: tu.Ptr("Shelf_2"), RoomId: tu.Ptr("Countertop"), XCoord: tu.Ptr(0.2), YCoord: tu.Ptr(0.2)},
							{Id: tu.Ptr("Shelf_3"), RoomId: tu.Ptr("Countertop"), XCoord: tu.Ptr(0.3), YCoord: tu.Ptr(0.3)},
						},
					},
					InverseMap: map[string][]int{
						"Lime":       {0, 3},
						"Lemon":      {1},
						"Grapefruit": {4, 2},
						"Blueberry":  {5, 8},
						"Gooseberry": {6},
						"Strawberry": {7},
					},
				},
			},
			individual: &Individual{
				Gene: [][]string{
					//    0        1         2           3          4          5         6         7         8
					// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
					{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
				},
			},
			want: 4.5649235352205375,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.domain.SchedulerData.ApplyMapping()
			got := tt.individual.teamProximityScore(tt.args.domain)
			assert.InDelta(t, tt.want, got, 0.0001, "individualTeamProximityScore() = %v, want %v", got, tt.want)
		})
	}
}

func Test_getUserPreferredResource(t *testing.T) {
	__test_domain.SchedulerData.ApplyMapping()
	type args struct {
		domain    *Domain
		userIndex int
	}
	tests := []struct {
		name string
		args args
		want *data.Resource
	}{
		{
			name: "Test 1",
			args: args{
				domain:    __test_domain,
				userIndex: 0,
			},
			want: __test_domain.SchedulerData.ResourcesMap["Shelf3"],
		},
		{
			name: "Test 2",
			args: args{
				domain:    __test_domain,
				userIndex: 2,
			},
			want: __test_domain.SchedulerData.ResourcesMap["Shelf1"],
		},
		{
			name: "Test 3",
			args: args{
				domain:    __test_domain,
				userIndex: 5,
			},
			want: __test_domain.SchedulerData.ResourcesMap["Shelf100"],
		},
		{
			name: "Test 4",
			args: args{
				domain:    __test_domain,
				userIndex: 6,
			},
			want: nil,
		},
		{
			name: "Test 5",
			args: args{
				domain:    __test_domain,
				userIndex: 7,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserPreferredResource(tt.args.domain, tt.args.userIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserPreferredResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userPreferredDeskProximity(t *testing.T) {
	__test_domain.SchedulerData.ApplyMapping()
	gene := [][]string{
		//    0        1         2           3          4          5         6         7         8
		// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
		{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
	}
	type args struct {
		indiv     *Individual
		domain    *Domain
		userIndex int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				indiv: &Individual{
					Gene: gene,
				},
				domain:    __test_domain,
				userIndex: 1,
			},
			want: 0.0,
		},
		{
			name: "Test 2",
			args: args{
				indiv: &Individual{
					Gene: gene,
				},
				domain:    __test_domain,
				userIndex: 0,
			},
			want: 2.828427125,
		},
		{
			name: "Test 3",
			args: args{
				indiv: &Individual{
					Gene: gene,
				},
				domain:    __test_domain,
				userIndex: 7,
			},
			want: -1.0,
		},
		{
			name: "Test 4",
			args: args{
				indiv: &Individual{
					Gene: gene,
				},
				domain:    __test_domain,
				userIndex: 6,
			},
			want: -1.0,
		},
		{
			name: "Test 5",
			args: args{
				indiv: &Individual{
					Gene: gene,
				},
				domain:    __test_domain,
				userIndex: 2,
			},
			want: -2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := userPreferredDeskProximity(tt.args.indiv, tt.args.domain, tt.args.userIndex)
			assert.InDeltaf(t, tt.want, got, 0.0001, "userPreferredDeskProximity() = %v, want %v", got, tt.want)
		})
	}
}

func TestIndividual_preferredDeskBonuses(t *testing.T) {
	__test_domain.SchedulerData.ApplyMapping()
	gene := [][]string{
		//    0        1         2           3          4          5         6         7         8
		// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
		{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
	}
	type args struct {
		domain *Domain
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
		want       float64
	}{
		{
			name: "Test 1",
			individual: &Individual{
				Gene: gene,
			},
			args: args{
				domain: __test_domain,
			},
			want: 0.1401337,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.individual.preferredDeskBonuses(tt.args.domain)
			assert.InDeltaf(t, tt.want, got, 0.0001, "Individual.preferredDeskBonuses() = %v, want %v", got, tt.want)
		})
	}
}

func Test_dailyFitness(t *testing.T) {
	__test_domain.SchedulerData.ApplyMapping()
	gene := [][]string{
		//    0        1         2           3          4          5         6         7         8
		// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
		{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
	}
	type args struct {
		domain     *Domain
		individual *Individual
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				domain: __test_domain,
				individual: &Individual{
					Gene:    gene,
					Fitness: -1.0,
				},
			},
			want: 6.715865257431512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyFitness(tt.args.domain, tt.args.individual)
			assert.InDeltaf(t, tt.want, got, 0.0001, "dailyFitness() = %v, want %v", got, tt.want)
		})
	}
}

func TestDailyFitness(t *testing.T) {
	gene := [][]string{
		//    0        1         2           3          4          5         6         7         8
		// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
		{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
	}
	__test_domain.SchedulerData.ApplyMapping()
	type args struct {
		domain      *Domain
		individuals Individuals
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test 1",
			args: args{
				domain: __test_domain,
				individuals: []*Individual{
					{Gene: gene, Fitness: -1.0},
					{Gene: gene, Fitness: -1.0},
				},
			},
			want: []float64{6.71586525743, 6.71586525743},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DailyFitness(tt.args.domain, tt.args.individuals)
			assert.InDeltaSlice(t, tt.want, got, 0.0001, "DailyFitness() = %v, want %v", got, tt.want)
			for i, fitness := range tt.want {
				assert.InDeltaf(t, fitness, tt.args.individuals[i].Fitness, 0.0001, "Expected to individual fitness to be %v, got %v", fitness, tt.args.individuals[i].Fitness)
			}
		})
	}
}

func Test_teamUsersCountByDay(t *testing.T) {
	// gene := [][]string{
	// 	{"Lime", "", "Lemon", "Strawberry", "Grapefruit", "Blueberry"},
	// 	{"", "Blueberry", "", "Lemon", "", },
	// 	{"Banana", "", "", "", "Gooseberry", },
	// }
	type args struct {
		domain    *Domain
		dailyMaps []map[string]int
	}
	tests := []struct {
		name string
		args args
		want []map[string]int
	}{
		{
			name: "Test 1",
			args: args{
				domain: __test_domain,
				dailyMaps: []map[string]int{
					{
						"Lime":       1,
						"Lemon":      1,
						"Strawberry": 1,
						"Grapefruit": 1,
						"Blueberry":  1,
					},
					{
						"Blueberry": 1,
						"Lemon":     1,
					},
					{
						"Banana":     1,
						"Gooseberry": 1,
					},
				},
			},
			want: []map[string]int{
				{
					"Cabbage":  3,
					"Broccoli": 2,
					"Lettuce":  2,
				},
				{
					"Cabbage":  1,
					"Broccoli": 2,
					"Lettuce":  1,
				},
				{
					"Cabbage":  1,
					"Broccoli": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := teamUsersCountByDay(tt.args.domain, tt.args.dailyMaps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsersCountByDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndividual_getUserCountMapsPerDay(t *testing.T) {
	gene := [][]string{
		{"Lime", "", "Lemon", "Strawberry", "Grapefruit", "Blueberry"},
		{"", "Blueberry", "", "Lemon", ""},
		{"Banana", "", "", "", "Gooseberry"},
	}
	tests := []struct {
		name       string
		individual *Individual
		want       []map[string]int
	}{
		{
			name: "Test 1",
			individual: &Individual{
				Gene: gene,
			},
			want: []map[string]int{
				{
					"Lime":       1,
					"Lemon":      1,
					"Strawberry": 1,
					"Grapefruit": 1,
					"Blueberry":  1,
				},
				{
					"Blueberry": 1,
					"Lemon":     1,
				},
				{
					"Banana":     1,
					"Gooseberry": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.individual.getUserCountMapsPerDay(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Individual.getUserCountMapsPerDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
