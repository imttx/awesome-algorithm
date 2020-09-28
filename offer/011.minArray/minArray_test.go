package offer

import (
	"testing"
)

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{numbers: []int{2, 2, 2, 0, 1}}, want: 0},
		{args: args{numbers: []int{3, 4, 5, 1, 2}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minArray2(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{numbers: []int{2, 2, 2, 0, 1}}, want: 0},
		{args: args{numbers: []int{3, 4, 5, 1, 2}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray2(tt.args.numbers); got != tt.want {
				t.Errorf("minArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}
