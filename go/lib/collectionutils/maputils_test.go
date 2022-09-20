package collectionutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapHasKey(t *testing.T) {
	type args[K comparable, V any] struct {
		_map map[K]V
		key  K
	}
	tests := []struct {
		name string
		args args[string, any]
		want bool
	}{
		{
			name: "Key that IS contained",
			args: args[string, any]{
				_map: map[string]any{
					"apple":  2,
					"orange": "3",
					"ananas": true,
				},
				key: "ananas",
			},
			want: true,
		},
		{
			name: "Key that IS NOT contained",
			args: args[string, any]{
				_map: map[string]any{
					"apple":  2,
					"orange": "3",
					"ananas": true,
				},
				key: "tomato",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MapHasKey(tt.args._map, tt.args.key), "Expected %v", tt.want)
		})
	}
}
