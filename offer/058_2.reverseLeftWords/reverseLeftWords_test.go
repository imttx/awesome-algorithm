package offer

import (
	"testing"
)

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "abcdefg", n: 2}, want: "cdefgab"},
		{args: args{s: "lrloseumgh", n: 6}, want: "umghlrlose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseLeftWords2(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "abcdefg", n: 2}, want: "cdefgab"},
		{args: args{s: "lrloseumgh", n: 6}, want: "umghlrlose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords2(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseLeftWords3(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "abcdefg", n: 2}, want: "cdefgab"},
		{args: args{s: "lrloseumgh", n: 6}, want: "umghlrlose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords3(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords3() = %v, want %v", got, tt.want)
			}
		})
	}
}
