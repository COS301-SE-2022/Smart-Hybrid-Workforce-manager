package ga

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dailySwapMutate(t *testing.T) {
	type args struct {
		domain     *Domain
		individual *Individual
		swapAmount int
	}
	tests := []struct {
		name         string
		args         args
		originalGene [][]string
	}{
		{
			name: "Normal swap mutate, 1 swap",
			args: args{
				domain: nil,
				individual: &Individual{
					Gene: [][]string{
						{"a", "b", "c", "d", "e", "f", "g"},
					},
					Fitness: 0.0,
				},
				swapAmount: 1,
			},
			originalGene: [][]string{
				{"a", "b", "c", "d", "e", "f", "g"},
			},
		},
		{
			name: "Normal swap mutate, 1 swap",
			args: args{
				domain: nil,
				individual: &Individual{
					Gene: [][]string{
						{"a", "b", "c", "d", "e", "f", "g"},
					},
					Fitness: 0.0,
				},
				swapAmount: 3,
			},
			originalGene: [][]string{
				{"a", "b", "c", "d", "e", "f", "g"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parentGene := tt.args.individual.Gene
			childGene := dailySwapMutate(tt.args.domain, tt.args.individual, tt.args.swapAmount).Gene
			assert.ElementsMatch(t, parentGene[0], childGene[0], "Expected elements in individual to stay same, started with %v ended with %v")
			assert.True(t, reflect.DeepEqual(parentGene, tt.originalGene), "Parent gene has been changed")
		})
	}
}

func Test_dailyPullMutateValid(t *testing.T) {
	type args struct {
		domain     *Domain
		individual *Individual
		pullAmount int
	}
	tests := []struct {
		name              string
		args              args
		originalGene      [][]string
		originalTerminals []string
	}{
		{
			name: "Normal swap mutate, 1 swap",
			args: args{
				domain: &Domain{
					Terminals: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				},
				individual: &Individual{
					Gene: [][]string{
						{"a", "b", "c", "d", "e", "f", "g"},
					},
					Fitness: 0.0,
				},
				pullAmount: 3,
			},
			originalGene: [][]string{
				{"a", "b", "c", "d", "e", "f", "g"},
			},
			originalTerminals: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parentGene := tt.args.individual.Gene
			childGene := dailyPullMutateValid(tt.args.domain, tt.args.individual, tt.args.pullAmount).Gene
			countMap := make(map[string]int)
			for _, v := range childGene[0] {
				countMap[v]++
				assert.LessOrEqual(t, countMap[v], 1, "Element {%v} appears %v times, when should only appear 0 or 1 times", v, countMap[v])
			}
			assert.True(t, reflect.DeepEqual(parentGene, tt.originalGene), "Parent gene has been changed")
			assert.True(t, reflect.DeepEqual(tt.args.domain.Terminals, tt.originalTerminals), "Parent individual terminals have been changed")
		})
	}
}

func Test_validWeeklySwap(t *testing.T) {
	rand.Seed(1)
	type args struct {
		indiv          *Individual
		mutationDegree int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				indiv: &Individual{
					Fitness: 0.0,
					Gene: [][]string{
						{"Apple", "", "Orange", "", "Pear", "Banana"},
						{"Apple", "Orange", "Pear", "Banana"},
						{"Apple", "", "Orange", "", "", "Banana"},
						{},
						{"Cucumber", "Brusselsprout"},
						{"", "", ""},
						{"", "Orange", "", "", "", "", "Banana"},
					},
				},
				mutationDegree: 0,
			},
		},
		{
			name: "Test 2",
			args: args{
				indiv: &Individual{
					Fitness: 0.0,
					Gene: [][]string{
						{"Apple", "", "Orange", "", "Pear", "Banana"},
						{"Apple", "Orange", "Pear", "Banana"},
						{"Apple", "", "Orange", "", "", "Banana"},
						{},
						{"Cucumber", "Brusselsprout"},
						{"", "", ""},
						{"", "Orange", "", "", "", "", "Banana"},
					},
				},
				mutationDegree: 1,
			},
		},
		{
			name: "Test 3",
			args: args{
				indiv: &Individual{
					Fitness: 0.0,
					Gene: [][]string{
						{"Apple", "", "Orange", "", "Pear", "Banana"},
						{"Apple", "Orange", "Pear", "Banana"},
						{"Apple", "", "Orange", "", "", "Banana"},
						{},
						{"Cucumber", "Brusselsprout"},
						{"", "", ""},
						{"", "Orange", "", "", "", "", "Banana"},
					},
				},
				mutationDegree: 400,
			},
		},
		{
			name: "Test 4",
			args: args{
				indiv: &Individual{
					Fitness: 0.0,
					Gene: [][]string{
						{"", "", "", "", "", ""},
						{"", "", "", ""},
						{"", "", "", "", "", ""},
						{},
						{"", ""},
						{"", "", ""},
						{"", "", "", "", "", "", ""},
					},
				},
				mutationDegree: 100,
			},
		},
		{
			name: "Test 4",
			args: args{
				indiv: &Individual{
					Fitness: 0.0,
					Gene: [][]string{
						{},
						{},
						{},
						{},
						{},
						{},
						{},
					},
				},
				mutationDegree: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			expectedIdCounts := make(map[string]int)
			for _, day := range tt.args.indiv.Gene {
				for _, id := range day {
					expectedIdCounts[id]++
				}
			}
			// Ensure that individuals are still valid
			got := validWeeklySwap(tt.args.indiv.Clone(), tt.args.mutationDegree)
			idCounts := make(map[string]int)
			for _, day := range got.Gene {
				dailyCounts := make(map[string]int)
				for _, id := range day {
					idCounts[id]++
					if id != "" {
						dailyCounts[id]++
						assert.Equalf(t, 1, dailyCounts[id], "Expected that id's appear at most once per day, %v appeared %v times", id, dailyCounts[id])
					}
				}
			}
			assert.Equalf(t, expectedIdCounts, idCounts, "Expected IDs and blanks to appear same amount of time in mutated individuals as original, expected counts %v, recieved %v", expectedIdCounts, idCounts)
		})
	}
}
