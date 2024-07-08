package main

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 3321,
			},
			want: 2999,
		},
		{
			name: "case2",
			args: args{
				n: 10,
			},
			want: 9,
		},
		{
			name: "case3",
			args: args{
				n: 1234,
			},
			want: 1234,
		},
		{
			name: "case4",
			args: args{
				n: 332,
			},
			want: 299,
		},
		{
			name: "case5",
			args: args{
				n: 333,
			},
			want: 333,
		},
		{
			name: "case6",
			args: args{
				n: 1231,
			},
			want: 1199,
		},
		{
			name: "case7",
			args: args{
				n: 1243,
			},
			want: 1239,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.n); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
