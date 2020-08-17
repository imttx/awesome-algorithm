package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{args: args{nums: []int{1, 2, 3, 4}, target: 6}, want: []int{1, 3}},
		{args: args{nums: []int{1, 2, 3, 4}, target: 5}, want: []int{1, 2}},
		{args: args{nums: []int{0, 8, 7, 3, 3, 4, 2}, target: 11}, want: []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
