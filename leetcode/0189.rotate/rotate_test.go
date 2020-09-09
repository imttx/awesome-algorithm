package leetcode

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}},
		{args: args{nums: []int{4, 5, 6, 7, 1, 2, 3}, k: 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.nums)
			rotate(tt.args.nums, tt.args.k)
			t.Log(tt.args.nums)
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.nums)
		})
	}
}

func Test_rotate2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}},
		{args: args{nums: []int{4, 5, 6, 7, 1, 2, 3}, k: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.nums)
			rotate2(tt.args.nums, tt.args.k)
			t.Log(tt.args.nums)
		})
	}
}
