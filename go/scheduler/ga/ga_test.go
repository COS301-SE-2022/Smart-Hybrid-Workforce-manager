package ga

import (
	"testing"
)

func TestIndividual_String(t *testing.T) {
	tests := []struct {
		name       string
		individual Individual
		notWant    string
	}{
		{
			name: "Test 1",
			individual: Individual{
				Gene: [][]string{
					{"A", "B", "C"},
					{"1", "2", "3"},
				},
			},
			notWant: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.individual.String(); got == tt.notWant {
				t.Errorf("Individual.String() = %v, want %v", got, tt.notWant)
			}
		})
	}
}

func TestIndividual_StringDomain(t *testing.T) {
	type args struct {
		domain Domain
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
		notWant    string
	}{
		{
			name: "Test 1",
			args: args{
				domain: Domain{},
			},
			individual: &Individual{
				Gene: [][]string{
					{"A", "B", "C"},
					{"1", "2", "3"},
				},
			},
			notWant: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.individual.StringDomain(tt.args.domain); got == tt.notWant {
				t.Errorf("Individual.StringDomain() = %v, want %v", got, tt.notWant)
			}
		})
	}
}

func Test_printGAGraphs(t *testing.T) {
	type args struct {
		multiplier    float64
		maxMultiplier float64
		avg           float64
		maxFitness    float64
		minFitness    float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				multiplier:    1.0,
				maxMultiplier: 1.0,
				avg:           1.0,
				maxFitness:    1.0,
				minFitness:    1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printGAGraphs(tt.args.multiplier, tt.args.maxMultiplier, tt.args.avg, tt.args.maxFitness, tt.args.minFitness)
		})
	}
}
