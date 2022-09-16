package ga

import (
	"testing"

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
