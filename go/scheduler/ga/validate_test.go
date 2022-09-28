package ga

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestIndividual_ValidateIndividual(t *testing.T) {
	type args struct {
		domain *Domain
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
	}{
		{
			name: "Test 1",
			args: args{
				domain: &Domain{
					Terminals: []string{"D", "F"},
				},
			},
			individual: &Individual{
				Gene: [][]string{
					{
						"A", "B", "C",
					},
				},
			},
		},
		{
			name: "Test 1",
			args: args{
				domain: &Domain{
					Terminals: []string{"D", "F"},
				},
			},
			individual: &Individual{
				Gene: [][]string{
					{
						"A", "B", "B", "C", "C",
					},
				},
			},
		},
		{
			name: "Test 3",
			args: args{
				domain: &Domain{
					Terminals: []string{"D", "F"},
				},
			},
			individual: &Individual{
				Gene: [][]string{
					{
						"A", "A", "A",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.individual.ValidateIndividual(tt.args.domain)
			assert.Truef(t, tt.individual.CheckIfValidDaily(), "Expected individual to be valid after call, but got %v", tt.individual)
		})
	}
}
