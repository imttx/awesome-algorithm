package offer

import (
	"testing"
)

func Test_findNumberIn2DArrayV1(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 5}, want: true},
		{args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 20}, want: false},
		{args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 17}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArrayV1(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArrayV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumberIn2DArrayV2(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 5}, want: true},
		{args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 20}, want: false},
		{args: args{matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, target: 17}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArrayV2(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArrayV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
