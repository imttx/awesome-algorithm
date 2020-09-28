package offer

import (
	"testing"
)

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{nums: []int{0, 0, 1, 2, 5}}, want: true},
		{args: args{nums: []int{1, 2, 4, 3, 5}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.nums); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStraight2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{nums: []int{0, 0, 1, 2, 5}}, want: true},
		{args: args{nums: []int{1, 2, 4, 3, 5}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight2(tt.args.nums); got != tt.want {
				t.Errorf("isStraight2() = %v, want %v", got, tt.want)
			}
		})
	}
}
