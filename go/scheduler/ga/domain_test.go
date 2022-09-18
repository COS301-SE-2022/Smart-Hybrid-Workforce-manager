package ga

import (
	"reflect"
	"scheduler/data"
	"testing"

	tu "lib/testutils"

	"github.com/stretchr/testify/assert"
)

func TestDomain_GetRandomUniqueTerminalArrays(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name   string
		domain *Domain
		args   args
	}{
		{
			name: "General uniqueness test",
			args: args{
				length: 4,
			},
			domain: &Domain{
				Terminals: []string{"1", "2", "3", "4", "5", "6"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Repeat 500 times to be sure
			for i := 0; i < 500; i++ {
				arr := tt.domain.GetRandomUniqueTerminalArrays(tt.args.length)
				// Check that all elements are unique
				countMap := make(map[string]int)
				for _, v := range arr {
					countMap[v]++
					assert.LessOrEqual(t, countMap[v], 1, "Element {%v} appears %v times, when should only appear 0 or 1 times", v, countMap[v])
				}
			}
		})
	}
}

func TestDomain_GetTeamUserIndices(t *testing.T) {
	tests := []struct {
		name   string
		domain *Domain
		want   map[string][]int
	}{
		{
			name: "Test 1",
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
				},
				InverseMap: map[string][]int{
					"Lime":       {0, 3},
					"Lemon":      {1},
					"Grapefruit": {4},
					"Blueberry":  {5, 8},
					"Gooseberry": {6},
					"Strawberry": {7},
				},
			},
			want: map[string][]int{
				"Cabbage":  {0, 3, 1, 4},
				"Broccoli": {5, 8, 6, 1},
				"Lettuce":  {7, 5, 8},
				"Eggplant": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.domain.GetTeamUserIndices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Domain.GetTeamUserIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
