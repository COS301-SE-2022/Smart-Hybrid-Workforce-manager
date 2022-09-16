package ga

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoPointSwap(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		xP1  int
		xP2  int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "Test1 filled arr",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  2,
				xP2:  6,
			},
			want:  []int{-1, 1, 20, 30, 40, 50, 6, 7, 8, 9},
			want1: []int{-10, 10, 2, 3, 4, 5, 60, 70, 80, 90},
		},
		{
			name: "Test2 crossover point at end",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  2,
				xP2:  10,
			},
			want:  []int{-1, 1, 20, 30, 40, 50, 60, 70, 80, 90},
			want1: []int{-10, 10, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test2 crossover point at start",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  0,
				xP2:  2,
			},
			want:  []int{-10, 10, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: []int{-1, 1, 20, 30, 40, 50, 60, 70, 80, 90},
		},
		{
			name: "Test2 crossover point 1 > crossover point 2",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  2,
				xP2:  1,
			},
			want:  []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := twoPointSwap(tt.args.arr1, tt.args.arr2, tt.args.xP1, tt.args.xP2)
			assert.True(t, reflect.DeepEqual(got, tt.want), "twoPointSwap() got = %v, want %v", got, tt.want)
			assert.True(t, reflect.DeepEqual(got1, tt.want1), "twoPointSwap() got = %v, want %v", got, tt.want)
		})
	}
}

func TestFindValid(t *testing.T) {
	type args struct {
		index          int
		parent         []string
		otherParentMap map[string]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Return mapped element",
			args: args{
				index:  0,
				parent: []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				otherParentMap: map[string]int{
					"c": 2,
					"d": 3,
					"z": 4,
					"e": 5,
				},
			},
			want: "m",
		},
		{
			name: "Return element at index",
			args: args{
				index:  1,
				parent: []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				otherParentMap: map[string]int{
					"c": 2,
					"d": 3,
					"z": 4,
					"e": 5,
				},
			},
			want: "f",
		},
		{
			name: "Return mapped element second test",
			args: args{
				index:  6,
				parent: []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				otherParentMap: map[string]int{
					"c": 2,
					"d": 3,
					"z": 4,
					"e": 5,
				},
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindValid(tt.args.index, tt.args.parent, tt.args.otherParentMap)
			assert.Equal(t, tt.want, got, "FindValid() = %v, want %v", got, tt.want)
		})
	}
}

func TestPMX(t *testing.T) {
	type args struct {
		p1  []string
		p2  []string
		xP1 int
		xP2 int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "No mapping required test, normal crossover",
			args: args{
				p1:  []string{"a", "b", "c", "d", "e", "f", "g", "h", "j"},
				p2:  []string{"h", "j", "f", "e", "d", "c", "g", "a", "b"},
				xP1: 2,
				xP2: 5,
			},
			want:  []string{"a", "b", "f", "e", "d", "c", "g", "h", "j"},
			want1: []string{"h", "j", "c", "d", "e", "f", "g", "a", "b"},
		},
		{
			name: "No mapping required test, crossover entire chromosome",
			args: args{
				p1:  []string{"a", "b", "c", "d", "e", "f", "g", "h", "j"},
				p2:  []string{"h", "j", "f", "e", "d", "c", "g", "a", "b"},
				xP1: 0,
				xP2: 9,
			},
			want:  []string{"h", "j", "f", "e", "d", "c", "g", "a", "b"},
			want1: []string{"a", "b", "c", "d", "e", "f", "g", "h", "j"},
		},
		{
			name: "Mapping required crossover",
			args: args{
				p1:  []string{"a", "b", "c", "d", "z", "e", "f", "g"},
				p2:  []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				xP1: 2,
				xP2: 6,
			},
			want:  []string{"z", "d", "e", "b", "a", "m", "f", "g"},
			want1: []string{"m", "f", "c", "d", "z", "e", "b", "g"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PMX(tt.args.p1, tt.args.p2, tt.args.xP1, tt.args.xP2)
			assert.True(t, reflect.DeepEqual(got, tt.want), "PMX() got = %v, want %v", got, tt.want)
			assert.True(t, reflect.DeepEqual(got1, tt.want1), "PMX() got1 = %v, want %v", got1, tt.want1)
		})
	}
}
