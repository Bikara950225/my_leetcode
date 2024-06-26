package main

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "bcabc",
			},
			want: "abc",
		},
		{
			args: args{
				s: "cbacdcbc",
			},
			want: "acdb",
		},
		{
			args: args{
				s: "cabd",
			},
			want: "cabd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLetters2() = %v, want %v", got, tt.want)
			}
		})
	}
}
