package main

import "testing"

func Test_verifyPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				preorder: []int{5, 2, 1, 3, 6},
			},
			want: true,
		},
		{
			args: args{
				preorder: []int{5, 2, 6, 1, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPreorder(tt.args.preorder); got != tt.want {
				t.Errorf("verifyPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
