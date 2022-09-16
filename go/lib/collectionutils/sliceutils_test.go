package collectionutils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy1DArr(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty arr",
			args: args{
				arr: []string{},
			},
			want: []string{},
		},
		{
			name: "Normal copy, non-empty arr",
			args: args{
				arr: []string{"pear", "banana", "apple"},
			},
			want: []string{"pear", "banana", "apple"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy1DArr(tt.args.arr)
			assert.True(t, reflect.DeepEqual(got, tt.want), "got=%v, expected=%v", got, tt.want)
		})
	}
}

func TestCopy2DArr(t *testing.T) {
	type args struct {
		arr [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Empty arr",
			args: args{
				arr: [][]string{},
			},
			want: [][]string{},
		},
		{
			name: "Normal copy, non-empty arr",
			args: args{
				arr: [][]string{
					{"pear", "banana", "apple"},
					{"cashew", "peanut"},
					{},
					{"pinenut"},
				},
			},
			want: [][]string{
				{"pear", "banana", "apple"},
				{"cashew", "peanut"},
				{},
				{"pinenut"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy2DArr(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy2DArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatten2DArr(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty arr",
			args: args{
				arr: [][]int{},
			},
			want: []int{},
		},
		{
			name: "Flatten non 2D arr",
			args: args{
				arr: [][]int{
					{1, 2, 3, 4},
					{50, 60},
					{},
					{700, 800},
				},
			},
			want: []int{1, 2, 3, 4, 50, 60, 700, 800},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten2DArr(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten2DArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartitionArray(t *testing.T) {
	type args struct {
		arr   []int
		sizes []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Empty arr, sizes={0}",
			args: args{
				arr:   []int{},
				sizes: []int{0},
			},
			want: [][]int{{}},
		},
		{
			name: "Empty arr, sizes={}",
			args: args{
				arr:   []int{},
				sizes: []int{},
			},
			want: [][]int{},
		},
		{
			name: "General partitioning",
			args: args{
				arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				sizes: []int{3, 2, 0, 4, 1},
			},
			want: [][]int{
				{0, 1, 2},
				{3, 4},
				{},
				{5, 6, 7, 8},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartitionArray(tt.args.arr, tt.args.sizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartitionArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty slice",
			args: args{
				s: []int{},
				e: 2,
			},
			want: false,
		},
		{
			name: "Contains",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				e: 2,
			},
			want: true,
		},
		{
			name: "Does not contain",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				e: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemElemenAtI(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Index > len of slice",
			args: args{
				slice: []int{1, 2},
				index: 3,
			},
			want: []int{1, 2},
		},
		{
			name: "Index < 0",
			args: args{
				slice: []int{1, 2},
				index: -1,
			},
			want: []int{1, 2},
		},
		{
			name: "Valid index, not last",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6},
				index: 2,
			},
			want: []int{1, 2, 6, 4, 5},
		},
		{
			name: "Last element",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6},
				index: 5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemElemenAtI(tt.args.slice, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemElemenAtI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSliceIntersection(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Both slices empty",
			args: args{
				slice1: []int{},
				slice2: []int{},
			},
			want: []int{},
		},
		{
			name: "Slice 1 empty",
			args: args{
				slice1: []int{},
				slice2: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			name: "Slice 2 empty",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{},
			},
			want: []int{},
		},
		{
			name: "Both contains distinct elements",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{4, 5, 6},
			},
			want: []int{},
		},
		{
			name: "Both contains some unique elements",
			args: args{
				slice1: []int{1, 2, 3, 4, 5},
				slice2: []int{4, 5, 6, 7, 8},
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntSliceIntersection(tt.args.slice1, tt.args.slice2)
			assert.ElementsMatch(t, got, tt.want, "Expected=%v, got=%v")
		})
	}
}

func TestSequentialSequence(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "start > end",
			args: args{
				start: 2,
				end:   1,
			},
			want: []int{},
		},
		{
			name: "start == end",
			args: args{
				start: 1,
				end:   1,
			},
			want: []int{},
		},
		{
			name: "Start=-1 end=5",
			args: args{
				start: -1,
				end:   5,
			},
			want: []int{-1, 0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SequentialSequence(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SequentialSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceIntersection(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Both slices empty",
			args: args{
				slice1: []string{},
				slice2: []string{},
			},
			want: []string{},
		},
		{
			name: "Slice 1 empty",
			args: args{
				slice1: []string{},
				slice2: []string{"1", "2", "3"},
			},
			want: []string{},
		},
		{
			name: "Slice 2 empty",
			args: args{
				slice1: []string{"1", "2", "3"},
				slice2: []string{},
			},
			want: []string{},
		},
		{
			name: "Both contains distinct elements",
			args: args{
				slice1: []string{"1", "2", "3"},
				slice2: []string{"4", "5", "6"},
			},
			want: []string{},
		},
		{
			name: "Both contains some unique elements",
			args: args{
				slice1: []string{"1", "2", "3", "4", "5"},
				slice2: []string{"4", "5", "6", "7", "8"},
			},
			want: []string{"4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceIntersection(tt.args.slice1, tt.args.slice2)
			assert.ElementsMatch(t, got, tt.want, "Expected=%v, got=%v")
		})
	}
}
