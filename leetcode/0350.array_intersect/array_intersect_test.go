package leetcode

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{9, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectV2(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectV2(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
