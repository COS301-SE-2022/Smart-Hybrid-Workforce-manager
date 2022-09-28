package ga

import (
	"testing"
)

func TestIndividual_CheckIfValidDaily(t *testing.T) {
	tests := []struct {
		name       string
		individual *Individual
		want       bool
	}{
		{
			name: "Test 1",
			individual: &Individual{
				Gene: [][]string{
					{
						"A", "B", "C",
					},
				},
			},
			want: true,
		},
		{
			name: "Test 1",
			individual: &Individual{
				Gene: [][]string{
					{
						"A", "B", "B", "C",
					},
				},
			},
			want: false,
		},
		{
			name: "Test 3",
			individual: &Individual{
				Gene: [][]string{
					{},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.individual.CheckIfValidDaily(); got != tt.want {
				t.Errorf("Individual.CheckIfValidDaily() = %v, want %v", got, tt.want)
			}
		})
	}
}
