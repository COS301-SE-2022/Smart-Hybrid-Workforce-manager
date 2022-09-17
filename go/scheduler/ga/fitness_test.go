package ga

import (
	"reflect"
	"testing"
)

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
