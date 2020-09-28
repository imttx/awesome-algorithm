package offer

import "testing"

func Test_replaceSpaceV1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "We are happy."}, want: "We%20are%20happy."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpaceV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
