package offer

import "testing"

func Test_findRepeatNumberV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{2, 3, 1, 0, 2, 5, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberV1(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumberV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatNumberV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{2, 3, 1, 0, 2, 5, 3}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberV2(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
