package data

import (
	"lib/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResource_GetDecorations(t *testing.T) {
	tests := []struct {
		name     string
		r        *Resource
		newDecor *string
		want     map[string]any
		newWant  map[string]any
	}{
		{
			name: "Test 1",
			r: &Resource{
				Decorations: testutils.Ptr(
					`{"capacity":5,"status":"good"}`,
				),
			},
			newDecor: testutils.Ptr(`{"capacity":6,"status":"good"}`),
			want: map[string]any{
				"capacity": 5.0,
				"status":   "good",
			},
			newWant: map[string]any{
				"capacity": 6.0,
				"status":   "good",
			},
		},
		{
			name: "Test 2",
			r: &Resource{
				Decorations: nil,
			},
			newDecor: testutils.Ptr(`{"capacity":6,"status":"good"}`),
			want:     map[string]any{},
			newWant: map[string]any{
				"capacity": 6.0,
				"status":   "good",
			},
		},
		{
			name: "Test 3",
			r: &Resource{
				Decorations: nil,
			},
			newDecor: nil,
			want:     map[string]any{},
			newWant:  map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.GetDecorations()
			assert.Equalf(t, tt.want, got, "Resource.ParseDecorations() = %v, want %v", got, tt.want)
			got = tt.r.GetDecorations()
			assert.Equalf(t, tt.want, got, "Resource.ParseDecorations() = %v, want %v", got, tt.want)
			tt.r.Decorations = tt.newDecor
			got = tt.r.GetDecorations()
			assert.Equalf(t, tt.newWant, got, "Resource.ParseDecorations() = %v, want %v", got, tt.newWant)
		})
	}
}

func TestResource_GetCapacity(t *testing.T) {
	tests := []struct {
		name string
		r    *Resource
		want int
	}{
		{
			name: "Test 1",
			r: &Resource{
				Decorations: testutils.Ptr(`{"capacity":5}`),
			},
			want: 5,
		},
		{
			name: "Test 2",
			r: &Resource{
				Decorations: testutils.Ptr(`{"capacity":"wrong format"}`),
			},
			want: -1,
		},
		{
			name: "Test 3",
			r: &Resource{
				Decorations: testutils.Ptr(`{}`),
			},
			want: -1,
		},
		{
			name: "Test 3",
			r: &Resource{
				Decorations: nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetCapacity(); got != tt.want {
				t.Errorf("Resource.GetCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}
